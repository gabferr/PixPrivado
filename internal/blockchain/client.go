// package blockchain // Corrigido o nome do pacote

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"log"
// 	"math/big"
// 	"pix-privado/internal/contracts"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// // Client for interacting with the blockchain.
// type Client struct {
// 	ethClient   *ethclient.Client
// 	contract    *contracts.Blockchain
// 	privateKey  *ecdsa.PrivateKey
// 	publicKey   *ecdsa.PublicKey
// 	fromAddress common.Address
// 	chainID     *big.Int
// }

// // NewClient creates a new blockchain client.
// // This function is now corrected to accept all necessary parameters.
// func NewClient(blockchainURL string, contractAddress string, privateKeyStr string, chainID *big.Int) (*Client, error) {
// 	// Conecta ao nó Ethereum via RPC
// 	ethClient, err := ethclient.Dial(blockchainURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Carrega o contrato
// 	contractAddressHex := common.HexToAddress(contractAddress)
// 	// Corrigido: Usando NewPixPrivado, o nome do seu contrato
// 	contract, err := contracts.NewBlockchain(contractAddressHex, ethClient)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Carrega a chave privada
// 	privateKey, err := crypto.HexToECDSA(privateKeyStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Obtém o endereço público da chave privada
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	return &Client{
// 		ethClient:   ethClient,
// 		contract:    contract,
// 		privateKey:  privateKey,
// 		publicKey:   publicKeyECDSA,
// 		fromAddress: fromAddress,
// 		chainID:     chainID,
// 	}, nil
// }

// // Transfer sends a transaction to the smart contract.
// func (c *Client) Transfer(ctx context.Context, toAddress common.Address, amount *big.Int, memo string) (common.Hash, error) {
// 	// A lógica para enviar a transação foi ajustada para usar o chainID
// 	// Isso evita o erro "Invalid chain id"
// 	nonce, err := c.ethClient.PendingNonceAt(ctx, c.fromAddress)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	// Corrigido: Usando a ChainID salva no Client
// 	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // 0 ETH
// 	auth.GasLimit = uint64(300000) // Limite de gás
// 	auth.GasPrice, err = c.ethClient.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	tx, err := c.contract.Transfer(auth, toAddress, amount, memo)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	return tx.Hash(), nil
// }

// // GetBalance retrieves the balance of an account from the smart contract.
// func (c *Client) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
// 	callOpts := &bind.CallOpts{Context: ctx}
// 	balance, err := c.contract.BalanceOf(callOpts, account)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return balance, nil
// }

package blockchain

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"pix-privado/internal/contracts"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client for interacting with the blockchain.
type Client struct {
	ethClient       *ethclient.Client
	contract        *contracts.Blockchain
	privateKey      *ecdsa.PrivateKey
	publicKey       *ecdsa.PublicKey
	fromAddress     common.Address
	chainID         *big.Int
	contractAddress common.Address // <-- ALTERAÇÃO 1: Adicionado campo para guardar o endereço do contrato
}

// NewClient creates a new blockchain client.
func NewClient(blockchainURL string, contractAddress string, privateKeyStr string, chainID *big.Int) (*Client, error) {
	ethClient, err := ethclient.Dial(blockchainURL)
	if err != nil {
		return nil, err
	}

	contractAddressHex := common.HexToAddress(contractAddress)
	contract, err := contracts.NewBlockchain(contractAddressHex, ethClient)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Client{
		ethClient:       ethClient,
		contract:        contract,
		privateKey:      privateKey,
		publicKey:       publicKeyECDSA,
		fromAddress:     fromAddress,
		chainID:         chainID,
		contractAddress: contractAddressHex, // <-- ALTERAÇÃO 2: Salvando o endereço no nosso client
	}, nil
}

// Transfer sends a transaction to the smart contract.
func (c *Client) Transfer(ctx context.Context, toAddress common.Address, amount *big.Int, memo string) (common.Hash, error) {
	// ... (função sem alterações)
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.fromAddress)
	if err != nil {
		return common.Hash{}, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return common.Hash{}, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice, err = c.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := c.contract.Transfer(auth, toAddress, amount, memo)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// GetBalance retrieves the balance of an account from the smart contract.
func (c *Client) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	// ... (função sem alterações)
	callOpts := &bind.CallOpts{Context: ctx}
	balance, err := c.contract.BalanceOf(callOpts, account)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// Em blockchain/client.go

// Substitua sua função SubscribeToEvents inteira por esta versão final:
func (c *Client) SubscribeToEvents(ctx context.Context, dbUpdater func(ctx context.Context, address common.Address, amount *big.Int) error) {
	contractAddress := c.contractAddress

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := c.ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Erro ao se inscrever nos eventos do contrato: %v", err)
	}

	log.Println("Sincronizador de eventos iniciado. Escutando por eventos 'Credit' e 'Transfer'...")

	contractAbi, err := abi.JSON(strings.NewReader(contracts.BlockchainABI))
	if err != nil {
		log.Fatalf("Erro ao fazer o parse do ABI do contrato: %v", err)
	}

	creditEventTopic := contractAbi.Events["Credit"].ID
	transferEventTopic := contractAbi.Events["Transfer"].ID

	for {
		select {
		case <-ctx.Done():
			log.Println("Contexto cancelado. Desligando o sincronizador de eventos.")
			return
		case err := <-sub.Err():
			log.Printf("Erro na subscrição de eventos: %v.", err)
		case vLog := <-logs:
			switch vLog.Topics[0] {
			case creditEventTopic:
				// <<< CORREÇÃO FINAL APLICADA AQUI >>>
				var creditData struct {
					Amount *big.Int
					Reason string // <-- CORRIGIDO de "Memo" para "Reason"
				}

				if err := contractAbi.UnpackIntoInterface(&creditData, "Credit", vLog.Data); err != nil {
					log.Printf("Erro ao decodificar dados do evento Credit: %v", err)
					continue
				}
				toAddress := common.BytesToAddress(vLog.Topics[1].Bytes())

				log.Printf("Evento 'Credit' recebido -> Para: %s, Valor: %s, Razão: %s", toAddress.Hex(), creditData.Amount.String(), creditData.Reason)

				if err := dbUpdater(ctx, toAddress, creditData.Amount); err != nil {
					log.Printf("ERRO no DB (Credit): %v", err)
				}

			case transferEventTopic:
				// Esta parte já estava correta, pois o evento Transfer usa "memo"
				var transferData struct {
					Amount *big.Int
					Memo   string
				}
				if err := contractAbi.UnpackIntoInterface(&transferData, "Transfer", vLog.Data); err != nil {
					log.Printf("Erro ao decodificar dados do evento Transfer: %v", err)
					continue
				}
				fromAddress := common.BytesToAddress(vLog.Topics[1].Bytes())
				toAddress := common.BytesToAddress(vLog.Topics[2].Bytes())

				log.Printf("Evento 'Transfer' recebido -> De: %s, Para: %s, Valor: %s, Memo: %s", fromAddress.Hex(), toAddress.Hex(), transferData.Amount.String(), transferData.Memo)

				amountNegative := new(big.Int).Neg(transferData.Amount)
				if err := dbUpdater(ctx, fromAddress, amountNegative); err != nil {
					log.Printf("ERRO no DB (Transfer/Débito): %v", err)
				}
				if err := dbUpdater(ctx, toAddress, transferData.Amount); err != nil {
					log.Printf("ERRO no DB (Transfer/Crédito): %v", err)
				}
			}
		}
	}
}
