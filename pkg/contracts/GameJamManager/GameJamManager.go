// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GameJamManager

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

// GameJamManagerABI is the input ABI used to generate the binding from.
const GameJamManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gameJamAddress\",\"type\":\"address\"}],\"name\":\"GameJamAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GameJamManagerAdded\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"gameJamList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isGameJamManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGameJamAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addNewGameJam\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// GameJamManager is an auto generated Go binding around an Ethereum contract.
type GameJamManager struct {
	GameJamManagerCaller     // Read-only binding to the contract
	GameJamManagerTransactor // Write-only binding to the contract
	GameJamManagerFilterer   // Log filterer for contract events
}

// GameJamManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GameJamManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GameJamManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GameJamManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameJamManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GameJamManagerSession struct {
	Contract     *GameJamManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameJamManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GameJamManagerCallerSession struct {
	Contract *GameJamManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GameJamManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GameJamManagerTransactorSession struct {
	Contract     *GameJamManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GameJamManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GameJamManagerRaw struct {
	Contract *GameJamManager // Generic contract binding to access the raw methods on
}

// GameJamManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GameJamManagerCallerRaw struct {
	Contract *GameJamManagerCaller // Generic read-only contract binding to access the raw methods on
}

// GameJamManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GameJamManagerTransactorRaw struct {
	Contract *GameJamManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGameJamManager creates a new instance of GameJamManager, bound to a specific deployed contract.
func NewGameJamManager(address common.Address, backend bind.ContractBackend) (*GameJamManager, error) {
	contract, err := bindGameJamManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameJamManager{GameJamManagerCaller: GameJamManagerCaller{contract: contract}, GameJamManagerTransactor: GameJamManagerTransactor{contract: contract}, GameJamManagerFilterer: GameJamManagerFilterer{contract: contract}}, nil
}

// NewGameJamManagerCaller creates a new read-only instance of GameJamManager, bound to a specific deployed contract.
func NewGameJamManagerCaller(address common.Address, caller bind.ContractCaller) (*GameJamManagerCaller, error) {
	contract, err := bindGameJamManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamManagerCaller{contract: contract}, nil
}

// NewGameJamManagerTransactor creates a new write-only instance of GameJamManager, bound to a specific deployed contract.
func NewGameJamManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GameJamManagerTransactor, error) {
	contract, err := bindGameJamManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameJamManagerTransactor{contract: contract}, nil
}

// NewGameJamManagerFilterer creates a new log filterer instance of GameJamManager, bound to a specific deployed contract.
func NewGameJamManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GameJamManagerFilterer, error) {
	contract, err := bindGameJamManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameJamManagerFilterer{contract: contract}, nil
}

// bindGameJamManager binds a generic wrapper to an already deployed contract.
func bindGameJamManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GameJamManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamManager *GameJamManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamManager.Contract.GameJamManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamManager *GameJamManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamManager.Contract.GameJamManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamManager *GameJamManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamManager.Contract.GameJamManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameJamManager *GameJamManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GameJamManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GameJamManager *GameJamManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameJamManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GameJamManager *GameJamManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameJamManager.Contract.contract.Transact(opts, method, params...)
}

// GameJamList is a free data retrieval call binding the contract method 0x96081f35.
//
// Solidity: function gameJamList(bytes32 ) constant returns(address)
func (_GameJamManager *GameJamManagerCaller) GameJamList(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GameJamManager.contract.Call(opts, out, "gameJamList", arg0)
	return *ret0, err
}

// GameJamList is a free data retrieval call binding the contract method 0x96081f35.
//
// Solidity: function gameJamList(bytes32 ) constant returns(address)
func (_GameJamManager *GameJamManagerSession) GameJamList(arg0 [32]byte) (common.Address, error) {
	return _GameJamManager.Contract.GameJamList(&_GameJamManager.CallOpts, arg0)
}

// GameJamList is a free data retrieval call binding the contract method 0x96081f35.
//
// Solidity: function gameJamList(bytes32 ) constant returns(address)
func (_GameJamManager *GameJamManagerCallerSession) GameJamList(arg0 [32]byte) (common.Address, error) {
	return _GameJamManager.Contract.GameJamList(&_GameJamManager.CallOpts, arg0)
}

// GetAllGameJamAddresses is a free data retrieval call binding the contract method 0xa92881d8.
//
// Solidity: function getAllGameJamAddresses() constant returns(address[])
func (_GameJamManager *GameJamManagerCaller) GetAllGameJamAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GameJamManager.contract.Call(opts, out, "getAllGameJamAddresses")
	return *ret0, err
}

