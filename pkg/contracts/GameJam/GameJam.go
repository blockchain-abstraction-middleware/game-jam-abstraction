// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GameJam

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

// GameJamABI is the input ABI used to generate the binding from.
const GameJamABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"competitor\",\"type\":\"address\"}],\"name\":\"CompetitorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GameJamAdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"GameJamFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"GameJameStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"competitorVotedFor\",\"type\":\"address\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"WinnerDeclared\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"competitors\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGameJamAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stage\",\"outputs\":[{\"internalType\":\"enumGameJamCommon.Stages\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initializeGameJam\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllCompetitorAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"competitor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"}],\"name\":\"addCompetitor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"competitor\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"payoutWinner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GameJam is an auto generated Go binding around an Ethereum contract.
type GameJam struct {
	GameJamCaller     // Read-only binding to the contract
	GameJamTransactor // Write-only binding to the contract
	GameJamFilterer   // Log filterer for contract events
}

// GameJamCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameJamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameJamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameJamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameJamSession struct {
	Contract     *GameJam          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameJamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameJamCallerSession struct {
	Contract *GameJamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GameJamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameJamTransactorSession struct {
	Contract     *GameJamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GameJamRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameJamRaw struct {
	Contract *GameJam // Generic contract binding to access the raw methods on
}

// GameJamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameJamCallerRaw struct {
	Contract *GameJamCaller // Generic read-only contract binding to access the raw methods on
}

// GameJamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameJamTransactorRaw struct {
	Contract *GameJamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameJam creates a new instance of GameJam, bound to a specific deployed contract.
func NewGameJam(address common.Address, backend bind.ContractBackend) (*GameJam, error) {
	contract, err := bindGameJam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameJam{GameJamCaller: GameJamCaller{contract: contract}, GameJamTransactor: GameJamTransactor{contract: contract}, GameJamFilterer: GameJamFilterer{contract: contract}}, nil
}

// NewGameJamCaller creates a new read-only instance of GameJam, bound to a specific deployed contract.
func NewGameJamCaller(address common.Address, caller bind.ContractCaller) (*GameJamCaller, error) {
	contract, err := bindGameJam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamCaller{contract: contract}, nil
}

// NewGameJamTransactor creates a new write-only instance of GameJam, bound to a specific deployed contract.
func NewGameJamTransactor(address common.Address, transactor bind.ContractTransactor) (*GameJamTransactor, error) {
	contract, err := bindGameJam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamTransactor{contract: contract}, nil
}

// NewGameJamFilterer creates a new log filterer instance of GameJam, bound to a specific deployed contract.
func NewGameJamFilterer(address common.Address, filterer bind.ContractFilterer) (*GameJamFilterer, error) {
	contract, err := bindGameJam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameJamFilterer{contract: contract}, nil
}

// bindGameJam binds a generic wrapper to an already deployed contract.
func bindGameJam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameJamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJam *GameJamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJam.Contract.GameJamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJam *GameJamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJam.Contract.GameJamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJam *GameJamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJam.Contract.GameJamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJam *GameJamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJam *GameJamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJam *GameJamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJam.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJam *GameJamCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GameJam.contract.Call(opts, out, "balance")
	return *ret0, err
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJam *GameJamSession) Balance() (*big.Int, error) {
	return _GameJam.Contract.Balance(&_GameJam.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() constant returns(uint256)
func (_GameJam *GameJamCallerSession) Balance() (*big.Int, error) {
	return _GameJam.Contract.Balance(&_GameJam.CallOpts)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJam *GameJamCaller) Competitors(opts *bind.CallOpts, arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	ret := new(struct {
		IpfsHash string
		Votes    *big.Int
	})
	out := ret
	err := _GameJam.contract.Call(opts, out, "competitors", arg0)
	return *ret, err
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJam *GameJamSession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJam.Contract.Competitors(&_GameJam.CallOpts, arg0)
}

// Competitors is a free data retrieval call binding the contract method 0xb7a076c6.
//
// Solidity: function competitors(address ) constant returns(string ipfsHash, uint256 votes)
func (_GameJam *GameJamCallerSession) Competitors(arg0 common.Address) (struct {
	IpfsHash string
	Votes    *big.Int
}, error) {
	return _GameJam.Contract.Competitors(&_GameJam.CallOpts, arg0)
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_GameJam *GameJamCaller) GetAllCompetitorAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GameJam.contract.Call(opts, out, "getAllCompetitorAddresses")
	return *ret0, err
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_GameJam *GameJamSession) GetAllCompetitorAddresses() ([]common.Address, error) {
	return _GameJam.Contract.GetAllCompetitorAddresses(&_GameJam.CallOpts)
}

// GetAllCompetitorAddresses is a free data retrieval call binding the contract method 0xba83e461.
//
// Solidity: function getAllCompetitorAddresses() constant returns(address[])
func (_GameJam *GameJamCallerSession) GetAllCompetitorAddresses() ([]common.Address, error) {
	return _GameJam.Contract.GetAllCompetitorAddresses(&_GameJam.CallOpts)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJam *GameJamCaller) IsGameJamAdmin(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameJam.contract.Call(opts, out, "isGameJamAdmin", account)
	return *ret0, err
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJam *GameJamSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJam.Contract.IsGameJamAdmin(&_GameJam.CallOpts, account)
}

// IsGameJamAdmin is a free data retrieval call binding the contract method 0x0f6feac2.
//
// Solidity: function isGameJamAdmin(address account) constant returns(bool)
func (_GameJam *GameJamCallerSession) IsGameJamAdmin(account common.Address) (bool, error) {
	return _GameJam.Contract.IsGameJamAdmin(&_GameJam.CallOpts, account)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJam *GameJamCaller) Stage(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _GameJam.contract.Call(opts, out, "stage")
	return *ret0, err
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJam *GameJamSession) Stage() (uint8, error) {
	return _GameJam.Contract.Stage(&_GameJam.CallOpts)
}

// Stage is a free data retrieval call binding the contract method 0xc040e6b8.
//
// Solidity: function stage() constant returns(uint8)
func (_GameJam *GameJamCallerSession) Stage() (uint8, error) {
	return _GameJam.Contract.Stage(&_GameJam.CallOpts)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_GameJam *GameJamTransactor) AddCompetitor(opts *bind.TransactOpts, competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "addCompetitor", competitor, ipfsHash)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_GameJam *GameJamSession) AddCompetitor(competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _GameJam.Contract.AddCompetitor(&_GameJam.TransactOpts, competitor, ipfsHash)
}

