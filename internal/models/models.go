package models

import (
	"time"

	"github.com/google/uuid"
)

// Estrutura para a tabela `accounts`
type Account struct {
	ID                uuid.UUID `db:"id"`
	OwnerID           uuid.UUID `db:"owner_id"`
	BlockchainAddress string    `db:"blockchain_address"`
	Status            string    `db:"status"`
	Balance           int64     `db:"balance"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

// Estrutura para a tabela `transactions`
type Transaction struct {
	ID            uuid.UUID `db:"id"`
	RequestID     string    `db:"request_id"`
	FromAccountID uuid.UUID `db:"from_account_id"` // Corrigido
	ToAccountID   uuid.UUID `db:"to_account_id"`   // Corrigido
	Amount        int64     `db:"amount"`
	Memo          string    `db:"memo"`
	Status        string    `db:"status"`
	BlockchainTx  string    `db:"blockchain_tx"`
	CreatedAt     time.Time `db:"created_at"`
}
