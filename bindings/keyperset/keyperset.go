// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keyperset

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

// KeypersetMetaData contains all meta data concerning the Keyperset contract.
var KeypersetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMembers\",\"inputs\":[{\"name\":\"newMembers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMember\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublisher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedToBroadcastEonKey\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublisher\",\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// KeypersetABI is the input ABI used to generate the binding from.
// Deprecated: Use KeypersetMetaData.ABI instead.
var KeypersetABI = KeypersetMetaData.ABI

// Keyperset is an auto generated Go binding around an Ethereum contract.
type Keyperset struct {
	KeypersetCaller     // Read-only binding to the contract
	KeypersetTransactor // Write-only binding to the contract
	KeypersetFilterer   // Log filterer for contract events
}

// KeypersetCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeypersetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeypersetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeypersetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeypersetSession struct {
	Contract     *Keyperset        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeypersetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeypersetCallerSession struct {
	Contract *KeypersetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KeypersetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeypersetTransactorSession struct {
	Contract     *KeypersetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KeypersetRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeypersetRaw struct {
	Contract *Keyperset // Generic contract binding to access the raw methods on
}

// KeypersetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeypersetCallerRaw struct {
	Contract *KeypersetCaller // Generic read-only contract binding to access the raw methods on
}

// KeypersetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeypersetTransactorRaw struct {
	Contract *KeypersetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyperset creates a new instance of Keyperset, bound to a specific deployed contract.
func NewKeyperset(address common.Address, backend bind.ContractBackend) (*Keyperset, error) {
	contract, err := bindKeyperset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keyperset{KeypersetCaller: KeypersetCaller{contract: contract}, KeypersetTransactor: KeypersetTransactor{contract: contract}, KeypersetFilterer: KeypersetFilterer{contract: contract}}, nil
}

// NewKeypersetCaller creates a new read-only instance of Keyperset, bound to a specific deployed contract.
func NewKeypersetCaller(address common.Address, caller bind.ContractCaller) (*KeypersetCaller, error) {
	contract, err := bindKeyperset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetCaller{contract: contract}, nil
}

// NewKeypersetTransactor creates a new write-only instance of Keyperset, bound to a specific deployed contract.
func NewKeypersetTransactor(address common.Address, transactor bind.ContractTransactor) (*KeypersetTransactor, error) {
	contract, err := bindKeyperset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetTransactor{contract: contract}, nil
}

// NewKeypersetFilterer creates a new log filterer instance of Keyperset, bound to a specific deployed contract.
func NewKeypersetFilterer(address common.Address, filterer bind.ContractFilterer) (*KeypersetFilterer, error) {
	contract, err := bindKeyperset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeypersetFilterer{contract: contract}, nil
}

// bindKeyperset binds a generic wrapper to an already deployed contract.
func bindKeyperset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeypersetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keyperset *KeypersetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keyperset.Contract.KeypersetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keyperset *KeypersetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyperset.Contract.KeypersetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keyperset *KeypersetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keyperset.Contract.KeypersetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keyperset *KeypersetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keyperset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keyperset *KeypersetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyperset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keyperset *KeypersetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keyperset.Contract.contract.Transact(opts, method, params...)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Keyperset *KeypersetCaller) GetMember(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getMember", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Keyperset *KeypersetSession) GetMember(index uint64) (common.Address, error) {
	return _Keyperset.Contract.GetMember(&_Keyperset.CallOpts, index)
}

