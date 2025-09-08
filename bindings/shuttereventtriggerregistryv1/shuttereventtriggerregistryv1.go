// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shuttereventtriggerregistryv1

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

// Shuttereventtriggerregistryv1MetaData contains all meta data concerning the Shuttereventtriggerregistryv1 contract.
var Shuttereventtriggerregistryv1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EventTriggerRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"expirationBlockNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610aaa8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063715018a6146100595780638129fc1c146100635780638da5cb5b1461006d578063d48b6d1e1461008b578063f2fde38b146100a7575b5f5ffd5b6100616100c3565b005b61006b6100d6565b005b610075610256565b6040516100829190610667565b60405180910390f35b6100a560048036038101906100a0919061083d565b61028b565b005b6100c160048036038101906100bc91906108e7565b6102ec565b005b6100cb610370565b6100d45f6103f7565b565b5f6100df6104c8565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff161480156101275750825b90505f60018367ffffffffffffffff1614801561015a57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610168575080155b1561019f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083156101ec576001855f0160086101000a81548160ff0219169083151502179055505b6101f5336104db565b831561024f575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516102469190610954565b60405180910390a15b5050505050565b5f5f6102606104ef565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b610293610370565b8367ffffffffffffffff167f06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db584338585436102ce919061099a565b6040516102de9493929190610a53565b60405180910390a250505050565b6102f4610370565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610364575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161035b9190610667565b60405180910390fd5b61036d816103f7565b50565b610378610516565b73ffffffffffffffffffffffffffffffffffffffff16610396610256565b73ffffffffffffffffffffffffffffffffffffffff16146103f5576103b9610516565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016103ec9190610667565b60405180910390fd5b565b5f6104006104ef565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f5f6104d261051d565b90508091505090565b6104e3610546565b6104ec81610586565b50565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b61054e61060a565b610584576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b61058e610546565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105fe575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016105f59190610667565b60405180910390fd5b610607816103f7565b50565b5f6106136104c8565b5f0160089054906101000a900460ff16905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61065182610628565b9050919050565b61066181610647565b82525050565b5f60208201905061067a5f830184610658565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6106ad81610691565b81146106b7575f5ffd5b50565b5f813590506106c8816106a4565b92915050565b5f819050919050565b6106e0816106ce565b81146106ea575f5ffd5b50565b5f813590506106fb816106d7565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61074f82610709565b810181811067ffffffffffffffff8211171561076e5761076d610719565b5b80604052505050565b5f610780610680565b905061078c8282610746565b919050565b5f67ffffffffffffffff8211156107ab576107aa610719565b5b6107b482610709565b9050602081019050919050565b828183375f83830152505050565b5f6107e16107dc84610791565b610777565b9050828152602081018484840111156107fd576107fc610705565b5b6108088482856107c1565b509392505050565b5f82601f83011261082457610823610701565b5b81356108348482602086016107cf565b91505092915050565b5f5f5f5f6080858703121561085557610854610689565b5b5f610862878288016106ba565b9450506020610873878288016106ed565b935050604085013567ffffffffffffffff8111156108945761089361068d565b5b6108a087828801610810565b92505060606108b1878288016106ba565b91505092959194509250565b6108c681610647565b81146108d0575f5ffd5b50565b5f813590506108e1816108bd565b92915050565b5f602082840312156108fc576108fb610689565b5b5f610909848285016108d3565b91505092915050565b5f819050919050565b5f819050919050565b5f61093e61093961093484610912565b61091b565b610691565b9050919050565b61094e81610924565b82525050565b5f6020820190506109675f830184610945565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6109a482610691565b91506109af83610691565b9250828201905067ffffffffffffffff8111156109cf576109ce61096d565b5b92915050565b6109de816106ce565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610a16826109e4565b610a2081856109ee565b9350610a308185602086016109fe565b610a3981610709565b840191505092915050565b610a4d81610691565b82525050565b5f608082019050610a665f8301876109d5565b610a736020830186610658565b8181036040830152610a858185610a0c565b9050610a946060830184610a44565b9594505050505056fea164736f6c634300081c000a",
}

// Shuttereventtriggerregistryv1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Shuttereventtriggerregistryv1MetaData.ABI instead.
var Shuttereventtriggerregistryv1ABI = Shuttereventtriggerregistryv1MetaData.ABI

// Shuttereventtriggerregistryv1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Shuttereventtriggerregistryv1MetaData.Bin instead.
var Shuttereventtriggerregistryv1Bin = Shuttereventtriggerregistryv1MetaData.Bin

// DeployShuttereventtriggerregistryv1 deploys a new Ethereum contract, binding an instance of Shuttereventtriggerregistryv1 to it.
func DeployShuttereventtriggerregistryv1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Shuttereventtriggerregistryv1, error) {
	parsed, err := Shuttereventtriggerregistryv1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Shuttereventtriggerregistryv1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shuttereventtriggerregistryv1{Shuttereventtriggerregistryv1Caller: Shuttereventtriggerregistryv1Caller{contract: contract}, Shuttereventtriggerregistryv1Transactor: Shuttereventtriggerregistryv1Transactor{contract: contract}, Shuttereventtriggerregistryv1Filterer: Shuttereventtriggerregistryv1Filterer{contract: contract}}, nil
}

