// package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"math/big"
// 	"net/http"
// 	"time"

// 	"pix-privado/internal/blockchain"
// 	"pix-privado/internal/database"
// 	"pix-privado/internal/models"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/go-redsync/redsync/v4"
// 	"github.com/google/uuid"
// )

// type TransferRequest struct {
// 	RequestID   string `json:"request_id"`
// 	FromAccount string `json:"from_account_id"`
// 	ToAccount   string `json:"to_account_id"`
// 	Amount      int64  `json:"amount"`
// 	Memo        string `json:"memo"`
// }

// // TransferHandler holds the dependencies for handling transfer requests.
// type TransferHandler struct {
// 	DBClient         *database.Client
// 	BlockchainClient *blockchain.Client
// 	RedsyncClient    *redsync.Redsync
// }

// // CreateTransfer handles the POST /transfers request.
// func (h *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
// 	var req TransferRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Parse UUIDs from strings
// 	fromAccountUUID, err := uuid.Parse(req.FromAccount)
// 	if err != nil {
// 		http.Error(w, "Invalid 'from_account_id'", http.StatusBadRequest)
// 		return
// 	}
// 	toAccountUUID, err := uuid.Parse(req.ToAccount)
// 	if err != nil {
// 		http.Error(w, "Invalid 'to_account_id'", http.StatusBadRequest)
// 		return
// 	}

// 	// 1. Idempotência e Lock
// 	mutex := h.RedsyncClient.NewMutex("transfer-lock:" + req.FromAccount)
// 	if err := mutex.Lock(); err != nil {
// 		http.Error(w, "Failed to acquire lock", http.StatusInternalServerError)
// 		log.Println("Error acquiring Redis lock:", err)
// 		return
// 	}
// 	defer func() {
// 		if _, err := mutex.Unlock(); err != nil {
// 			log.Println("Error releasing Redis lock:", err)
// 		}
// 	}()

// 	exists, err := h.DBClient.CheckIfRequestExists(r.Context(), req.RequestID)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Println("Error checking for existing request:", err)
// 		return
// 	}
// 	if exists {
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"status":  "success",
// 			"message": "Request already processed",
// 		})
// 		return
// 	}

// 	// 2. Validar saldo no Postgres
// 	balance, err := h.DBClient.GetBalance(r.Context(), fromAccountUUID)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Println("Error getting account balance:", err)
// 		return
// 	}
// 	if balance < req.Amount {
// 		http.Error(w, "Insufficient balance", http.StatusPreconditionFailed)
// 		return
// 	}

// 	// 3. Salvar no banco (status pendente)
// 	transaction := models.Transaction{
// 		ID:            uuid.New(),
// 		RequestID:     req.RequestID,
// 		FromAccountID: fromAccountUUID,
// 		ToAccountID:   toAccountUUID,
// 		Amount:        req.Amount,
// 		Memo:          req.Memo,
// 		Status:        "pending",
// 		CreatedAt:     time.Now(),
// 	}

// 	if err := h.DBClient.CreateTransaction(r.Context(), &transaction); err != nil {
// 		http.Error(w, "Failed to create transaction record", http.StatusInternalServerError)
// 		log.Println("Error creating transaction in DB:", err)
// 		return
// 	}

// 	// 4. Enviar transação para a blockchain via RPC
// 	// Buscar o endereço de blockchain da conta de destino no banco de dados.
// 	// Nota: não é necessário buscar o endereço de origem, pois ele é derivado da privateKey.
// 	toAcc, err := h.DBClient.GetAccountByUUID(r.Context(), toAccountUUID)
// 	if err != nil {
// 		http.Error(w, "Internal server error: to account not found", http.StatusInternalServerError)
// 		return
// 	}

// 	// Converter o endereço string para o tipo common.Address
// 	toAddress := common.HexToAddress(toAcc.BlockchainAddress)
// 	// Converter a quantidade int64 para *big.Int
// 	amount := big.NewInt(transaction.Amount)

// 	// Chame a função Transfer com os argumentos corrigidos
// 	txHash, err := h.BlockchainClient.Transfer(
// 		r.Context(),
// 		toAddress,
// 		amount,
// 		transaction.Memo,
// 	)
// 	if err != nil {
// 		// Marcar a transação como falhada no DB e retornar erro
// 		h.DBClient.UpdateTransactionStatus(r.Context(), transaction.ID, "failed", "")
// 		http.Error(w, "Failed to send blockchain transaction", http.StatusInternalServerError)
// 		log.Println("Error sending blockchain transaction:", err)
// 		return
// 	}

// 	// A transação foi enviada. Aguardaremos a confirmação em uma goroutine
// 	txHashStr := txHash.String()
// 	go func(ctx context.Context, transactionID uuid.UUID, txHash string) {
// 		log.Println("Waiting for blockchain confirmation for tx:", txHash)
// 		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
// 		defer cancel()
// 		select {
// 		case <-ctx.Done():
// 			log.Println("Timeout waiting for transaction confirmation")
// 			h.DBClient.UpdateTransactionStatus(ctx, transactionID, "failed", txHash)
// 			return
// 		case <-time.After(30 * time.Second):
// 		}
// 		err := h.DBClient.UpdateTransactionStatus(ctx, transactionID, "confirmed", txHash)
// 		if err != nil {
// 			log.Println("Error updating transaction status to 'confirmed':", err)
// 			return
// 		}
// 		log.Printf("Transaction %s confirmed with hash %s\n", transactionID.String(), txHash)
// 	}(r.Context(), transaction.ID, txHashStr)

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"status":  "processing",
// 		"message": "Transaction sent to blockchain, awaiting confirmation",
// 		"tx_hash": txHashStr,
// 	})
// }

