// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keybroadcastcontract

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

// KeybroadcastcontractMetaData contains all meta data concerning the Keybroadcastcontract contract.
var KeybroadcastcontractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keyperSetManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcastEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEonKey\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EonKeyBroadcast\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"key\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyHaveKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowed\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610b77380380610b77833981810160405281019061003191906100d5565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610100565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a48261007b565b9050919050565b6100b48161009a565b81146100be575f5ffd5b50565b5f815190506100cf816100ab565b92915050565b5f602082840312156100ea576100e9610077565b5b5f6100f7848285016100c1565b91505092915050565b610a6a8061010d5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80638a0b8b2814610038578063daade8e814610068575b5f5ffd5b610052600480360381019061004d91906103e4565b610084565b60405161005f919061047f565b60405180910390f35b610082600480360381019061007d91906105cb565b610138565b005b60605f5f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2080546100b590610652565b80601f01602080910402602001604051908101604052809291908181526020018280546100e190610652565b801561012c5780601f106101035761010080835404028352916020019161012c565b820191905f5260205f20905b81548152906001019060200180831161010f57829003601f168201915b50505050509050919050565b5f815103610172576040517f76d4e1e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8467ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2080546101a290610652565b905011156101dc576040517f58d819a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f90f3bed846040518263ffffffff1660e01b81526004016102379190610691565b602060405180830381865afa158015610252573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102769190610704565b90508073ffffffffffffffffffffffffffffffffffffffff1663cde1532d336040518263ffffffff1660e01b81526004016102b1919061073e565b602060405180830381865afa1580156102cc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102f0919061078c565b610326576040517f3d693ada00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f5f8567ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f2090816103579190610960565b507ff0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad37688383604051610389929190610a2f565b60405180910390a1505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff82169050919050565b6103c3816103a7565b81146103cd575f5ffd5b50565b5f813590506103de816103ba565b92915050565b5f602082840312156103f9576103f861039f565b5b5f610406848285016103d0565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6104518261040f565b61045b8185610419565b935061046b818560208601610429565b61047481610437565b840191505092915050565b5f6020820190508181035f8301526104978184610447565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104dd82610437565b810181811067ffffffffffffffff821117156104fc576104fb6104a7565b5b80604052505050565b5f61050e610396565b905061051a82826104d4565b919050565b5f67ffffffffffffffff821115610539576105386104a7565b5b61054282610437565b9050602081019050919050565b828183375f83830152505050565b5f61056f61056a8461051f565b610505565b90508281526020810184848401111561058b5761058a6104a3565b5b61059684828561054f565b509392505050565b5f82601f8301126105b2576105b161049f565b5b81356105c284826020860161055d565b91505092915050565b5f5f604083850312156105e1576105e061039f565b5b5f6105ee858286016103d0565b925050602083013567ffffffffffffffff81111561060f5761060e6103a3565b5b61061b8582860161059e565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061066957607f821691505b60208210810361067c5761067b610625565b5b50919050565b61068b816103a7565b82525050565b5f6020820190506106a45f830184610682565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106d3826106aa565b9050919050565b6106e3816106c9565b81146106ed575f5ffd5b50565b5f815190506106fe816106da565b92915050565b5f602082840312156107195761071861039f565b5b5f610726848285016106f0565b91505092915050565b610738816106c9565b82525050565b5f6020820190506107515f83018461072f565b92915050565b5f8115159050919050565b61076b81610757565b8114610775575f5ffd5b50565b5f8151905061078681610762565b92915050565b5f602082840312156107a1576107a061039f565b5b5f6107ae84828501610778565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026108137fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826107d8565b61081d86836107d8565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61086161085c61085784610835565b61083e565b610835565b9050919050565b5f819050919050565b61087a83610847565b61088e61088682610868565b8484546107e4565b825550505050565b5f5f905090565b6108a5610896565b6108b0818484610871565b505050565b5b818110156108d3576108c85f8261089d565b6001810190506108b6565b5050565b601f821115610918576108e9816107b7565b6108f2846107c9565b81016020851015610901578190505b61091561090d856107c9565b8301826108b5565b50505b505050565b5f82821c905092915050565b5f6109385f198460080261091d565b1980831691505092915050565b5f6109508383610929565b9150826002028217905092915050565b6109698261040f565b67ffffffffffffffff811115610982576109816104a7565b5b61098c8254610652565b6109978282856108d7565b5f60209050601f8311600181146109c8575f84156109b6578287015190505b6109c08582610945565b865550610a27565b601f1984166109d6866107b7565b5f5b828110156109fd578489015182556001820191506020850194506020810190506109d8565b86831015610a1a5784890151610a16601f891682610929565b8355505b6001600288020188555050505b505050505050565b5f604082019050610a425f830185610682565b8181036020830152610a548184610447565b9050939250505056fea164736f6c634300081c000a",
}

// KeybroadcastcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use KeybroadcastcontractMetaData.ABI instead.
var KeybroadcastcontractABI = KeybroadcastcontractMetaData.ABI

// KeybroadcastcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KeybroadcastcontractMetaData.Bin instead.
var KeybroadcastcontractBin = KeybroadcastcontractMetaData.Bin

// DeployKeybroadcastcontract deploys a new Ethereum contract, binding an instance of Keybroadcastcontract to it.
func DeployKeybroadcastcontract(auth *bind.TransactOpts, backend bind.ContractBackend, keyperSetManagerAddress common.Address) (common.Address, *types.Transaction, *Keybroadcastcontract, error) {
	parsed, err := KeybroadcastcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeybroadcastcontractBin), backend, keyperSetManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Keybroadcastcontract{KeybroadcastcontractCaller: KeybroadcastcontractCaller{contract: contract}, KeybroadcastcontractTransactor: KeybroadcastcontractTransactor{contract: contract}, KeybroadcastcontractFilterer: KeybroadcastcontractFilterer{contract: contract}}, nil
}

// Keybroadcastcontract is an auto generated Go binding around an Ethereum contract.
type Keybroadcastcontract struct {
	KeybroadcastcontractCaller     // Read-only binding to the contract
	KeybroadcastcontractTransactor // Write-only binding to the contract
	KeybroadcastcontractFilterer   // Log filterer for contract events
}

// KeybroadcastcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeybroadcastcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeybroadcastcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeybroadcastcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeybroadcastcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeybroadcastcontractSession struct {
	Contract     *Keybroadcastcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KeybroadcastcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeybroadcastcontractCallerSession struct {
	Contract *KeybroadcastcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// KeybroadcastcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeybroadcastcontractTransactorSession struct {
	Contract     *KeybroadcastcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// KeybroadcastcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeybroadcastcontractRaw struct {
	Contract *Keybroadcastcontract // Generic contract binding to access the raw methods on
}

// KeybroadcastcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeybroadcastcontractCallerRaw struct {
	Contract *KeybroadcastcontractCaller // Generic read-only contract binding to access the raw methods on
}

// KeybroadcastcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeybroadcastcontractTransactorRaw struct {
	Contract *KeybroadcastcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeybroadcastcontract creates a new instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontract(address common.Address, backend bind.ContractBackend) (*Keybroadcastcontract, error) {
	contract, err := bindKeybroadcastcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Keybroadcastcontract{KeybroadcastcontractCaller: KeybroadcastcontractCaller{contract: contract}, KeybroadcastcontractTransactor: KeybroadcastcontractTransactor{contract: contract}, KeybroadcastcontractFilterer: KeybroadcastcontractFilterer{contract: contract}}, nil
}

// NewKeybroadcastcontractCaller creates a new read-only instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractCaller(address common.Address, caller bind.ContractCaller) (*KeybroadcastcontractCaller, error) {
	contract, err := bindKeybroadcastcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractCaller{contract: contract}, nil
}

// NewKeybroadcastcontractTransactor creates a new write-only instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*KeybroadcastcontractTransactor, error) {
	contract, err := bindKeybroadcastcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractTransactor{contract: contract}, nil
}

// NewKeybroadcastcontractFilterer creates a new log filterer instance of Keybroadcastcontract, bound to a specific deployed contract.
func NewKeybroadcastcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*KeybroadcastcontractFilterer, error) {
	contract, err := bindKeybroadcastcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractFilterer{contract: contract}, nil
}

