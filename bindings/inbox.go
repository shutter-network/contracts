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

// InboxTransaction is an auto generated low-level Go binding around an user-defined struct.
type InboxTransaction struct {
	EncryptedTransaction []byte
	Sender               common.Address
	GasLimit             uint64
	CumulativeGasLimit   uint64
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_blockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initializer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCK_GAS_LIMIT_SETTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clear\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBlockGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransactions\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInbox.Transaction[]\",\"components\":[{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"dao\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockGasLimit\",\"inputs\":[{\"name\":\"newBlockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptedTransaction\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EncryptedTransactionSubmitted\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"block\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockGasLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferEtherFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedInitializer\",\"inputs\":[]}]",
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

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) BLOCKGASLIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "BLOCK_GAS_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Bindings.Contract.BLOCKGASLIMITSETTERROLE(&_Bindings.CallOpts)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Bindings.Contract.BLOCKGASLIMITSETTERROLE(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bindings.Contract.DEFAULTADMINROLE(&_Bindings.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) PAUSERROLE() ([32]byte, error) {
	return _Bindings.Contract.PAUSERROLE(&_Bindings.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Bindings.Contract.PAUSERROLE(&_Bindings.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) SEQUENCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "SEQUENCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) SEQUENCERROLE() ([32]byte, error) {
	return _Bindings.Contract.SEQUENCERROLE(&_Bindings.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) SEQUENCERROLE() ([32]byte, error) {
	return _Bindings.Contract.SEQUENCERROLE(&_Bindings.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Bindings *BindingsCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Bindings *BindingsSession) WITHDRAWROLE() ([32]byte, error) {
	return _Bindings.Contract.WITHDRAWROLE(&_Bindings.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Bindings *BindingsCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Bindings.Contract.WITHDRAWROLE(&_Bindings.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Bindings *BindingsCaller) GetBlockGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBlockGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Bindings *BindingsSession) GetBlockGasLimit() (uint64, error) {
	return _Bindings.Contract.GetBlockGasLimit(&_Bindings.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Bindings *BindingsCallerSession) GetBlockGasLimit() (uint64, error) {
	return _Bindings.Contract.GetBlockGasLimit(&_Bindings.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bindings *BindingsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bindings.Contract.GetRoleAdmin(&_Bindings.CallOpts, role)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Bindings *BindingsCaller) GetTransactions(opts *bind.CallOpts, blockNumber uint64) ([]InboxTransaction, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getTransactions", blockNumber)

	if err != nil {
		return *new([]InboxTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]InboxTransaction)).(*[]InboxTransaction)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Bindings *BindingsSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Bindings.Contract.GetTransactions(&_Bindings.CallOpts, blockNumber)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Bindings *BindingsCallerSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Bindings.Contract.GetTransactions(&_Bindings.CallOpts, blockNumber)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bindings *BindingsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bindings.Contract.HasRole(&_Bindings.CallOpts, role, account)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Bindings *BindingsCaller) Initializer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "initializer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Bindings *BindingsSession) Initializer() (common.Address, error) {
	return _Bindings.Contract.Initializer(&_Bindings.CallOpts)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Bindings *BindingsCallerSession) Initializer() (common.Address, error) {
	return _Bindings.Contract.Initializer(&_Bindings.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bindings *BindingsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bindings *BindingsSession) Paused() (bool, error) {
	return _Bindings.Contract.Paused(&_Bindings.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bindings *BindingsCallerSession) Paused() (bool, error) {
	return _Bindings.Contract.Paused(&_Bindings.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bindings *BindingsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bindings.Contract.SupportsInterface(&_Bindings.CallOpts, interfaceId)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Bindings *BindingsTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Bindings *BindingsSession) Clear() (*types.Transaction, error) {
	return _Bindings.Contract.Clear(&_Bindings.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Bindings *BindingsTransactorSession) Clear() (*types.Transaction, error) {
	return _Bindings.Contract.Clear(&_Bindings.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.GrantRole(&_Bindings.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Bindings *BindingsTransactor) Initialize(opts *bind.TransactOpts, dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "initialize", dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Bindings *BindingsSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Bindings *BindingsTransactorSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Initialize(&_Bindings.TransactOpts, dao, sequencer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bindings *BindingsTransactorSession) Pause() (*types.Transaction, error) {
	return _Bindings.Contract.Pause(&_Bindings.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bindings *BindingsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RenounceRole(&_Bindings.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bindings *BindingsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.RevokeRole(&_Bindings.TransactOpts, role, account)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Bindings *BindingsTransactor) SetBlockGasLimit(opts *bind.TransactOpts, newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setBlockGasLimit", newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Bindings *BindingsSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Bindings.Contract.SetBlockGasLimit(&_Bindings.TransactOpts, newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Bindings *BindingsTransactorSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Bindings.Contract.SetBlockGasLimit(&_Bindings.TransactOpts, newBlockGasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Bindings *BindingsTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitEncryptedTransaction", blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Bindings *BindingsSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitEncryptedTransaction(&_Bindings.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Bindings *BindingsTransactorSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitEncryptedTransaction(&_Bindings.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsSession) Unpause() (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bindings *BindingsTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bindings.Contract.Unpause(&_Bindings.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Bindings *BindingsTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Bindings *BindingsSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Bindings *BindingsTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.Withdraw(&_Bindings.TransactOpts, recipient)
}

// BindingsEncryptedTransactionSubmittedIterator is returned from FilterEncryptedTransactionSubmitted and is used to iterate over the raw logs and unpacked data for EncryptedTransactionSubmitted events raised by the Bindings contract.
type BindingsEncryptedTransactionSubmittedIterator struct {
	Event *BindingsEncryptedTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingsEncryptedTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsEncryptedTransactionSubmitted)
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
		it.Event = new(BindingsEncryptedTransactionSubmitted)
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
func (it *BindingsEncryptedTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsEncryptedTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsEncryptedTransactionSubmitted represents a EncryptedTransactionSubmitted event raised by the Bindings contract.
type BindingsEncryptedTransactionSubmitted struct {
	Index                uint64
	Block                uint64
	EncryptedTransaction []byte
	Sender               common.Address
	GasLimit             uint64
	CumulativeGasLimit   uint64
	Fee                  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterEncryptedTransactionSubmitted is a free log retrieval operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Bindings *BindingsFilterer) FilterEncryptedTransactionSubmitted(opts *bind.FilterOpts, index []uint64, block []uint64) (*BindingsEncryptedTransactionSubmittedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return &BindingsEncryptedTransactionSubmittedIterator{contract: _Bindings.contract, event: "EncryptedTransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchEncryptedTransactionSubmitted is a free log subscription operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Bindings *BindingsFilterer) WatchEncryptedTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsEncryptedTransactionSubmitted, index []uint64, block []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsEncryptedTransactionSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
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

// ParseEncryptedTransactionSubmitted is a log parse operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Bindings *BindingsFilterer) ParseEncryptedTransactionSubmitted(log types.Log) (*BindingsEncryptedTransactionSubmitted, error) {
	event := new(BindingsEncryptedTransactionSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the Bindings contract.
type BindingsFeesWithdrawnIterator struct {
	Event *BindingsFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *BindingsFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsFeesWithdrawn)
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
		it.Event = new(BindingsFeesWithdrawn)
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
func (it *BindingsFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsFeesWithdrawn represents a FeesWithdrawn event raised by the Bindings contract.
type BindingsFeesWithdrawn struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Bindings *BindingsFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*BindingsFeesWithdrawnIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &BindingsFeesWithdrawnIterator{contract: _Bindings.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Bindings *BindingsFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *BindingsFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsFeesWithdrawn)
				if err := _Bindings.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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

// ParseFeesWithdrawn is a log parse operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Bindings *BindingsFilterer) ParseFeesWithdrawn(log types.Log) (*BindingsFeesWithdrawn, error) {
	event := new(BindingsFeesWithdrawn)
	if err := _Bindings.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bindings contract.
type BindingsPausedIterator struct {
	Event *BindingsPaused // Event containing the contract specifics and raw log

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
func (it *BindingsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsPaused)
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
		it.Event = new(BindingsPaused)
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
func (it *BindingsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsPaused represents a Paused event raised by the Bindings contract.
type BindingsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bindings *BindingsFilterer) FilterPaused(opts *bind.FilterOpts) (*BindingsPausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BindingsPausedIterator{contract: _Bindings.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bindings *BindingsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingsPaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsPaused)
				if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParsePaused(log types.Log) (*BindingsPaused, error) {
	event := new(BindingsPaused)
	if err := _Bindings.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bindings contract.
type BindingsRoleAdminChangedIterator struct {
	Event *BindingsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BindingsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleAdminChanged)
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
		it.Event = new(BindingsRoleAdminChanged)
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
func (it *BindingsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleAdminChanged represents a RoleAdminChanged event raised by the Bindings contract.
type BindingsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BindingsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleAdminChangedIterator{contract: _Bindings.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bindings *BindingsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BindingsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleAdminChanged)
				if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleAdminChanged(log types.Log) (*BindingsRoleAdminChanged, error) {
	event := new(BindingsRoleAdminChanged)
	if err := _Bindings.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bindings contract.
type BindingsRoleGrantedIterator struct {
	Event *BindingsRoleGranted // Event containing the contract specifics and raw log

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
func (it *BindingsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleGranted)
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
		it.Event = new(BindingsRoleGranted)
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
func (it *BindingsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleGranted represents a RoleGranted event raised by the Bindings contract.
type BindingsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleGrantedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleGrantedIterator{contract: _Bindings.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BindingsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleGranted)
				if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleGranted(log types.Log) (*BindingsRoleGranted, error) {
	event := new(BindingsRoleGranted)
	if err := _Bindings.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bindings contract.
type BindingsRoleRevokedIterator struct {
	Event *BindingsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BindingsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRoleRevoked)
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
		it.Event = new(BindingsRoleRevoked)
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
func (it *BindingsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRoleRevoked represents a RoleRevoked event raised by the Bindings contract.
type BindingsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BindingsRoleRevokedIterator, error) {

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

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRoleRevokedIterator{contract: _Bindings.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bindings *BindingsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BindingsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRoleRevoked)
				if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseRoleRevoked(log types.Log) (*BindingsRoleRevoked, error) {
	event := new(BindingsRoleRevoked)
	if err := _Bindings.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bindings contract.
type BindingsUnpausedIterator struct {
	Event *BindingsUnpaused // Event containing the contract specifics and raw log

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
func (it *BindingsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsUnpaused)
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
		it.Event = new(BindingsUnpaused)
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
func (it *BindingsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsUnpaused represents a Unpaused event raised by the Bindings contract.
type BindingsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bindings *BindingsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BindingsUnpausedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BindingsUnpausedIterator{contract: _Bindings.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bindings *BindingsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BindingsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsUnpaused)
				if err := _Bindings.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bindings *BindingsFilterer) ParseUnpaused(log types.Log) (*BindingsUnpaused, error) {
	event := new(BindingsUnpaused)
	if err := _Bindings.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
