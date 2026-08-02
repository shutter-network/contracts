// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecieskeyregistry

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

// EcieskeyregistryMetaData contains all meta data concerning the Ecieskeyregistry contract.
var EcieskeyregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getKey\",\"inputs\":[{\"name\":\"keyper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperAt\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeyperCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyperSetManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractKeyperSetManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerKey\",\"inputs\":[{\"name\":\"keyperSetIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"keyperIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"eciesPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"KeyRegistered\",\"inputs\":[{\"name\":\"keyper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"eciesPublicKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyECIESPublicKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyperSetNotFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAMember\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051610f04380380610f04833981810160405281019061003191906100c9565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f4565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b6100a88161008e565b81146100b2575f5ffd5b50565b5f815190506100c38161009f565b92915050565b5f602082840312156100de576100dd61006b565b5b5f6100eb848285016100b5565b91505092915050565b608051610df16101135f395f81816101fb01526102200152610df15ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806393790f44146100595780639ce385a214610089578063d4b2cedd146100b9578063df713e19146100d7578063ecad263c146100f3575b5f5ffd5b610073600480360381019061006e91906106b0565b610111565b604051610080919061074b565b60405180910390f35b6100a3600480360381019061009e919061079e565b6101de565b6040516100b091906107d8565b60405180910390f35b6100c16101f9565b6040516100ce919061084c565b60405180910390f35b6100f160048036038101906100ec9190610903565b61021d565b005b6100fb61052b565b6040516101089190610983565b60405180910390f35b606060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805461015b906109c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610187906109c9565b80156101d25780601f106101a9576101008083540402835291602001916101d2565b820191905f5260205f20905b8154815290600101906020018083116101b557829003601f168201915b50505050509050919050565b5f6101f2825f61053a90919063ffffffff16565b9050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f90f3bed866040518263ffffffff1660e01b81526004016102779190610a08565b602060405180830381865afa158015610292573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102b69190610a35565b90508073ffffffffffffffffffffffffffffffffffffffff16638d4e40836040518163ffffffff1660e01b8152600401602060405180830381865afa158015610301573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103259190610a95565b61035b576040517feac631aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16632e8e6cad866040518263ffffffff1660e01b81526004016103ab9190610a08565b602060405180830381865afa1580156103c6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ea9190610a35565b73ffffffffffffffffffffffffffffffffffffffff1614610437576040517f2818e88800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8383905003610473576040517f0f154c2800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610486335f61055190919063ffffffff16565b50828260025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2091826104d3929190610c8e565b503373ffffffffffffffffffffffffffffffffffffffff167f9b368ac41b4759b00feab16152db5d57270846c8d258ce8ab77ea58800b288f2848460405161051c929190610d95565b60405180910390a25050505050565b5f6105355f61057e565b905090565b5f610547835f0183610591565b5f1c905092915050565b5f610576835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6105b8565b905092915050565b5f61058a825f0161061f565b9050919050565b5f825f0182815481106105a7576105a6610db7565b5b905f5260205f200154905092915050565b5f6105c3838361062e565b61061557825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050610619565b5f90505b92915050565b5f815f01805490509050919050565b5f5f836001015f8481526020019081526020015f20541415905092915050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61067f82610656565b9050919050565b61068f81610675565b8114610699575f5ffd5b50565b5f813590506106aa81610686565b92915050565b5f602082840312156106c5576106c461064e565b5b5f6106d28482850161069c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61071d826106db565b61072781856106e5565b93506107378185602086016106f5565b61074081610703565b840191505092915050565b5f6020820190508181035f8301526107638184610713565b905092915050565b5f819050919050565b61077d8161076b565b8114610787575f5ffd5b50565b5f8135905061079881610774565b92915050565b5f602082840312156107b3576107b261064e565b5b5f6107c08482850161078a565b91505092915050565b6107d281610675565b82525050565b5f6020820190506107eb5f8301846107c9565b92915050565b5f819050919050565b5f61081461080f61080a84610656565b6107f1565b610656565b9050919050565b5f610825826107fa565b9050919050565b5f6108368261081b565b9050919050565b6108468161082c565b82525050565b5f60208201905061085f5f83018461083d565b92915050565b5f67ffffffffffffffff82169050919050565b61088181610865565b811461088b575f5ffd5b50565b5f8135905061089c81610878565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126108c3576108c26108a2565b5b8235905067ffffffffffffffff8111156108e0576108df6108a6565b5b6020830191508360018202830111156108fc576108fb6108aa565b5b9250929050565b5f5f5f5f6060858703121561091b5761091a61064e565b5b5f6109288782880161088e565b94505060206109398782880161088e565b935050604085013567ffffffffffffffff81111561095a57610959610652565b5b610966878288016108ae565b925092505092959194509250565b61097d8161076b565b82525050565b5f6020820190506109965f830184610974565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806109e057607f821691505b6020821081036109f3576109f261099c565b5b50919050565b610a0281610865565b82525050565b5f602082019050610a1b5f8301846109f9565b92915050565b5f81519050610a2f81610686565b92915050565b5f60208284031215610a4a57610a4961064e565b5b5f610a5784828501610a21565b91505092915050565b5f8115159050919050565b610a7481610a60565b8114610a7e575f5ffd5b50565b5f81519050610a8f81610a6b565b92915050565b5f60208284031215610aaa57610aa961064e565b5b5f610ab784828501610a81565b91505092915050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610b537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b18565b610b5d8683610b18565b95508019841693508086168417925050509392505050565b5f610b8f610b8a610b858461076b565b6107f1565b61076b565b9050919050565b5f819050919050565b610ba883610b75565b610bbc610bb482610b96565b848454610b24565b825550505050565b5f5f905090565b610bd3610bc4565b610bde818484610b9f565b505050565b5b81811015610c0157610bf65f82610bcb565b600181019050610be4565b5050565b601f821115610c4657610c1781610af7565b610c2084610b09565b81016020851015610c2f578190505b610c43610c3b85610b09565b830182610be3565b50505b505050565b5f82821c905092915050565b5f610c665f1984600802610c4b565b1980831691505092915050565b5f610c7e8383610c57565b9150826002028217905092915050565b610c988383610ac0565b67ffffffffffffffff811115610cb157610cb0610aca565b5b610cbb82546109c9565b610cc6828285610c05565b5f601f831160018114610cf3575f8415610ce1578287013590505b610ceb8582610c73565b865550610d52565b601f198416610d0186610af7565b5f5b82811015610d2857848901358255600182019150602085019450602081019050610d03565b86831015610d455784890135610d41601f891682610c57565b8355505b6001600288020188555050505b50505050505050565b828183375f83830152505050565b5f610d7483856106e5565b9350610d81838584610d5b565b610d8a83610703565b840190509392505050565b5f6020820190508181035f830152610dae818486610d69565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c634300081c000a",
}

// EcieskeyregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EcieskeyregistryMetaData.ABI instead.
var EcieskeyregistryABI = EcieskeyregistryMetaData.ABI

// EcieskeyregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EcieskeyregistryMetaData.Bin instead.
var EcieskeyregistryBin = EcieskeyregistryMetaData.Bin

// DeployEcieskeyregistry deploys a new Ethereum contract, binding an instance of Ecieskeyregistry to it.
func DeployEcieskeyregistry(auth *bind.TransactOpts, backend bind.ContractBackend, keyperSetManagerAddress common.Address) (common.Address, *types.Transaction, *Ecieskeyregistry, error) {
	parsed, err := EcieskeyregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EcieskeyregistryBin), backend, keyperSetManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ecieskeyregistry{EcieskeyregistryCaller: EcieskeyregistryCaller{contract: contract}, EcieskeyregistryTransactor: EcieskeyregistryTransactor{contract: contract}, EcieskeyregistryFilterer: EcieskeyregistryFilterer{contract: contract}}, nil
}

// Ecieskeyregistry is an auto generated Go binding around an Ethereum contract.
type Ecieskeyregistry struct {
	EcieskeyregistryCaller     // Read-only binding to the contract
	EcieskeyregistryTransactor // Write-only binding to the contract
	EcieskeyregistryFilterer   // Log filterer for contract events
}

// EcieskeyregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcieskeyregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcieskeyregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcieskeyregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcieskeyregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcieskeyregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcieskeyregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcieskeyregistrySession struct {
	Contract     *Ecieskeyregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcieskeyregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcieskeyregistryCallerSession struct {
	Contract *EcieskeyregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EcieskeyregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcieskeyregistryTransactorSession struct {
	Contract     *EcieskeyregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EcieskeyregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcieskeyregistryRaw struct {
	Contract *Ecieskeyregistry // Generic contract binding to access the raw methods on
}

// EcieskeyregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcieskeyregistryCallerRaw struct {
	Contract *EcieskeyregistryCaller // Generic read-only contract binding to access the raw methods on
}

// EcieskeyregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcieskeyregistryTransactorRaw struct {
	Contract *EcieskeyregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcieskeyregistry creates a new instance of Ecieskeyregistry, bound to a specific deployed contract.
func NewEcieskeyregistry(address common.Address, backend bind.ContractBackend) (*Ecieskeyregistry, error) {
	contract, err := bindEcieskeyregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ecieskeyregistry{EcieskeyregistryCaller: EcieskeyregistryCaller{contract: contract}, EcieskeyregistryTransactor: EcieskeyregistryTransactor{contract: contract}, EcieskeyregistryFilterer: EcieskeyregistryFilterer{contract: contract}}, nil
}

// NewEcieskeyregistryCaller creates a new read-only instance of Ecieskeyregistry, bound to a specific deployed contract.
func NewEcieskeyregistryCaller(address common.Address, caller bind.ContractCaller) (*EcieskeyregistryCaller, error) {
	contract, err := bindEcieskeyregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcieskeyregistryCaller{contract: contract}, nil
}

// NewEcieskeyregistryTransactor creates a new write-only instance of Ecieskeyregistry, bound to a specific deployed contract.
func NewEcieskeyregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EcieskeyregistryTransactor, error) {
	contract, err := bindEcieskeyregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcieskeyregistryTransactor{contract: contract}, nil
}

// NewEcieskeyregistryFilterer creates a new log filterer instance of Ecieskeyregistry, bound to a specific deployed contract.
func NewEcieskeyregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EcieskeyregistryFilterer, error) {
	contract, err := bindEcieskeyregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcieskeyregistryFilterer{contract: contract}, nil
}

// bindEcieskeyregistry binds a generic wrapper to an already deployed contract.
func bindEcieskeyregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcieskeyregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecieskeyregistry *EcieskeyregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecieskeyregistry.Contract.EcieskeyregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecieskeyregistry *EcieskeyregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.EcieskeyregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecieskeyregistry *EcieskeyregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.EcieskeyregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ecieskeyregistry *EcieskeyregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ecieskeyregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ecieskeyregistry *EcieskeyregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ecieskeyregistry *EcieskeyregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.contract.Transact(opts, method, params...)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address keyper) view returns(bytes)
func (_Ecieskeyregistry *EcieskeyregistryCaller) GetKey(opts *bind.CallOpts, keyper common.Address) ([]byte, error) {
	var out []interface{}
	err := _Ecieskeyregistry.contract.Call(opts, &out, "getKey", keyper)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address keyper) view returns(bytes)
func (_Ecieskeyregistry *EcieskeyregistrySession) GetKey(keyper common.Address) ([]byte, error) {
	return _Ecieskeyregistry.Contract.GetKey(&_Ecieskeyregistry.CallOpts, keyper)
}

// GetKey is a free data retrieval call binding the contract method 0x93790f44.
//
// Solidity: function getKey(address keyper) view returns(bytes)
func (_Ecieskeyregistry *EcieskeyregistryCallerSession) GetKey(keyper common.Address) ([]byte, error) {
	return _Ecieskeyregistry.Contract.GetKey(&_Ecieskeyregistry.CallOpts, keyper)
}

