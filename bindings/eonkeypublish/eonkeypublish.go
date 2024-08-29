// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eonkeypublish

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

// EonkeypublishMetaData contains all meta data concerning the Eonkeypublish contract.
var EonkeypublishMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyperSet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_broadcaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eonKeyConfirmed\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasKeyperVoted\",\"inputs\":[{\"name\":\"keyper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishEonKey\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keyperId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EonVoteRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
}

// EonkeypublishABI is the input ABI used to generate the binding from.
// Deprecated: Use EonkeypublishMetaData.ABI instead.
var EonkeypublishABI = EonkeypublishMetaData.ABI

// Eonkeypublish is an auto generated Go binding around an Ethereum contract.
type Eonkeypublish struct {
	EonkeypublishCaller     // Read-only binding to the contract
	EonkeypublishTransactor // Write-only binding to the contract
	EonkeypublishFilterer   // Log filterer for contract events
}

// EonkeypublishCaller is an auto generated read-only Go binding around an Ethereum contract.
type EonkeypublishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EonkeypublishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EonkeypublishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EonkeypublishSession struct {
	Contract     *Eonkeypublish    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EonkeypublishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EonkeypublishCallerSession struct {
	Contract *EonkeypublishCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EonkeypublishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EonkeypublishTransactorSession struct {
	Contract     *EonkeypublishTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EonkeypublishRaw is an auto generated low-level Go binding around an Ethereum contract.
type EonkeypublishRaw struct {
	Contract *Eonkeypublish // Generic contract binding to access the raw methods on
}

// EonkeypublishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EonkeypublishCallerRaw struct {
	Contract *EonkeypublishCaller // Generic read-only contract binding to access the raw methods on
}

// EonkeypublishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EonkeypublishTransactorRaw struct {
	Contract *EonkeypublishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEonkeypublish creates a new instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublish(address common.Address, backend bind.ContractBackend) (*Eonkeypublish, error) {
	contract, err := bindEonkeypublish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eonkeypublish{EonkeypublishCaller: EonkeypublishCaller{contract: contract}, EonkeypublishTransactor: EonkeypublishTransactor{contract: contract}, EonkeypublishFilterer: EonkeypublishFilterer{contract: contract}}, nil
}

// NewEonkeypublishCaller creates a new read-only instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishCaller(address common.Address, caller bind.ContractCaller) (*EonkeypublishCaller, error) {
	contract, err := bindEonkeypublish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishCaller{contract: contract}, nil
}

// NewEonkeypublishTransactor creates a new write-only instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishTransactor(address common.Address, transactor bind.ContractTransactor) (*EonkeypublishTransactor, error) {
	contract, err := bindEonkeypublish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishTransactor{contract: contract}, nil
}

// NewEonkeypublishFilterer creates a new log filterer instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishFilterer(address common.Address, filterer bind.ContractFilterer) (*EonkeypublishFilterer, error) {
	contract, err := bindEonkeypublish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishFilterer{contract: contract}, nil
}

// bindEonkeypublish binds a generic wrapper to an already deployed contract.
func bindEonkeypublish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EonkeypublishMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eonkeypublish *EonkeypublishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eonkeypublish.Contract.EonkeypublishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eonkeypublish *EonkeypublishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.EonkeypublishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eonkeypublish *EonkeypublishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.EonkeypublishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eonkeypublish *EonkeypublishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eonkeypublish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eonkeypublish *EonkeypublishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eonkeypublish *EonkeypublishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.contract.Transact(opts, method, params...)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishCaller) EonKeyConfirmed(opts *bind.CallOpts, eonKey []byte) (bool, error) {
	var out []interface{}
	err := _Eonkeypublish.contract.Call(opts, &out, "eonKeyConfirmed", eonKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Eonkeypublish.Contract.EonKeyConfirmed(&_Eonkeypublish.CallOpts, eonKey)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishCallerSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Eonkeypublish.Contract.EonKeyConfirmed(&_Eonkeypublish.CallOpts, eonKey)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishCaller) HasKeyperVoted(opts *bind.CallOpts, keyper common.Address) (bool, error) {
	var out []interface{}
	err := _Eonkeypublish.contract.Call(opts, &out, "hasKeyperVoted", keyper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _Eonkeypublish.Contract.HasKeyperVoted(&_Eonkeypublish.CallOpts, keyper)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishCallerSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _Eonkeypublish.Contract.HasKeyperVoted(&_Eonkeypublish.CallOpts, keyper)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishTransactor) PublishEonKey(opts *bind.TransactOpts, eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.contract.Transact(opts, "publishEonKey", eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.PublishEonKey(&_Eonkeypublish.TransactOpts, eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishTransactorSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.PublishEonKey(&_Eonkeypublish.TransactOpts, eonKey, keyperId)
}

// EonkeypublishEonVoteRegisteredIterator is returned from FilterEonVoteRegistered and is used to iterate over the raw logs and unpacked data for EonVoteRegistered events raised by the Eonkeypublish contract.
type EonkeypublishEonVoteRegisteredIterator struct {
	Event *EonkeypublishEonVoteRegistered // Event containing the contract specifics and raw log

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
func (it *EonkeypublishEonVoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EonkeypublishEonVoteRegistered)
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
		it.Event = new(EonkeypublishEonVoteRegistered)
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
func (it *EonkeypublishEonVoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EonkeypublishEonVoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EonkeypublishEonVoteRegistered represents a EonVoteRegistered event raised by the Eonkeypublish contract.
type EonkeypublishEonVoteRegistered struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonVoteRegistered is a free log retrieval operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Eonkeypublish *EonkeypublishFilterer) FilterEonVoteRegistered(opts *bind.FilterOpts) (*EonkeypublishEonVoteRegisteredIterator, error) {

	logs, sub, err := _Eonkeypublish.contract.FilterLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return &EonkeypublishEonVoteRegisteredIterator{contract: _Eonkeypublish.contract, event: "EonVoteRegistered", logs: logs, sub: sub}, nil
}

// WatchEonVoteRegistered is a free log subscription operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Eonkeypublish *EonkeypublishFilterer) WatchEonVoteRegistered(opts *bind.WatchOpts, sink chan<- *EonkeypublishEonVoteRegistered) (event.Subscription, error) {

	logs, sub, err := _Eonkeypublish.contract.WatchLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EonkeypublishEonVoteRegistered)
				if err := _Eonkeypublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
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
func (_Eonkeypublish *EonkeypublishFilterer) ParseEonVoteRegistered(log types.Log) (*EonkeypublishEonVoteRegistered, error) {
	event := new(EonkeypublishEonVoteRegistered)
	if err := _Eonkeypublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
