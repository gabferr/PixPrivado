package blockchain // Corrigido o nome do pacote

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"pix-privado/internal/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client for interacting with the blockchain.
type Client struct {
	ethClient   *ethclient.Client
	contract    *contracts.Blockchain
	privateKey  *ecdsa.PrivateKey
	publicKey   *ecdsa.PublicKey
	fromAddress common.Address
	chainID     *big.Int
}

// NewClient creates a new blockchain client.
// This function is now corrected to accept all necessary parameters.
func NewClient(blockchainURL string, contractAddress string, privateKeyStr string, chainID *big.Int) (*Client, error) {
	// Conecta ao nó Ethereum via RPC
	ethClient, err := ethclient.Dial(blockchainURL)
	if err != nil {
		return nil, err
	}

	// Carrega o contrato
	contractAddressHex := common.HexToAddress(contractAddress)
	// Corrigido: Usando NewPixPrivado, o nome do seu contrato
	contract, err := contracts.NewBlockchain(contractAddressHex, ethClient)
	if err != nil {
		return nil, err
	}

	// Carrega a chave privada
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	// Obtém o endereço público da chave privada
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Client{
		ethClient:   ethClient,
		contract:    contract,
		privateKey:  privateKey,
		publicKey:   publicKeyECDSA,
		fromAddress: fromAddress,
		chainID:     chainID,
	}, nil
}

// Transfer sends a transaction to the smart contract.
func (c *Client) Transfer(ctx context.Context, toAddress common.Address, amount *big.Int, memo string) (common.Hash, error) {
	// A lógica para enviar a transação foi ajustada para usar o chainID
	// Isso evita o erro "Invalid chain id"
	nonce, err := c.ethClient.PendingNonceAt(ctx, c.fromAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// Corrigido: Usando a ChainID salva no Client
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return common.Hash{}, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // 0 ETH
	auth.GasLimit = uint64(300000) // Limite de gás
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
	callOpts := &bind.CallOpts{Context: ctx}
	balance, err := c.contract.BalanceOf(callOpts, account)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
