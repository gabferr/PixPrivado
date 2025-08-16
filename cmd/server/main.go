// package main

// import (
// 	"context"
// 	"log"
// 	"math/big"
// 	"net/http"
// 	"os"

// 	"pix-privado/internal/blockchain"
// 	"pix-privado/internal/database"
// 	"pix-privado/internal/handlers"
// 	"pix-privado/internal/models"
// 	"strconv"

// 	"github.com/go-redsync/redsync/v4"
// 	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	"github.com/redis/go-redis/v9"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	// Carregar variáveis de ambiente
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}

// 	// Inicializar clientes
// 	ctx := context.Background()

// 	// --- LÓGICA PARA CRIAR AS TABELAS COM GORM ---
// 	// A DSN do banco de dados (DATABASE_URL)
// 	pgDSN := os.Getenv("DATABASE_URL")
// 	if pgDSN == "" {
// 		log.Fatal("DATABASE_URL not set")
// 	}

// 	// Conexão com GORM
// 	db, err := gorm.Open(postgres.Open(pgDSN), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database with GORM: %v", err)
// 	}

// 	// Realizar a migração automática para criar as tabelas
// 	// GORM usa as structs do models para definir o esquema das tabelas
// 	log.Println("Migrando as tabelas do banco de dados...")
// 	err = db.AutoMigrate(&models.Account{}, &models.Transaction{})
// 	if err != nil {
// 		log.Fatalf("Failed to migrate database tables: %v", err)
// 	}
// 	log.Println("Migração do banco de dados concluída com sucesso.")
// 	// --- FIM DA LÓGICA DE MIGRAÇÃO ---

// 	dbClient, err := database.NewClient(ctx, pgDSN)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	defer dbClient.Close()

// 	// Inicializar o cliente Redis para o redsync
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr: os.Getenv("REDIS_URL"),
// 	})
// 	pool := goredis.NewPool(redisClient)
// 	rs := redsync.New(pool)

// 	// --- ATUALIZAÇÃO: Obter o Chain ID e passar para o cliente da blockchain ---
// 	// O Chain ID do Ganache geralmente é 1337. O Hardhat usa 31337 por padrão.
// 	// Para evitar o conflito, vamos usar um Chain ID fixo para desenvolvimento.
// 	// Alternativamente, você pode ler de uma variável de ambiente.
// 	var chainID int64 = 1337
// 	if os.Getenv("CHAIN_ID") != "" {
// 		chainID, err = strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
// 		if err != nil {
// 			log.Fatalf("Failed to parse CHAIN_ID: %v", err)
// 		}
// 	}

// 	blockchainURL := os.Getenv("BLOCKCHAIN_URL")
// 	contractAddress := os.Getenv("CONTRACT_ADDRESS")
// 	privateKey := os.Getenv("OPERATOR_ROLE_PRIVATE_KEY")
// 	blockchainClient, err := blockchain.NewClient(blockchainURL, contractAddress, privateKey, big.NewInt(chainID))
// 	if err != nil {
// 		log.Fatalf("Failed to connect to blockchain: %v", err)
// 	}
// 	// --- FIM DA ATUALIZAÇÃO ---

// 	// Criar handlers
// 	transferHandler := &handlers.TransferHandler{
// 		DBClient:         dbClient,
// 		BlockchainClient: blockchainClient,
// 		RedsyncClient:    rs,
// 	}
// 	balanceHandler := &handlers.BalanceHandler{
// 		DBClient: dbClient,
// 	}

// 	// Configurar o router
// 	r := mux.NewRouter()
// 	r.HandleFunc("/transfers", transferHandler.CreateTransfer).Methods("POST")
// 	r.HandleFunc("/accounts/{id}/balance", balanceHandler.GetBalance).Methods("GET")

// 	log.Println("Servidor rodando na porta 8080...")
// 	http.ListenAndServe(":8080", r)
// }

package main

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"os"

	"pix-privado/internal/blockchain"
	"pix-privado/internal/database"
	"pix-privado/internal/handlers"
	"pix-privado/internal/models"
	"strconv"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Inicializar clientes
	ctx := context.Background()

	// --- LÓGICA PARA CRIAR AS TABELAS COM GORM ---
	pgDSN := os.Getenv("DATABASE_URL")
	if pgDSN == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(pgDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database with GORM: %v", err)
	}

	log.Println("Migrando as tabelas do banco de dados...")
	err = db.AutoMigrate(&models.Account{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate database tables: %v", err)
	}
	log.Println("Migração do banco de dados concluída com sucesso.")

	dbClient, err := database.NewClient(ctx, pgDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbClient.Close()

	// Inicializar o cliente Redis para o redsync
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	pool := goredis.NewPool(redisClient)
	rs := redsync.New(pool)

	// --- Configuração do cliente da blockchain ---
	var chainID int64 = 1337
	if os.Getenv("CHAIN_ID") != "" {
		chainID, err = strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse CHAIN_ID: %v", err)
		}
	}

	blockchainURL := os.Getenv("BLOCKCHAIN_URL")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	privateKey := os.Getenv("OPERATOR_ROLE_PRIVATE_KEY")
	blockchainClient, err := blockchain.NewClient(blockchainURL, contractAddress, privateKey, big.NewInt(chainID))
	if err != nil {
		log.Fatalf("Failed to connect to blockchain: %v", err)
	}

	// <<< INÍCIO DA ATUALIZAÇÃO >>>
	// =========================================================================
	// INICIAR O SINCRONIZADOR DE EVENTOS EM SEGUNDO PLANO (GOROUTINE)
	// =========================================================================
	go func() {
		log.Println("Iniciando o sincronizador de eventos da blockchain em segundo plano...")
		// Esta chamada irá bloquear esta goroutine e ficará escutando por eventos
		// enquanto o resto da aplicação (servidor web) continua normalmente.
		blockchainClient.SubscribeToEvents(ctx, dbClient.UpdateBalanceByAddress)
	}()
	// =========================================================================
	// <<< FIM DA ATUALIZAÇÃO >>>

	// Criar handlers
	transferHandler := &handlers.TransferHandler{
		DBClient:         dbClient,
		BlockchainClient: blockchainClient,
		RedsyncClient:    rs,
	}
	balanceHandler := &handlers.BalanceHandler{
		DBClient: dbClient,
	}

	// Configurar o router
	r := mux.NewRouter()
	r.HandleFunc("/transfers", transferHandler.CreateTransfer).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", balanceHandler.GetBalance).Methods("GET")

	// Iniciar o servidor HTTP
	// Esta é uma chamada bloqueante, que manterá a aplicação viva,
	// permitindo que nossa goroutine do sincronizador continue rodando.
	log.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	}
}