// GetAllGameJamAddresses is a free data retrieval call binding the contract method 0xa92881d8.
//
// Solidity: function getAllGameJamAddresses() constant returns(address[])
func (_GameJamManager *GameJamManagerSession) GetAllGameJamAddresses() ([]common.Address, error) {
	return _GameJamManager.Contract.GetAllGameJamAddresses(&_GameJamManager.CallOpts)
}

// GetAllGameJamAddresses is a free data retrieval call binding the contract method 0xa92881d8.
//
// Solidity: function getAllGameJamAddresses() constant returns(address[])
func (_GameJamManager *GameJamManagerCallerSession) GetAllGameJamAddresses() ([]common.Address, error) {
	return _GameJamManager.Contract.GetAllGameJamAddresses(&_GameJamManager.CallOpts)
}

// IsGameJamManager is a free data retrieval call binding the contract method 0x0a3e1f8f.
//
// Solidity: function isGameJamManager(address account) constant returns(bool)
func (_GameJamManager *GameJamManagerCaller) IsGameJamManager(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GameJamManager.contract.Call(opts, out, "isGameJamManager", account)
	return *ret0, err
}

// IsGameJamManager is a free data retrieval call binding the contract method 0x0a3e1f8f.
//
// Solidity: function isGameJamManager(address account) constant returns(bool)
func (_GameJamManager *GameJamManagerSession) IsGameJamManager(account common.Address) (bool, error) {
	return _GameJamManager.Contract.IsGameJamManager(&_GameJamManager.CallOpts, account)
}

// IsGameJamManager is a free data retrieval call binding the contract method 0x0a3e1f8f.
//
// Solidity: function isGameJamManager(address account) constant returns(bool)
func (_GameJamManager *GameJamManagerCallerSession) IsGameJamManager(account common.Address) (bool, error) {
	return _GameJamManager.Contract.IsGameJamManager(&_GameJamManager.CallOpts, account)
}

// AddNewGameJam is a paid mutator transaction binding the contract method 0xeddccc24.
//
// Solidity: function addNewGameJam(bytes32 _name, address admin) returns(bytes32)
func (_GameJamManager *GameJamManagerTransactor) AddNewGameJam(opts *bind.TransactOpts, _name [32]byte, admin common.Address) (*types.Transaction, error) {
	return _GameJamManager.contract.Transact(opts, "addNewGameJam", _name, admin)
}

// AddNewGameJam is a paid mutator transaction binding the contract method 0xeddccc24.
//
// Solidity: function addNewGameJam(bytes32 _name, address admin) returns(bytes32)
func (_GameJamManager *GameJamManagerSession) AddNewGameJam(_name [32]byte, admin common.Address) (*types.Transaction, error) {
	return _GameJamManager.Contract.AddNewGameJam(&_GameJamManager.TransactOpts, _name, admin)
}

// AddNewGameJam is a paid mutator transaction binding the contract method 0xeddccc24.
//
// Solidity: function addNewGameJam(bytes32 _name, address admin) returns(bytes32)
func (_GameJamManager *GameJamManagerTransactorSession) AddNewGameJam(_name [32]byte, admin common.Address) (*types.Transaction, error) {
	return _GameJamManager.Contract.AddNewGameJam(&_GameJamManager.TransactOpts, _name, admin)
}

