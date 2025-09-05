// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eventtriggertesthelper

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

// EventtriggertesthelperMetaData contains all meta data concerning the Eventtriggertesthelper contract.
var EventtriggertesthelperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"trigger\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Trigger\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"data1\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"data2\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101af8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063838860701461002d575b5f5ffd5b61004760048036038101906100429190610108565b610049565b005b828467ffffffffffffffff167f4e62cf06e16b508286e9b2a7e026a6b8d6e4f0a4e863a7b9c127bad45a6d0868848460405161008692919061017b565b60405180910390a350505050565b5f5ffd5b5f67ffffffffffffffff82169050919050565b6100b481610098565b81146100be575f5ffd5b50565b5f813590506100cf816100ab565b92915050565b5f819050919050565b6100e7816100d5565b81146100f1575f5ffd5b50565b5f81359050610102816100de565b92915050565b5f5f5f5f608085870312156101205761011f610094565b5b5f61012d878288016100c1565b945050602061013e878288016100f4565b935050604061014f878288016100f4565b9250506060610160878288016100f4565b91505092959194509250565b610175816100d5565b82525050565b5f60408201905061018e5f83018561016c565b61019b602083018461016c565b939250505056fea164736f6c634300081c000a",
}

// EventtriggertesthelperABI is the input ABI used to generate the binding from.
// Deprecated: Use EventtriggertesthelperMetaData.ABI instead.
var EventtriggertesthelperABI = EventtriggertesthelperMetaData.ABI

// EventtriggertesthelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventtriggertesthelperMetaData.Bin instead.
var EventtriggertesthelperBin = EventtriggertesthelperMetaData.Bin

// DeployEventtriggertesthelper deploys a new Ethereum contract, binding an instance of Eventtriggertesthelper to it.
func DeployEventtriggertesthelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eventtriggertesthelper, error) {
	parsed, err := EventtriggertesthelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventtriggertesthelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eventtriggertesthelper{EventtriggertesthelperCaller: EventtriggertesthelperCaller{contract: contract}, EventtriggertesthelperTransactor: EventtriggertesthelperTransactor{contract: contract}, EventtriggertesthelperFilterer: EventtriggertesthelperFilterer{contract: contract}}, nil
}

// Eventtriggertesthelper is an auto generated Go binding around an Ethereum contract.
type Eventtriggertesthelper struct {
	EventtriggertesthelperCaller     // Read-only binding to the contract
	EventtriggertesthelperTransactor // Write-only binding to the contract
	EventtriggertesthelperFilterer   // Log filterer for contract events
}

// EventtriggertesthelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventtriggertesthelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventtriggertesthelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventtriggertesthelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventtriggertesthelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventtriggertesthelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventtriggertesthelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventtriggertesthelperSession struct {
	Contract     *Eventtriggertesthelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventtriggertesthelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventtriggertesthelperCallerSession struct {
	Contract *EventtriggertesthelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// EventtriggertesthelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventtriggertesthelperTransactorSession struct {
	Contract     *EventtriggertesthelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// EventtriggertesthelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventtriggertesthelperRaw struct {
	Contract *Eventtriggertesthelper // Generic contract binding to access the raw methods on
}

// EventtriggertesthelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventtriggertesthelperCallerRaw struct {
	Contract *EventtriggertesthelperCaller // Generic read-only contract binding to access the raw methods on
}

// EventtriggertesthelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventtriggertesthelperTransactorRaw struct {
	Contract *EventtriggertesthelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventtriggertesthelper creates a new instance of Eventtriggertesthelper, bound to a specific deployed contract.
func NewEventtriggertesthelper(address common.Address, backend bind.ContractBackend) (*Eventtriggertesthelper, error) {
	contract, err := bindEventtriggertesthelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eventtriggertesthelper{EventtriggertesthelperCaller: EventtriggertesthelperCaller{contract: contract}, EventtriggertesthelperTransactor: EventtriggertesthelperTransactor{contract: contract}, EventtriggertesthelperFilterer: EventtriggertesthelperFilterer{contract: contract}}, nil
}

// NewEventtriggertesthelperCaller creates a new read-only instance of Eventtriggertesthelper, bound to a specific deployed contract.
func NewEventtriggertesthelperCaller(address common.Address, caller bind.ContractCaller) (*EventtriggertesthelperCaller, error) {
	contract, err := bindEventtriggertesthelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventtriggertesthelperCaller{contract: contract}, nil
}

// NewEventtriggertesthelperTransactor creates a new write-only instance of Eventtriggertesthelper, bound to a specific deployed contract.
func NewEventtriggertesthelperTransactor(address common.Address, transactor bind.ContractTransactor) (*EventtriggertesthelperTransactor, error) {
	contract, err := bindEventtriggertesthelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventtriggertesthelperTransactor{contract: contract}, nil
}

// NewEventtriggertesthelperFilterer creates a new log filterer instance of Eventtriggertesthelper, bound to a specific deployed contract.
func NewEventtriggertesthelperFilterer(address common.Address, filterer bind.ContractFilterer) (*EventtriggertesthelperFilterer, error) {
	contract, err := bindEventtriggertesthelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventtriggertesthelperFilterer{contract: contract}, nil
}

