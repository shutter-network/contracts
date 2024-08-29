// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keybroadcastcontract

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

// KeybroadcastcontractMetaData contains all meta data concerning the Keybroadcastcontract contract.
var KeybroadcastcontractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcastEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EonKeyBroadcast\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyHaveKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
}

// KeybroadcastcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use KeybroadcastcontractMetaData.ABI instead.
var KeybroadcastcontractABI = KeybroadcastcontractMetaData.ABI

// Keybroadcastcontract is an auto generated Go binding around an Ethereum contract.
type Keybroadcastcontract struct {
	KeybroadcastcontractCaller     // Read-only binding to the contract
	KeybroadcastcontractTransactor // Write-only binding to the contract
	KeybroadcastcontractFilterer   // Log filterer for contract events
}

// KeybroadcastcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeybroadcastcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeybroadcastcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeybroadcastcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeybroadcastcontractSession struct {
	Contract     *Keybroadcastcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeybroadcastcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeybroadcastcontractCallerSession struct {
	Contract *KeybroadcastcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// KeybroadcastcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeybroadcastcontractTransactorSession struct {
	Contract     *KeybroadcastcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// KeybroadcastcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeybroadcastcontractRaw struct {
	Contract *Keybroadcastcontract // Generic contract binding to access the raw methods on
}

// KeybroadcastcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeybroadcastcontractCallerRaw struct {
	Contract *KeybroadcastcontractCaller // Generic read-only contract binding to access the raw methods on
}

// KeybroadcastcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeybroadcastcontractTransactorRaw struct {
	Contract *KeybroadcastcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeybroadcastcontract creates a new instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontract(address common.Address, backend bind.ContractBackend) (*Keybroadcastcontract, error) {
	contract, err := bindKeybroadcastcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keybroadcastcontract{KeybroadcastcontractCaller: KeybroadcastcontractCaller{contract: contract}, KeybroadcastcontractTransactor: KeybroadcastcontractTransactor{contract: contract}, KeybroadcastcontractFilterer: KeybroadcastcontractFilterer{contract: contract}}, nil
}

// NewKeybroadcastcontractCaller creates a new read-only instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractCaller(address common.Address, caller bind.ContractCaller) (*KeybroadcastcontractCaller, error) {
	contract, err := bindKeybroadcastcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractCaller{contract: contract}, nil
}

// NewKeybroadcastcontractTransactor creates a new write-only instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*KeybroadcastcontractTransactor, error) {
	contract, err := bindKeybroadcastcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractTransactor{contract: contract}, nil
}

// NewKeybroadcastcontractFilterer creates a new log filterer instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*KeybroadcastcontractFilterer, error) {
	contract, err := bindKeybroadcastcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractFilterer{contract: contract}, nil
}

// bindKeybroadcastcontract binds a generic wrapper to an already deployed contract.
func bindKeybroadcastcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeybroadcastcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keybroadcastcontract *KeybroadcastcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keybroadcastcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keybroadcastcontract *KeybroadcastcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keybroadcastcontract *KeybroadcastcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.contract.Transact(opts, method, params...)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractCaller) GetEonKey(opts *bind.CallOpts, eon uint64) ([]byte, error) {
	var out []interface{}
	err := _Keybroadcastcontract.contract.Call(opts, &out, "getEonKey", eon)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Keybroadcastcontract.Contract.GetEonKey(&_Keybroadcastcontract.CallOpts, eon)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractCallerSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Keybroadcastcontract.Contract.GetEonKey(&_Keybroadcastcontract.CallOpts, eon)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractTransactor) BroadcastEonKey(opts *bind.TransactOpts, eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.contract.Transact(opts, "broadcastEonKey", eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.BroadcastEonKey(&_Keybroadcastcontract.TransactOpts, eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractTransactorSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.BroadcastEonKey(&_Keybroadcastcontract.TransactOpts, eon, key)
}

// KeybroadcastcontractEonKeyBroadcastIterator is returned from FilterEonKeyBroadcast and is used to iterate over the raw logs and unpacked data for EonKeyBroadcast events raised by the Keybroadcastcontract contract.
type KeybroadcastcontractEonKeyBroadcastIterator struct {
	Event *KeybroadcastcontractEonKeyBroadcast // Event containing the contract specifics and raw log

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
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeybroadcastcontractEonKeyBroadcast)
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
		it.Event = new(KeybroadcastcontractEonKeyBroadcast)
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
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeybroadcastcontractEonKeyBroadcast represents a EonKeyBroadcast event raised by the Keybroadcastcontract contract.
type KeybroadcastcontractEonKeyBroadcast struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonKeyBroadcast is a free log retrieval operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) FilterEonKeyBroadcast(opts *bind.FilterOpts) (*KeybroadcastcontractEonKeyBroadcastIterator, error) {

	logs, sub, err := _Keybroadcastcontract.contract.FilterLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractEonKeyBroadcastIterator{contract: _Keybroadcastcontract.contract, event: "EonKeyBroadcast", logs: logs, sub: sub}, nil
}

// WatchEonKeyBroadcast is a free log subscription operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) WatchEonKeyBroadcast(opts *bind.WatchOpts, sink chan<- *KeybroadcastcontractEonKeyBroadcast) (event.Subscription, error) {

	logs, sub, err := _Keybroadcastcontract.contract.WatchLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeybroadcastcontractEonKeyBroadcast)
				if err := _Keybroadcastcontract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
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
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) ParseEonKeyBroadcast(log types.Log) (*KeybroadcastcontractEonKeyBroadcast, error) {
	event := new(KeybroadcastcontractEonKeyBroadcast)
	if err := _Keybroadcastcontract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
