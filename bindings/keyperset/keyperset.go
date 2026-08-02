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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMembers\",\"inputs\":[{\"name\":\"newMembers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDKGContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMember\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNumMembers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPublisher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedToBroadcastEonKey\",\"inputs\":[{\"name\":\"a\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFinalized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDKGContract\",\"inputs\":[{\"name\":\"_dkgContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFinalized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPublisher\",\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610081575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100789190610196565b60405180910390fd5b6100908161009660201b60201c565b506101af565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018082610157565b9050919050565b61019081610176565b82525050565b5f6020820190506101a95f830184610187565b92915050565b610d02806101bc5f395ff3fe608060405234801561000f575f5ffd5b50600436106100fd575f3560e01c80638d4e408311610095578063cde1532d11610064578063cde1532d1461024b578063dbf4ab4e1461027b578063e75235b814610299578063f2fde38b146102b7576100fd565b80638d4e4083146101d55780638da5cb5b146101f35780639eab525314610211578063cab636611461022f576100fd565b80632e8e6cad116100d15780632e8e6cad146101635780636ab47aae146101935780636f4d469b146101af578063715018a6146101cb576100fd565b806266f0a81461010157806317c4de351461011f57806317d5430a1461013b5780631de7725314610159575b5f5ffd5b6101096102d3565b60405161011691906109b0565b60405180910390f35b61013960048036038101906101349190610a0e565b6102fb565b005b610143610374565b6040516101509190610a48565b60405180910390f35b610161610380565b005b61017d60048036038101906101789190610a0e565b6103a4565b60405161018a91906109b0565b60405180910390f35b6101ad60048036038101906101a89190610a8b565b6103f2565b005b6101c960048036038101906101c49190610b17565b610483565b005b6101d3610594565b005b6101dd6105a7565b6040516101ea9190610b7c565b60405180910390f35b6101fb6105bc565b60405161020891906109b0565b60405180910390f35b6102196105e3565b6040516102269190610c4c565b60405180910390f35b61024960048036038101906102449190610a8b565b61066e565b005b61026560048036038101906102609190610a8b565b610700565b6040516102729190610b7c565b60405180910390f35b610283610759565b60405161029091906109b0565b60405180910390f35b6102a1610782565b6040516102ae9190610a48565b60405180910390f35b6102d160048036038101906102cc9190610a8b565b61079e565b005b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610303610822565b5f60149054906101000a900460ff1615610349576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060025f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b5f600180549050905090565b610388610822565b60015f60146101000a81548160ff021916908315150217905550565b5f60018267ffffffffffffffff16815481106103c3576103c2610c6c565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6103fa610822565b5f60149054906101000a900460ff1615610440576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61048b610822565b5f60149054906101000a900460ff16156104d1576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f90505b828290508167ffffffffffffffff16101561058f57600183838367ffffffffffffffff1681811061050a57610509610c6c565b5b905060200201602081019061051f9190610a8b565b908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061058790610cc6565b9150506104d6565b505050565b61059c610822565b6105a55f6108a9565b565b5f5f60149054906101000a900460ff16905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600180548060200260200160405190810160405280929190818152602001828054801561066457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161061b575b5050505050905090565b610676610822565b5f60149054906101000a900460ff16156106bc576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600260086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f600260089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050919050565b5f600260089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60025f9054906101000a900467ffffffffffffffff16905090565b6107a6610822565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610816575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161080d91906109b0565b60405180910390fd5b61081f816108a9565b50565b61082a61096a565b73ffffffffffffffffffffffffffffffffffffffff166108486105bc565b73ffffffffffffffffffffffffffffffffffffffff16146108a75761086b61096a565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161089e91906109b0565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61099a82610971565b9050919050565b6109aa81610990565b82525050565b5f6020820190506109c35f8301846109a1565b92915050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6109ed816109d1565b81146109f7575f5ffd5b50565b5f81359050610a08816109e4565b92915050565b5f60208284031215610a2357610a226109c9565b5b5f610a30848285016109fa565b91505092915050565b610a42816109d1565b82525050565b5f602082019050610a5b5f830184610a39565b92915050565b610a6a81610990565b8114610a74575f5ffd5b50565b5f81359050610a8581610a61565b92915050565b5f60208284031215610aa057610a9f6109c9565b5b5f610aad84828501610a77565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610ad757610ad6610ab6565b5b8235905067ffffffffffffffff811115610af457610af3610aba565b5b602083019150836020820283011115610b1057610b0f610abe565b5b9250929050565b5f5f60208385031215610b2d57610b2c6109c9565b5b5f83013567ffffffffffffffff811115610b4a57610b496109cd565b5b610b5685828601610ac2565b92509250509250929050565b5f8115159050919050565b610b7681610b62565b82525050565b5f602082019050610b8f5f830184610b6d565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610bc781610990565b82525050565b5f610bd88383610bbe565b60208301905092915050565b5f602082019050919050565b5f610bfa82610b95565b610c048185610b9f565b9350610c0f83610baf565b805f5b83811015610c3f578151610c268882610bcd565b9750610c3183610be4565b925050600181019050610c12565b5085935050505092915050565b5f6020820190508181035f830152610c648184610bf0565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610cd0826109d1565b915067ffffffffffffffff8203610cea57610ce9610c99565b5b60018201905091905056fea164736f6c634300081c000a",
}

