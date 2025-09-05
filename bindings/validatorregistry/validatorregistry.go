// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorregistry

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

// IValidatorRegistryUpdate is an auto generated low-level Go binding around an user-defined struct.
type IValidatorRegistryUpdate struct {
	Message   []byte
	Signature []byte
}

// ValidatorregistryMetaData contains all meta data concerning the Validatorregistry contract.
var ValidatorregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNumUpdates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUpdate\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Update\",\"components\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Updated\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061095b8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003e575f3560e01c8062c7c0191461004257806332cb25be146100605780633f37dce214610090575b5f5ffd5b61004a6100ac565b60405161005791906102e6565b60405180910390f35b61007a6004803603810190610075919061033a565b6100b7565b6040516100879190610416565b60405180910390f35b6100aa60048036038101906100a59190610562565b61020f565b005b5f5f80549050905090565b6100bf6102b4565b5f82815481106100d2576100d16105d8565b5b905f5260205f2090600202016040518060400160405290815f820180546100f890610632565b80601f016020809104026020016040519081016040528092919081815260200182805461012490610632565b801561016f5780601f106101465761010080835404028352916020019161016f565b820191905f5260205f20905b81548152906001019060200180831161015257829003601f168201915b5050505050815260200160018201805461018890610632565b80601f01602080910402602001604051908101604052809291908181526020018280546101b490610632565b80156101ff5780601f106101d6576101008083540402835291602001916101ff565b820191905f5260205f20905b8154815290600101906020018083116101e257829003601f168201915b5050505050815250509050919050565b5f604051806040016040528084815260200183815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01908161025e9190610802565b5060208201518160010190816102749190610802565b5050507f9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a8782826040516102a8929190610919565b60405180910390a15050565b604051806040016040528060608152602001606081525090565b5f819050919050565b6102e0816102ce565b82525050565b5f6020820190506102f95f8301846102d7565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b610319816102ce565b8114610323575f5ffd5b50565b5f8135905061033481610310565b92915050565b5f6020828403121561034f5761034e610308565b5b5f61035c84828501610326565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6103a782610365565b6103b1818561036f565b93506103c181856020860161037f565b6103ca8161038d565b840191505092915050565b5f604083015f8301518482035f8601526103ef828261039d565b91505060208301518482036020860152610409828261039d565b9150508091505092915050565b5f6020820190508181035f83015261042e81846103d5565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6104748261038d565b810181811067ffffffffffffffff821117156104935761049261043e565b5b80604052505050565b5f6104a56102ff565b90506104b1828261046b565b919050565b5f67ffffffffffffffff8211156104d0576104cf61043e565b5b6104d98261038d565b9050602081019050919050565b828183375f83830152505050565b5f610506610501846104b6565b61049c565b9050828152602081018484840111156105225761052161043a565b5b61052d8482856104e6565b509392505050565b5f82601f83011261054957610548610436565b5b81356105598482602086016104f4565b91505092915050565b5f5f6040838503121561057857610577610308565b5b5f83013567ffffffffffffffff8111156105955761059461030c565b5b6105a185828601610535565b925050602083013567ffffffffffffffff8111156105c2576105c161030c565b5b6105ce85828601610535565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061064957607f821691505b60208210810361065c5761065b610605565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106be7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610683565b6106c88683610683565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6107036106fe6106f9846102ce565b6106e0565b6102ce565b9050919050565b5f819050919050565b61071c836106e9565b6107306107288261070a565b84845461068f565b825550505050565b5f5f905090565b610747610738565b610752818484610713565b505050565b5b818110156107755761076a5f8261073f565b600181019050610758565b5050565b601f8211156107ba5761078b81610662565b61079484610674565b810160208510156107a3578190505b6107b76107af85610674565b830182610757565b50505b505050565b5f82821c905092915050565b5f6107da5f19846008026107bf565b1980831691505092915050565b5f6107f283836107cb565b9150826002028217905092915050565b61080b82610365565b67ffffffffffffffff8111156108245761082361043e565b5b61082e8254610632565b610839828285610779565b5f60209050601f83116001811461086a575f8415610858578287015190505b61086285826107e7565b8655506108c9565b601f19841661087886610662565b5f5b8281101561089f5784890151825560018201915060208501945060208101905061087a565b868310156108bc57848901516108b8601f8916826107cb565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b5f6108eb82610365565b6108f581856108d1565b935061090581856020860161037f565b61090e8161038d565b840191505092915050565b5f6040820190508181035f83015261093181856108e1565b9050818103602083015261094581846108e1565b9050939250505056fea164736f6c634300081c000a",
}

// ValidatorregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorregistryMetaData.ABI instead.
var ValidatorregistryABI = ValidatorregistryMetaData.ABI

// ValidatorregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorregistryMetaData.Bin instead.
var ValidatorregistryBin = ValidatorregistryMetaData.Bin

// DeployValidatorregistry deploys a new Ethereum contract, binding an instance of Validatorregistry to it.
func DeployValidatorregistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validatorregistry, error) {
	parsed, err := ValidatorregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorregistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validatorregistry{ValidatorregistryCaller: ValidatorregistryCaller{contract: contract}, ValidatorregistryTransactor: ValidatorregistryTransactor{contract: contract}, ValidatorregistryFilterer: ValidatorregistryFilterer{contract: contract}}, nil
}

