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

// KeyperSetMetaData contains all meta data concerning the KeyperSet contract.
var KeyperSetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMembers\",\"inputs\":[{\"name\":\"newMembers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMember\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublisher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedToBroadcastEonKey\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublisher\",\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// KeyperSetABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyperSetMetaData.ABI instead.
var KeyperSetABI = KeyperSetMetaData.ABI

// KeyperSet is an auto generated Go binding around an Ethereum contract.
type KeyperSet struct {
	KeyperSetCaller     // Read-only binding to the contract
	KeyperSetTransactor // Write-only binding to the contract
	KeyperSetFilterer   // Log filterer for contract events
}

// KeyperSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyperSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyperSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyperSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyperSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyperSetSession struct {
	Contract     *KeyperSet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyperSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyperSetCallerSession struct {
	Contract *KeyperSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KeyperSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyperSetTransactorSession struct {
	Contract     *KeyperSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KeyperSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyperSetRaw struct {
	Contract *KeyperSet // Generic contract binding to access the raw methods on
}

// KeyperSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyperSetCallerRaw struct {
	Contract *KeyperSetCaller // Generic read-only contract binding to access the raw methods on
}

// KeyperSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyperSetTransactorRaw struct {
	Contract *KeyperSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyperSet creates a new instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSet(address common.Address, backend bind.ContractBackend) (*KeyperSet, error) {
	contract, err := bindKeyperSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyperSet{KeyperSetCaller: KeyperSetCaller{contract: contract}, KeyperSetTransactor: KeyperSetTransactor{contract: contract}, KeyperSetFilterer: KeyperSetFilterer{contract: contract}}, nil
}

// NewKeyperSetCaller creates a new read-only instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetCaller(address common.Address, caller bind.ContractCaller) (*KeyperSetCaller, error) {
	contract, err := bindKeyperSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetCaller{contract: contract}, nil
}

// NewKeyperSetTransactor creates a new write-only instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyperSetTransactor, error) {
	contract, err := bindKeyperSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyperSetTransactor{contract: contract}, nil
}

// NewKeyperSetFilterer creates a new log filterer instance of KeyperSet, bound to a specific deployed contract.
func NewKeyperSetFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyperSetFilterer, error) {
	contract, err := bindKeyperSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyperSetFilterer{contract: contract}, nil
}

