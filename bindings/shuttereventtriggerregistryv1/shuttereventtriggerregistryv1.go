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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"EventTriggerRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"expirationBlockNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610042575f5ffd5b5061005161005660201b60201c565b6101d1565b5f61006561015460201b60201c565b9050805f0160089054906101000a900460ff16156100af576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101515767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161014891906101b8565b60405180910390a15b50565b5f5f61016461016d60201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b6101b281610196565b82525050565b5f6020820190506101cb5f8301846101a9565b92915050565b6080516113436101f75f395f81816104d40152818161052901526106e301526113435ff3fe60806040526004361061007a575f3560e01c80638da5cb5b1161004d5780638da5cb5b146100f0578063ad3cb1cc1461011a578063d48b6d1e14610144578063f2fde38b1461016c5761007a565b80634f1ef2861461007e57806352d1902d1461009a578063715018a6146100c45780638129fc1c146100da575b5f5ffd5b61009860048036038101906100939190610ef4565b610194565b005b3480156100a5575f5ffd5b506100ae6101b3565b6040516100bb9190610f66565b60405180910390f35b3480156100cf575f5ffd5b506100d86101e4565b005b3480156100e5575f5ffd5b506100ee6101f7565b005b3480156100fb575f5ffd5b5061010461037f565b6040516101119190610f8e565b60405180910390f35b348015610125575f5ffd5b5061012e6103b4565b60405161013b9190611007565b60405180910390f35b34801561014f575f5ffd5b5061016a6004803603810190610165919061108e565b6103ed565b005b348015610177575f5ffd5b50610192600480360381019061018d919061110e565b61044e565b005b61019c6104d2565b6101a5826105b8565b6101af82826105c3565b5050565b5f6101bc6106e1565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b6101ec610768565b6101f55f6107ef565b565b5f6102006108c0565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff161480156102485750825b90505f60018367ffffffffffffffff1614801561027b57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610289575080155b156102c0576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561030d576001855f0160086101000a81548160ff0219169083151502179055505b610316336108d3565b61031e6108e7565b8315610378575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2600160405161036f919061117b565b60405180910390a15b5050505050565b5f5f6103896108f1565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6103f5610768565b8367ffffffffffffffff167f06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5843385854361043091906111c1565b604051610440949392919061125d565b60405180910390a250505050565b610456610768565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104c6575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104bd9190610f8e565b60405180910390fd5b6104cf816107ef565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16148061057f57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610566610918565b73ffffffffffffffffffffffffffffffffffffffff1614155b156105b6576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6105c0610768565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561062b57506040513d601f19601f8201168201806040525081019061062891906112bb565b60015b61066c57816040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016106639190610f8e565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b81146106d257806040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004016106c99190610f66565b60405180910390fd5b6106dc838361096b565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610766576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6107706109dd565b73ffffffffffffffffffffffffffffffffffffffff1661078e61037f565b73ffffffffffffffffffffffffffffffffffffffff16146107ed576107b16109dd565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016107e49190610f8e565b60405180910390fd5b565b5f6107f86108f1565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f5f6108ca6109e4565b90508091505090565b6108db610a0d565b6108e481610a4d565b50565b6108ef610a0d565b565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f6109447f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ad1565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61097482610ada565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f815111156109d0576109ca8282610ba3565b506109d9565b6109d8610c23565b5b5050565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b610a15610c5f565b610a4b576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610a55610a0d565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ac5575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610abc9190610f8e565b60405180910390fd5b610ace816107ef565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b03610b3557806040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610b2c9190610f8e565b60405180910390fd5b80610b617f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b610ad1565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff1684604051610bcc9190611320565b5f60405180830381855af49150503d805f8114610c04576040519150601f19603f3d011682016040523d82523d5f602084013e610c09565b606091505b5091509150610c19858383610c7d565b9250505092915050565b5f341115610c5d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f610c686108c0565b5f0160089054906101000a900460ff16905090565b606082610c9257610c8d82610d0a565b610d02565b5f8251148015610cb857505f8473ffffffffffffffffffffffffffffffffffffffff163b145b15610cfa57836040517f9996b315000000000000000000000000000000000000000000000000000000008152600401610cf19190610f8e565b60405180910390fd5b819050610d03565b5b9392505050565b5f81511115610d1b57805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d8782610d5e565b9050919050565b610d9781610d7d565b8114610da1575f5ffd5b50565b5f81359050610db281610d8e565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e0682610dc0565b810181811067ffffffffffffffff82111715610e2557610e24610dd0565b5b80604052505050565b5f610e37610d4d565b9050610e438282610dfd565b919050565b5f67ffffffffffffffff821115610e6257610e61610dd0565b5b610e6b82610dc0565b9050602081019050919050565b828183375f83830152505050565b5f610e98610e9384610e48565b610e2e565b905082815260208101848484011115610eb457610eb3610dbc565b5b610ebf848285610e78565b509392505050565b5f82601f830112610edb57610eda610db8565b5b8135610eeb848260208601610e86565b91505092915050565b5f5f60408385031215610f0a57610f09610d56565b5b5f610f1785828601610da4565b925050602083013567ffffffffffffffff811115610f3857610f37610d5a565b5b610f4485828601610ec7565b9150509250929050565b5f819050919050565b610f6081610f4e565b82525050565b5f602082019050610f795f830184610f57565b92915050565b610f8881610d7d565b82525050565b5f602082019050610fa15f830184610f7f565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610fd982610fa7565b610fe38185610fb1565b9350610ff3818560208601610fc1565b610ffc81610dc0565b840191505092915050565b5f6020820190508181035f83015261101f8184610fcf565b905092915050565b5f67ffffffffffffffff82169050919050565b61104381611027565b811461104d575f5ffd5b50565b5f8135905061105e8161103a565b92915050565b61106d81610f4e565b8114611077575f5ffd5b50565b5f8135905061108881611064565b92915050565b5f5f5f5f608085870312156110a6576110a5610d56565b5b5f6110b387828801611050565b94505060206110c48782880161107a565b935050604085013567ffffffffffffffff8111156110e5576110e4610d5a565b5b6110f187828801610ec7565b925050606061110287828801611050565b91505092959194509250565b5f6020828403121561112357611122610d56565b5b5f61113084828501610da4565b91505092915050565b5f819050919050565b5f819050919050565b5f61116561116061115b84611139565b611142565b611027565b9050919050565b6111758161114b565b82525050565b5f60208201905061118e5f83018461116c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111cb82611027565b91506111d683611027565b9250828201905067ffffffffffffffff8111156111f6576111f5611194565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f611220826111fc565b61122a8185611206565b935061123a818560208601610fc1565b61124381610dc0565b840191505092915050565b61125781611027565b82525050565b5f6080820190506112705f830187610f57565b61127d6020830186610f7f565b818103604083015261128f8185611216565b905061129e606083018461124e565b95945050505050565b5f815190506112b581611064565b92915050565b5f602082840312156112d0576112cf610d56565b5b5f6112dd848285016112a7565b91505092915050565b5f81905092915050565b5f6112fa826111fc565b61130481856112e6565b9350611314818560208601610fc1565b80840191505092915050565b5f61132b82846112f0565b91508190509291505056fea164736f6c634300081c000a",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Caller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Shuttereventtriggerregistryv1.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) UPGRADEINTERFACEVERSION() (string, error) {
	return _Shuttereventtriggerregistryv1.Contract.UPGRADEINTERFACEVERSION(&_Shuttereventtriggerregistryv1.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1CallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Shuttereventtriggerregistryv1.Contract.UPGRADEINTERFACEVERSION(&_Shuttereventtriggerregistryv1.CallOpts)
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Shuttereventtriggerregistryv1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) ProxiableUUID() ([32]byte, error) {
	return _Shuttereventtriggerregistryv1.Contract.ProxiableUUID(&_Shuttereventtriggerregistryv1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _Shuttereventtriggerregistryv1.Contract.ProxiableUUID(&_Shuttereventtriggerregistryv1.CallOpts)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.UpgradeToAndCall(&_Shuttereventtriggerregistryv1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Shuttereventtriggerregistryv1.Contract.UpgradeToAndCall(&_Shuttereventtriggerregistryv1.TransactOpts, newImplementation, data)
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

// Shuttereventtriggerregistryv1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1UpgradedIterator struct {
	Event *Shuttereventtriggerregistryv1Upgraded // Event containing the contract specifics and raw log

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
func (it *Shuttereventtriggerregistryv1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Shuttereventtriggerregistryv1Upgraded)
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
		it.Event = new(Shuttereventtriggerregistryv1Upgraded)
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
func (it *Shuttereventtriggerregistryv1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Shuttereventtriggerregistryv1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Shuttereventtriggerregistryv1Upgraded represents a Upgraded event raised by the Shuttereventtriggerregistryv1 contract.
type Shuttereventtriggerregistryv1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Shuttereventtriggerregistryv1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistryv1UpgradedIterator{contract: _Shuttereventtriggerregistryv1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Shuttereventtriggerregistryv1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Shuttereventtriggerregistryv1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Shuttereventtriggerregistryv1Upgraded)
				if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Shuttereventtriggerregistryv1 *Shuttereventtriggerregistryv1Filterer) ParseUpgraded(log types.Log) (*Shuttereventtriggerregistryv1Upgraded, error) {
	event := new(Shuttereventtriggerregistryv1Upgraded)
	if err := _Shuttereventtriggerregistryv1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