// Shuttereventtriggerregistryv1 is an auto generated Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1 struct {
	Shuttereventtriggerregistryv1Caller     // Read-only binding to the contract
	Shuttereventtriggerregistryv1Transactor // Write-only binding to the contract
	Shuttereventtriggerregistryv1Filterer   // Log filterer for contract events
}

// Shuttereventtriggerregistryv1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Shuttereventtriggerregistryv1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Shuttereventtriggerregistryv1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Shuttereventtriggerregistryv1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Shuttereventtriggerregistryv1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Shuttereventtriggerregistryv1Session struct {
	Contract     *Shuttereventtriggerregistryv1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Shuttereventtriggerregistryv1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Shuttereventtriggerregistryv1CallerSession struct {
	Contract *Shuttereventtriggerregistryv1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// Shuttereventtriggerregistryv1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Shuttereventtriggerregistryv1TransactorSession struct {
	Contract     *Shuttereventtriggerregistryv1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// Shuttereventtriggerregistryv1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1Raw struct {
	Contract *Shuttereventtriggerregistryv1 // Generic contract binding to access the raw methods on
}

// Shuttereventtriggerregistryv1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1CallerRaw struct {
	Contract *Shuttereventtriggerregistryv1Caller // Generic read-only contract binding to access the raw methods on
}

// Shuttereventtriggerregistryv1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Shuttereventtriggerregistryv1TransactorRaw struct {
	Contract *Shuttereventtriggerregistryv1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewShuttereventtriggerregistryv1 creates a new instance of Shuttereventtriggerregistryv1, bound to a specific deployed contract.
func NewShuttereventtriggerregistryv1(address common.Address, backend bind.ContractBackend) (*Shuttereventtriggerregistryv1, error) {
	contract, err := bindShuttereventtriggerregistryv1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1{Shuttereventtriggerregistryv1Caller: Shuttereventtriggerregistryv1Caller{contract: contract}, Shuttereventtriggerregistryv1Transactor: Shuttereventtriggerregistryv1Transactor{contract: contract}, Shuttereventtriggerregistryv1Filterer: Shuttereventtriggerregistryv1Filterer{contract: contract}}, nil
}

// NewShuttereventtriggerregistryv1Caller creates a new read-only instance of Shuttereventtriggerregistryv1, bound to a specific deployed contract.
func NewShuttereventtriggerregistryv1Caller(address common.Address, caller bind.ContractCaller) (*Shuttereventtriggerregistryv1Caller, error) {
	contract, err := bindShuttereventtriggerregistryv1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1Caller{contract: contract}, nil
}

// NewShuttereventtriggerregistryv1Transactor creates a new write-only instance of Shuttereventtriggerregistryv1, bound to a specific deployed contract.
func NewShuttereventtriggerregistryv1Transactor(address common.Address, transactor bind.ContractTransactor) (*Shuttereventtriggerregistryv1Transactor, error) {
	contract, err := bindShuttereventtriggerregistryv1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1Transactor{contract: contract}, nil
}

// NewShuttereventtriggerregistryv1Filterer creates a new log filterer instance of Shuttereventtriggerregistryv1, bound to a specific deployed contract.
func NewShuttereventtriggerregistryv1Filterer(address common.Address, filterer bind.ContractFilterer) (*Shuttereventtriggerregistryv1Filterer, error) {
	contract, err := bindShuttereventtriggerregistryv1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1Filterer{contract: contract}, nil
}