// bindKeyperSet binds a generic wrapper to an already deployed contract.
func bindKeyperSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyperSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSet *KeyperSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSet.Contract.KeyperSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSet *KeyperSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.Contract.KeyperSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSet *KeyperSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSet.Contract.KeyperSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyperSet *KeyperSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyperSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyperSet *KeyperSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyperSet *KeyperSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyperSet.Contract.contract.Transact(opts, method, params...)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetCaller) GetMember(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getMember", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetSession) GetMember(index uint64) (common.Address, error) {
	return _KeyperSet.Contract.GetMember(&_KeyperSet.CallOpts, index)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_KeyperSet *KeyperSetCallerSession) GetMember(index uint64) (common.Address, error) {
	return _KeyperSet.Contract.GetMember(&_KeyperSet.CallOpts, index)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetSession) GetMembers() ([]common.Address, error) {
	return _KeyperSet.Contract.GetMembers(&_KeyperSet.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_KeyperSet *KeyperSetCallerSession) GetMembers() ([]common.Address, error) {
	return _KeyperSet.Contract.GetMembers(&_KeyperSet.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetCaller) GetNumMembers(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getNumMembers")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetSession) GetNumMembers() (uint64, error) {
	return _KeyperSet.Contract.GetNumMembers(&_KeyperSet.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_KeyperSet *KeyperSetCallerSession) GetNumMembers() (uint64, error) {
	return _KeyperSet.Contract.GetNumMembers(&_KeyperSet.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_KeyperSet *KeyperSetCaller) GetPublisher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getPublisher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_KeyperSet *KeyperSetSession) GetPublisher() (common.Address, error) {
	return _KeyperSet.Contract.GetPublisher(&_KeyperSet.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_KeyperSet *KeyperSetCallerSession) GetPublisher() (common.Address, error) {
	return _KeyperSet.Contract.GetPublisher(&_KeyperSet.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetCaller) GetThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetSession) GetThreshold() (uint64, error) {
	return _KeyperSet.Contract.GetThreshold(&_KeyperSet.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_KeyperSet *KeyperSetCallerSession) GetThreshold() (uint64, error) {
	return _KeyperSet.Contract.GetThreshold(&_KeyperSet.CallOpts)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetCaller) IsAllowedToBroadcastEonKey(opts *bind.CallOpts, a common.Address) (bool, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "isAllowedToBroadcastEonKey", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _KeyperSet.Contract.IsAllowedToBroadcastEonKey(&_KeyperSet.CallOpts, a)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_KeyperSet *KeyperSetCallerSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _KeyperSet.Contract.IsAllowedToBroadcastEonKey(&_KeyperSet.CallOpts, a)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetSession) IsFinalized() (bool, error) {
	return _KeyperSet.Contract.IsFinalized(&_KeyperSet.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_KeyperSet *KeyperSetCallerSession) IsFinalized() (bool, error) {
	return _KeyperSet.Contract.IsFinalized(&_KeyperSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyperSet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetSession) Owner() (common.Address, error) {
	return _KeyperSet.Contract.Owner(&_KeyperSet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyperSet *KeyperSetCallerSession) Owner() (common.Address, error) {
	return _KeyperSet.Contract.Owner(&_KeyperSet.CallOpts)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetTransactor) AddMembers(opts *bind.TransactOpts, newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "addMembers", newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.AddMembers(&_KeyperSet.TransactOpts, newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_KeyperSet *KeyperSetTransactorSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.AddMembers(&_KeyperSet.TransactOpts, newMembers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyperSet.Contract.RenounceOwnership(&_KeyperSet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KeyperSet *KeyperSetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KeyperSet.Contract.RenounceOwnership(&_KeyperSet.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetTransactor) SetFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setFinalized")
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetSession) SetFinalized() (*types.Transaction, error) {
	return _KeyperSet.Contract.SetFinalized(&_KeyperSet.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_KeyperSet *KeyperSetTransactorSession) SetFinalized() (*types.Transaction, error) {
	return _KeyperSet.Contract.SetFinalized(&_KeyperSet.TransactOpts)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_KeyperSet *KeyperSetTransactor) SetPublisher(opts *bind.TransactOpts, _publisher common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setPublisher", _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_KeyperSet *KeyperSetSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetPublisher(&_KeyperSet.TransactOpts, _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_KeyperSet *KeyperSetTransactorSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetPublisher(&_KeyperSet.TransactOpts, _publisher)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetTransactor) SetThreshold(opts *bind.TransactOpts, _threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetThreshold(&_KeyperSet.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_KeyperSet *KeyperSetTransactorSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _KeyperSet.Contract.SetThreshold(&_KeyperSet.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.TransferOwnership(&_KeyperSet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyperSet *KeyperSetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyperSet.Contract.TransferOwnership(&_KeyperSet.TransactOpts, newOwner)
}

// KeyperSetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyperSet contract.
type KeyperSetOwnershipTransferredIterator struct {
	Event *KeyperSetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyperSetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyperSetOwnershipTransferred)
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
		it.Event = new(KeyperSetOwnershipTransferred)
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
func (it *KeyperSetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyperSetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyperSetOwnershipTransferred represents a OwnershipTransferred event raised by the KeyperSet contract.
type KeyperSetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyperSet *KeyperSetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyperSetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyperSet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyperSetOwnershipTransferredIterator{contract: _KeyperSet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyperSet *KeyperSetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyperSetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyperSet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyperSetOwnershipTransferred)
				if err := _KeyperSet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeyperSet *KeyperSetFilterer) ParseOwnershipTransferred(log types.Log) (*KeyperSetOwnershipTransferred, error) {
	event := new(KeyperSetOwnershipTransferred)
	if err := _KeyperSet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