// GameJamManagerGameJamAddedIterator is returned from FilterGameJamAdded and is used to iterate over the raw logs and unpacked data for GameJamAdded events raised by the GameJamManager contract.
type GameJamManagerGameJamAddedIterator struct {
	Event *GameJamManagerGameJamAdded // Event containing the contract specifics and raw log

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
func (it *GameJamManagerGameJamAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamManagerGameJamAdded)
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
		it.Event = new(GameJamManagerGameJamAdded)
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
func (it *GameJamManagerGameJamAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamManagerGameJamAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamManagerGameJamAdded represents a GameJamAdded event raised by the GameJamManager contract.
type GameJamManagerGameJamAdded struct {
	GameJamAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGameJamAdded is a free log retrieval operation binding the contract event 0xf314602fc17c3119a1d591fa560483ddb604778c0ef2b4821be21a7a2781e95f.
//
// Solidity: event GameJamAdded(address indexed gameJamAddress)
func (_GameJamManager *GameJamManagerFilterer) FilterGameJamAdded(opts *bind.FilterOpts, gameJamAddress []common.Address) (*GameJamManagerGameJamAddedIterator, error) {

	var gameJamAddressRule []interface{}
	for _, gameJamAddressItem := range gameJamAddress {
		gameJamAddressRule = append(gameJamAddressRule, gameJamAddressItem)
	}

	logs, sub, err := _GameJamManager.contract.FilterLogs(opts, "GameJamAdded", gameJamAddressRule)
	if err != nil {
		return nil, err
	}
	return &GameJamManagerGameJamAddedIterator{contract: _GameJamManager.contract, event: "GameJamAdded", logs: logs, sub: sub}, nil
}

// WatchGameJamAdded is a free log subscription operation binding the contract event 0xf314602fc17c3119a1d591fa560483ddb604778c0ef2b4821be21a7a2781e95f.
//
// Solidity: event GameJamAdded(address indexed gameJamAddress)
func (_GameJamManager *GameJamManagerFilterer) WatchGameJamAdded(opts *bind.WatchOpts, sink chan<- *GameJamManagerGameJamAdded, gameJamAddress []common.Address) (event.Subscription, error) {

	var gameJamAddressRule []interface{}
	for _, gameJamAddressItem := range gameJamAddress {
		gameJamAddressRule = append(gameJamAddressRule, gameJamAddressItem)
	}

	logs, sub, err := _GameJamManager.contract.WatchLogs(opts, "GameJamAdded", gameJamAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamManagerGameJamAdded)
				if err := _GameJamManager.contract.UnpackLog(event, "GameJamAdded", log); err != nil {
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

// ParseGameJamAdded is a log parse operation binding the contract event 0xf314602fc17c3119a1d591fa560483ddb604778c0ef2b4821be21a7a2781e95f.
//
// Solidity: event GameJamAdded(address indexed gameJamAddress)
func (_GameJamManager *GameJamManagerFilterer) ParseGameJamAdded(log types.Log) (*GameJamManagerGameJamAdded, error) {
	event := new(GameJamManagerGameJamAdded)
	if err := _GameJamManager.contract.UnpackLog(event, "GameJamAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GameJamManagerGameJamManagerAddedIterator is returned from FilterGameJamManagerAdded and is used to iterate over the raw logs and unpacked data for GameJamManagerAdded events raised by the GameJamManager contract.
type GameJamManagerGameJamManagerAddedIterator struct {
	Event *GameJamManagerGameJamManagerAdded // Event containing the contract specifics and raw log

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
func (it *GameJamManagerGameJamManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameJamManagerGameJamManagerAdded)
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
		it.Event = new(GameJamManagerGameJamManagerAdded)
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
func (it *GameJamManagerGameJamManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameJamManagerGameJamManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameJamManagerGameJamManagerAdded represents a GameJamManagerAdded event raised by the GameJamManager contract.
type GameJamManagerGameJamManagerAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGameJamManagerAdded is a free log retrieval operation binding the contract event 0x10e5a1aa6f2296a041c44f80adf7d836b963a23e6622c82092efa6d996095a8b.
//
// Solidity: event GameJamManagerAdded(address indexed account)
func (_GameJamManager *GameJamManagerFilterer) FilterGameJamManagerAdded(opts *bind.FilterOpts, account []common.Address) (*GameJamManagerGameJamManagerAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GameJamManager.contract.FilterLogs(opts, "GameJamManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &GameJamManagerGameJamManagerAddedIterator{contract: _GameJamManager.contract, event: "GameJamManagerAdded", logs: logs, sub: sub}, nil
}

// WatchGameJamManagerAdded is a free log subscription operation binding the contract event 0x10e5a1aa6f2296a041c44f80adf7d836b963a23e6622c82092efa6d996095a8b.
//
// Solidity: event GameJamManagerAdded(address indexed account)
func (_GameJamManager *GameJamManagerFilterer) WatchGameJamManagerAdded(opts *bind.WatchOpts, sink chan<- *GameJamManagerGameJamManagerAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _GameJamManager.contract.WatchLogs(opts, "GameJamManagerAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameJamManagerGameJamManagerAdded)
				if err := _GameJamManager.contract.UnpackLog(event, "GameJamManagerAdded", log); err != nil {
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

// ParseGameJamManagerAdded is a log parse operation binding the contract event 0x10e5a1aa6f2296a041c44f80adf7d836b963a23e6622c82092efa6d996095a8b.
//
// Solidity: event GameJamManagerAdded(address indexed account)
func (_GameJamManager *GameJamManagerFilterer) ParseGameJamManagerAdded(log types.Log) (*GameJamManagerGameJamManagerAdded, error) {
	event := new(GameJamManagerGameJamManagerAdded)
	if err := _GameJamManager.contract.UnpackLog(event, "GameJamManagerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}