package handlers

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"pix-privado/internal/blockchain"
	"pix-privado/internal/database"
	"pix-privado/internal/models"
	"strconv" // <-- Adicionado para conversões
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redsync/redsync/v4"
	"github.com/google/uuid"
)

type TransferRequest struct {
	RequestID   string `json:"request_id"`
	FromAccount string `json:"from_account_id"`
	ToAccount   string `json:"to_account_id"`
	Amount      int64  `json:"amount"`
	Memo        string `json:"memo"`
}

// TransferHandler holds the dependencies for handling transfer requests.
type TransferHandler struct {
	DBClient         *database.Client
	BlockchainClient *blockchain.Client
	RedsyncClient    *redsync.Redsync
}

// CreateTransfer handles the POST /transfers request.
func (h *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromAccountUUID, err := uuid.Parse(req.FromAccount)
	if err != nil {
		http.Error(w, "Invalid 'from_account_id'", http.StatusBadRequest)
		return
	}
	toAccountUUID, err := uuid.Parse(req.ToAccount)
	if err != nil {
		http.Error(w, "Invalid 'to_account_id'", http.StatusBadRequest)
		return
	}

	// 1. Idempotência e Lock
	mutex := h.RedsyncClient.NewMutex("transfer-lock:" + req.FromAccount)
	if err := mutex.Lock(); err != nil {
		http.Error(w, "Failed to acquire lock", http.StatusInternalServerError)
		log.Println("Error acquiring Redis lock:", err)
		return
	}
	defer func() {
		if _, err := mutex.Unlock(); err != nil {
			log.Println("Error releasing Redis lock:", err)
		}
	}()

	exists, err := h.DBClient.CheckIfRequestExists(r.Context(), req.RequestID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error checking for existing request:", err)
		return
	}
	if exists {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Request already processed",
		})
		return
	}

	// 2. Validar saldo no Postgres
	// <<< CORREÇÃO 1: Validação de Saldo >>>
	// A função GetBalance agora retorna uma string. Precisamos comparar usando big.Int.
	balanceStr, err := h.DBClient.GetBalance(r.Context(), fromAccountUUID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error getting account balance:", err)
		return
	}

	balanceBigInt, ok := new(big.Int).SetString(balanceStr, 10)
	if !ok {
		http.Error(w, "Internal server error: invalid balance format in DB", http.StatusInternalServerError)
		return
	}

	reqAmountBigInt := big.NewInt(req.Amount)

	// Compara os dois valores: se balanceBigInt for menor que reqAmountBigInt, Cmp retorna -1.
	if balanceBigInt.Cmp(reqAmountBigInt) < 0 {
		http.Error(w, "Insufficient balance", http.StatusPreconditionFailed)
		return
	}

	// 3. Salvar no banco (status pendente)
	// <<< CORREÇÃO 2: Criação da Transação >>>
	// Convertemos o Amount (int64) do request para string antes de salvar no modelo.
	transaction := models.Transaction{
		ID:            uuid.New(),
		RequestID:     req.RequestID,
		FromAccountID: fromAccountUUID,
		ToAccountID:   toAccountUUID,
		Amount:        strconv.FormatInt(req.Amount, 10),
		Memo:          req.Memo,
		Status:        "pending",
		CreatedAt:     time.Now(),
	}

	if err := h.DBClient.CreateTransaction(r.Context(), &transaction); err != nil {
		http.Error(w, "Failed to create transaction record", http.StatusInternalServerError)
		log.Println("Error creating transaction in DB:", err)
		return
	}

	// 4. Enviar transação para a blockchain via RPC
	toAcc, err := h.DBClient.GetAccountByUUID(r.Context(), toAccountUUID)
	if err != nil {
		http.Error(w, "Internal server error: to account not found", http.StatusInternalServerError)
		return
	}

	toAddress := common.HexToAddress(toAcc.BlockchainAddress)

	// <<< CORREÇÃO 3: Envio para a Blockchain >>>
	// A `transaction.Amount` já é uma string, então a convertemos para *big.Int.
	amount, ok := new(big.Int).SetString(transaction.Amount, 10)
	if !ok {
		http.Error(w, "Internal server error: invalid transaction amount format", http.StatusInternalServerError)
		return
	}

	txHash, err := h.BlockchainClient.Transfer(r.Context(), toAddress, amount, transaction.Memo)
	if err != nil {
		h.DBClient.UpdateTransactionStatus(r.Context(), transaction.ID, "failed", "")
		http.Error(w, "Failed to send blockchain transaction", http.StatusInternalServerError)
		log.Println("Error sending blockchain transaction:", err)
		return
	}

	// O resto da sua lógica de confirmação assíncrona está boa.
	txHashStr := txHash.String()
	go func(ctx context.Context, transactionID uuid.UUID, txHash string) {
		log.Println("Waiting for blockchain confirmation for tx:", txHash)
		// ... (sua lógica de goroutine continua aqui)
	}(r.Context(), transaction.ID, txHashStr)

	w.WriteHeader(http.StatusAccepted) // Status 202 Accepted é mais apropriado para operações assíncronas.
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "processing",
		"message":        "Transaction sent to blockchain, awaiting confirmation",
		"transaction_id": transaction.ID.String(),
		"tx_hash":        txHashStr,
	})
}
