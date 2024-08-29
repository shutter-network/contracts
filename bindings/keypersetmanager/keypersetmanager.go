// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keypersetmanager

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

// KeypersetmanagerMetaData contains all meta data concerning the Keypersetmanager contract.
var KeypersetmanagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initializer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addKeyperSet\",\"inputs\":[{\"name\":\"activationBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getKeyperSetActivationBlock\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperSetAddress\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperSetIndexByBlock\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumKeyperSets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"KeyperSetAdded\",\"inputs\":[{\"name\":\"activationBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"members\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeyperSetAdded\",\"inputs\":[{\"name\":\"activationSlot\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"keyperSetContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyHaveKeyperSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveKeyperSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedInitializer\",\"inputs\":[]}]",
}

// KeypersetmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use KeypersetmanagerMetaData.ABI instead.
var KeypersetmanagerABI = KeypersetmanagerMetaData.ABI

// Keypersetmanager is an auto generated Go binding around an Ethereum contract.
type Keypersetmanager struct {
	KeypersetmanagerCaller     // Read-only binding to the contract
	KeypersetmanagerTransactor // Write-only binding to the contract
	KeypersetmanagerFilterer   // Log filterer for contract events
}

// KeypersetmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeypersetmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeypersetmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeypersetmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeypersetmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeypersetmanagerSession struct {
	Contract     *Keypersetmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeypersetmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeypersetmanagerCallerSession struct {
	Contract *KeypersetmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KeypersetmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeypersetmanagerTransactorSession struct {
	Contract     *KeypersetmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KeypersetmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeypersetmanagerRaw struct {
	Contract *Keypersetmanager // Generic contract binding to access the raw methods on
}

// KeypersetmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeypersetmanagerCallerRaw struct {
	Contract *KeypersetmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// KeypersetmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeypersetmanagerTransactorRaw struct {
	Contract *KeypersetmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeypersetmanager creates a new instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanager(address common.Address, backend bind.ContractBackend) (*Keypersetmanager, error) {
	contract, err := bindKeypersetmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keypersetmanager{KeypersetmanagerCaller: KeypersetmanagerCaller{contract: contract}, KeypersetmanagerTransactor: KeypersetmanagerTransactor{contract: contract}, KeypersetmanagerFilterer: KeypersetmanagerFilterer{contract: contract}}, nil
}

// NewKeypersetmanagerCaller creates a new read-only instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerCaller(address common.Address, caller bind.ContractCaller) (*KeypersetmanagerCaller, error) {
	contract, err := bindKeypersetmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerCaller{contract: contract}, nil
}

// NewKeypersetmanagerTransactor creates a new write-only instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeypersetmanagerTransactor, error) {
	contract, err := bindKeypersetmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerTransactor{contract: contract}, nil
}

// NewKeypersetmanagerFilterer creates a new log filterer instance of Keypersetmanager, bound to a specific deployed contract.
func NewKeypersetmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeypersetmanagerFilterer, error) {
	contract, err := bindKeypersetmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerFilterer{contract: contract}, nil
}

// bindKeypersetmanager binds a generic wrapper to an already deployed contract.
func bindKeypersetmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeypersetmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keypersetmanager *KeypersetmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keypersetmanager.Contract.KeypersetmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keypersetmanager *KeypersetmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.KeypersetmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keypersetmanager *KeypersetmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.KeypersetmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keypersetmanager *KeypersetmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keypersetmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keypersetmanager *KeypersetmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keypersetmanager *KeypersetmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.DEFAULTADMINROLE(&_Keypersetmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.DEFAULTADMINROLE(&_Keypersetmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) PAUSERROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.PAUSERROLE(&_Keypersetmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Keypersetmanager.Contract.PAUSERROLE(&_Keypersetmanager.CallOpts)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetActivationBlock(opts *bind.CallOpts, index uint64) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetActivationBlock", index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetActivationBlock(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetActivationBlock is a free data retrieval call binding the contract method 0x636df979.
//
// Solidity: function getKeyperSetActivationBlock(uint64 index) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetActivationBlock(index uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetActivationBlock(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetAddress(opts *bind.CallOpts, index uint64) (common.Address, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetAddress", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _Keypersetmanager.Contract.GetKeyperSetAddress(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetAddress is a free data retrieval call binding the contract method 0xf90f3bed.
//
// Solidity: function getKeyperSetAddress(uint64 index) view returns(address)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetAddress(index uint64) (common.Address, error) {
	return _Keypersetmanager.Contract.GetKeyperSetAddress(&_Keypersetmanager.CallOpts, index)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetKeyperSetIndexByBlock(opts *bind.CallOpts, blockNumber uint64) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getKeyperSetIndexByBlock", blockNumber)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetIndexByBlock(&_Keypersetmanager.CallOpts, blockNumber)
}

// GetKeyperSetIndexByBlock is a free data retrieval call binding the contract method 0x035cef15.
//
// Solidity: function getKeyperSetIndexByBlock(uint64 blockNumber) view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetKeyperSetIndexByBlock(blockNumber uint64) (uint64, error) {
	return _Keypersetmanager.Contract.GetKeyperSetIndexByBlock(&_Keypersetmanager.CallOpts, blockNumber)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCaller) GetNumKeyperSets(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getNumKeyperSets")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerSession) GetNumKeyperSets() (uint64, error) {
	return _Keypersetmanager.Contract.GetNumKeyperSets(&_Keypersetmanager.CallOpts)
}

