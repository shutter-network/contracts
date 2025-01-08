// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shutterregistry

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
)

// ShutterregistryMetaData contains all meta data concerning the Shutterregistry contract.
var ShutterregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registrations\",\"inputs\":[{\"name\":\"identity\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"IdentityRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIdentityPrefix\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TimestampInThePast\",\"inputs\":[]}]",
}

// ShutterregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ShutterregistryMetaData.ABI instead.
var ShutterregistryABI = ShutterregistryMetaData.ABI

// Shutterregistry is an auto generated Go binding around an Ethereum contract.
type Shutterregistry struct {
	ShutterregistryCaller     // Read-only binding to the contract
	ShutterregistryTransactor // Write-only binding to the contract
	ShutterregistryFilterer   // Log filterer for contract events
}

// ShutterregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShutterregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShutterregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShutterregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShutterregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShutterregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShutterregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShutterregistrySession struct {
	Contract     *Shutterregistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShutterregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShutterregistryCallerSession struct {
	Contract *ShutterregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ShutterregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShutterregistryTransactorSession struct {
	Contract     *ShutterregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ShutterregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShutterregistryRaw struct {
	Contract *Shutterregistry // Generic contract binding to access the raw methods on
}

// ShutterregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShutterregistryCallerRaw struct {
	Contract *ShutterregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ShutterregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShutterregistryTransactorRaw struct {
	Contract *ShutterregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShutterregistry creates a new instance of Shutterregistry, bound to a specific deployed contract.
func NewShutterregistry(address common.Address, backend bind.ContractBackend) (*Shutterregistry, error) {
	contract, err := bindShutterregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shutterregistry{ShutterregistryCaller: ShutterregistryCaller{contract: contract}, ShutterregistryTransactor: ShutterregistryTransactor{contract: contract}, ShutterregistryFilterer: ShutterregistryFilterer{contract: contract}}, nil
}

// NewShutterregistryCaller creates a new read-only instance of Shutterregistry, bound to a specific deployed contract.
func NewShutterregistryCaller(address common.Address, caller bind.ContractCaller) (*ShutterregistryCaller, error) {
	contract, err := bindShutterregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShutterregistryCaller{contract: contract}, nil
}

// NewShutterregistryTransactor creates a new write-only instance of Shutterregistry, bound to a specific deployed contract.
func NewShutterregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ShutterregistryTransactor, error) {
	contract, err := bindShutterregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShutterregistryTransactor{contract: contract}, nil
}

// NewShutterregistryFilterer creates a new log filterer instance of Shutterregistry, bound to a specific deployed contract.
func NewShutterregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ShutterregistryFilterer, error) {
	contract, err := bindShutterregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShutterregistryFilterer{contract: contract}, nil
}

// bindShutterregistry binds a generic wrapper to an already deployed contract.
func bindShutterregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShutterregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shutterregistry *ShutterregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shutterregistry.Contract.ShutterregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shutterregistry *ShutterregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shutterregistry.Contract.ShutterregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shutterregistry *ShutterregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shutterregistry.Contract.ShutterregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shutterregistry *ShutterregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shutterregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shutterregistry *ShutterregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shutterregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shutterregistry *ShutterregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shutterregistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shutterregistry *ShutterregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shutterregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shutterregistry *ShutterregistrySession) Owner() (common.Address, error) {
	return _Shutterregistry.Contract.Owner(&_Shutterregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shutterregistry *ShutterregistryCallerSession) Owner() (common.Address, error) {
	return _Shutterregistry.Contract.Owner(&_Shutterregistry.CallOpts)
}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 timestamp)
func (_Shutterregistry *ShutterregistryCaller) Registrations(opts *bind.CallOpts, identity [32]byte) (struct {
	Eon       uint64
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _Shutterregistry.contract.Call(opts, &out, "registrations", identity)

	outstruct := new(struct {
		Eon       uint64
		Timestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Eon = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 timestamp)
func (_Shutterregistry *ShutterregistrySession) Registrations(identity [32]byte) (struct {
	Eon       uint64
	Timestamp uint64
}, error) {
	return _Shutterregistry.Contract.Registrations(&_Shutterregistry.CallOpts, identity)
}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 timestamp)
func (_Shutterregistry *ShutterregistryCallerSession) Registrations(identity [32]byte) (struct {
	Eon       uint64
	Timestamp uint64
}, error) {
	return _Shutterregistry.Contract.Registrations(&_Shutterregistry.CallOpts, identity)
}

// Register is a paid mutator transaction binding the contract method 0xeaac3573.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, uint64 timestamp) returns()
func (_Shutterregistry *ShutterregistryTransactor) Register(opts *bind.TransactOpts, eon uint64, identityPrefix [32]byte, timestamp uint64) (*types.Transaction, error) {
	return _Shutterregistry.contract.Transact(opts, "register", eon, identityPrefix, timestamp)
}

// Register is a paid mutator transaction binding the contract method 0xeaac3573.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, uint64 timestamp) returns()
func (_Shutterregistry *ShutterregistrySession) Register(eon uint64, identityPrefix [32]byte, timestamp uint64) (*types.Transaction, error) {
	return _Shutterregistry.Contract.Register(&_Shutterregistry.TransactOpts, eon, identityPrefix, timestamp)
}

// Register is a paid mutator transaction binding the contract method 0xeaac3573.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, uint64 timestamp) returns()
func (_Shutterregistry *ShutterregistryTransactorSession) Register(eon uint64, identityPrefix [32]byte, timestamp uint64) (*types.Transaction, error) {
	return _Shutterregistry.Contract.Register(&_Shutterregistry.TransactOpts, eon, identityPrefix, timestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shutterregistry *ShutterregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shutterregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shutterregistry *ShutterregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Shutterregistry.Contract.RenounceOwnership(&_Shutterregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shutterregistry *ShutterregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Shutterregistry.Contract.RenounceOwnership(&_Shutterregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shutterregistry *ShutterregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Shutterregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shutterregistry *ShutterregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shutterregistry.Contract.TransferOwnership(&_Shutterregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shutterregistry *ShutterregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shutterregistry.Contract.TransferOwnership(&_Shutterregistry.TransactOpts, newOwner)
}

// ShutterregistryIdentityRegisteredIterator is returned from FilterIdentityRegistered and is used to iterate over the raw logs and unpacked data for IdentityRegistered events raised by the Shutterregistry contract.
type ShutterregistryIdentityRegisteredIterator struct {
	Event *ShutterregistryIdentityRegistered // Event containing the contract specifics and raw log

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
func (it *ShutterregistryIdentityRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShutterregistryIdentityRegistered)
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
		it.Event = new(ShutterregistryIdentityRegistered)
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
func (it *ShutterregistryIdentityRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShutterregistryIdentityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShutterregistryIdentityRegistered represents a IdentityRegistered event raised by the Shutterregistry contract.
type ShutterregistryIdentityRegistered struct {
	Eon            uint64
	IdentityPrefix [32]byte
	Sender         common.Address
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIdentityRegistered is a free log retrieval operation binding the contract event 0xa254e5a8078a79959db8c7203cbd84a28d6e1b9750d0c4fd743a1a069c21b05b.
//
// Solidity: event IdentityRegistered(uint64 eon, bytes32 identityPrefix, address sender, uint64 timestamp)
func (_Shutterregistry *ShutterregistryFilterer) FilterIdentityRegistered(opts *bind.FilterOpts) (*ShutterregistryIdentityRegisteredIterator, error) {

	logs, sub, err := _Shutterregistry.contract.FilterLogs(opts, "IdentityRegistered")
	if err != nil {
		return nil, err
	}
	return &ShutterregistryIdentityRegisteredIterator{contract: _Shutterregistry.contract, event: "IdentityRegistered", logs: logs, sub: sub}, nil
}

// WatchIdentityRegistered is a free log subscription operation binding the contract event 0xa254e5a8078a79959db8c7203cbd84a28d6e1b9750d0c4fd743a1a069c21b05b.
//
// Solidity: event IdentityRegistered(uint64 eon, bytes32 identityPrefix, address sender, uint64 timestamp)
func (_Shutterregistry *ShutterregistryFilterer) WatchIdentityRegistered(opts *bind.WatchOpts, sink chan<- *ShutterregistryIdentityRegistered) (event.Subscription, error) {

	logs, sub, err := _Shutterregistry.contract.WatchLogs(opts, "IdentityRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShutterregistryIdentityRegistered)
				if err := _Shutterregistry.contract.UnpackLog(event, "IdentityRegistered", log); err != nil {
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

// ParseIdentityRegistered is a log parse operation binding the contract event 0xa254e5a8078a79959db8c7203cbd84a28d6e1b9750d0c4fd743a1a069c21b05b.
//
// Solidity: event IdentityRegistered(uint64 eon, bytes32 identityPrefix, address sender, uint64 timestamp)
func (_Shutterregistry *ShutterregistryFilterer) ParseIdentityRegistered(log types.Log) (*ShutterregistryIdentityRegistered, error) {
	event := new(ShutterregistryIdentityRegistered)
	if err := _Shutterregistry.contract.UnpackLog(event, "IdentityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShutterregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Shutterregistry contract.
type ShutterregistryOwnershipTransferredIterator struct {
	Event *ShutterregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShutterregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShutterregistryOwnershipTransferred)
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
		it.Event = new(ShutterregistryOwnershipTransferred)
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
func (it *ShutterregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShutterregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShutterregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Shutterregistry contract.
type ShutterregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shutterregistry *ShutterregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShutterregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shutterregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShutterregistryOwnershipTransferredIterator{contract: _Shutterregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shutterregistry *ShutterregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShutterregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shutterregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShutterregistryOwnershipTransferred)
				if err := _Shutterregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shutterregistry *ShutterregistryFilterer) ParseOwnershipTransferred(log types.Log) (*ShutterregistryOwnershipTransferred, error) {
	event := new(ShutterregistryOwnershipTransferred)
	if err := _Shutterregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
