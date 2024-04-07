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

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcastEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EonKeyBroadcast\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyHaveKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Bindings *BindingsCaller) GetEonKey(opts *bind.CallOpts, eon uint64) ([]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getEonKey", eon)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Bindings *BindingsSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Bindings.Contract.GetEonKey(&_Bindings.CallOpts, eon)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Bindings *BindingsCallerSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Bindings.Contract.GetEonKey(&_Bindings.CallOpts, eon)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Bindings *BindingsTransactor) BroadcastEonKey(opts *bind.TransactOpts, eon uint64, key []byte) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "broadcastEonKey", eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Bindings *BindingsSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Bindings.Contract.BroadcastEonKey(&_Bindings.TransactOpts, eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Bindings *BindingsTransactorSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Bindings.Contract.BroadcastEonKey(&_Bindings.TransactOpts, eon, key)
}

// BindingsEonKeyBroadcastIterator is returned from FilterEonKeyBroadcast and is used to iterate over the raw logs and unpacked data for EonKeyBroadcast events raised by the Bindings contract.
type BindingsEonKeyBroadcastIterator struct {
	Event *BindingsEonKeyBroadcast // Event containing the contract specifics and raw log

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
func (it *BindingsEonKeyBroadcastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEonKeyBroadcast)
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
		it.Event = new(BindingsEonKeyBroadcast)
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
func (it *BindingsEonKeyBroadcastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEonKeyBroadcastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEonKeyBroadcast represents a EonKeyBroadcast event raised by the Bindings contract.
type BindingsEonKeyBroadcast struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonKeyBroadcast is a free log retrieval operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) FilterEonKeyBroadcast(opts *bind.FilterOpts) (*BindingsEonKeyBroadcastIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return &BindingsEonKeyBroadcastIterator{contract: _Bindings.contract, event: "EonKeyBroadcast", logs: logs, sub: sub}, nil
}

// WatchEonKeyBroadcast is a free log subscription operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) WatchEonKeyBroadcast(opts *bind.WatchOpts, sink chan<- *BindingsEonKeyBroadcast) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEonKeyBroadcast)
				if err := _Bindings.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
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

// ParseEonKeyBroadcast is a log parse operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Bindings *BindingsFilterer) ParseEonKeyBroadcast(log types.Log) (*BindingsEonKeyBroadcast, error) {
	event := new(BindingsEonKeyBroadcast)
	if err := _Bindings.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