// bindKeybroadcastcontract binds a generic wrapper to an already deployed contract.
func bindKeybroadcastcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeybroadcastcontractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keybroadcastcontract *KeybroadcastcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.KeybroadcastcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Keybroadcastcontract *KeybroadcastcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Keybroadcastcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Keybroadcastcontract *KeybroadcastcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Keybroadcastcontract *KeybroadcastcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.contract.Transact(opts, method, params...)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractCaller) GetEonKey(opts *bind.CallOpts, eon uint64) ([]byte, error) {
	var out []interface{}
	err := _Keybroadcastcontract.contract.Call(opts, &out, "getEonKey", eon)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Keybroadcastcontract.Contract.GetEonKey(&_Keybroadcastcontract.CallOpts, eon)
}

// GetEonKey is a free data retrieval call binding the contract method 0x8a0b8b28.
//
// Solidity: function getEonKey(uint64 eon) view returns(bytes)
func (_Keybroadcastcontract *KeybroadcastcontractCallerSession) GetEonKey(eon uint64) ([]byte, error) {
	return _Keybroadcastcontract.Contract.GetEonKey(&_Keybroadcastcontract.CallOpts, eon)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractTransactor) BroadcastEonKey(opts *bind.TransactOpts, eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.contract.Transact(opts, "broadcastEonKey", eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.BroadcastEonKey(&_Keybroadcastcontract.TransactOpts, eon, key)
}

// BroadcastEonKey is a paid mutator transaction binding the contract method 0xdaade8e8.
//
// Solidity: function broadcastEonKey(uint64 eon, bytes key) returns()
func (_Keybroadcastcontract *KeybroadcastcontractTransactorSession) BroadcastEonKey(eon uint64, key []byte) (*types.Transaction, error) {
	return _Keybroadcastcontract.Contract.BroadcastEonKey(&_Keybroadcastcontract.TransactOpts, eon, key)
}

// KeybroadcastcontractEonKeyBroadcastIterator is returned from FilterEonKeyBroadcast and is used to iterate over the raw logs and unpacked data for EonKeyBroadcast events raised by the Keybroadcastcontract contract.
type KeybroadcastcontractEonKeyBroadcastIterator struct {
	Event *KeybroadcastcontractEonKeyBroadcast // Event containing the contract specifics and raw log

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
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeybroadcastcontractEonKeyBroadcast)
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
		it.Event = new(KeybroadcastcontractEonKeyBroadcast)
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
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeybroadcastcontractEonKeyBroadcastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeybroadcastcontractEonKeyBroadcast represents a EonKeyBroadcast event raised by the Keybroadcastcontract contract.
type KeybroadcastcontractEonKeyBroadcast struct {
	Eon uint64
	Key []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEonKeyBroadcast is a free log retrieval operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) FilterEonKeyBroadcast(opts *bind.FilterOpts) (*KeybroadcastcontractEonKeyBroadcastIterator, error) {

	logs, sub, err := _Keybroadcastcontract.contract.FilterLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return &KeybroadcastcontractEonKeyBroadcastIterator{contract: _Keybroadcastcontract.contract, event: "EonKeyBroadcast", logs: logs, sub: sub}, nil
}

// WatchEonKeyBroadcast is a free log subscription operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) WatchEonKeyBroadcast(opts *bind.WatchOpts, sink chan<- *KeybroadcastcontractEonKeyBroadcast) (event.Subscription, error) {

	logs, sub, err := _Keybroadcastcontract.contract.WatchLogs(opts, "EonKeyBroadcast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeybroadcastcontractEonKeyBroadcast)
				if err := _Keybroadcastcontract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
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

// ParseEonKeyBroadcast is a log parse operation binding the contract event 0xf0dbcf46bf98296dd97ce9cb1ef117ac6fd1b2f126741125433174d56dad3768.
//
// Solidity: event EonKeyBroadcast(uint64 eon, bytes key)
func (_Keybroadcastcontract *KeybroadcastcontractFilterer) ParseEonKeyBroadcast(log types.Log) (*KeybroadcastcontractEonKeyBroadcast, error) {
	event := new(KeybroadcastcontractEonKeyBroadcast)
	if err := _Keybroadcastcontract.contract.UnpackLog(event, "EonKeyBroadcast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
