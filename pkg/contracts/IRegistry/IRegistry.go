// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IRegistry

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

// IRegistryABI is the input ABI used to generate the binding from.
const IRegistryABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IRegistry is an auto generated Go binding around an Ethereum contract.
type IRegistry struct {
	IRegistryCaller     // Read-only binding to the contract
	IRegistryTransactor // Write-only binding to the contract
	IRegistryFilterer   // Log filterer for contract events
}

// IRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRegistrySession struct {
	Contract     *IRegistry        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRegistryCallerSession struct {
	Contract *IRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRegistryTransactorSession struct {
	Contract     *IRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRegistryRaw struct {
	Contract *IRegistry // Generic contract binding to access the raw methods on
}

// IRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRegistryCallerRaw struct {
	Contract *IRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRegistryTransactorRaw struct {
	Contract *IRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistry creates a new instance of IRegistry, bound to a specific deployed contract.
func NewIRegistry(address common.Address, backend bind.ContractBackend) (*IRegistry, error) {
	contract, err := bindIRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistry{IRegistryCaller: IRegistryCaller{contract: contract}, IRegistryTransactor: IRegistryTransactor{contract: contract}, IRegistryFilterer: IRegistryFilterer{contract: contract}}, nil
}

// NewIRegistryCaller creates a new read-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryCaller(address common.Address, caller bind.ContractCaller) (*IRegistryCaller, error) {
	contract, err := bindIRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryCaller{contract: contract}, nil
}

// NewIRegistryTransactor creates a new write-only instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistryTransactor, error) {
	contract, err := bindIRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistryTransactor{contract: contract}, nil
}

// NewIRegistryFilterer creates a new log filterer instance of IRegistry, bound to a specific deployed contract.
func NewIRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistryFilterer, error) {
	contract, err := bindIRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistryFilterer{contract: contract}, nil
}

// bindIRegistry binds a generic wrapper to an already deployed contract.
func bindIRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.IRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.IRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistry *IRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistry *IRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistry *IRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x04433bbc.
//
// Solidity: function getContractAddress(string _name) constant returns(address)
func (_IRegistry *IRegistryCaller) GetContractAddress(opts *bind.CallOpts, _name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IRegistry.contract.Call(opts, out, "getContractAddress", _name)
	return *ret0, err
}

// GetContractAddress is a free data retrieval call binding the contract method 0x04433bbc.
//
// Solidity: function getContractAddress(string _name) constant returns(address)
func (_IRegistry *IRegistrySession) GetContractAddress(_name string) (common.Address, error) {
	return _IRegistry.Contract.GetContractAddress(&_IRegistry.CallOpts, _name)
}

// GetContractAddress is a free data retrieval call binding the contract method 0x04433bbc.
//
// Solidity: function getContractAddress(string _name) constant returns(address)
func (_IRegistry *IRegistryCallerSession) GetContractAddress(_name string) (common.Address, error) {
	return _IRegistry.Contract.GetContractAddress(&_IRegistry.CallOpts, _name)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x38aed2f7.
//
// Solidity: function updateContractAddress(string _name, address _address) returns(address)
func (_IRegistry *IRegistryTransactor) UpdateContractAddress(opts *bind.TransactOpts, _name string, _address common.Address) (*types.Transaction, error) {
	return _IRegistry.contract.Transact(opts, "updateContractAddress", _name, _address)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x38aed2f7.
//
// Solidity: function updateContractAddress(string _name, address _address) returns(address)
func (_IRegistry *IRegistrySession) UpdateContractAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _IRegistry.Contract.UpdateContractAddress(&_IRegistry.TransactOpts, _name, _address)
}

// UpdateContractAddress is a paid mutator transaction binding the contract method 0x38aed2f7.
//
// Solidity: function updateContractAddress(string _name, address _address) returns(address)
func (_IRegistry *IRegistryTransactorSession) UpdateContractAddress(_name string, _address common.Address) (*types.Transaction, error) {
	return _IRegistry.Contract.UpdateContractAddress(&_IRegistry.TransactOpts, _name, _address)
}
