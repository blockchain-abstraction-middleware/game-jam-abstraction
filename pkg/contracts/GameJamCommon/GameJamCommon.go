// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GameJamCommon

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

// GameJamCommonABI is the input ABI used to generate the binding from.
const GameJamCommonABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"competitors\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stage\",\"outputs\":[{\"internalType\":\"enumGameJamCommon.Stages\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGameJamAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GameJamCommon is an auto generated Go binding around an Ethereum contract.
type GameJamCommon struct {
	GameJamCommonCaller     // Read-only binding to the contract
	GameJamCommonTransactor // Write-only binding to the contract
	GameJamCommonFilterer   // Log filterer for contract events
}

// GameJamCommonCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameJamCommonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamCommonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameJamCommonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamCommonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameJamCommonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamCommonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameJamCommonSession struct {
	Contract     *GameJamCommon    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameJamCommonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameJamCommonCallerSession struct {
	Contract *GameJamCommonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// GameJamCommonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameJamCommonTransactorSession struct {
	Contract     *GameJamCommonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GameJamCommonRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameJamCommonRaw struct {
	Contract *GameJamCommon // Generic contract binding to access the raw methods on
}

// GameJamCommonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameJamCommonCallerRaw struct {
	Contract *GameJamCommonCaller // Generic read-only contract binding to access the raw methods on
}

// GameJamCommonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameJamCommonTransactorRaw struct {
	Contract *GameJamCommonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameJamCommon creates a new instance of GameJamCommon, bound to a specific deployed contract.
func NewGameJamCommon(address common.Address, backend bind.ContractBackend) (*GameJamCommon, error) {
	contract, err := bindGameJamCommon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameJamCommon{GameJamCommonCaller: GameJamCommonCaller{contract: contract}, GameJamCommonTransactor: GameJamCommonTransactor{contract: contract}, GameJamCommonFilterer: GameJamCommonFilterer{contract: contract}}, nil
}

// NewGameJamCommonCaller creates a new read-only instance of GameJamCommon, bound to a specific deployed contract.
func NewGameJamCommonCaller(address common.Address, caller bind.ContractCaller) (*GameJamCommonCaller, error) {
	contract, err := bindGameJamCommon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamCommonCaller{contract: contract}, nil
}

// NewGameJamCommonTransactor creates a new write-only instance of GameJamCommon, bound to a specific deployed contract.
func NewGameJamCommonTransactor(address common.Address, transactor bind.ContractTransactor) (*GameJamCommonTransactor, error) {
	contract, err := bindGameJamCommon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamCommonTransactor{contract: contract}, nil
}

// NewGameJamCommonFilterer creates a new log filterer instance of GameJamCommon, bound to a specific deployed contract.
func NewGameJamCommonFilterer(address common.Address, filterer bind.ContractFilterer) (*GameJamCommonFilterer, error) {
	contract, err := bindGameJamCommon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameJamCommonFilterer{contract: contract}, nil
}

// bindGameJamCommon binds a generic wrapper to an already deployed contract.
func bindGameJamCommon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameJamCommonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamCommon *GameJamCommonRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamCommon.Contract.GameJamCommonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamCommon *GameJamCommonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamCommon.Contract.GameJamCommonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamCommon *GameJamCommonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamCommon.Contract.GameJamCommonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamCommon *GameJamCommonCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamCommon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamCommon *GameJamCommonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamCommon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamCommon *GameJamCommonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamCommon.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamCommon *GameJamCommonCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GameJamCommon.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamCommon *GameJamCommonSession) Balance() (*big.Int, error) {
	return _GameJamCommon.Contract.Balance(&_GameJamCommon.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJamCommon *GameJamCommonCallerSession) Balance() (*big.Int, error) {
	return _GameJamCommon.Contract.Balance(&_GameJamCommon.CallOpts)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamCommon *GameJamCommonCaller) Competitors(opts *bind.CallOpts, arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	ret := new(struct {
		IpfsHash string
		Votes    *big.Int
	})
	out := ret
	err := _GameJamCommon.contract.Call(opts, out, "competitors", arg0)
	return *ret, err
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamCommon *GameJamCommonSession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJamCommon.Contract.Competitors(&_GameJamCommon.CallOpts, arg0)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJamCommon *GameJamCommonCallerSession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJamCommon.Contract.Competitors(&_GameJamCommon.CallOpts, arg0)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamCommon *GameJamCommonCaller) IsGameJamAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameJamCommon.contract.Call(opts, out, "isGameJamAdmin", account)
	return *ret0, err
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamCommon *GameJamCommonSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJamCommon.Contract.IsGameJamAdmin(&_GameJamCommon.CallOpts, account)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJamCommon *GameJamCommonCallerSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJamCommon.Contract.IsGameJamAdmin(&_GameJamCommon.CallOpts, account)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamCommon *GameJamCommonCaller) Stage(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GameJamCommon.contract.Call(opts, out, "stage")
	return *ret0, err
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamCommon *GameJamCommonSession) Stage() (uint8, error) {
	return _GameJamCommon.Contract.Stage(&_GameJamCommon.CallOpts)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJamCommon *GameJamCommonCallerSession) Stage() (uint8, error) {
	return _GameJamCommon.Contract.Stage(&_GameJamCommon.CallOpts)
}
