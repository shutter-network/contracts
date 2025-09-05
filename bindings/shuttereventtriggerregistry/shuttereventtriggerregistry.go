// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shuttereventtriggerregistry

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

// ShuttereventtriggerregistryMetaData contains all meta data concerning the Shuttereventtriggerregistry contract.
var ShuttereventtriggerregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registrations\",\"inputs\":[{\"name\":\"identity\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"triggerDefinitionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EventTriggerRegistered\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"triggerDefinition\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"ttl\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIdentityPrefix\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TTLTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TimestampInThePast\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610081575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100789190610196565b60405180910390fd5b6100908161009660201b60201c565b506101af565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018082610157565b9050919050565b61019081610176565b82525050565b5f6020820190506101a95f830184610187565b92915050565b610a94806101bc5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063715018a6146100595780638da5cb5b14610063578063d48b6d1e14610081578063da7c6a421461009d578063f2fde38b146100cf575b5f5ffd5b6100616100eb565b005b61006b6100fe565b6040516100789190610604565b60405180910390f35b61009b600480360381019061009691906107da565b610125565b005b6100b760048036038101906100b2919061085a565b6103a6565b6040516100c6939291906108a3565b60405180910390f35b6100e960048036038101906100e49190610902565b6103f2565b005b6100f3610476565b6100fc5f6104fd565b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b438167ffffffffffffffff161015610169576040517f84f8e55900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f1b83036101a4576040517f63a4021d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83336040516020016101b8929190610992565b6040516020818303038152906040528051906020012090505f60015f8381526020019081526020015f2090505f815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff16146102d4578267ffffffffffffffff16815f0160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1610610270576040517fb5f2184000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060010154846040516020016102869190610a1d565b60405160208183030381529060405280519060200120146102d3576040517f3a81d6fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b82815f0160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555085815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836040516020016103389190610a1d565b6040516020818303038152906040528051906020012081600101819055508567ffffffffffffffff167f06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5863387876040516103969493929190610a3d565b60405180910390a2505050505050565b6001602052805f5260405f205f91509050805f015f9054906101000a900467ffffffffffffffff1690805f0160089054906101000a900467ffffffffffffffff16908060010154905083565b6103fa610476565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361046a575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104619190610604565b60405180910390fd5b610473816104fd565b50565b61047e6105be565b73ffffffffffffffffffffffffffffffffffffffff1661049c6100fe565b73ffffffffffffffffffffffffffffffffffffffff16146104fb576104bf6105be565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016104f29190610604565b60405180910390fd5b565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f33905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105ee826105c5565b9050919050565b6105fe816105e4565b82525050565b5f6020820190506106175f8301846105f5565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b61064a8161062e565b8114610654575f5ffd5b50565b5f8135905061066581610641565b92915050565b5f819050919050565b61067d8161066b565b8114610687575f5ffd5b50565b5f8135905061069881610674565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6106ec826106a6565b810181811067ffffffffffffffff8211171561070b5761070a6106b6565b5b80604052505050565b5f61071d61061d565b905061072982826106e3565b919050565b5f67ffffffffffffffff821115610748576107476106b6565b5b610751826106a6565b9050602081019050919050565b828183375f83830152505050565b5f61077e6107798461072e565b610714565b90508281526020810184848401111561079a576107996106a2565b5b6107a584828561075e565b509392505050565b5f82601f8301126107c1576107c061069e565b5b81356107d184826020860161076c565b91505092915050565b5f5f5f5f608085870312156107f2576107f1610626565b5b5f6107ff87828801610657565b94505060206108108782880161068a565b935050604085013567ffffffffffffffff8111156108315761083061062a565b5b61083d878288016107ad565b925050606061084e87828801610657565b91505092959194509250565b5f6020828403121561086f5761086e610626565b5b5f61087c8482850161068a565b91505092915050565b61088e8161062e565b82525050565b61089d8161066b565b82525050565b5f6060820190506108b65f830186610885565b6108c36020830185610885565b6108d06040830184610894565b949350505050565b6108e1816105e4565b81146108eb575f5ffd5b50565b5f813590506108fc816108d8565b92915050565b5f6020828403121561091757610916610626565b5b5f610924848285016108ee565b91505092915050565b5f819050919050565b6109476109428261066b565b61092d565b82525050565b5f8160601b9050919050565b5f6109638261094d565b9050919050565b5f61097482610959565b9050919050565b61098c610987826105e4565b61096a565b82525050565b5f61099d8285610936565b6020820191506109ad828461097b565b6014820191508190509392505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6109ef826109bd565b6109f981856109c7565b9350610a098185602086016109d7565b610a12816106a6565b840191505092915050565b5f6020820190508181035f830152610a3581846109e5565b905092915050565b5f608082019050610a505f830187610894565b610a5d60208301866105f5565b8181036040830152610a6f81856109e5565b9050610a7e6060830184610885565b9594505050505056fea164736f6c634300081c000a",
}

// ShuttereventtriggerregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ShuttereventtriggerregistryMetaData.ABI instead.
var ShuttereventtriggerregistryABI = ShuttereventtriggerregistryMetaData.ABI

// ShuttereventtriggerregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShuttereventtriggerregistryMetaData.Bin instead.
var ShuttereventtriggerregistryBin = ShuttereventtriggerregistryMetaData.Bin

// DeployShuttereventtriggerregistry deploys a new Ethereum contract, binding an instance of Shuttereventtriggerregistry to it.
func DeployShuttereventtriggerregistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Shuttereventtriggerregistry, error) {
	parsed, err := ShuttereventtriggerregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShuttereventtriggerregistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shuttereventtriggerregistry{ShuttereventtriggerregistryCaller: ShuttereventtriggerregistryCaller{contract: contract}, ShuttereventtriggerregistryTransactor: ShuttereventtriggerregistryTransactor{contract: contract}, ShuttereventtriggerregistryFilterer: ShuttereventtriggerregistryFilterer{contract: contract}}, nil
}

// Shuttereventtriggerregistry is an auto generated Go binding around an Ethereum contract.
type Shuttereventtriggerregistry struct {
	ShuttereventtriggerregistryCaller     // Read-only binding to the contract
	ShuttereventtriggerregistryTransactor // Write-only binding to the contract
	ShuttereventtriggerregistryFilterer   // Log filterer for contract events
}

// ShuttereventtriggerregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShuttereventtriggerregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttereventtriggerregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShuttereventtriggerregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttereventtriggerregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShuttereventtriggerregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShuttereventtriggerregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShuttereventtriggerregistrySession struct {
	Contract     *Shuttereventtriggerregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ShuttereventtriggerregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShuttereventtriggerregistryCallerSession struct {
	Contract *ShuttereventtriggerregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ShuttereventtriggerregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShuttereventtriggerregistryTransactorSession struct {
	Contract     *ShuttereventtriggerregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ShuttereventtriggerregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShuttereventtriggerregistryRaw struct {
	Contract *Shuttereventtriggerregistry // Generic contract binding to access the raw methods on
}

// ShuttereventtriggerregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShuttereventtriggerregistryCallerRaw struct {
	Contract *ShuttereventtriggerregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ShuttereventtriggerregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShuttereventtriggerregistryTransactorRaw struct {
	Contract *ShuttereventtriggerregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShuttereventtriggerregistry creates a new instance of Shuttereventtriggerregistry, bound to a specific deployed contract.
func NewShuttereventtriggerregistry(address common.Address, backend bind.ContractBackend) (*Shuttereventtriggerregistry, error) {
	contract, err := bindShuttereventtriggerregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shuttereventtriggerregistry{ShuttereventtriggerregistryCaller: ShuttereventtriggerregistryCaller{contract: contract}, ShuttereventtriggerregistryTransactor: ShuttereventtriggerregistryTransactor{contract: contract}, ShuttereventtriggerregistryFilterer: ShuttereventtriggerregistryFilterer{contract: contract}}, nil
}

// NewShuttereventtriggerregistryCaller creates a new read-only instance of Shuttereventtriggerregistry, bound to a specific deployed contract.
func NewShuttereventtriggerregistryCaller(address common.Address, caller bind.ContractCaller) (*ShuttereventtriggerregistryCaller, error) {
	contract, err := bindShuttereventtriggerregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShuttereventtriggerregistryCaller{contract: contract}, nil
}

// NewShuttereventtriggerregistryTransactor creates a new write-only instance of Shuttereventtriggerregistry, bound to a specific deployed contract.
func NewShuttereventtriggerregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ShuttereventtriggerregistryTransactor, error) {
	contract, err := bindShuttereventtriggerregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShuttereventtriggerregistryTransactor{contract: contract}, nil
}

// NewShuttereventtriggerregistryFilterer creates a new log filterer instance of Shuttereventtriggerregistry, bound to a specific deployed contract.
func NewShuttereventtriggerregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ShuttereventtriggerregistryFilterer, error) {
	contract, err := bindShuttereventtriggerregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShuttereventtriggerregistryFilterer{contract: contract}, nil
}

// bindShuttereventtriggerregistry binds a generic wrapper to an already deployed contract.
func bindShuttereventtriggerregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShuttereventtriggerregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shuttereventtriggerregistry.Contract.ShuttereventtriggerregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.ShuttereventtriggerregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.ShuttereventtriggerregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shuttereventtriggerregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Shuttereventtriggerregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistrySession) Owner() (common.Address, error) {
	return _Shuttereventtriggerregistry.Contract.Owner(&_Shuttereventtriggerregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryCallerSession) Owner() (common.Address, error) {
	return _Shuttereventtriggerregistry.Contract.Owner(&_Shuttereventtriggerregistry.CallOpts)
}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 ttl, bytes32 triggerDefinitionHash)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryCaller) Registrations(opts *bind.CallOpts, identity [32]byte) (struct {
	Eon                   uint64
	Ttl                   uint64
	TriggerDefinitionHash [32]byte
}, error) {
	var out []interface{}
	err := _Shuttereventtriggerregistry.contract.Call(opts, &out, "registrations", identity)

	outstruct := new(struct {
		Eon                   uint64
		Ttl                   uint64
		TriggerDefinitionHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Eon = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Ttl = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.TriggerDefinitionHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 ttl, bytes32 triggerDefinitionHash)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistrySession) Registrations(identity [32]byte) (struct {
	Eon                   uint64
	Ttl                   uint64
	TriggerDefinitionHash [32]byte
}, error) {
	return _Shuttereventtriggerregistry.Contract.Registrations(&_Shuttereventtriggerregistry.CallOpts, identity)
}

// Registrations is a free data retrieval call binding the contract method 0xda7c6a42.
//
// Solidity: function registrations(bytes32 identity) view returns(uint64 eon, uint64 ttl, bytes32 triggerDefinitionHash)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryCallerSession) Registrations(identity [32]byte) (struct {
	Eon                   uint64
	Ttl                   uint64
	TriggerDefinitionHash [32]byte
}, error) {
	return _Shuttereventtriggerregistry.Contract.Registrations(&_Shuttereventtriggerregistry.CallOpts, identity)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactor) Register(opts *bind.TransactOpts, eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.contract.Transact(opts, "register", eon, identityPrefix, triggerDefinition, ttl)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistrySession) Register(eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.Register(&_Shuttereventtriggerregistry.TransactOpts, eon, identityPrefix, triggerDefinition, ttl)
}

// Register is a paid mutator transaction binding the contract method 0xd48b6d1e.
//
// Solidity: function register(uint64 eon, bytes32 identityPrefix, bytes triggerDefinition, uint64 ttl) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactorSession) Register(eon uint64, identityPrefix [32]byte, triggerDefinition []byte, ttl uint64) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.Register(&_Shuttereventtriggerregistry.TransactOpts, eon, identityPrefix, triggerDefinition, ttl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.RenounceOwnership(&_Shuttereventtriggerregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.RenounceOwnership(&_Shuttereventtriggerregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.TransferOwnership(&_Shuttereventtriggerregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Shuttereventtriggerregistry.Contract.TransferOwnership(&_Shuttereventtriggerregistry.TransactOpts, newOwner)
}

