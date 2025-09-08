// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eonkeypublish

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

// EonkeypublishMetaData contains all meta data concerning the Eonkeypublish contract.
var EonkeypublishMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_keyperSet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_broadcaster\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eonKeyConfirmed\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasKeyperVoted\",\"inputs\":[{\"name\":\"keyper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishEonKey\",\"inputs\":[{\"name\":\"eonKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"keyperId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EonVoteRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyVoted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610ddc380380610ddc8339818101604052810190610031919061017d565b8260025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050506101cd565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61010f826100e6565b9050919050565b61011f81610105565b8114610129575f5ffd5b50565b5f8151905061013a81610116565b92915050565b5f67ffffffffffffffff82169050919050565b61015c81610140565b8114610166575f5ffd5b50565b5f8151905061017781610153565b92915050565b5f5f5f60608486031215610194576101936100e2565b5b5f6101a18682870161012c565b93505060206101b28682870161012c565b92505060406101c386828701610169565b9150509250925092565b610c02806101da5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c8063517e1cf714610043578063a1dd75ba14610073578063b118b6ed1461008f575b5f5ffd5b61005d600480360381019061005891906107c0565b6100bf565b60405161006a9190610821565b60405180910390f35b61008d60048036038101906100889190610877565b6101b0565b005b6100a960048036038101906100a4919061092b565b610621565b6040516100b69190610821565b60405180910390f35b5f5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e75235b86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561012b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061014f919061096a565b67ffffffffffffffff1690505f5f5f858051906020012081526020019081526020015f205f9054906101000a900467ffffffffffffffff169050818167ffffffffffffffff16106101a5576001925050506101ab565b5f925050505b919050565b5f8251036101ea576040517f76d4e1e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d4e40836040518163ffffffff1660e01b8152600401602060405180830381865afa158015610254573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061027891906109bf565b6102ae576040517feac631aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632e8e6cad836040518263ffffffff1660e01b815260040161031f91906109f9565b602060405180830381865afa15801561033a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035e9190610a26565b73ffffffffffffffffffffffffffffffffffffffff16146103ab576040517f3d693ada00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610502576001805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f5f838051906020012081526020019081526020015f205f81819054906101000a900467ffffffffffffffff168092919061048790610a7e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550507fd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a600360149054906101000a900467ffffffffffffffff16836040516104f5929190610b0d565b60405180910390a1610534565b6040517f7c9a1cf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61053d826100bf565b1561061c5760035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663daade8e8600360149054906101000a900467ffffffffffffffff16846040518363ffffffff1660e01b81526004016105b4929190610b0d565b5f604051808303815f87803b1580156105cb575f5ffd5b505af19250505080156105dc575060015b610617576105e8610b47565b806308c379a00361060d57506105fc610b66565b80610607575061060f565b5061061d565b505b3d5f5f3e3d5ffd5b61061d565b5b5050565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6106d28261068c565b810181811067ffffffffffffffff821117156106f1576106f061069c565b5b80604052505050565b5f610703610673565b905061070f82826106c9565b919050565b5f67ffffffffffffffff82111561072e5761072d61069c565b5b6107378261068c565b9050602081019050919050565b828183375f83830152505050565b5f61076461075f84610714565b6106fa565b9050828152602081018484840111156107805761077f610688565b5b61078b848285610744565b509392505050565b5f82601f8301126107a7576107a6610684565b5b81356107b7848260208601610752565b91505092915050565b5f602082840312156107d5576107d461067c565b5b5f82013567ffffffffffffffff8111156107f2576107f1610680565b5b6107fe84828501610793565b91505092915050565b5f8115159050919050565b61081b81610807565b82525050565b5f6020820190506108345f830184610812565b92915050565b5f67ffffffffffffffff82169050919050565b6108568161083a565b8114610860575f5ffd5b50565b5f813590506108718161084d565b92915050565b5f5f6040838503121561088d5761088c61067c565b5b5f83013567ffffffffffffffff8111156108aa576108a9610680565b5b6108b685828601610793565b92505060206108c785828601610863565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108fa826108d1565b9050919050565b61090a816108f0565b8114610914575f5ffd5b50565b5f8135905061092581610901565b92915050565b5f602082840312156109405761093f61067c565b5b5f61094d84828501610917565b91505092915050565b5f815190506109648161084d565b92915050565b5f6020828403121561097f5761097e61067c565b5b5f61098c84828501610956565b91505092915050565b61099e81610807565b81146109a8575f5ffd5b50565b5f815190506109b981610995565b92915050565b5f602082840312156109d4576109d361067c565b5b5f6109e1848285016109ab565b91505092915050565b6109f38161083a565b82525050565b5f602082019050610a0c5f8301846109ea565b92915050565b5f81519050610a2081610901565b92915050565b5f60208284031215610a3b57610a3a61067c565b5b5f610a4884828501610a12565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610a888261083a565b915067ffffffffffffffff8203610aa257610aa1610a51565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610adf82610aad565b610ae98185610ab7565b9350610af9818560208601610ac7565b610b028161068c565b840191505092915050565b5f604082019050610b205f8301856109ea565b8181036020830152610b328184610ad5565b90509392505050565b5f8160e01c9050919050565b5f60033d1115610b635760045f5f3e610b605f51610b3b565b90505b90565b5f60443d10610bf257610b77610673565b60043d036004823e80513d602482011167ffffffffffffffff82111715610b9f575050610bf2565b808201805167ffffffffffffffff811115610bbd5750505050610bf2565b80602083010160043d038501811115610bda575050505050610bf2565b610be9826020018501866106c9565b82955050505050505b9056fea164736f6c634300081c000a",
}

