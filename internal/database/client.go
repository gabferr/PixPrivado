package database

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"pix-privado/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Client holds the database connection pool.
type Client struct {
	pool *pgxpool.Pool
}

// NewClient creates a new database connection pool.
func NewClient(ctx context.Context, dsn string) (*Client, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &Client{pool: pool}, nil
}

// Close closes the database connection pool.
func (c *Client) Close() {
	c.pool.Close()
}

// A função GetBalance deve agora escanear e retornar uma string.
func (c *Client) GetBalance(ctx context.Context, accountID uuid.UUID) (string, error) {
	var balance string // <-- O tipo da variável de destino agora é string
	query := "SELECT balance FROM accounts WHERE id = $1"

	// O Scan irá ler o valor NUMERIC do banco e colocá-lo na string.
	err := c.pool.QueryRow(ctx, query, accountID).Scan(&balance)
	if err != nil {
		return "", err // <-- Retorna uma string vazia em caso de erro
	}

	return balance, nil
}

// GetAccountByUUID retrieves an account from the database by its UUID.
func (c *Client) GetAccountByUUID(ctx context.Context, accountID uuid.UUID) (*models.Account, error) {
	var account models.Account
	query := "SELECT id, owner_id, blockchain_address, status, created_at, updated_at FROM accounts WHERE id = $1"
	err := c.pool.QueryRow(ctx, query, accountID).Scan(&account.ID, &account.OwnerID, &account.BlockchainAddress, &account.Status, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}
	return &account, nil
}

// =============================================================================
// NOVA FUNÇÃO ADICIONADA AQUI
// =============================================================================

// UpdateBalanceByAddress atualiza o saldo de uma conta com base no seu endereço blockchain.
// O valor (amount) pode ser positivo (crédito) ou negativo (débito).
func (c *Client) UpdateBalanceByAddress(ctx context.Context, address common.Address, amount *big.Int) error {
	// A coluna `blockchain_address` deve ser do tipo TEXT ou VARCHAR e conter o endereço em formato hexadecimal.
	// É uma boa prática armazenar endereços em minúsculas para consistência.
	addressHex := address.Hex()

	// A coluna `balance` deve ser do tipo NUMERIC para acomodar valores grandes da blockchain.
	// A operação `balance = balance + $1` é atômica no banco de dados.
	query := `UPDATE accounts SET balance = balance + $1 WHERE blockchain_address = $2`

	// Convertemos o *big.Int para string para passá-lo ao driver do banco de dados.
	// Esta é a maneira mais segura de garantir que não haja perda de precisão.
	cmdTag, err := c.pool.Exec(ctx, query, amount.String(), addressHex)
	if err != nil {
		return fmt.Errorf("erro ao atualizar saldo para o endereço %s: %w", addressHex, err)
	}

	if cmdTag.RowsAffected() == 0 {
		log.Printf("Aviso: Nenhuma conta encontrada com o endereço %s para atualizar o saldo.", addressHex)
		// Pode ser útil retornar um erro específico se a conta não for encontrada.
		// return fmt.Errorf("nenhuma conta encontrada com o endereço %s", addressHex)
	} else {
		log.Printf("Sincronizador: Saldo da conta %s atualizado com o valor %s.", addressHex, amount.String())
	}

	return nil
}

// =============================================================================

// CheckIfRequestExists checks for idempotency in the transactions table.
func (c *Client) CheckIfRequestExists(ctx context.Context, requestID string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM transactions WHERE request_id = $1)"
	err := c.pool.QueryRow(ctx, query, requestID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CreateTransaction creates a new transaction record in the database.
func (c *Client) CreateTransaction(ctx context.Context, tx *models.Transaction) error {
	query := "INSERT INTO transactions (id, request_id, from_account_id, to_account_id, amount, memo, status, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := c.pool.Exec(ctx, query, tx.ID, tx.RequestID, tx.FromAccountID, tx.ToAccountID, tx.Amount, tx.Memo, tx.Status, tx.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

// UpdateTransactionStatus updates the status and blockchain_tx for a transaction.
func (c *Client) UpdateTransactionStatus(ctx context.Context, transactionID uuid.UUID, status string, txHash string) error {
	query := "UPDATE transactions SET status = $1, blockchain_tx = $2, updated_at = NOW() WHERE id = $3"
	_, err := c.pool.Exec(ctx, query, status, txHash, transactionID)
	if err != nil {
		return fmt.Errorf("failed to update transaction status: %w", err)
	}
	return nil
}