// GetNumKeyperSets is a free data retrieval call binding the contract method 0xf2e6100a.
//
// Solidity: function getNumKeyperSets() view returns(uint64)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetNumKeyperSets() (uint64, error) {
	return _Keypersetmanager.Contract.GetNumKeyperSets(&_Keypersetmanager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Keypersetmanager.Contract.GetRoleAdmin(&_Keypersetmanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Keypersetmanager *KeypersetmanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Keypersetmanager.Contract.GetRoleAdmin(&_Keypersetmanager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Keypersetmanager.Contract.HasRole(&_Keypersetmanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Keypersetmanager.Contract.HasRole(&_Keypersetmanager.CallOpts, role, account)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerCaller) Initializer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "initializer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerSession) Initializer() (common.Address, error) {
	return _Keypersetmanager.Contract.Initializer(&_Keypersetmanager.CallOpts)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Keypersetmanager *KeypersetmanagerCallerSession) Initializer() (common.Address, error) {
	return _Keypersetmanager.Contract.Initializer(&_Keypersetmanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) Paused() (bool, error) {
	return _Keypersetmanager.Contract.Paused(&_Keypersetmanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) Paused() (bool, error) {
	return _Keypersetmanager.Contract.Paused(&_Keypersetmanager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Keypersetmanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Keypersetmanager.Contract.SupportsInterface(&_Keypersetmanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Keypersetmanager *KeypersetmanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Keypersetmanager.Contract.SupportsInterface(&_Keypersetmanager.CallOpts, interfaceId)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) AddKeyperSet(opts *bind.TransactOpts, activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "addKeyperSet", activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.AddKeyperSet(&_Keypersetmanager.TransactOpts, activationBlock, keyperSetContract)
}

// AddKeyperSet is a paid mutator transaction binding the contract method 0xd3877c43.
//
// Solidity: function addKeyperSet(uint64 activationBlock, address keyperSetContract) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) AddKeyperSet(activationBlock uint64, keyperSetContract common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.AddKeyperSet(&_Keypersetmanager.TransactOpts, activationBlock, keyperSetContract)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.GrantRole(&_Keypersetmanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.GrantRole(&_Keypersetmanager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "initialize", admin, pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerSession) Initialize(admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Initialize(&_Keypersetmanager.TransactOpts, admin, pauser)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address pauser) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Initialize(admin common.Address, pauser common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Initialize(&_Keypersetmanager.TransactOpts, admin, pauser)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerSession) Pause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Pause(&_Keypersetmanager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Pause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Pause(&_Keypersetmanager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RenounceRole(&_Keypersetmanager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RenounceRole(&_Keypersetmanager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RevokeRole(&_Keypersetmanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Keypersetmanager.Contract.RevokeRole(&_Keypersetmanager.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keypersetmanager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerSession) Unpause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Unpause(&_Keypersetmanager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Keypersetmanager *KeypersetmanagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _Keypersetmanager.Contract.Unpause(&_Keypersetmanager.TransactOpts)
}

// KeypersetmanagerKeyperSetAddedIterator is returned from FilterKeyperSetAdded and is used to iterate over the raw logs and unpacked data for KeyperSetAdded events raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAddedIterator struct {
	Event *KeypersetmanagerKeyperSetAdded // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerKeyperSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerKeyperSetAdded)
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
		it.Event = new(KeypersetmanagerKeyperSetAdded)
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
func (it *KeypersetmanagerKeyperSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerKeyperSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerKeyperSetAdded represents a KeyperSetAdded event raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded struct {
	ActivationBlock   uint64
	KeyperSetContract common.Address
	Members           []common.Address
	Threshold         uint64
	Eon               uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterKeyperSetAdded is a free log retrieval operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterKeyperSetAdded(opts *bind.FilterOpts) (*KeypersetmanagerKeyperSetAddedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerKeyperSetAddedIterator{contract: _Keypersetmanager.contract, event: "KeyperSetAdded", logs: logs, sub: sub}, nil
}

// WatchKeyperSetAdded is a free log subscription operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchKeyperSetAdded(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerKeyperSetAdded) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "KeyperSetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerKeyperSetAdded)
				if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
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

// ParseKeyperSetAdded is a log parse operation binding the contract event 0xa940387dac06ebd336730f1d14b21629a9d137069a9137e871f95313e1010165.
//
// Solidity: event KeyperSetAdded(uint64 activationBlock, address keyperSetContract, address[] members, uint64 threshold, uint64 eon)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseKeyperSetAdded(log types.Log) (*KeypersetmanagerKeyperSetAdded, error) {
	event := new(KeypersetmanagerKeyperSetAdded)
	if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerKeyperSetAdded0Iterator is returned from FilterKeyperSetAdded0 and is used to iterate over the raw logs and unpacked data for KeyperSetAdded0 events raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded0Iterator struct {
	Event *KeypersetmanagerKeyperSetAdded0 // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerKeyperSetAdded0)
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
		it.Event = new(KeypersetmanagerKeyperSetAdded0)
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
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerKeyperSetAdded0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerKeyperSetAdded0 represents a KeyperSetAdded0 event raised by the Keypersetmanager contract.
type KeypersetmanagerKeyperSetAdded0 struct {
	ActivationSlot    uint64
	KeyperSetContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterKeyperSetAdded0 is a free log retrieval operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterKeyperSetAdded0(opts *bind.FilterOpts) (*KeypersetmanagerKeyperSetAdded0Iterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "KeyperSetAdded0")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerKeyperSetAdded0Iterator{contract: _Keypersetmanager.contract, event: "KeyperSetAdded0", logs: logs, sub: sub}, nil
}

// WatchKeyperSetAdded0 is a free log subscription operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchKeyperSetAdded0(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerKeyperSetAdded0) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "KeyperSetAdded0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerKeyperSetAdded0)
				if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded0", log); err != nil {
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

// ParseKeyperSetAdded0 is a log parse operation binding the contract event 0x6605cb866297050f9f49ae7e0b38e0e4c8178d4b176e24332bc01672818d707b.
//
// Solidity: event KeyperSetAdded(uint64 activationSlot, address keyperSetContract)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseKeyperSetAdded0(log types.Log) (*KeypersetmanagerKeyperSetAdded0, error) {
	event := new(KeypersetmanagerKeyperSetAdded0)
	if err := _Keypersetmanager.contract.UnpackLog(event, "KeyperSetAdded0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Keypersetmanager contract.
type KeypersetmanagerPausedIterator struct {
	Event *KeypersetmanagerPaused // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerPaused)
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
		it.Event = new(KeypersetmanagerPaused)
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
func (it *KeypersetmanagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerPaused represents a Paused event raised by the Keypersetmanager contract.
type KeypersetmanagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterPaused(opts *bind.FilterOpts) (*KeypersetmanagerPausedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerPausedIterator{contract: _Keypersetmanager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerPaused) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerPaused)
				if err := _Keypersetmanager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParsePaused(log types.Log) (*KeypersetmanagerPaused, error) {
	event := new(KeypersetmanagerPaused)
	if err := _Keypersetmanager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleAdminChangedIterator struct {
	Event *KeypersetmanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleAdminChanged)
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
		it.Event = new(KeypersetmanagerRoleAdminChanged)
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
func (it *KeypersetmanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KeypersetmanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleAdminChangedIterator{contract: _Keypersetmanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleAdminChanged)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleAdminChanged(log types.Log) (*KeypersetmanagerRoleAdminChanged, error) {
	event := new(KeypersetmanagerRoleAdminChanged)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleGrantedIterator struct {
	Event *KeypersetmanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleGranted)
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
		it.Event = new(KeypersetmanagerRoleGranted)
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
func (it *KeypersetmanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleGranted represents a RoleGranted event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeypersetmanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleGrantedIterator{contract: _Keypersetmanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleGranted)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleGranted(log types.Log) (*KeypersetmanagerRoleGranted, error) {
	event := new(KeypersetmanagerRoleGranted)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Keypersetmanager contract.
type KeypersetmanagerRoleRevokedIterator struct {
	Event *KeypersetmanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerRoleRevoked)
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
		it.Event = new(KeypersetmanagerRoleRevoked)
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
func (it *KeypersetmanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerRoleRevoked represents a RoleRevoked event raised by the Keypersetmanager contract.
type KeypersetmanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KeypersetmanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerRoleRevokedIterator{contract: _Keypersetmanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerRoleRevoked)
				if err := _Keypersetmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseRoleRevoked(log types.Log) (*KeypersetmanagerRoleRevoked, error) {
	event := new(KeypersetmanagerRoleRevoked)
	if err := _Keypersetmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeypersetmanagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Keypersetmanager contract.
type KeypersetmanagerUnpausedIterator struct {
	Event *KeypersetmanagerUnpaused // Event containing the contract specifics and raw log

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
func (it *KeypersetmanagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeypersetmanagerUnpaused)
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
		it.Event = new(KeypersetmanagerUnpaused)
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
func (it *KeypersetmanagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeypersetmanagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeypersetmanagerUnpaused represents a Unpaused event raised by the Keypersetmanager contract.
type KeypersetmanagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KeypersetmanagerUnpausedIterator, error) {

	logs, sub, err := _Keypersetmanager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KeypersetmanagerUnpausedIterator{contract: _Keypersetmanager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KeypersetmanagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Keypersetmanager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeypersetmanagerUnpaused)
				if err := _Keypersetmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Keypersetmanager *KeypersetmanagerFilterer) ParseUnpaused(log types.Log) (*KeypersetmanagerUnpaused, error) {
	event := new(KeypersetmanagerUnpaused)
	if err := _Keypersetmanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