// EonkeypublishABI is the input ABI used to generate the binding from.
// Deprecated: Use EonkeypublishMetaData.ABI instead.
var EonkeypublishABI = EonkeypublishMetaData.ABI

// EonkeypublishBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EonkeypublishMetaData.Bin instead.
var EonkeypublishBin = EonkeypublishMetaData.Bin

// DeployEonkeypublish deploys a new Ethereum contract, binding an instance of Eonkeypublish to it.
func DeployEonkeypublish(auth *bind.TransactOpts, backend bind.ContractBackend, _keyperSet common.Address, _broadcaster common.Address, _eon uint64) (common.Address, *types.Transaction, *Eonkeypublish, error) {
	parsed, err := EonkeypublishMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EonkeypublishBin), backend, _keyperSet, _broadcaster, _eon)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eonkeypublish{EonkeypublishCaller: EonkeypublishCaller{contract: contract}, EonkeypublishTransactor: EonkeypublishTransactor{contract: contract}, EonkeypublishFilterer: EonkeypublishFilterer{contract: contract}}, nil
}

// Eonkeypublish is an auto generated Go binding around an Ethereum contract.
type Eonkeypublish struct {
	EonkeypublishCaller     // Read-only binding to the contract
	EonkeypublishTransactor // Write-only binding to the contract
	EonkeypublishFilterer   // Log filterer for contract events
}

// EonkeypublishCaller is an auto generated read-only Go binding around an Ethereum contract.
type EonkeypublishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EonkeypublishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EonkeypublishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EonkeypublishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EonkeypublishSession struct {
	Contract     *Eonkeypublish    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EonkeypublishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EonkeypublishCallerSession struct {
	Contract *EonkeypublishCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EonkeypublishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EonkeypublishTransactorSession struct {
	Contract     *EonkeypublishTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EonkeypublishRaw is an auto generated low-level Go binding around an Ethereum contract.
type EonkeypublishRaw struct {
	Contract *Eonkeypublish // Generic contract binding to access the raw methods on
}

// EonkeypublishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EonkeypublishCallerRaw struct {
	Contract *EonkeypublishCaller // Generic read-only contract binding to access the raw methods on
}

// EonkeypublishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EonkeypublishTransactorRaw struct {
	Contract *EonkeypublishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEonkeypublish creates a new instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublish(address common.Address, backend bind.ContractBackend) (*Eonkeypublish, error) {
	contract, err := bindEonkeypublish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eonkeypublish{EonkeypublishCaller: EonkeypublishCaller{contract: contract}, EonkeypublishTransactor: EonkeypublishTransactor{contract: contract}, EonkeypublishFilterer: EonkeypublishFilterer{contract: contract}}, nil
}

// NewEonkeypublishCaller creates a new read-only instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishCaller(address common.Address, caller bind.ContractCaller) (*EonkeypublishCaller, error) {
	contract, err := bindEonkeypublish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishCaller{contract: contract}, nil
}

// NewEonkeypublishTransactor creates a new write-only instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishTransactor(address common.Address, transactor bind.ContractTransactor) (*EonkeypublishTransactor, error) {
	contract, err := bindEonkeypublish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishTransactor{contract: contract}, nil
}

// NewEonkeypublishFilterer creates a new log filterer instance of Eonkeypublish, bound to a specific deployed contract.
func NewEonkeypublishFilterer(address common.Address, filterer bind.ContractFilterer) (*EonkeypublishFilterer, error) {
	contract, err := bindEonkeypublish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EonkeypublishFilterer{contract: contract}, nil
}