// ShuttereventtriggerregistryEventTriggerRegisteredIterator is returned from FilterEventTriggerRegistered and is used to iterate over the raw logs and unpacked data for EventTriggerRegistered events raised by the Shuttereventtriggerregistry contract.
type ShuttereventtriggerregistryEventTriggerRegisteredIterator struct {
	Event *ShuttereventtriggerregistryEventTriggerRegistered // Event containing the contract specifics and raw log

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
func (it *ShuttereventtriggerregistryEventTriggerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShuttereventtriggerregistryEventTriggerRegistered)
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
		it.Event = new(ShuttereventtriggerregistryEventTriggerRegistered)
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
func (it *ShuttereventtriggerregistryEventTriggerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShuttereventtriggerregistryEventTriggerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShuttereventtriggerregistryEventTriggerRegistered represents a EventTriggerRegistered event raised by the Shuttereventtriggerregistry contract.
type ShuttereventtriggerregistryEventTriggerRegistered struct {
	Eon               uint64
	IdentityPrefix    [32]byte
	Sender            common.Address
	TriggerDefinition []byte
	Ttl               uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEventTriggerRegistered is a free log retrieval operation binding the contract event 0x06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5.
//
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 ttl)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) FilterEventTriggerRegistered(opts *bind.FilterOpts, eon []uint64) (*ShuttereventtriggerregistryEventTriggerRegisteredIterator, error) {

	var eonRule []interface{}
	for _, eonItem := range eon {
		eonRule = append(eonRule, eonItem)
	}

	logs, sub, err := _Shuttereventtriggerregistry.contract.FilterLogs(opts, "EventTriggerRegistered", eonRule)
	if err != nil {
		return nil, err
	}
	return &ShuttereventtriggerregistryEventTriggerRegisteredIterator{contract: _Shuttereventtriggerregistry.contract, event: "EventTriggerRegistered", logs: logs, sub: sub}, nil
}

// WatchEventTriggerRegistered is a free log subscription operation binding the contract event 0x06809c10c5d53027eb51ca4ec6da8fbfd54d79f41cdefaab2cb142d10b039db5.
//
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 ttl)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) WatchEventTriggerRegistered(opts *bind.WatchOpts, sink chan<- *ShuttereventtriggerregistryEventTriggerRegistered, eon []uint64) (event.Subscription, error) {

	var eonRule []interface{}
	for _, eonItem := range eon {
		eonRule = append(eonRule, eonItem)
	}

	logs, sub, err := _Shuttereventtriggerregistry.contract.WatchLogs(opts, "EventTriggerRegistered", eonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShuttereventtriggerregistryEventTriggerRegistered)
				if err := _Shuttereventtriggerregistry.contract.UnpackLog(event, "EventTriggerRegistered", log); err != nil {
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
// Solidity: event EventTriggerRegistered(uint64 indexed eon, bytes32 identityPrefix, address sender, bytes triggerDefinition, uint64 ttl)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) ParseEventTriggerRegistered(log types.Log) (*ShuttereventtriggerregistryEventTriggerRegistered, error) {
	event := new(ShuttereventtriggerregistryEventTriggerRegistered)
	if err := _Shuttereventtriggerregistry.contract.UnpackLog(event, "EventTriggerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShuttereventtriggerregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Shuttereventtriggerregistry contract.
type ShuttereventtriggerregistryOwnershipTransferredIterator struct {
	Event *ShuttereventtriggerregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShuttereventtriggerregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShuttereventtriggerregistryOwnershipTransferred)
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
		it.Event = new(ShuttereventtriggerregistryOwnershipTransferred)
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
func (it *ShuttereventtriggerregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShuttereventtriggerregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShuttereventtriggerregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Shuttereventtriggerregistry contract.
type ShuttereventtriggerregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShuttereventtriggerregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shuttereventtriggerregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShuttereventtriggerregistryOwnershipTransferredIterator{contract: _Shuttereventtriggerregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShuttereventtriggerregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Shuttereventtriggerregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShuttereventtriggerregistryOwnershipTransferred)
				if err := _Shuttereventtriggerregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Shuttereventtriggerregistry *ShuttereventtriggerregistryFilterer) ParseOwnershipTransferred(log types.Log) (*ShuttereventtriggerregistryOwnershipTransferred, error) {
	event := new(ShuttereventtriggerregistryOwnershipTransferred)
	if err := _Shuttereventtriggerregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
