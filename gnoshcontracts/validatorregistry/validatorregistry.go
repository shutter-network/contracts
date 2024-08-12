// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorregistry

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

// IValidatorRegistryUpdate is an auto generated low-level Go binding around an user-defined struct.
type IValidatorRegistryUpdate struct {
	Message   []byte
	Signature []byte
}

// ValidatorregistryMetaData contains all meta data concerning the Validatorregistry contract.
var ValidatorregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNumUpdates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpdate\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Update\",\"components\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Updated\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// ValidatorregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorregistryMetaData.ABI instead.
var ValidatorregistryABI = ValidatorregistryMetaData.ABI

// Validatorregistry is an auto generated Go binding around an Ethereum contract.
type Validatorregistry struct {
	ValidatorregistryCaller     // Read-only binding to the contract
	ValidatorregistryTransactor // Write-only binding to the contract
	ValidatorregistryFilterer   // Log filterer for contract events
}

// ValidatorregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorregistrySession struct {
	Contract     *Validatorregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorregistryCallerSession struct {
	Contract *ValidatorregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorregistryTransactorSession struct {
	Contract     *ValidatorregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorregistryRaw struct {
	Contract *Validatorregistry // Generic contract binding to access the raw methods on
}

// ValidatorregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorregistryCallerRaw struct {
	Contract *ValidatorregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactorRaw struct {
	Contract *ValidatorregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorregistry creates a new instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistry(address common.Address, backend bind.ContractBackend) (*Validatorregistry, error) {
	contract, err := bindValidatorregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorregistry{ValidatorregistryCaller: ValidatorregistryCaller{contract: contract}, ValidatorregistryTransactor: ValidatorregistryTransactor{contract: contract}, ValidatorregistryFilterer: ValidatorregistryFilterer{contract: contract}}, nil
}

// NewValidatorregistryCaller creates a new read-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorregistryCaller, error) {
	contract, err := bindValidatorregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryCaller{contract: contract}, nil
}

// NewValidatorregistryTransactor creates a new write-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorregistryTransactor, error) {
	contract, err := bindValidatorregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryTransactor{contract: contract}, nil
}

// NewValidatorregistryFilterer creates a new log filterer instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorregistryFilterer, error) {
	contract, err := bindValidatorregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryFilterer{contract: contract}, nil
}

// bindValidatorregistry binds a generic wrapper to an already deployed contract.
func bindValidatorregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.ValidatorregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transact(opts, method, params...)
}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) GetNumUpdates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getNumUpdates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) GetNumUpdates() (*big.Int, error) {
	return _Validatorregistry.Contract.GetNumUpdates(&_Validatorregistry.CallOpts)
}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) GetNumUpdates() (*big.Int, error) {
	return _Validatorregistry.Contract.GetNumUpdates(&_Validatorregistry.CallOpts)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistryCaller) GetUpdate(opts *bind.CallOpts, i *big.Int) (IValidatorRegistryUpdate, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getUpdate", i)

	if err != nil {
		return *new(IValidatorRegistryUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IValidatorRegistryUpdate)).(*IValidatorRegistryUpdate)

	return out0, err

}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistrySession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Validatorregistry.Contract.GetUpdate(&_Validatorregistry.CallOpts, i)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistryCallerSession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Validatorregistry.Contract.GetUpdate(&_Validatorregistry.CallOpts, i)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistryTransactor) Update(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "update", message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistrySession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Update(&_Validatorregistry.TransactOpts, message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Update(&_Validatorregistry.TransactOpts, message, signature)
}

// ValidatorregistryUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the Validatorregistry contract.
type ValidatorregistryUpdatedIterator struct {
	Event *ValidatorregistryUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryUpdated)
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
		it.Event = new(ValidatorregistryUpdated)
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
func (it *ValidatorregistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryUpdated represents a Updated event raised by the Validatorregistry contract.
type ValidatorregistryUpdated struct {
	Message   []byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) FilterUpdated(opts *bind.FilterOpts) (*ValidatorregistryUpdatedIterator, error) {

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryUpdatedIterator{contract: _Validatorregistry.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorregistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryUpdated)
				if err := _Validatorregistry.contract.UnpackLog(event, "Updated", log); err != nil {
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

// ParseUpdated is a log parse operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) ParseUpdated(log types.Log) (*ValidatorregistryUpdated, error) {
	event := new(ValidatorregistryUpdated)
	if err := _Validatorregistry.contract.UnpackLog(event, "Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