// GetKeyperAt is a free data retrieval call binding the contract method 0x9ce385a2.
//
// Solidity: function getKeyperAt(uint256 index) view returns(address)
func (_Ecieskeyregistry *EcieskeyregistryCaller) GetKeyperAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ecieskeyregistry.contract.Call(opts, &out, "getKeyperAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKeyperAt is a free data retrieval call binding the contract method 0x9ce385a2.
//
// Solidity: function getKeyperAt(uint256 index) view returns(address)
func (_Ecieskeyregistry *EcieskeyregistrySession) GetKeyperAt(index *big.Int) (common.Address, error) {
	return _Ecieskeyregistry.Contract.GetKeyperAt(&_Ecieskeyregistry.CallOpts, index)
}

// GetKeyperAt is a free data retrieval call binding the contract method 0x9ce385a2.
//
// Solidity: function getKeyperAt(uint256 index) view returns(address)
func (_Ecieskeyregistry *EcieskeyregistryCallerSession) GetKeyperAt(index *big.Int) (common.Address, error) {
	return _Ecieskeyregistry.Contract.GetKeyperAt(&_Ecieskeyregistry.CallOpts, index)
}

// GetKeyperCount is a free data retrieval call binding the contract method 0xecad263c.
//
// Solidity: function getKeyperCount() view returns(uint256)
func (_Ecieskeyregistry *EcieskeyregistryCaller) GetKeyperCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ecieskeyregistry.contract.Call(opts, &out, "getKeyperCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeyperCount is a free data retrieval call binding the contract method 0xecad263c.
//
// Solidity: function getKeyperCount() view returns(uint256)
func (_Ecieskeyregistry *EcieskeyregistrySession) GetKeyperCount() (*big.Int, error) {
	return _Ecieskeyregistry.Contract.GetKeyperCount(&_Ecieskeyregistry.CallOpts)
}

// GetKeyperCount is a free data retrieval call binding the contract method 0xecad263c.
//
// Solidity: function getKeyperCount() view returns(uint256)
func (_Ecieskeyregistry *EcieskeyregistryCallerSession) GetKeyperCount() (*big.Int, error) {
	return _Ecieskeyregistry.Contract.GetKeyperCount(&_Ecieskeyregistry.CallOpts)
}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Ecieskeyregistry *EcieskeyregistryCaller) KeyperSetManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ecieskeyregistry.contract.Call(opts, &out, "keyperSetManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Ecieskeyregistry *EcieskeyregistrySession) KeyperSetManager() (common.Address, error) {
	return _Ecieskeyregistry.Contract.KeyperSetManager(&_Ecieskeyregistry.CallOpts)
}

// KeyperSetManager is a free data retrieval call binding the contract method 0xd4b2cedd.
//
// Solidity: function keyperSetManager() view returns(address)
func (_Ecieskeyregistry *EcieskeyregistryCallerSession) KeyperSetManager() (common.Address, error) {
	return _Ecieskeyregistry.Contract.KeyperSetManager(&_Ecieskeyregistry.CallOpts)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xdf713e19.
//
// Solidity: function registerKey(uint64 keyperSetIndex, uint64 keyperIndex, bytes eciesPublicKey) returns()
func (_Ecieskeyregistry *EcieskeyregistryTransactor) RegisterKey(opts *bind.TransactOpts, keyperSetIndex uint64, keyperIndex uint64, eciesPublicKey []byte) (*types.Transaction, error) {
	return _Ecieskeyregistry.contract.Transact(opts, "registerKey", keyperSetIndex, keyperIndex, eciesPublicKey)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xdf713e19.
//
// Solidity: function registerKey(uint64 keyperSetIndex, uint64 keyperIndex, bytes eciesPublicKey) returns()
func (_Ecieskeyregistry *EcieskeyregistrySession) RegisterKey(keyperSetIndex uint64, keyperIndex uint64, eciesPublicKey []byte) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.RegisterKey(&_Ecieskeyregistry.TransactOpts, keyperSetIndex, keyperIndex, eciesPublicKey)
}

// RegisterKey is a paid mutator transaction binding the contract method 0xdf713e19.
//
// Solidity: function registerKey(uint64 keyperSetIndex, uint64 keyperIndex, bytes eciesPublicKey) returns()
func (_Ecieskeyregistry *EcieskeyregistryTransactorSession) RegisterKey(keyperSetIndex uint64, keyperIndex uint64, eciesPublicKey []byte) (*types.Transaction, error) {
	return _Ecieskeyregistry.Contract.RegisterKey(&_Ecieskeyregistry.TransactOpts, keyperSetIndex, keyperIndex, eciesPublicKey)
}

// EcieskeyregistryKeyRegisteredIterator is returned from FilterKeyRegistered and is used to iterate over the raw logs and unpacked data for KeyRegistered events raised by the Ecieskeyregistry contract.
type EcieskeyregistryKeyRegisteredIterator struct {
	Event *EcieskeyregistryKeyRegistered // Event containing the contract specifics and raw log

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
func (it *EcieskeyregistryKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcieskeyregistryKeyRegistered)
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
		it.Event = new(EcieskeyregistryKeyRegistered)
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
func (it *EcieskeyregistryKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcieskeyregistryKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcieskeyregistryKeyRegistered represents a KeyRegistered event raised by the Ecieskeyregistry contract.
type EcieskeyregistryKeyRegistered struct {
	Keyper         common.Address
	EciesPublicKey []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyRegistered is a free log retrieval operation binding the contract event 0x9b368ac41b4759b00feab16152db5d57270846c8d258ce8ab77ea58800b288f2.
//
// Solidity: event KeyRegistered(address indexed keyper, bytes eciesPublicKey)
func (_Ecieskeyregistry *EcieskeyregistryFilterer) FilterKeyRegistered(opts *bind.FilterOpts, keyper []common.Address) (*EcieskeyregistryKeyRegisteredIterator, error) {

	var keyperRule []interface{}
	for _, keyperItem := range keyper {
		keyperRule = append(keyperRule, keyperItem)
	}

	logs, sub, err := _Ecieskeyregistry.contract.FilterLogs(opts, "KeyRegistered", keyperRule)
	if err != nil {
		return nil, err
	}
	return &EcieskeyregistryKeyRegisteredIterator{contract: _Ecieskeyregistry.contract, event: "KeyRegistered", logs: logs, sub: sub}, nil
}

// WatchKeyRegistered is a free log subscription operation binding the contract event 0x9b368ac41b4759b00feab16152db5d57270846c8d258ce8ab77ea58800b288f2.
//
// Solidity: event KeyRegistered(address indexed keyper, bytes eciesPublicKey)
func (_Ecieskeyregistry *EcieskeyregistryFilterer) WatchKeyRegistered(opts *bind.WatchOpts, sink chan<- *EcieskeyregistryKeyRegistered, keyper []common.Address) (event.Subscription, error) {

	var keyperRule []interface{}
	for _, keyperItem := range keyper {
		keyperRule = append(keyperRule, keyperItem)
	}

	logs, sub, err := _Ecieskeyregistry.contract.WatchLogs(opts, "KeyRegistered", keyperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcieskeyregistryKeyRegistered)
				if err := _Ecieskeyregistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
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

// ParseKeyRegistered is a log parse operation binding the contract event 0x9b368ac41b4759b00feab16152db5d57270846c8d258ce8ab77ea58800b288f2.
//
// Solidity: event KeyRegistered(address indexed keyper, bytes eciesPublicKey)
func (_Ecieskeyregistry *EcieskeyregistryFilterer) ParseKeyRegistered(log types.Log) (*EcieskeyregistryKeyRegistered, error) {
	event := new(EcieskeyregistryKeyRegistered)
	if err := _Ecieskeyregistry.contract.UnpackLog(event, "KeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
