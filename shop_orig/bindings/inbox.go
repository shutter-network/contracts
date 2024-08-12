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

// InboxMetaData contains all meta data concerning the Inbox contract.
var InboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_blockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"initializer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLOCK_GAS_LIMIT_SETTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clear\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBlockGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransactions\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structInbox.Transaction[]\",\"components\":[{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"dao\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequencer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockGasLimit\",\"inputs\":[{\"name\":\"newBlockGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptedTransaction\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EncryptedTransactionSubmitted\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"block\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cumulativeGasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockGasLimitExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferEtherFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedInitializer\",\"inputs\":[]}]",
}

// InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxMetaData.ABI instead.
var InboxABI = InboxMetaData.ABI

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) BLOCKGASLIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "BLOCK_GAS_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// BLOCKGASLIMITSETTERROLE is a free data retrieval call binding the contract method 0x91de737e.
//
// Solidity: function BLOCK_GAS_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) BLOCKGASLIMITSETTERROLE() ([32]byte, error) {
	return _Inbox.Contract.BLOCKGASLIMITSETTERROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Inbox.Contract.DEFAULTADMINROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Inbox.Contract.PAUSERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) SEQUENCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "SEQUENCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// SEQUENCERROLE is a free data retrieval call binding the contract method 0x4842855c.
//
// Solidity: function SEQUENCER_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) SEQUENCERROLE() ([32]byte, error) {
	return _Inbox.Contract.SEQUENCERROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCaller) WITHDRAWROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "WITHDRAW_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// WITHDRAWROLE is a free data retrieval call binding the contract method 0xe02023a1.
//
// Solidity: function WITHDRAW_ROLE() view returns(bytes32)
func (_Inbox *InboxCallerSession) WITHDRAWROLE() ([32]byte, error) {
	return _Inbox.Contract.WITHDRAWROLE(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCaller) GetBlockGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getBlockGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint64)
