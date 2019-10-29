// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IGameJam

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

// IGameJamABI is the input ABI used to generate the binding from.
const IGameJamABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initializeGameJam\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGameJamAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllCompetitorAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"competitor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"addCompetitor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"competitor\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payoutWinner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGameJam is an auto generated Go binding around an Ethereum contract.
type IGameJam struct {
	IGameJamCaller     // Read-only binding to the contract
	IGameJamTransactor // Write-only binding to the contract
	IGameJamFilterer   // Log filterer for contract events
}

// IGameJamCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGameJamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGameJamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGameJamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGameJamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGameJamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGameJamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGameJamSession struct {
	Contract     *IGameJam         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGameJamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGameJamCallerSession struct {
	Contract *IGameJamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IGameJamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGameJamTransactorSession struct {
	Contract     *IGameJamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGameJamRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGameJamRaw struct {
	Contract *IGameJam // Generic contract binding to access the raw methods on
}

// IGameJamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGameJamCallerRaw struct {
	Contract *IGameJamCaller // Generic read-only contract binding to access the raw methods on
}

// IGameJamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGameJamTransactorRaw struct {
	Contract *IGameJamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGameJam creates a new instance of IGameJam, bound to a specific deployed contract.
func NewIGameJam(address common.Address, backend bind.ContractBackend) (*IGameJam, error) {
	contract, err := bindIGameJam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGameJam{IGameJamCaller: IGameJamCaller{contract: contract}, IGameJamTransactor: IGameJamTransactor{contract: contract}, IGameJamFilterer: IGameJamFilterer{contract: contract}}, nil
}

// NewIGameJamCaller creates a new read-only instance of IGameJam, bound to a specific deployed contract.
func NewIGameJamCaller(address common.Address, caller bind.ContractCaller) (*IGameJamCaller, error) {
	contract, err := bindIGameJam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGameJamCaller{contract: contract}, nil
}

// NewIGameJamTransactor creates a new write-only instance of IGameJam, bound to a specific deployed contract.
func NewIGameJamTransactor(address common.Address, transactor bind.ContractTransactor) (*IGameJamTransactor, error) {
	contract, err := bindIGameJam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGameJamTransactor{contract: contract}, nil
}

// NewIGameJamFilterer creates a new log filterer instance of IGameJam, bound to a specific deployed contract.
func NewIGameJamFilterer(address common.Address, filterer bind.ContractFilterer) (*IGameJamFilterer, error) {
	contract, err := bindIGameJam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGameJamFilterer{contract: contract}, nil
}

// bindIGameJam binds a generic wrapper to an already deployed contract.
func bindIGameJam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGameJamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGameJam *IGameJamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGameJam.Contract.IGameJamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGameJam *IGameJamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGameJam.Contract.IGameJamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGameJam *IGameJamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGameJam.Contract.IGameJamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGameJam *IGameJamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGameJam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGameJam *IGameJamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGameJam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGameJam *IGameJamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGameJam.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_IGameJam *IGameJamCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IGameJam.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_IGameJam *IGameJamSession) Balance() (*big.Int, error) {
	return _IGameJam.Contract.Balance(&_IGameJam.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_IGameJam *IGameJamCallerSession) Balance() (*big.Int, error) {
	return _IGameJam.Contract.Balance(&_IGameJam.CallOpts)
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_IGameJam *IGameJamCaller) GetAllCompetitorAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IGameJam.contract.Call(opts, out, "getAllCompetitorAddresses")
	return *ret0, err
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_IGameJam *IGameJamSession) GetAllCompetitorAddresses() ([]common.Address, error) {
	return _IGameJam.Contract.GetAllCompetitorAddresses(&_IGameJam.CallOpts)
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_IGameJam *IGameJamCallerSession) GetAllCompetitorAddresses() ([]common.Address, error) {
	return _IGameJam.Contract.GetAllCompetitorAddresses(&_IGameJam.CallOpts)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_IGameJam *IGameJamCaller) IsGameJamAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IGameJam.contract.Call(opts, out, "isGameJamAdmin", account)
	return *ret0, err
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_IGameJam *IGameJamSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _IGameJam.Contract.IsGameJamAdmin(&_IGameJam.CallOpts, account)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_IGameJam *IGameJamCallerSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _IGameJam.Contract.IsGameJamAdmin(&_IGameJam.CallOpts, account)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_IGameJam *IGameJamTransactor) AddCompetitor(opts *bind.TransactOpts, competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "addCompetitor", competitor, ipfsHash)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_IGameJam *IGameJamSession) AddCompetitor(competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _IGameJam.Contract.AddCompetitor(&_IGameJam.TransactOpts, competitor, ipfsHash)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_IGameJam *IGameJamTransactorSession) AddCompetitor(competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _IGameJam.Contract.AddCompetitor(&_IGameJam.TransactOpts, competitor, ipfsHash)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_IGameJam *IGameJamTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_IGameJam *IGameJamSession) Finish() (*types.Transaction, error) {
	return _IGameJam.Contract.Finish(&_IGameJam.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_IGameJam *IGameJamTransactorSession) Finish() (*types.Transaction, error) {
	return _IGameJam.Contract.Finish(&_IGameJam.TransactOpts)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_IGameJam *IGameJamTransactor) InitializeGameJam(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "initializeGameJam", admin)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_IGameJam *IGameJamSession) InitializeGameJam(admin common.Address) (*types.Transaction, error) {
	return _IGameJam.Contract.InitializeGameJam(&_IGameJam.TransactOpts, admin)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_IGameJam *IGameJamTransactorSession) InitializeGameJam(admin common.Address) (*types.Transaction, error) {
	return _IGameJam.Contract.InitializeGameJam(&_IGameJam.TransactOpts, admin)
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_IGameJam *IGameJamTransactor) PayoutWinner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "payoutWinner")
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_IGameJam *IGameJamSession) PayoutWinner() (*types.Transaction, error) {
	return _IGameJam.Contract.PayoutWinner(&_IGameJam.TransactOpts)
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_IGameJam *IGameJamTransactorSession) PayoutWinner() (*types.Transaction, error) {
	return _IGameJam.Contract.PayoutWinner(&_IGameJam.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_IGameJam *IGameJamTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_IGameJam *IGameJamSession) Start() (*types.Transaction, error) {
	return _IGameJam.Contract.Start(&_IGameJam.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_IGameJam *IGameJamTransactorSession) Start() (*types.Transaction, error) {
	return _IGameJam.Contract.Start(&_IGameJam.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_IGameJam *IGameJamTransactor) Vote(opts *bind.TransactOpts, competitor common.Address) (*types.Transaction, error) {
	return _IGameJam.contract.Transact(opts, "vote", competitor)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_IGameJam *IGameJamSession) Vote(competitor common.Address) (*types.Transaction, error) {
	return _IGameJam.Contract.Vote(&_IGameJam.TransactOpts, competitor)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_IGameJam *IGameJamTransactorSession) Vote(competitor common.Address) (*types.Transaction, error) {
	return _IGameJam.Contract.Vote(&_IGameJam.TransactOpts, competitor)
}