// AddCompetitor is a paid mutator transaction binding the contract method 0x62436371.
//
// Solidity: function addCompetitor(address competitor, string ipfsHash) returns()
func (_GameJam *GameJamTransactorSession) AddCompetitor(competitor common.Address, ipfsHash string) (*types.Transaction, error) {
	return _GameJam.Contract.AddCompetitor(&_GameJam.TransactOpts, competitor, ipfsHash)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_GameJam *GameJamTransactor) Finish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "finish")
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_GameJam *GameJamSession) Finish() (*types.Transaction, error) {
	return _GameJam.Contract.Finish(&_GameJam.TransactOpts)
}

// Finish is a paid mutator transaction binding the contract method 0xd56b2889.
//
// Solidity: function finish() returns()
func (_GameJam *GameJamTransactorSession) Finish() (*types.Transaction, error) {
	return _GameJam.Contract.Finish(&_GameJam.TransactOpts)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_GameJam *GameJamTransactor) InitializeGameJam(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "initializeGameJam", admin)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_GameJam *GameJamSession) InitializeGameJam(admin common.Address) (*types.Transaction, error) {
	return _GameJam.Contract.InitializeGameJam(&_GameJam.TransactOpts, admin)
}

// InitializeGameJam is a paid mutator transaction binding the contract method 0x34675aeb.
//
// Solidity: function initializeGameJam(address admin) returns()
func (_GameJam *GameJamTransactorSession) InitializeGameJam(admin common.Address) (*types.Transaction, error) {
	return _GameJam.Contract.InitializeGameJam(&_GameJam.TransactOpts, admin)
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_GameJam *GameJamTransactor) PayoutWinner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "payoutWinner")
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_GameJam *GameJamSession) PayoutWinner() (*types.Transaction, error) {
	return _GameJam.Contract.PayoutWinner(&_GameJam.TransactOpts)
}