// GetMember is a free data retrieval call binding the contract method 0x2e8e6cad.
//
// Solidity: function getMember(uint64 index) view returns(address)
func (_Keyperset *KeypersetCallerSession) GetMember(index uint64) (common.Address, error) {
	return _Keyperset.Contract.GetMember(&_Keyperset.CallOpts, index)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Keyperset *KeypersetCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Keyperset *KeypersetSession) GetMembers() ([]common.Address, error) {
	return _Keyperset.Contract.GetMembers(&_Keyperset.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_Keyperset *KeypersetCallerSession) GetMembers() ([]common.Address, error) {
	return _Keyperset.Contract.GetMembers(&_Keyperset.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Keyperset *KeypersetCaller) GetNumMembers(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getNumMembers")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Keyperset *KeypersetSession) GetNumMembers() (uint64, error) {
	return _Keyperset.Contract.GetNumMembers(&_Keyperset.CallOpts)
}

// GetNumMembers is a free data retrieval call binding the contract method 0x17d5430a.
//
// Solidity: function getNumMembers() view returns(uint64)
func (_Keyperset *KeypersetCallerSession) GetNumMembers() (uint64, error) {
	return _Keyperset.Contract.GetNumMembers(&_Keyperset.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Keyperset *KeypersetCaller) GetPublisher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getPublisher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Keyperset *KeypersetSession) GetPublisher() (common.Address, error) {
	return _Keyperset.Contract.GetPublisher(&_Keyperset.CallOpts)
}

// GetPublisher is a free data retrieval call binding the contract method 0xdbf4ab4e.
//
// Solidity: function getPublisher() view returns(address)
func (_Keyperset *KeypersetCallerSession) GetPublisher() (common.Address, error) {
	return _Keyperset.Contract.GetPublisher(&_Keyperset.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Keyperset *KeypersetCaller) GetThreshold(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Keyperset *KeypersetSession) GetThreshold() (uint64, error) {
	return _Keyperset.Contract.GetThreshold(&_Keyperset.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint64)
func (_Keyperset *KeypersetCallerSession) GetThreshold() (uint64, error) {
	return _Keyperset.Contract.GetThreshold(&_Keyperset.CallOpts)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Keyperset *KeypersetCaller) IsAllowedToBroadcastEonKey(opts *bind.CallOpts, a common.Address) (bool, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "isAllowedToBroadcastEonKey", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Keyperset *KeypersetSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _Keyperset.Contract.IsAllowedToBroadcastEonKey(&_Keyperset.CallOpts, a)
}

// IsAllowedToBroadcastEonKey is a free data retrieval call binding the contract method 0xcde1532d.
//
// Solidity: function isAllowedToBroadcastEonKey(address a) view returns(bool)
func (_Keyperset *KeypersetCallerSession) IsAllowedToBroadcastEonKey(a common.Address) (bool, error) {
	return _Keyperset.Contract.IsAllowedToBroadcastEonKey(&_Keyperset.CallOpts, a)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Keyperset *KeypersetCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Keyperset *KeypersetSession) IsFinalized() (bool, error) {
	return _Keyperset.Contract.IsFinalized(&_Keyperset.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Keyperset *KeypersetCallerSession) IsFinalized() (bool, error) {
	return _Keyperset.Contract.IsFinalized(&_Keyperset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Keyperset *KeypersetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Keyperset *KeypersetSession) Owner() (common.Address, error) {
	return _Keyperset.Contract.Owner(&_Keyperset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Keyperset *KeypersetCallerSession) Owner() (common.Address, error) {
	return _Keyperset.Contract.Owner(&_Keyperset.CallOpts)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Keyperset *KeypersetTransactor) AddMembers(opts *bind.TransactOpts, newMembers []common.Address) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "addMembers", newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Keyperset *KeypersetSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.AddMembers(&_Keyperset.TransactOpts, newMembers)
}

// AddMembers is a paid mutator transaction binding the contract method 0x6f4d469b.
//
// Solidity: function addMembers(address[] newMembers) returns()
func (_Keyperset *KeypersetTransactorSession) AddMembers(newMembers []common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.AddMembers(&_Keyperset.TransactOpts, newMembers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Keyperset *KeypersetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Keyperset *KeypersetSession) RenounceOwnership() (*types.Transaction, error) {
	return _Keyperset.Contract.RenounceOwnership(&_Keyperset.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Keyperset *KeypersetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Keyperset.Contract.RenounceOwnership(&_Keyperset.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Keyperset *KeypersetTransactor) SetFinalized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "setFinalized")
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Keyperset *KeypersetSession) SetFinalized() (*types.Transaction, error) {
	return _Keyperset.Contract.SetFinalized(&_Keyperset.TransactOpts)
}

// SetFinalized is a paid mutator transaction binding the contract method 0x1de77253.
//
// Solidity: function setFinalized() returns()
func (_Keyperset *KeypersetTransactorSession) SetFinalized() (*types.Transaction, error) {
	return _Keyperset.Contract.SetFinalized(&_Keyperset.TransactOpts)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Keyperset *KeypersetTransactor) SetPublisher(opts *bind.TransactOpts, _publisher common.Address) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "setPublisher", _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Keyperset *KeypersetSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.SetPublisher(&_Keyperset.TransactOpts, _publisher)
}

// SetPublisher is a paid mutator transaction binding the contract method 0xcab63661.
//
// Solidity: function setPublisher(address _publisher) returns()
func (_Keyperset *KeypersetTransactorSession) SetPublisher(_publisher common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.SetPublisher(&_Keyperset.TransactOpts, _publisher)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Keyperset *KeypersetTransactor) SetThreshold(opts *bind.TransactOpts, _threshold uint64) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Keyperset *KeypersetSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _Keyperset.Contract.SetThreshold(&_Keyperset.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x17c4de35.
//
// Solidity: function setThreshold(uint64 _threshold) returns()
func (_Keyperset *KeypersetTransactorSession) SetThreshold(_threshold uint64) (*types.Transaction, error) {
	return _Keyperset.Contract.SetThreshold(&_Keyperset.TransactOpts, _threshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Keyperset *KeypersetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Keyperset *KeypersetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.TransferOwnership(&_Keyperset.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Keyperset *KeypersetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.TransferOwnership(&_Keyperset.TransactOpts, newOwner)
}

// KeypersetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Keyperset contract.
type KeypersetOwnershipTransferredIterator struct {
	Event *KeypersetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeypersetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetOwnershipTransferred)
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
		it.Event = new(KeypersetOwnershipTransferred)
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
func (it *KeypersetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetOwnershipTransferred represents a OwnershipTransferred event raised by the Keyperset contract.
type KeypersetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Keyperset *KeypersetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeypersetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Keyperset.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetOwnershipTransferredIterator{contract: _Keyperset.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Keyperset *KeypersetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeypersetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Keyperset.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetOwnershipTransferred)
				if err := _Keyperset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Keyperset *KeypersetFilterer) ParseOwnershipTransferred(log types.Log) (*KeypersetOwnershipTransferred, error) {
	event := new(KeypersetOwnershipTransferred)
	if err := _Keyperset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
