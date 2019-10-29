// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GameJamProxy

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GameJamProxyABI is the input ABI used to generate the binding from.
const GameJamProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gameJamContractAddress\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"competitors\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gameJamContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGameJamAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stage\",\"outputs\":[{\"internalType\":\"enumGameJamCommon.Stages\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GameJamProxy is an auto generated Go binding around an Ethereum contract.
type GameJamProxy struct {
	GameJamProxyCaller     // Read-only binding to the contract
	GameJamProxyTransactor // Write-only binding to the contract
	GameJamProxyFilterer   // Log filterer for contract events
}

// GameJamProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameJamProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameJamProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameJamProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameJamProxySession struct {
	Contract     *GameJamProxy     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameJamProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameJamProxyCallerSession struct {
	Contract *GameJamProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GameJamProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameJamProxyTransactorSession struct {
	Contract     *GameJamProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GameJamProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameJamProxyRaw struct {
	Contract *GameJamProxy // Generic contract binding to access the raw methods on
}

// GameJamProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameJamProxyCallerRaw struct {
	Contract *GameJamProxyCaller // Generic read-only contract binding to access the raw methods on
}

// GameJamProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameJamProxyTransactorRaw struct {
	Contract *GameJamProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameJamProxy creates a new instance of GameJamProxy, bound to a specific deployed contract.
func NewGameJamProxy(address common.Address, backend bind.ContractBackend) (*GameJamProxy, error) {
	contract, err := bindGameJamProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameJamProxy{GameJamProxyCaller: GameJamProxyCaller{contract: contract}, GameJamProxyTransactor: GameJamProxyTransactor{contract: contract}, GameJamProxyFilterer: GameJamProxyFilterer{contract: contract}}, nil
}

// NewGameJamProxyCaller creates a new read-only instance of GameJamProxy, bound to a specific deployed contract.
func NewGameJamProxyCaller(address common.Address, caller bind.ContractCaller) (*GameJamProxyCaller, error) {
	contract, err := bindGameJamProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamProxyCaller{contract: contract}, nil
}

// NewGameJamProxyTransactor creates a new write-only instance of GameJamProxy, bound to a specific deployed contract.
func NewGameJamProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*GameJamProxyTransactor, error) {
	contract, err := bindGameJamProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamProxyTransactor{contract: contract}, nil
}

// NewGameJamProxyFilterer creates a new log filterer instance of GameJamProxy, bound to a specific deployed contract.
func NewGameJamProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*GameJamProxyFilterer, error) {
	contract, err := bindGameJamProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameJamProxyFilterer{contract: contract}, nil
}

// bindGameJamProxy binds a generic wrapper to an already deployed contract.
func bindGameJamProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameJamProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamProxy *GameJamProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamProxy.Contract.GameJamProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamProxy *GameJamProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamProxy.Contract.GameJamProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamProxy *GameJamProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamProxy.Contract.GameJamProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamProxy *GameJamProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamProxy *GameJamProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamProxy *GameJamProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamProxy.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamProxy *GameJamProxyCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GameJamProxy.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamProxy *GameJamProxySession) Balance() (*big.Int, error) {
	return _GameJamProxy.Contract.Balance(&_GameJamProxy.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamProxy *GameJamProxyCallerSession) Balance() (*big.Int, error) {
	return _GameJamProxy.Contract.Balance(&_GameJamProxy.CallOpts)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamProxy *GameJamProxyCaller) Competitors(opts *bind.CallOpts, arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	ret := new(struct {
		IpfsHash string
		Votes    *big.Int
	})
	out := ret
	err := _GameJamProxy.contract.Call(opts, out, "competitors", arg0)
	return *ret, err
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamProxy *GameJamProxySession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJamProxy.Contract.Competitors(&_GameJamProxy.CallOpts, arg0)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamProxy *GameJamProxyCallerSession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJamProxy.Contract.Competitors(&_GameJamProxy.CallOpts, arg0)
}

// GameJamContractAddress is a free data retrieval call binding the contract method 0x88215064.
//
// Solidity: function gameJamContractAddress() constant returns(address)
func (_GameJamProxy *GameJamProxyCaller) GameJamContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GameJamProxy.contract.Call(opts, out, "gameJamContractAddress")
	return *ret0, err
}

// GameJamContractAddress is a free data retrieval call binding the contract method 0x88215064.
//
// Solidity: function gameJamContractAddress() constant returns(address)
func (_GameJamProxy *GameJamProxySession) GameJamContractAddress() (common.Address, error) {
	return _GameJamProxy.Contract.GameJamContractAddress(&_GameJamProxy.CallOpts)
}

// GameJamContractAddress is a free data retrieval call binding the contract method 0x88215064.
//
// Solidity: function gameJamContractAddress() constant returns(address)
func (_GameJamProxy *GameJamProxyCallerSession) GameJamContractAddress() (common.Address, error) {
	return _GameJamProxy.Contract.GameJamContractAddress(&_GameJamProxy.CallOpts)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamProxy *GameJamProxyCaller) IsGameJamAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameJamProxy.contract.Call(opts, out, "isGameJamAdmin", account)
	return *ret0, err
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamProxy *GameJamProxySession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJamProxy.Contract.IsGameJamAdmin(&_GameJamProxy.CallOpts, account)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamProxy *GameJamProxyCallerSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJamProxy.Contract.IsGameJamAdmin(&_GameJamProxy.CallOpts, account)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamProxy *GameJamProxyCaller) Stage(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GameJamProxy.contract.Call(opts, out, "stage")
	return *ret0, err
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamProxy *GameJamProxySession) Stage() (uint8, error) {
	return _GameJamProxy.Contract.Stage(&_GameJamProxy.CallOpts)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamProxy *GameJamProxyCallerSession) Stage() (uint8, error) {
	return _GameJamProxy.Contract.Stage(&_GameJamProxy.CallOpts)
}
