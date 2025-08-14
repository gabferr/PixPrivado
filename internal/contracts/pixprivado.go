// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"credit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"debit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Blockchain.Contract.DEFAULTADMINROLE(&_Blockchain.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Blockchain.Contract.DEFAULTADMINROLE(&_Blockchain.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainSession) OPERATORROLE() ([32]byte, error) {
	return _Blockchain.Contract.OPERATORROLE(&_Blockchain.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Blockchain *BlockchainCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Blockchain.Contract.OPERATORROLE(&_Blockchain.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_Blockchain *BlockchainCaller) BalanceOf(opts *bind.CallOpts, acc common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "balanceOf", acc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_Blockchain *BlockchainSession) BalanceOf(acc common.Address) (*big.Int, error) {
	return _Blockchain.Contract.BalanceOf(&_Blockchain.CallOpts, acc)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) BalanceOf(acc common.Address) (*big.Int, error) {
	return _Blockchain.Contract.BalanceOf(&_Blockchain.CallOpts, acc)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockchain *BlockchainCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockchain *BlockchainSession) Decimals() (uint8, error) {
	return _Blockchain.Contract.Decimals(&_Blockchain.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Blockchain *BlockchainCallerSession) Decimals() (uint8, error) {
	return _Blockchain.Contract.Decimals(&_Blockchain.CallOpts)
}

// Frozen is a free data retrieval call binding the contract method 0xd0516650.
//
// Solidity: function frozen(address ) view returns(bool)
func (_Blockchain *BlockchainCaller) Frozen(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "frozen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Frozen is a free data retrieval call binding the contract method 0xd0516650.
//
// Solidity: function frozen(address ) view returns(bool)
func (_Blockchain *BlockchainSession) Frozen(arg0 common.Address) (bool, error) {
	return _Blockchain.Contract.Frozen(&_Blockchain.CallOpts, arg0)
}

// Frozen is a free data retrieval call binding the contract method 0xd0516650.
//
// Solidity: function frozen(address ) view returns(bool)
func (_Blockchain *BlockchainCallerSession) Frozen(arg0 common.Address) (bool, error) {
	return _Blockchain.Contract.Frozen(&_Blockchain.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blockchain *BlockchainCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blockchain *BlockchainSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Blockchain.Contract.GetRoleAdmin(&_Blockchain.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Blockchain *BlockchainCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Blockchain.Contract.GetRoleAdmin(&_Blockchain.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blockchain *BlockchainCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blockchain *BlockchainSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Blockchain.Contract.HasRole(&_Blockchain.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Blockchain *BlockchainCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Blockchain.Contract.HasRole(&_Blockchain.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blockchain *BlockchainCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blockchain *BlockchainSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Blockchain.Contract.SupportsInterface(&_Blockchain.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Blockchain *BlockchainCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Blockchain.Contract.SupportsInterface(&_Blockchain.CallOpts, interfaceId)
}

// Credit is a paid mutator transaction binding the contract method 0x25e017af.
//
// Solidity: function credit(address to, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainTransactor) Credit(opts *bind.TransactOpts, to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "credit", to, amount, reason)
}

// Credit is a paid mutator transaction binding the contract method 0x25e017af.
//
// Solidity: function credit(address to, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainSession) Credit(to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.Contract.Credit(&_Blockchain.TransactOpts, to, amount, reason)
}

// Credit is a paid mutator transaction binding the contract method 0x25e017af.
//
// Solidity: function credit(address to, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainTransactorSession) Credit(to common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.Contract.Credit(&_Blockchain.TransactOpts, to, amount, reason)
}

// Debit is a paid mutator transaction binding the contract method 0x5a275810.
//
// Solidity: function debit(address from, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainTransactor) Debit(opts *bind.TransactOpts, from common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "debit", from, amount, reason)
}

// Debit is a paid mutator transaction binding the contract method 0x5a275810.
//
// Solidity: function debit(address from, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainSession) Debit(from common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.Contract.Debit(&_Blockchain.TransactOpts, from, amount, reason)
}

// Debit is a paid mutator transaction binding the contract method 0x5a275810.
//
// Solidity: function debit(address from, uint256 amount, string reason) returns()
func (_Blockchain *BlockchainTransactorSession) Debit(from common.Address, amount *big.Int, reason string) (*types.Transaction, error) {
	return _Blockchain.Contract.Debit(&_Blockchain.TransactOpts, from, amount, reason)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.GrantRole(&_Blockchain.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.GrantRole(&_Blockchain.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blockchain *BlockchainTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blockchain *BlockchainSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceRole(&_Blockchain.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Blockchain *BlockchainTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceRole(&_Blockchain.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RevokeRole(&_Blockchain.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Blockchain *BlockchainTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.RevokeRole(&_Blockchain.TransactOpts, role, account)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(address acc, bool f) returns()
func (_Blockchain *BlockchainTransactor) SetFrozen(opts *bind.TransactOpts, acc common.Address, f bool) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "setFrozen", acc, f)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(address acc, bool f) returns()
func (_Blockchain *BlockchainSession) SetFrozen(acc common.Address, f bool) (*types.Transaction, error) {
	return _Blockchain.Contract.SetFrozen(&_Blockchain.TransactOpts, acc, f)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(address acc, bool f) returns()
func (_Blockchain *BlockchainTransactorSession) SetFrozen(acc common.Address, f bool) (*types.Transaction, error) {
	return _Blockchain.Contract.SetFrozen(&_Blockchain.TransactOpts, acc, f)
}

// Transfer is a paid mutator transaction binding the contract method 0x56b8c724.
//
// Solidity: function transfer(address to, uint256 amount, string memo) returns()
func (_Blockchain *BlockchainTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int, memo string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "transfer", to, amount, memo)
}

// Transfer is a paid mutator transaction binding the contract method 0x56b8c724.
//
// Solidity: function transfer(address to, uint256 amount, string memo) returns()
func (_Blockchain *BlockchainSession) Transfer(to common.Address, amount *big.Int, memo string) (*types.Transaction, error) {
	return _Blockchain.Contract.Transfer(&_Blockchain.TransactOpts, to, amount, memo)
}

// Transfer is a paid mutator transaction binding the contract method 0x56b8c724.
//
// Solidity: function transfer(address to, uint256 amount, string memo) returns()
func (_Blockchain *BlockchainTransactorSession) Transfer(to common.Address, amount *big.Int, memo string) (*types.Transaction, error) {
	return _Blockchain.Contract.Transfer(&_Blockchain.TransactOpts, to, amount, memo)
}

// BlockchainCreditIterator is returned from FilterCredit and is used to iterate over the raw logs and unpacked data for Credit events raised by the Blockchain contract.
type BlockchainCreditIterator struct {
	Event *BlockchainCredit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainCredit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainCredit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainCredit represents a Credit event raised by the Blockchain contract.
type BlockchainCredit struct {
	To     common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCredit is a free log retrieval operation binding the contract event 0x580bfe81f02c66d781c958d1b432d9657db8a7fa963ad93776841e15a89ea2ed.
//
// Solidity: event Credit(address indexed to, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) FilterCredit(opts *bind.FilterOpts, to []common.Address) (*BlockchainCreditIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Credit", toRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainCreditIterator{contract: _Blockchain.contract, event: "Credit", logs: logs, sub: sub}, nil
}

// WatchCredit is a free log subscription operation binding the contract event 0x580bfe81f02c66d781c958d1b432d9657db8a7fa963ad93776841e15a89ea2ed.
//
// Solidity: event Credit(address indexed to, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) WatchCredit(opts *bind.WatchOpts, sink chan<- *BlockchainCredit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Credit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainCredit)
				if err := _Blockchain.contract.UnpackLog(event, "Credit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCredit is a log parse operation binding the contract event 0x580bfe81f02c66d781c958d1b432d9657db8a7fa963ad93776841e15a89ea2ed.
//
// Solidity: event Credit(address indexed to, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) ParseCredit(log types.Log) (*BlockchainCredit, error) {
	event := new(BlockchainCredit)
	if err := _Blockchain.contract.UnpackLog(event, "Credit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainDebitIterator is returned from FilterDebit and is used to iterate over the raw logs and unpacked data for Debit events raised by the Blockchain contract.
type BlockchainDebitIterator struct {
	Event *BlockchainDebit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainDebitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainDebit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainDebit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainDebitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainDebitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainDebit represents a Debit event raised by the Blockchain contract.
type BlockchainDebit struct {
	From   common.Address
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDebit is a free log retrieval operation binding the contract event 0xf8bcbb21a2bb72d1a49565365420fd02eb47b91a28e3dca8d844aa623fc2d657.
//
// Solidity: event Debit(address indexed from, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) FilterDebit(opts *bind.FilterOpts, from []common.Address) (*BlockchainDebitIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Debit", fromRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainDebitIterator{contract: _Blockchain.contract, event: "Debit", logs: logs, sub: sub}, nil
}

// WatchDebit is a free log subscription operation binding the contract event 0xf8bcbb21a2bb72d1a49565365420fd02eb47b91a28e3dca8d844aa623fc2d657.
//
// Solidity: event Debit(address indexed from, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) WatchDebit(opts *bind.WatchOpts, sink chan<- *BlockchainDebit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Debit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainDebit)
				if err := _Blockchain.contract.UnpackLog(event, "Debit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDebit is a log parse operation binding the contract event 0xf8bcbb21a2bb72d1a49565365420fd02eb47b91a28e3dca8d844aa623fc2d657.
//
// Solidity: event Debit(address indexed from, uint256 amount, string reason)
func (_Blockchain *BlockchainFilterer) ParseDebit(log types.Log) (*BlockchainDebit, error) {
	event := new(BlockchainDebit)
	if err := _Blockchain.contract.UnpackLog(event, "Debit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainFrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the Blockchain contract.
type BlockchainFrozenIterator struct {
	Event *BlockchainFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainFrozen represents a Frozen event raised by the Blockchain contract.
type BlockchainFrozen struct {
	Account  common.Address
	IsFrozen bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x713eb400302cebac61f82eb8de5051d38458517ffac43ae45f4a9fd5d09ee698.
//
// Solidity: event Frozen(address indexed account, bool isFrozen)
func (_Blockchain *BlockchainFilterer) FilterFrozen(opts *bind.FilterOpts, account []common.Address) (*BlockchainFrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainFrozenIterator{contract: _Blockchain.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x713eb400302cebac61f82eb8de5051d38458517ffac43ae45f4a9fd5d09ee698.
//
// Solidity: event Frozen(address indexed account, bool isFrozen)
func (_Blockchain *BlockchainFilterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *BlockchainFrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainFrozen)
				if err := _Blockchain.contract.UnpackLog(event, "Frozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFrozen is a log parse operation binding the contract event 0x713eb400302cebac61f82eb8de5051d38458517ffac43ae45f4a9fd5d09ee698.
//
// Solidity: event Frozen(address indexed account, bool isFrozen)
func (_Blockchain *BlockchainFilterer) ParseFrozen(log types.Log) (*BlockchainFrozen, error) {
	event := new(BlockchainFrozen)
	if err := _Blockchain.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Blockchain contract.
type BlockchainRoleAdminChangedIterator struct {
	Event *BlockchainRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRoleAdminChanged represents a RoleAdminChanged event raised by the Blockchain contract.
type BlockchainRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Blockchain *BlockchainFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BlockchainRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRoleAdminChangedIterator{contract: _Blockchain.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Blockchain *BlockchainFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BlockchainRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRoleAdminChanged)
				if err := _Blockchain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Blockchain *BlockchainFilterer) ParseRoleAdminChanged(log types.Log) (*BlockchainRoleAdminChanged, error) {
	event := new(BlockchainRoleAdminChanged)
	if err := _Blockchain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Blockchain contract.
type BlockchainRoleGrantedIterator struct {
	Event *BlockchainRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRoleGranted represents a RoleGranted event raised by the Blockchain contract.
type BlockchainRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlockchainRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRoleGrantedIterator{contract: _Blockchain.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BlockchainRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRoleGranted)
				if err := _Blockchain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) ParseRoleGranted(log types.Log) (*BlockchainRoleGranted, error) {
	event := new(BlockchainRoleGranted)
	if err := _Blockchain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Blockchain contract.
type BlockchainRoleRevokedIterator struct {
	Event *BlockchainRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainRoleRevoked represents a RoleRevoked event raised by the Blockchain contract.
type BlockchainRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlockchainRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainRoleRevokedIterator{contract: _Blockchain.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BlockchainRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainRoleRevoked)
				if err := _Blockchain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Blockchain *BlockchainFilterer) ParseRoleRevoked(log types.Log) (*BlockchainRoleRevoked, error) {
	event := new(BlockchainRoleRevoked)
	if err := _Blockchain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Blockchain contract.
type BlockchainTransferIterator struct {
	Event *BlockchainTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockchainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockchainTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockchainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTransfer represents a Transfer event raised by the Blockchain contract.
type BlockchainTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Memo   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount, string memo)
func (_Blockchain *BlockchainFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BlockchainTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransferIterator{contract: _Blockchain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount, string memo)
func (_Blockchain *BlockchainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BlockchainTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTransfer)
				if err := _Blockchain.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xcd6e659e4c2e75c3bfe47fecaccf39aeb368116a0ee52afb532e07f6cba6c0d1.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount, string memo)
func (_Blockchain *BlockchainFilterer) ParseTransfer(log types.Log) (*BlockchainTransfer, error) {
	event := new(BlockchainTransfer)
	if err := _Blockchain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