// PayoutWinner is a paid mutator transaction binding the contract method 0x07552463.
//
// Solidity: function payoutWinner() returns()
func (_GameJam *GameJamTransactorSession) PayoutWinner() (*types.Transaction, error) {
	return _GameJam.Contract.PayoutWinner(&_GameJam.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_GameJam *GameJamTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_GameJam *GameJamSession) Start() (*types.Transaction, error) {
	return _GameJam.Contract.Start(&_GameJam.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_GameJam *GameJamTransactorSession) Start() (*types.Transaction, error) {
	return _GameJam.Contract.Start(&_GameJam.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_GameJam *GameJamTransactor) Vote(opts *bind.TransactOpts, competitor common.Address) (*types.Transaction, error) {
	return _GameJam.contract.Transact(opts, "vote", competitor)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_GameJam *GameJamSession) Vote(competitor common.Address) (*types.Transaction, error) {
	return _GameJam.Contract.Vote(&_GameJam.TransactOpts, competitor)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(address competitor) returns()
func (_GameJam *GameJamTransactorSession) Vote(competitor common.Address) (*types.Transaction, error) {
	return _GameJam.Contract.Vote(&_GameJam.TransactOpts, competitor)
}

// GameJamCompetitorAddedIterator is returned from FilterCompetitorAdded and is used to iterate over the raw logs and unpacked data for CompetitorAdded events raised by the GameJam contract.
type GameJamCompetitorAddedIterator struct {
	Event *GameJamCompetitorAdded // Event containing the contract specifics and raw log

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
func (it *GameJamCompetitorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamCompetitorAdded)
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
		it.Event = new(GameJamCompetitorAdded)
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
func (it *GameJamCompetitorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamCompetitorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamCompetitorAdded represents a CompetitorAdded event raised by the GameJam contract.
type GameJamCompetitorAdded struct {
	Competitor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCompetitorAdded is a free log retrieval operation binding the contract event 0x5c2db1fe8e7be28012487253743d81109ac5a3ffced6c0d2bd7db7bf65277bfb.
//
// Solidity: event CompetitorAdded(address competitor)
func (_GameJam *GameJamFilterer) FilterCompetitorAdded(opts *bind.FilterOpts) (*GameJamCompetitorAddedIterator, error) {

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "CompetitorAdded")
	if err != nil {
		return nil, err
	}
	return &GameJamCompetitorAddedIterator{contract: _GameJam.contract, event: "CompetitorAdded", logs: logs, sub: sub}, nil
}

// WatchCompetitorAdded is a free log subscription operation binding the contract event 0x5c2db1fe8e7be28012487253743d81109ac5a3ffced6c0d2bd7db7bf65277bfb.
//
// Solidity: event CompetitorAdded(address competitor)
func (_GameJam *GameJamFilterer) WatchCompetitorAdded(opts *bind.WatchOpts, sink chan<- *GameJamCompetitorAdded) (event.Subscription, error) {

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "CompetitorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamCompetitorAdded)
				if err := _GameJam.contract.UnpackLog(event, "CompetitorAdded", log); err != nil {
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

// ParseCompetitorAdded is a log parse operation binding the contract event 0x5c2db1fe8e7be28012487253743d81109ac5a3ffced6c0d2bd7db7bf65277bfb.
//
// Solidity: event CompetitorAdded(address competitor)
func (_GameJam *GameJamFilterer) ParseCompetitorAdded(log types.Log) (*GameJamCompetitorAdded, error) {
	event := new(GameJamCompetitorAdded)
	if err := _GameJam.contract.UnpackLog(event, "CompetitorAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamGameJamAdminAddedIterator is returned from FilterGameJamAdminAdded and is used to iterate over the raw logs and unpacked data for GameJamAdminAdded events raised by the GameJam contract.
type GameJamGameJamAdminAddedIterator struct {
	Event *GameJamGameJamAdminAdded // Event containing the contract specifics and raw log

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
func (it *GameJamGameJamAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamGameJamAdminAdded)
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
		it.Event = new(GameJamGameJamAdminAdded)
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
func (it *GameJamGameJamAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamGameJamAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamGameJamAdminAdded represents a GameJamAdminAdded event raised by the GameJam contract.
type GameJamGameJamAdminAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameJamAdminAdded is a free log retrieval operation binding the contract event 0x9d667829fd9a3d77e814b7334ea10a9442f6f29357d304db86d1fa5f4b305a24.
//
// Solidity: event GameJamAdminAdded(address indexed account)
func (_GameJam *GameJamFilterer) FilterGameJamAdminAdded(opts *bind.FilterOpts, account []common.Address) (*GameJamGameJamAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "GameJamAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &GameJamGameJamAdminAddedIterator{contract: _GameJam.contract, event: "GameJamAdminAdded", logs: logs, sub: sub}, nil
}

// WatchGameJamAdminAdded is a free log subscription operation binding the contract event 0x9d667829fd9a3d77e814b7334ea10a9442f6f29357d304db86d1fa5f4b305a24.
//
// Solidity: event GameJamAdminAdded(address indexed account)
func (_GameJam *GameJamFilterer) WatchGameJamAdminAdded(opts *bind.WatchOpts, sink chan<- *GameJamGameJamAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "GameJamAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamGameJamAdminAdded)
				if err := _GameJam.contract.UnpackLog(event, "GameJamAdminAdded", log); err != nil {
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

// ParseGameJamAdminAdded is a log parse operation binding the contract event 0x9d667829fd9a3d77e814b7334ea10a9442f6f29357d304db86d1fa5f4b305a24.
//
// Solidity: event GameJamAdminAdded(address indexed account)
func (_GameJam *GameJamFilterer) ParseGameJamAdminAdded(log types.Log) (*GameJamGameJamAdminAdded, error) {
	event := new(GameJamGameJamAdminAdded)
	if err := _GameJam.contract.UnpackLog(event, "GameJamAdminAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamGameJamFinishedIterator is returned from FilterGameJamFinished and is used to iterate over the raw logs and unpacked data for GameJamFinished events raised by the GameJam contract.
type GameJamGameJamFinishedIterator struct {
	Event *GameJamGameJamFinished // Event containing the contract specifics and raw log

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
func (it *GameJamGameJamFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamGameJamFinished)
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
		it.Event = new(GameJamGameJamFinished)
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
func (it *GameJamGameJamFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamGameJamFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamGameJamFinished represents a GameJamFinished event raised by the GameJam contract.
type GameJamGameJamFinished struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameJamFinished is a free log retrieval operation binding the contract event 0x5bcf61278175a9b1e3c7cca2a9117825ebe2477d611ecc0f35575bc6629579f1.
//
// Solidity: event GameJamFinished(address winner)
func (_GameJam *GameJamFilterer) FilterGameJamFinished(opts *bind.FilterOpts) (*GameJamGameJamFinishedIterator, error) {

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "GameJamFinished")
	if err != nil {
		return nil, err
	}
	return &GameJamGameJamFinishedIterator{contract: _GameJam.contract, event: "GameJamFinished", logs: logs, sub: sub}, nil
}

// WatchGameJamFinished is a free log subscription operation binding the contract event 0x5bcf61278175a9b1e3c7cca2a9117825ebe2477d611ecc0f35575bc6629579f1.
//
// Solidity: event GameJamFinished(address winner)
func (_GameJam *GameJamFilterer) WatchGameJamFinished(opts *bind.WatchOpts, sink chan<- *GameJamGameJamFinished) (event.Subscription, error) {

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "GameJamFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamGameJamFinished)
				if err := _GameJam.contract.UnpackLog(event, "GameJamFinished", log); err != nil {
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

// ParseGameJamFinished is a log parse operation binding the contract event 0x5bcf61278175a9b1e3c7cca2a9117825ebe2477d611ecc0f35575bc6629579f1.
//
// Solidity: event GameJamFinished(address winner)
func (_GameJam *GameJamFilterer) ParseGameJamFinished(log types.Log) (*GameJamGameJamFinished, error) {
	event := new(GameJamGameJamFinished)
	if err := _GameJam.contract.UnpackLog(event, "GameJamFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamGameJameStartedIterator is returned from FilterGameJameStarted and is used to iterate over the raw logs and unpacked data for GameJameStarted events raised by the GameJam contract.
type GameJamGameJameStartedIterator struct {
	Event *GameJamGameJameStarted // Event containing the contract specifics and raw log

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
func (it *GameJamGameJameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamGameJameStarted)
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
		it.Event = new(GameJamGameJameStarted)
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
func (it *GameJamGameJameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamGameJameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamGameJameStarted represents a GameJameStarted event raised by the GameJam contract.
type GameJamGameJameStarted struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGameJameStarted is a free log retrieval operation binding the contract event 0x8191bc0927e9ff0bd2b5c4d0975cbe0a05c45afbf2ec01a0d3f176a997596aeb.
//
// Solidity: event GameJameStarted(uint256 startTime)
func (_GameJam *GameJamFilterer) FilterGameJameStarted(opts *bind.FilterOpts) (*GameJamGameJameStartedIterator, error) {

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "GameJameStarted")
	if err != nil {
		return nil, err
	}
	return &GameJamGameJameStartedIterator{contract: _GameJam.contract, event: "GameJameStarted", logs: logs, sub: sub}, nil
}

// WatchGameJameStarted is a free log subscription operation binding the contract event 0x8191bc0927e9ff0bd2b5c4d0975cbe0a05c45afbf2ec01a0d3f176a997596aeb.
//
// Solidity: event GameJameStarted(uint256 startTime)
func (_GameJam *GameJamFilterer) WatchGameJameStarted(opts *bind.WatchOpts, sink chan<- *GameJamGameJameStarted) (event.Subscription, error) {

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "GameJameStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamGameJameStarted)
				if err := _GameJam.contract.UnpackLog(event, "GameJameStarted", log); err != nil {
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

// ParseGameJameStarted is a log parse operation binding the contract event 0x8191bc0927e9ff0bd2b5c4d0975cbe0a05c45afbf2ec01a0d3f176a997596aeb.
//
// Solidity: event GameJameStarted(uint256 startTime)
func (_GameJam *GameJamFilterer) ParseGameJameStarted(log types.Log) (*GameJamGameJameStarted, error) {
	event := new(GameJamGameJameStarted)
	if err := _GameJam.contract.UnpackLog(event, "GameJameStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the GameJam contract.
type GameJamVoteCastIterator struct {
	Event *GameJamVoteCast // Event containing the contract specifics and raw log

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
func (it *GameJamVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamVoteCast)
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
		it.Event = new(GameJamVoteCast)
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
func (it *GameJamVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamVoteCast represents a VoteCast event raised by the GameJam contract.
type GameJamVoteCast struct {
	CompetitorVotedFor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb87a61d78bd15221c59b983f6ef032fbde0631e66bfe18d1421fe5a73c16741d.
//
// Solidity: event VoteCast(address competitorVotedFor)
func (_GameJam *GameJamFilterer) FilterVoteCast(opts *bind.FilterOpts) (*GameJamVoteCastIterator, error) {

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &GameJamVoteCastIterator{contract: _GameJam.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb87a61d78bd15221c59b983f6ef032fbde0631e66bfe18d1421fe5a73c16741d.
//
// Solidity: event VoteCast(address competitorVotedFor)
func (_GameJam *GameJamFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *GameJamVoteCast) (event.Subscription, error) {

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamVoteCast)
				if err := _GameJam.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb87a61d78bd15221c59b983f6ef032fbde0631e66bfe18d1421fe5a73c16741d.
//
// Solidity: event VoteCast(address competitorVotedFor)
func (_GameJam *GameJamFilterer) ParseVoteCast(log types.Log) (*GameJamVoteCast, error) {
	event := new(GameJamVoteCast)
	if err := _GameJam.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamWinnerDeclaredIterator is returned from FilterWinnerDeclared and is used to iterate over the raw logs and unpacked data for WinnerDeclared events raised by the GameJam contract.
type GameJamWinnerDeclaredIterator struct {
	Event *GameJamWinnerDeclared // Event containing the contract specifics and raw log

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
func (it *GameJamWinnerDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamWinnerDeclared)
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
		it.Event = new(GameJamWinnerDeclared)
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
func (it *GameJamWinnerDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamWinnerDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamWinnerDeclared represents a WinnerDeclared event raised by the GameJam contract.
type GameJamWinnerDeclared struct {
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinnerDeclared is a free log retrieval operation binding the contract event 0x2d66076f24af48812dc2fdd69ac3d62bc732d6766ce65fdf93a45675faba22bd.
//
// Solidity: event WinnerDeclared(address winner)
func (_GameJam *GameJamFilterer) FilterWinnerDeclared(opts *bind.FilterOpts) (*GameJamWinnerDeclaredIterator, error) {

	logs, sub, err := _GameJam.contract.FilterLogs(opts, "WinnerDeclared")
	if err != nil {
		return nil, err
	}
	return &GameJamWinnerDeclaredIterator{contract: _GameJam.contract, event: "WinnerDeclared", logs: logs, sub: sub}, nil
}

// WatchWinnerDeclared is a free log subscription operation binding the contract event 0x2d66076f24af48812dc2fdd69ac3d62bc732d6766ce65fdf93a45675faba22bd.
//
// Solidity: event WinnerDeclared(address winner)
func (_GameJam *GameJamFilterer) WatchWinnerDeclared(opts *bind.WatchOpts, sink chan<- *GameJamWinnerDeclared) (event.Subscription, error) {

	logs, sub, err := _GameJam.contract.WatchLogs(opts, "WinnerDeclared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamWinnerDeclared)
				if err := _GameJam.contract.UnpackLog(event, "WinnerDeclared", log); err != nil {
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

// ParseWinnerDeclared is a log parse operation binding the contract event 0x2d66076f24af48812dc2fdd69ac3d62bc732d6766ce65fdf93a45675faba22bd.
//
// Solidity: event WinnerDeclared(address winner)
func (_GameJam *GameJamFilterer) ParseWinnerDeclared(log types.Log) (*GameJamWinnerDeclared, error) {
	event := new(GameJamWinnerDeclared)
	if err := _GameJam.contract.UnpackLog(event, "WinnerDeclared", log); err != nil {
		return nil, err
	}
	return event, nil
}