func (_Inbox *InboxCallerSession) GetBlockGasLimit() (uint64, error) {
	return _Inbox.Contract.GetBlockGasLimit(&_Inbox.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Inbox *InboxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Inbox.Contract.GetRoleAdmin(&_Inbox.CallOpts, role)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxCaller) GetTransactions(opts *bind.CallOpts, blockNumber uint64) ([]InboxTransaction, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "getTransactions", blockNumber)

	if err != nil {
		return *new([]InboxTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new([]InboxTransaction)).(*[]InboxTransaction)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// GetTransactions is a free data retrieval call binding the contract method 0x552fd4aa.
//
// Solidity: function getTransactions(uint64 blockNumber) view returns((bytes,address,uint64,uint64)[])
func (_Inbox *InboxCallerSession) GetTransactions(blockNumber uint64) ([]InboxTransaction, error) {
	return _Inbox.Contract.GetTransactions(&_Inbox.CallOpts, blockNumber)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Inbox *InboxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Inbox.Contract.HasRole(&_Inbox.CallOpts, role, account)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxCaller) Initializer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "initializer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxSession) Initializer() (common.Address, error) {
	return _Inbox.Contract.Initializer(&_Inbox.CallOpts)
}

// Initializer is a free data retrieval call binding the contract method 0x9ce110d7.
//
// Solidity: function initializer() view returns(address)
func (_Inbox *InboxCallerSession) Initializer() (common.Address, error) {
	return _Inbox.Contract.Initializer(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCallerSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Inbox *InboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Inbox.Contract.SupportsInterface(&_Inbox.CallOpts, interfaceId)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns()
func (_Inbox *InboxTransactorSession) Clear() (*types.Transaction, error) {
	return _Inbox.Contract.Clear(&_Inbox.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.GrantRole(&_Inbox.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, dao, sequencer)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address dao, address sequencer) returns()
func (_Inbox *InboxTransactorSession) Initialize(dao common.Address, sequencer common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, dao, sequencer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Inbox *InboxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RenounceRole(&_Inbox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Inbox *InboxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.RevokeRole(&_Inbox.TransactOpts, role, account)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactor) SetBlockGasLimit(opts *bind.TransactOpts, newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setBlockGasLimit", newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xae0dad60.
//
// Solidity: function setBlockGasLimit(uint64 newBlockGasLimit) returns()
func (_Inbox *InboxTransactorSession) SetBlockGasLimit(newBlockGasLimit uint64) (*types.Transaction, error) {
	return _Inbox.Contract.SetBlockGasLimit(&_Inbox.TransactOpts, newBlockGasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "submitEncryptedTransaction", blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0xa44e7cfe.
//
// Solidity: function submitEncryptedTransaction(uint64 blockNumber, bytes encryptedTransaction, uint64 gasLimit, address excessFeeRecipient) payable returns()
func (_Inbox *InboxTransactorSession) SubmitEncryptedTransaction(blockNumber uint64, encryptedTransaction []byte, gasLimit uint64, excessFeeRecipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SubmitEncryptedTransaction(&_Inbox.TransactOpts, blockNumber, encryptedTransaction, gasLimit, excessFeeRecipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Inbox *InboxTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Withdraw(&_Inbox.TransactOpts, recipient)
}

// InboxEncryptedTransactionSubmittedIterator is returned from FilterEncryptedTransactionSubmitted and is used to iterate over the raw logs and unpacked data for EncryptedTransactionSubmitted events raised by the Inbox contract.
type InboxEncryptedTransactionSubmittedIterator struct {
	Event *InboxEncryptedTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *InboxEncryptedTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxEncryptedTransactionSubmitted)
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
		it.Event = new(InboxEncryptedTransactionSubmitted)
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
func (it *InboxEncryptedTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxEncryptedTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxEncryptedTransactionSubmitted represents a EncryptedTransactionSubmitted event raised by the Inbox contract.
type InboxEncryptedTransactionSubmitted struct {
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
func (_Inbox *InboxFilterer) FilterEncryptedTransactionSubmitted(opts *bind.FilterOpts, index []uint64, block []uint64) (*InboxEncryptedTransactionSubmittedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return &InboxEncryptedTransactionSubmittedIterator{contract: _Inbox.contract, event: "EncryptedTransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchEncryptedTransactionSubmitted is a free log subscription operation binding the contract event 0x0ac44a68e7048007d82fa113a6c5d84bec8a110ed9e953d4f234fab3c9ecac53.
//
// Solidity: event EncryptedTransactionSubmitted(uint64 indexed index, uint64 indexed block, bytes encryptedTransaction, address sender, uint64 gasLimit, uint64 cumulativeGasLimit, uint256 fee)
func (_Inbox *InboxFilterer) WatchEncryptedTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *InboxEncryptedTransactionSubmitted, index []uint64, block []uint64) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var blockRule []interface{}
	for _, blockItem := range block {
		blockRule = append(blockRule, blockItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "EncryptedTransactionSubmitted", indexRule, blockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxEncryptedTransactionSubmitted)
				if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseEncryptedTransactionSubmitted(log types.Log) (*InboxEncryptedTransactionSubmitted, error) {
	event := new(InboxEncryptedTransactionSubmitted)
	if err := _Inbox.contract.UnpackLog(event, "EncryptedTransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxFeesWithdrawnIterator is returned from FilterFeesWithdrawn and is used to iterate over the raw logs and unpacked data for FeesWithdrawn events raised by the Inbox contract.
type InboxFeesWithdrawnIterator struct {
	Event *InboxFeesWithdrawn // Event containing the contract specifics and raw log

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
func (it *InboxFeesWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxFeesWithdrawn)
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
		it.Event = new(InboxFeesWithdrawn)
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
func (it *InboxFeesWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxFeesWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxFeesWithdrawn represents a FeesWithdrawn event raised by the Inbox contract.
type InboxFeesWithdrawn struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesWithdrawn is a free log retrieval operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) FilterFeesWithdrawn(opts *bind.FilterOpts) (*InboxFeesWithdrawnIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return &InboxFeesWithdrawnIterator{contract: _Inbox.contract, event: "FeesWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeesWithdrawn is a free log subscription operation binding the contract event 0x792248b395a0ac81e65e6d79494b5382c8de690233f36c2e5e672f77044887c7.
//
// Solidity: event FeesWithdrawn(address recipient)
func (_Inbox *InboxFilterer) WatchFeesWithdrawn(opts *bind.WatchOpts, sink chan<- *InboxFeesWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "FeesWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxFeesWithdrawn)
				if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseFeesWithdrawn(log types.Log) (*InboxFeesWithdrawn, error) {
	event := new(InboxFeesWithdrawn)
	if err := _Inbox.contract.UnpackLog(event, "FeesWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Inbox contract.
type InboxPausedIterator struct {
	Event *InboxPaused // Event containing the contract specifics and raw log

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
func (it *InboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPaused)
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
		it.Event = new(InboxPaused)
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
func (it *InboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPaused represents a Paused event raised by the Inbox contract.
type InboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) FilterPaused(opts *bind.FilterOpts) (*InboxPausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InboxPausedIterator{contract: _Inbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InboxPaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPaused)
				if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParsePaused(log types.Log) (*InboxPaused, error) {
	event := new(InboxPaused)
	if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Inbox contract.
type InboxRoleAdminChangedIterator struct {
	Event *InboxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *InboxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleAdminChanged)
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
		it.Event = new(InboxRoleAdminChanged)
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
func (it *InboxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleAdminChanged represents a RoleAdminChanged event raised by the Inbox contract.
type InboxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*InboxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleAdminChangedIterator{contract: _Inbox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Inbox *InboxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *InboxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleAdminChanged)
				if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleAdminChanged(log types.Log) (*InboxRoleAdminChanged, error) {
	event := new(InboxRoleAdminChanged)
	if err := _Inbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Inbox contract.
type InboxRoleGrantedIterator struct {
	Event *InboxRoleGranted // Event containing the contract specifics and raw log

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
func (it *InboxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleGranted)
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
		it.Event = new(InboxRoleGranted)
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
func (it *InboxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleGranted represents a RoleGranted event raised by the Inbox contract.
type InboxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleGrantedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleGrantedIterator{contract: _Inbox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *InboxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleGranted)
				if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleGranted(log types.Log) (*InboxRoleGranted, error) {
	event := new(InboxRoleGranted)
	if err := _Inbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Inbox contract.
type InboxRoleRevokedIterator struct {
	Event *InboxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *InboxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxRoleRevoked)
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
		it.Event = new(InboxRoleRevoked)
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
func (it *InboxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxRoleRevoked represents a RoleRevoked event raised by the Inbox contract.
type InboxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InboxRoleRevokedIterator, error) {

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

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InboxRoleRevokedIterator{contract: _Inbox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Inbox *InboxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *InboxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxRoleRevoked)
				if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseRoleRevoked(log types.Log) (*InboxRoleRevoked, error) {
	event := new(InboxRoleRevoked)
	if err := _Inbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Inbox contract.
type InboxUnpausedIterator struct {
	Event *InboxUnpaused // Event containing the contract specifics and raw log

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
func (it *InboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxUnpaused)
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
		it.Event = new(InboxUnpaused)
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
func (it *InboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxUnpaused represents a Unpaused event raised by the Inbox contract.
type InboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InboxUnpausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InboxUnpausedIterator{contract: _Inbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxUnpaused)
				if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseUnpaused(log types.Log) (*InboxUnpaused, error) {
	event := new(InboxUnpaused)
	if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