// Validatorregistry is an auto generated Go binding around an Ethereum contract.
type Validatorregistry struct {
	ValidatorregistryCaller     // Read-only binding to the contract
	ValidatorregistryTransactor // Write-only binding to the contract
	ValidatorregistryFilterer   // Log filterer for contract events
}

// ValidatorregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorregistrySession struct {
	Contract     *Validatorregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorregistryCallerSession struct {
	Contract *ValidatorregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorregistryTransactorSession struct {
	Contract     *ValidatorregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorregistryRaw struct {
	Contract *Validatorregistry // Generic contract binding to access the raw methods on
}

// ValidatorregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorregistryCallerRaw struct {
	Contract *ValidatorregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactorRaw struct {
	Contract *ValidatorregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorregistry creates a new instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistry(address common.Address, backend bind.ContractBackend) (*Validatorregistry, error) {
	contract, err := bindValidatorregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorregistry{ValidatorregistryCaller: ValidatorregistryCaller{contract: contract}, ValidatorregistryTransactor: ValidatorregistryTransactor{contract: contract}, ValidatorregistryFilterer: ValidatorregistryFilterer{contract: contract}}, nil
}

// NewValidatorregistryCaller creates a new read-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorregistryCaller, error) {
	contract, err := bindValidatorregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryCaller{contract: contract}, nil
}

// NewValidatorregistryTransactor creates a new write-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorregistryTransactor, error) {
	contract, err := bindValidatorregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryTransactor{contract: contract}, nil
}

// NewValidatorregistryFilterer creates a new log filterer instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorregistryFilterer, error) {
	contract, err := bindValidatorregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryFilterer{contract: contract}, nil
}

// bindValidatorregistry binds a generic wrapper to an already deployed contract.
func bindValidatorregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.ValidatorregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transact(opts, method, params...)
}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) GetNumUpdates(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getNumUpdates")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) GetNumUpdates() (*big.Int, error) {
	return _Validatorregistry.Contract.GetNumUpdates(&_Validatorregistry.CallOpts)
}

// GetNumUpdates is a free data retrieval call binding the contract method 0x00c7c019.
//
// Solidity: function getNumUpdates() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) GetNumUpdates() (*big.Int, error) {
	return _Validatorregistry.Contract.GetNumUpdates(&_Validatorregistry.CallOpts)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistryCaller) GetUpdate(opts *bind.CallOpts, i *big.Int) (IValidatorRegistryUpdate, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getUpdate", i)

	if err != nil {
		return *new(IValidatorRegistryUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IValidatorRegistryUpdate)).(*IValidatorRegistryUpdate)

	return out0, err

}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistrySession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Validatorregistry.Contract.GetUpdate(&_Validatorregistry.CallOpts, i)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 i) view returns((bytes,bytes))
func (_Validatorregistry *ValidatorregistryCallerSession) GetUpdate(i *big.Int) (IValidatorRegistryUpdate, error) {
	return _Validatorregistry.Contract.GetUpdate(&_Validatorregistry.CallOpts, i)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistryTransactor) Update(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "update", message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistrySession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Update(&_Validatorregistry.TransactOpts, message, signature)
}

// Update is a paid mutator transaction binding the contract method 0x3f37dce2.
//
// Solidity: function update(bytes message, bytes signature) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) Update(message []byte, signature []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Update(&_Validatorregistry.TransactOpts, message, signature)
}

// ValidatorregistryUpdatedIterator is returned from FilterUpdated and is used to iterate over the raw logs and unpacked data for Updated events raised by the Validatorregistry contract.
type ValidatorregistryUpdatedIterator struct {
	Event *ValidatorregistryUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryUpdated)
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
		it.Event = new(ValidatorregistryUpdated)
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
func (it *ValidatorregistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryUpdated represents a Updated event raised by the Validatorregistry contract.
type ValidatorregistryUpdated struct {
	Message   []byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdated is a free log retrieval operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) FilterUpdated(opts *bind.FilterOpts) (*ValidatorregistryUpdatedIterator, error) {

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryUpdatedIterator{contract: _Validatorregistry.contract, event: "Updated", logs: logs, sub: sub}, nil
}

// WatchUpdated is a free log subscription operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) WatchUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorregistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryUpdated)
				if err := _Validatorregistry.contract.UnpackLog(event, "Updated", log); err != nil {
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

// ParseUpdated is a log parse operation binding the contract event 0x9796f15c93411b364b7f09bb591d0f77aa4dc399cf0481b8de1c3ce9f10a3a87.
//
// Solidity: event Updated(bytes message, bytes signature)
func (_Validatorregistry *ValidatorregistryFilterer) ParseUpdated(log types.Log) (*ValidatorregistryUpdated, error) {
	event := new(ValidatorregistryUpdated)
	if err := _Validatorregistry.contract.UnpackLog(event, "Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
