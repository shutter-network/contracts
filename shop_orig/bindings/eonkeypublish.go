// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// EonKeyPublishMetaData contains all meta data concerning the EonKeyPublish contract.
var EonKeyPublishMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyperSet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_broadcaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eonKeyConfirmed\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasKeyperVoted\",\"inputs\":[{\"name\":\"keyper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishEonKey\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keyperId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EonVoteRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
}

// EonKeyPublishABI is the input ABI used to generate the binding from.
// Deprecated: Use EonKeyPublishMetaData.ABI instead.
var EonKeyPublishABI = EonKeyPublishMetaData.ABI

// EonKeyPublish is an auto generated Go binding around an Ethereum contract.
type EonKeyPublish struct {
	EonKeyPublishCaller     // Read-only binding to the contract
	EonKeyPublishTransactor // Write-only binding to the contract
	EonKeyPublishFilterer   // Log filterer for contract events
}

// EonKeyPublishCaller is an auto generated read-only Go binding around an Ethereum contract.
type EonKeyPublishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonKeyPublishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EonKeyPublishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonKeyPublishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EonKeyPublishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonKeyPublishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EonKeyPublishSession struct {
	Contract     *EonKeyPublish    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EonKeyPublishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EonKeyPublishCallerSession struct {
	Contract *EonKeyPublishCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EonKeyPublishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EonKeyPublishTransactorSession struct {
	Contract     *EonKeyPublishTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EonKeyPublishRaw is an auto generated low-level Go binding around an Ethereum contract.
type EonKeyPublishRaw struct {
	Contract *EonKeyPublish // Generic contract binding to access the raw methods on
}

// EonKeyPublishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EonKeyPublishCallerRaw struct {
	Contract *EonKeyPublishCaller // Generic read-only contract binding to access the raw methods on
}

// EonKeyPublishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EonKeyPublishTransactorRaw struct {
	Contract *EonKeyPublishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEonKeyPublish creates a new instance of EonKeyPublish, bound to a specific deployed contract.
func NewEonKeyPublish(address common.Address, backend bind.ContractBackend) (*EonKeyPublish, error) {
	contract, err := bindEonKeyPublish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EonKeyPublish{EonKeyPublishCaller: EonKeyPublishCaller{contract: contract}, EonKeyPublishTransactor: EonKeyPublishTransactor{contract: contract}, EonKeyPublishFilterer: EonKeyPublishFilterer{contract: contract}}, nil
}

// NewEonKeyPublishCaller creates a new read-only instance of EonKeyPublish, bound to a specific deployed contract.
func NewEonKeyPublishCaller(address common.Address, caller bind.ContractCaller) (*EonKeyPublishCaller, error) {
	contract, err := bindEonKeyPublish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EonKeyPublishCaller{contract: contract}, nil
}

// NewEonKeyPublishTransactor creates a new write-only instance of EonKeyPublish, bound to a specific deployed contract.
func NewEonKeyPublishTransactor(address common.Address, transactor bind.ContractTransactor) (*EonKeyPublishTransactor, error) {
	contract, err := bindEonKeyPublish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EonKeyPublishTransactor{contract: contract}, nil
}

// NewEonKeyPublishFilterer creates a new log filterer instance of EonKeyPublish, bound to a specific deployed contract.
func NewEonKeyPublishFilterer(address common.Address, filterer bind.ContractFilterer) (*EonKeyPublishFilterer, error) {
	contract, err := bindEonKeyPublish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EonKeyPublishFilterer{contract: contract}, nil
}

// bindEonKeyPublish binds a generic wrapper to an already deployed contract.
func bindEonKeyPublish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EonKeyPublishMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EonKeyPublish *EonKeyPublishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EonKeyPublish.Contract.EonKeyPublishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EonKeyPublish *EonKeyPublishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.EonKeyPublishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EonKeyPublish *EonKeyPublishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.EonKeyPublishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EonKeyPublish *EonKeyPublishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EonKeyPublish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EonKeyPublish *EonKeyPublishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EonKeyPublish *EonKeyPublishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.contract.Transact(opts, method, params...)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_EonKeyPublish *EonKeyPublishCaller) EonKeyConfirmed(opts *bind.CallOpts, eonKey []byte) (bool, error) {
	var out []interface{}
	err := _EonKeyPublish.contract.Call(opts, &out, "eonKeyConfirmed", eonKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_EonKeyPublish *EonKeyPublishSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _EonKeyPublish.Contract.EonKeyConfirmed(&_EonKeyPublish.CallOpts, eonKey)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_EonKeyPublish *EonKeyPublishCallerSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _EonKeyPublish.Contract.EonKeyConfirmed(&_EonKeyPublish.CallOpts, eonKey)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_EonKeyPublish *EonKeyPublishCaller) HasKeyperVoted(opts *bind.CallOpts, keyper common.Address) (bool, error) {
	var out []interface{}
	err := _EonKeyPublish.contract.Call(opts, &out, "hasKeyperVoted", keyper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_EonKeyPublish *EonKeyPublishSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _EonKeyPublish.Contract.HasKeyperVoted(&_EonKeyPublish.CallOpts, keyper)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_EonKeyPublish *EonKeyPublishCallerSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _EonKeyPublish.Contract.HasKeyperVoted(&_EonKeyPublish.CallOpts, keyper)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_EonKeyPublish *EonKeyPublishTransactor) PublishEonKey(opts *bind.TransactOpts, eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _EonKeyPublish.contract.Transact(opts, "publishEonKey", eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_EonKeyPublish *EonKeyPublishSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.PublishEonKey(&_EonKeyPublish.TransactOpts, eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_EonKeyPublish *EonKeyPublishTransactorSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _EonKeyPublish.Contract.PublishEonKey(&_EonKeyPublish.TransactOpts, eonKey, keyperId)
}

// EonKeyPublishEonVoteRegisteredIterator is returned from FilterEonVoteRegistered and is used to iterate over the raw logs and unpacked data for EonVoteRegistered events raised by the EonKeyPublish contract.
type EonKeyPublishEonVoteRegisteredIterator struct {
	Event *EonKeyPublishEonVoteRegistered // Event containing the contract specifics and raw log

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
func (it *EonKeyPublishEonVoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EonKeyPublishEonVoteRegistered)
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
		it.Event = new(EonKeyPublishEonVoteRegistered)
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
func (it *EonKeyPublishEonVoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EonKeyPublishEonVoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EonKeyPublishEonVoteRegistered represents a EonVoteRegistered event raised by the EonKeyPublish contract.
type EonKeyPublishEonVoteRegistered struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonVoteRegistered is a free log retrieval operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_EonKeyPublish *EonKeyPublishFilterer) FilterEonVoteRegistered(opts *bind.FilterOpts) (*EonKeyPublishEonVoteRegisteredIterator, error) {

	logs, sub, err := _EonKeyPublish.contract.FilterLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return &EonKeyPublishEonVoteRegisteredIterator{contract: _EonKeyPublish.contract, event: "EonVoteRegistered", logs: logs, sub: sub}, nil
}

// WatchEonVoteRegistered is a free log subscription operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_EonKeyPublish *EonKeyPublishFilterer) WatchEonVoteRegistered(opts *bind.WatchOpts, sink chan<- *EonKeyPublishEonVoteRegistered) (event.Subscription, error) {

	logs, sub, err := _EonKeyPublish.contract.WatchLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EonKeyPublishEonVoteRegistered)
				if err := _EonKeyPublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
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

// ParseEonVoteRegistered is a log parse operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_EonKeyPublish *EonKeyPublishFilterer) ParseEonVoteRegistered(log types.Log) (*EonKeyPublishEonVoteRegistered, error) {
	event := new(EonKeyPublishEonVoteRegistered)
	if err := _EonKeyPublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