// bindShuttereventtriggerregistryv1 binds a generic wrapper to an already deployed contract.
func bindShuttereventtriggerregistryv1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Shuttereventtriggerregistryv1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shuttereventtriggerregistryv1.Contract.Shuttereventtriggerregistryv1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Shuttereventtriggerregistryv1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Shuttereventtriggerregistryv1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shuttereventtriggerregistryv1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shuttereventtriggerregistryv1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) Owner() (common.Address, error) {
	return _Shuttereventtriggerregistryv1.Contract.Owner(&_Shuttereventtriggerregistryv1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1CallerSession) Owner() (common.Address, error) {
	return _Shuttereventtriggerregistryv1.Contract.Owner(&_Shuttereventtriggerregistryv1.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Transactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) Initialize() (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Initialize(&_Shuttereventtriggerregistryv1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorSession) Initialize() (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Initialize(&_Shuttereventtriggerregistryv1.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Transactor) Register(opts *bind.TransactOpts, eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.contract.Transact(opts, "register", eon, identityPrefix, triggerDefinition, ttl)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) Register(eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Register(&_Shuttereventtriggerregistryv1.TransactOpts, eon, identityPrefix, triggerDefinition, ttl)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorSession) Register(eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.Register(&_Shuttereventtriggerregistryv1.TransactOpts, eon, identityPrefix, triggerDefinition, ttl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) RenounceOwnership() (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.RenounceOwnership(&_Shuttereventtriggerregistryv1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.RenounceOwnership(&_Shuttereventtriggerregistryv1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.TransferOwnership(&_Shuttereventtriggerregistryv1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.TransferOwnership(&_Shuttereventtriggerregistryv1.TransactOpts, newOwner)
}

// Shuttereventtriggerregistryv1EventTriggerRegisteredIterator is returned from FilterEventTriggerRegistered and is used to iterate over the raw logs and unpacked data for EventTriggerRegistered events raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1EventTriggerRegisteredIterator struct {
	Event *Shuttereventtriggerregistryv1EventTriggerRegistered // Event containing the contract specifics and raw log

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
func (it *Shuttereventtriggerregistryv1EventTriggerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Shuttereventtriggerregistryv1EventTriggerRegistered)
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
		it.Event = new(Shuttereventtriggerregistryv1EventTriggerRegistered)
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
func (it *Shuttereventtriggerregistryv1EventTriggerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Shuttereventtriggerregistryv1EventTriggerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Shuttereventtriggerregistryv1EventTriggerRegistered represents a EventTriggerRegistered event raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1EventTriggerRegistered struct {
	Eon                   uint64
	IdentityPrefix        [32]byte
	Sender                common.Address
	TriggerDefinition     []byte
	ExpirationBlockNumber uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterEventTriggerRegistered is a free log retrieval operation binding the contract event 0x06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5.
//
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 expirationBlockNumber)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) FilterEventTriggerRegistered(opts *bind.FilterOpts, eon []uint64) (*Shuttereventtriggerregistryv1EventTriggerRegisteredIterator, error) {

	var eonRule []interface{}
	for _, eonItem := range eon {
		eonRule = append(eonRule, eonItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.FilterLogs(opts, "EventTriggerRegistered", eonRule)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1EventTriggerRegisteredIterator{contract: _Shuttereventtriggerregistryv1.contract, event: "EventTriggerRegistered", logs: logs, sub: sub}, nil
}

// WatchEventTriggerRegistered is a free log subscription operation binding the contract event 0x06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5.
//
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 expirationBlockNumber)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) WatchEventTriggerRegistered(opts *bind.WatchOpts, sink chan<- *Shuttereventtriggerregistryv1EventTriggerRegistered, eon []uint64) (event.Subscription, error) {

	var eonRule []interface{}
	for _, eonItem := range eon {
		eonRule = append(eonRule, eonItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.WatchLogs(opts, "EventTriggerRegistered", eonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Shuttereventtriggerregistryv1EventTriggerRegistered)
				if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "EventTriggerRegistered", log); err != nil {
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

// ParseEventTriggerRegistered is a log parse operation binding the contract event 0x06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5.
//
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 expirationBlockNumber)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) ParseEventTriggerRegistered(log types.Log) (*Shuttereventtriggerregistryv1EventTriggerRegistered, error) {
	event := new(Shuttereventtriggerregistryv1EventTriggerRegistered)
	if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "EventTriggerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Shuttereventtriggerregistryv1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1InitializedIterator struct {
	Event *Shuttereventtriggerregistryv1Initialized // Event containing the contract specifics and raw log

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
func (it *Shuttereventtriggerregistryv1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Shuttereventtriggerregistryv1Initialized)
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
		it.Event = new(Shuttereventtriggerregistryv1Initialized)
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
func (it *Shuttereventtriggerregistryv1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Shuttereventtriggerregistryv1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Shuttereventtriggerregistryv1Initialized represents a Initialized event raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) FilterInitialized(opts *bind.FilterOpts) (*Shuttereventtriggerregistryv1InitializedIterator, error) {

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1InitializedIterator{contract: _Shuttereventtriggerregistryv1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Shuttereventtriggerregistryv1Initialized) (event.Subscription, error) {

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Shuttereventtriggerregistryv1Initialized)
				if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) ParseInitialized(log types.Log) (*Shuttereventtriggerregistryv1Initialized, error) {
	event := new(Shuttereventtriggerregistryv1Initialized)
	if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Shuttereventtriggerregistryv1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1OwnershipTransferredIterator struct {
	Event *Shuttereventtriggerregistryv1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Shuttereventtriggerregistryv1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Shuttereventtriggerregistryv1OwnershipTransferred)
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
		it.Event = new(Shuttereventtriggerregistryv1OwnershipTransferred)
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
func (it *Shuttereventtriggerregistryv1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Shuttereventtriggerregistryv1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Shuttereventtriggerregistryv1OwnershipTransferred represents a OwnershipTransferred event raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Shuttereventtriggerregistryv1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1OwnershipTransferredIterator{contract: _Shuttereventtriggerregistryv1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Shuttereventtriggerregistryv1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Shuttereventtriggerregistryv1OwnershipTransferred)
				if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) ParseOwnershipTransferred(log types.Log) (*Shuttereventtriggerregistryv1OwnershipTransferred, error) {
	event := new(Shuttereventtriggerregistryv1OwnershipTransferred)
	if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