// bindEventtriggertesthelper binds a generic wrapper to an already deployed contract.
func bindEventtriggertesthelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventtriggertesthelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventtriggertesthelper *EventtriggertesthelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventtriggertesthelper.Contract.EventtriggertesthelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventtriggertesthelper *EventtriggertesthelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.EventtriggertesthelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventtriggertesthelper *EventtriggertesthelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.EventtriggertesthelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventtriggertesthelper *EventtriggertesthelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventtriggertesthelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventtriggertesthelper *EventtriggertesthelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventtriggertesthelper *EventtriggertesthelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.contract.Transact(opts, method, params...)
}

// Trigger is a paid mutator transaction binding the contract method 0x83886070.
//
// Solidity: function trigger(uint64 topic1, bytes32 topic2, bytes32 data1, bytes32 data2) returns()
func (_Eventtriggertesthelper *EventtriggertesthelperTransactor) Trigger(opts *bind.TransactOpts, topic1 uint64, topic2 [32]byte, data1 [32]byte, data2 [32]byte) (*types.Transaction, error) {
	return _Eventtriggertesthelper.contract.Transact(opts, "trigger", topic1, topic2, data1, data2)
}

// Trigger is a paid mutator transaction binding the contract method 0x83886070.
//
// Solidity: function trigger(uint64 topic1, bytes32 topic2, bytes32 data1, bytes32 data2) returns()
func (_Eventtriggertesthelper *EventtriggertesthelperSession) Trigger(topic1 uint64, topic2 [32]byte, data1 [32]byte, data2 [32]byte) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.Trigger(&_Eventtriggertesthelper.TransactOpts, topic1, topic2, data1, data2)
}

// Trigger is a paid mutator transaction binding the contract method 0x83886070.
//
// Solidity: function trigger(uint64 topic1, bytes32 topic2, bytes32 data1, bytes32 data2) returns()
func (_Eventtriggertesthelper *EventtriggertesthelperTransactorSession) Trigger(topic1 uint64, topic2 [32]byte, data1 [32]byte, data2 [32]byte) (*types.Transaction, error) {
	return _Eventtriggertesthelper.Contract.Trigger(&_Eventtriggertesthelper.TransactOpts, topic1, topic2, data1, data2)
}

// EventtriggertesthelperTriggerIterator is returned from FilterTrigger and is used to iterate over the raw logs and unpacked data for Trigger events raised by the Eventtriggertesthelper contract.
type EventtriggertesthelperTriggerIterator struct {
	Event *EventtriggertesthelperTrigger // Event containing the contract specifics and raw log

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
func (it *EventtriggertesthelperTriggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventtriggertesthelperTrigger)
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
		it.Event = new(EventtriggertesthelperTrigger)
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
func (it *EventtriggertesthelperTriggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventtriggertesthelperTriggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventtriggertesthelperTrigger represents a Trigger event raised by the Eventtriggertesthelper contract.
type EventtriggertesthelperTrigger struct {
	Topic1 uint64
	Topic2 [32]byte
	Data1  [32]byte
	Data2  [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrigger is a free log retrieval operation binding the contract event 0x4e62cf06e16b508286e9b2a7e026a6b8d6e4f0a4e863a7b9c127bad45a6d0868.
//
// Solidity: event Trigger(uint64 indexed topic1, bytes32 indexed topic2, bytes32 data1, bytes32 data2)
func (_Eventtriggertesthelper *EventtriggertesthelperFilterer) FilterTrigger(opts *bind.FilterOpts, topic1 []uint64, topic2 [][32]byte) (*EventtriggertesthelperTriggerIterator, error) {

	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Eventtriggertesthelper.contract.FilterLogs(opts, "Trigger", topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return &EventtriggertesthelperTriggerIterator{contract: _Eventtriggertesthelper.contract, event: "Trigger", logs: logs, sub: sub}, nil
}

// WatchTrigger is a free log subscription operation binding the contract event 0x4e62cf06e16b508286e9b2a7e026a6b8d6e4f0a4e863a7b9c127bad45a6d0868.
//
// Solidity: event Trigger(uint64 indexed topic1, bytes32 indexed topic2, bytes32 data1, bytes32 data2)
func (_Eventtriggertesthelper *EventtriggertesthelperFilterer) WatchTrigger(opts *bind.WatchOpts, sink chan<- *EventtriggertesthelperTrigger, topic1 []uint64, topic2 [][32]byte) (event.Subscription, error) {

	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _Eventtriggertesthelper.contract.WatchLogs(opts, "Trigger", topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventtriggertesthelperTrigger)
				if err := _Eventtriggertesthelper.contract.UnpackLog(event, "Trigger", log); err != nil {
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

// ParseTrigger is a log parse operation binding the contract event 0x4e62cf06e16b508286e9b2a7e026a6b8d6e4f0a4e863a7b9c127bad45a6d0868.
//
// Solidity: event Trigger(uint64 indexed topic1, bytes32 indexed topic2, bytes32 data1, bytes32 data2)
func (_Eventtriggertesthelper *EventtriggertesthelperFilterer) ParseTrigger(log types.Log) (*EventtriggertesthelperTrigger, error) {
	event := new(EventtriggertesthelperTrigger)
	if err := _Eventtriggertesthelper.contract.UnpackLog(event, "Trigger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