// KeypersetABI is the input ABI used to generate the binding from.
// Deprecated: Use KeypersetMetaData.ABI instead.
var KeypersetABI = KeypersetMetaData.ABI

// KeypersetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeypersetMetaData.Bin instead.
var KeypersetBin = KeypersetMetaData.Bin

// DeployKeyperset deploys a new Ethereum contract, binding an instance of Keyperset to it.
func DeployKeyperset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Keyperset, error) {
	parsed, err := KeypersetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeypersetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Keyperset{KeypersetCaller: KeypersetCaller{contract: contract}, KeypersetTransactor: KeypersetTransactor{contract: contract}, KeypersetFilterer: KeypersetFilterer{contract: contract}}, nil
}

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

// GetDKGContract is a free data retrieval call binding the contract method 0x0066f0a8.
//
// Solidity: function getDKGContract() view returns(address)
func (_Keyperset *KeypersetCaller) GetDKGContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Keyperset.contract.Call(opts, &out, "getDKGContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDKGContract is a free data retrieval call binding the contract method 0x0066f0a8.
//
// Solidity: function getDKGContract() view returns(address)
func (_Keyperset *KeypersetSession) GetDKGContract() (common.Address, error) {
	return _Keyperset.Contract.GetDKGContract(&_Keyperset.CallOpts)
}

// GetDKGContract is a free data retrieval call binding the contract method 0x0066f0a8.
//
// Solidity: function getDKGContract() view returns(address)
func (_Keyperset *KeypersetCallerSession) GetDKGContract() (common.Address, error) {
	return _Keyperset.Contract.GetDKGContract(&_Keyperset.CallOpts)
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

// SetDKGContract is a paid mutator transaction binding the contract method 0x6ab47aae.
//
// Solidity: function setDKGContract(address _dkgContract) returns()
func (_Keyperset *KeypersetTransactor) SetDKGContract(opts *bind.TransactOpts, _dkgContract common.Address) (*types.Transaction, error) {
	return _Keyperset.contract.Transact(opts, "setDKGContract", _dkgContract)
}

// SetDKGContract is a paid mutator transaction binding the contract method 0x6ab47aae.
//
// Solidity: function setDKGContract(address _dkgContract) returns()
func (_Keyperset *KeypersetSession) SetDKGContract(_dkgContract common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.SetDKGContract(&_Keyperset.TransactOpts, _dkgContract)
}

// SetDKGContract is a paid mutator transaction binding the contract method 0x6ab47aae.
//
// Solidity: function setDKGContract(address _dkgContract) returns()
func (_Keyperset *KeypersetTransactorSession) SetDKGContract(_dkgContract common.Address) (*types.Transaction, error) {
	return _Keyperset.Contract.SetDKGContract(&_Keyperset.TransactOpts, _dkgContract)
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
