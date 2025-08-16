// package models

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// // Estrutura para a tabela `accounts`
// type Account struct {
// 	ID                uuid.UUID `db:"id"`
// 	OwnerID           uuid.UUID `db:"owner_id"`
// 	BlockchainAddress string    `db:"blockchain_address"`
// 	Status            string    `db:"status"`
// 	Balance           int64     `db:"balance"`
// 	CreatedAt         time.Time `db:"created_at"`
// 	UpdatedAt         time.Time `db:"updated_at"`
// }

// // Estrutura para a tabela `transactions`
// type Transaction struct {
// 	ID            uuid.UUID `db:"id"`
// 	RequestID     string    `db:"request_id"`
// 	FromAccountID uuid.UUID `db:"from_account_id"` // Corrigido
// 	ToAccountID   uuid.UUID `db:"to_account_id"`   // Corrigido
// 	Amount        int64     `db:"amount"`
// 	Memo          string    `db:"memo"`
// 	Status        string    `db:"status"`
// 	BlockchainTx  string    `db:"blockchain_tx"`
// 	CreatedAt     time.Time `db:"created_at"`
// }

// Arquivo: internal/models/models.go (ou similar)

package models

import (
	"time"

	"github.com/google/uuid"
)

// Estrutura para a tabela `accounts`
type Account struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" db:"id"`
	OwnerID           uuid.UUID `gorm:"type:uuid;not null" db:"owner_id"`
	BlockchainAddress string    `gorm:"unique;not null" db:"blockchain_address"`
	Status            string    `gorm:"not null" db:"status"`
	// ALTERADO: De int64 para string, com tipo de coluna NUMERIC para suportar valores da blockchain.
	Balance   string    `gorm:"type:numeric;not null;default:0" db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Estrutura para a tabela `transactions`
type Transaction struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" db:"id"`
	RequestID     string    `gorm:"unique;not null" db:"request_id"`
	FromAccountID uuid.UUID `gorm:"type:uuid;not null" db:"from_account_id"`
	ToAccountID   uuid.UUID `gorm:"type:uuid;not null" db:"to_account_id"`
	// ALTERADO: De int64 para string, para consistência com o saldo e valores da blockchain.
	Amount       string    `gorm:"type:numeric;not null" db:"amount"`
	Memo         string    `db:"memo"`
	Status       string    `gorm:"not null" db:"status"`
	BlockchainTx string    `db:"blockchain_tx"` // Hash da transação, pode ser nulo inicialmente
	CreatedAt    time.Time `db:"created_at"`
}