// bindEonkeypublish binds a generic wrapper to an already deployed contract.
func bindEonkeypublish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EonkeypublishMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eonkeypublish *EonkeypublishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eonkeypublish.Contract.EonkeypublishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eonkeypublish *EonkeypublishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.EonkeypublishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eonkeypublish *EonkeypublishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.EonkeypublishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eonkeypublish *EonkeypublishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eonkeypublish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eonkeypublish *EonkeypublishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eonkeypublish *EonkeypublishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.contract.Transact(opts, method, params...)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishCaller) EonKeyConfirmed(opts *bind.CallOpts, eonKey []byte) (bool, error) {
	var out []interface{}
	err := _Eonkeypublish.contract.Call(opts, &out, "eonKeyConfirmed", eonKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Eonkeypublish.Contract.EonKeyConfirmed(&_Eonkeypublish.CallOpts, eonKey)
}

// EonKeyConfirmed is a free data retrieval call binding the contract method 0x517e1cf7.
//
// Solidity: function eonKeyConfirmed(bytes eonKey) view returns(bool)
func (_Eonkeypublish *EonkeypublishCallerSession) EonKeyConfirmed(eonKey []byte) (bool, error) {
	return _Eonkeypublish.Contract.EonKeyConfirmed(&_Eonkeypublish.CallOpts, eonKey)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishCaller) HasKeyperVoted(opts *bind.CallOpts, keyper common.Address) (bool, error) {
	var out []interface{}
	err := _Eonkeypublish.contract.Call(opts, &out, "hasKeyperVoted", keyper)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _Eonkeypublish.Contract.HasKeyperVoted(&_Eonkeypublish.CallOpts, keyper)
}

// HasKeyperVoted is a free data retrieval call binding the contract method 0xb118b6ed.
//
// Solidity: function hasKeyperVoted(address keyper) view returns(bool)
func (_Eonkeypublish *EonkeypublishCallerSession) HasKeyperVoted(keyper common.Address) (bool, error) {
	return _Eonkeypublish.Contract.HasKeyperVoted(&_Eonkeypublish.CallOpts, keyper)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishTransactor) PublishEonKey(opts *bind.TransactOpts, eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.contract.Transact(opts, "publishEonKey", eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.PublishEonKey(&_Eonkeypublish.TransactOpts, eonKey, keyperId)
}

// PublishEonKey is a paid mutator transaction binding the contract method 0xa1dd75ba.
//
// Solidity: function publishEonKey(bytes eonKey, uint64 keyperId) returns()
func (_Eonkeypublish *EonkeypublishTransactorSession) PublishEonKey(eonKey []byte, keyperId uint64) (*types.Transaction, error) {
	return _Eonkeypublish.Contract.PublishEonKey(&_Eonkeypublish.TransactOpts, eonKey, keyperId)
}

// EonkeypublishEonVoteRegisteredIterator is returned from FilterEonVoteRegistered and is used to iterate over the raw logs and unpacked data for EonVoteRegistered events raised by the Eonkeypublish contract.
type EonkeypublishEonVoteRegisteredIterator struct {
	Event *EonkeypublishEonVoteRegistered // Event containing the contract specifics and raw log

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
func (it *EonkeypublishEonVoteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EonkeypublishEonVoteRegistered)
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
		it.Event = new(EonkeypublishEonVoteRegistered)
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
func (it *EonkeypublishEonVoteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EonkeypublishEonVoteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EonkeypublishEonVoteRegistered represents a EonVoteRegistered event raised by the Eonkeypublish contract.
type EonkeypublishEonVoteRegistered struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonVoteRegistered is a free log retrieval operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Eonkeypublish *EonkeypublishFilterer) FilterEonVoteRegistered(opts *bind.FilterOpts) (*EonkeypublishEonVoteRegisteredIterator, error) {

	logs, sub, err := _Eonkeypublish.contract.FilterLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return &EonkeypublishEonVoteRegisteredIterator{contract: _Eonkeypublish.contract, event: "EonVoteRegistered", logs: logs, sub: sub}, nil
}

// WatchEonVoteRegistered is a free log subscription operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Eonkeypublish *EonkeypublishFilterer) WatchEonVoteRegistered(opts *bind.WatchOpts, sink chan<- *EonkeypublishEonVoteRegistered) (event.Subscription, error) {

	logs, sub, err := _Eonkeypublish.contract.WatchLogs(opts, "EonVoteRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EonkeypublishEonVoteRegistered)
				if err := _Eonkeypublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
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

// ParseEonVoteRegistered is a log parse operation binding the contract event 0xd521fd602f7dce253c2e07c489358eb9e00e3af335a362168a0300d30e401a8a.
//
// Solidity: event EonVoteRegistered(uint64 eon, bytes key)
func (_Eonkeypublish *EonkeypublishFilterer) ParseEonVoteRegistered(log types.Log) (*EonkeypublishEonVoteRegistered, error) {
	event := new(EonkeypublishEonVoteRegistered)
	if err := _Eonkeypublish.contract.UnpackLog(event, "EonVoteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
