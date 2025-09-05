// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sequencer

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

// SequencerMetaData contains all meta data concerning the Sequencer contract.
var SequencerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getTxCountForEon\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitDecryptionProgress\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitEncryptedTransaction\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DecryptionProgressSubmitted\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionSubmitted\",\"inputs\":[{\"name\":\"eon\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"txIndex\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"identityPrefix\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encryptedTransaction\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b506107528061001c5f395ff3fe608060405260043610610033575f3560e01c80632d32522e146100375780636a69d2e11461005f5780637bbd164b1461007b575b5f5ffd5b348015610042575f5ffd5b5061005d6004803603810190610058919061039c565b6100b7565b005b61007960048036038101906100749190610486565b6100f1565b005b348015610086575f5ffd5b506100a1600480360381019061009c9190610506565b61020f565b6040516100ae9190610540565b60405180910390f35b7fa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59816040516100e691906105b9565b60405180910390a150565b80486100fd9190610606565b341015610136576040517f025dbdd400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff16905060018161017e9190610647565b5f5f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fa7f1b5467be46c45249fb93063cceef96c63ddad03819246bc7770e32d4f5b7d858286338787604051610200969594939291906106df565b60405180910390a15050505050565b5f5f5f8367ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff169050919050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102ae82610268565b810181811067ffffffffffffffff821117156102cd576102cc610278565b5b80604052505050565b5f6102df61024f565b90506102eb82826102a5565b919050565b5f67ffffffffffffffff82111561030a57610309610278565b5b61031382610268565b9050602081019050919050565b828183375f83830152505050565b5f61034061033b846102f0565b6102d6565b90508281526020810184848401111561035c5761035b610264565b5b610367848285610320565b509392505050565b5f82601f83011261038357610382610260565b5b813561039384826020860161032e565b91505092915050565b5f602082840312156103b1576103b0610258565b5b5f82013567ffffffffffffffff8111156103ce576103cd61025c565b5b6103da8482850161036f565b91505092915050565b5f67ffffffffffffffff82169050919050565b6103ff816103e3565b8114610409575f5ffd5b50565b5f8135905061041a816103f6565b92915050565b5f819050919050565b61043281610420565b811461043c575f5ffd5b50565b5f8135905061044d81610429565b92915050565b5f819050919050565b61046581610453565b811461046f575f5ffd5b50565b5f813590506104808161045c565b92915050565b5f5f5f5f6080858703121561049e5761049d610258565b5b5f6104ab8782880161040c565b94505060206104bc8782880161043f565b935050604085013567ffffffffffffffff8111156104dd576104dc61025c565b5b6104e98782880161036f565b92505060606104fa87828801610472565b91505092959194509250565b5f6020828403121561051b5761051a610258565b5b5f6105288482850161040c565b91505092915050565b61053a816103e3565b82525050565b5f6020820190506105535f830184610531565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61058b82610559565b6105958185610563565b93506105a5818560208601610573565b6105ae81610268565b840191505092915050565b5f6020820190508181035f8301526105d18184610581565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61061082610453565b915061061b83610453565b925082820261062981610453565b915082820484148315176106405761063f6105d9565b5b5092915050565b5f610651826103e3565b915061065c836103e3565b9250828201905067ffffffffffffffff81111561067c5761067b6105d9565b5b92915050565b61068b81610420565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106ba82610691565b9050919050565b6106ca816106b0565b82525050565b6106d981610453565b82525050565b5f60c0820190506106f25f830189610531565b6106ff6020830188610531565b61070c6040830187610682565b61071960608301866106c1565b818103608083015261072b8185610581565b905061073a60a08301846106d0565b97965050505050505056fea164736f6c634300081c000a",
}

// SequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerMetaData.ABI instead.
var SequencerABI = SequencerMetaData.ABI

// SequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerMetaData.Bin instead.
var SequencerBin = SequencerMetaData.Bin

// DeploySequencer deploys a new Ethereum contract, binding an instance of Sequencer to it.
func DeploySequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sequencer, error) {
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// Sequencer is an auto generated Go binding around an Ethereum contract.
type Sequencer struct {
	SequencerCaller     // Read-only binding to the contract
	SequencerTransactor // Write-only binding to the contract
	SequencerFilterer   // Log filterer for contract events
}

// SequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSession struct {
	Contract     *Sequencer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerCallerSession struct {
	Contract *SequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerTransactorSession struct {
	Contract     *SequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerRaw struct {
	Contract *Sequencer // Generic contract binding to access the raw methods on
}

// SequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerCallerRaw struct {
	Contract *SequencerCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerTransactorRaw struct {
	Contract *SequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencer creates a new instance of Sequencer, bound to a specific deployed contract.
func NewSequencer(address common.Address, backend bind.ContractBackend) (*Sequencer, error) {
	contract, err := bindSequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sequencer{SequencerCaller: SequencerCaller{contract: contract}, SequencerTransactor: SequencerTransactor{contract: contract}, SequencerFilterer: SequencerFilterer{contract: contract}}, nil
}

// NewSequencerCaller creates a new read-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerCaller(address common.Address, caller bind.ContractCaller) (*SequencerCaller, error) {
	contract, err := bindSequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerCaller{contract: contract}, nil
}

// NewSequencerTransactor creates a new write-only instance of Sequencer, bound to a specific deployed contract.
func NewSequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerTransactor, error) {
	contract, err := bindSequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerTransactor{contract: contract}, nil
}

// NewSequencerFilterer creates a new log filterer instance of Sequencer, bound to a specific deployed contract.
func NewSequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerFilterer, error) {
	contract, err := bindSequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerFilterer{contract: contract}, nil
}

// bindSequencer binds a generic wrapper to an already deployed contract.
func bindSequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.SequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.SequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sequencer *SequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sequencer *SequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sequencer *SequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sequencer.Contract.contract.Transact(opts, method, params...)
}

// GetTxCountForEon is a free data retrieval call binding the contract method 0x7bbd164b.
//
// Solidity: function getTxCountForEon(uint64 eon) view returns(uint64)
func (_Sequencer *SequencerCaller) GetTxCountForEon(opts *bind.CallOpts, eon uint64) (uint64, error) {
	var out []interface{}
	err := _Sequencer.contract.Call(opts, &out, "getTxCountForEon", eon)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTxCountForEon is a free data retrieval call binding the contract method 0x7bbd164b.
//
// Solidity: function getTxCountForEon(uint64 eon) view returns(uint64)
func (_Sequencer *SequencerSession) GetTxCountForEon(eon uint64) (uint64, error) {
	return _Sequencer.Contract.GetTxCountForEon(&_Sequencer.CallOpts, eon)
}

// GetTxCountForEon is a free data retrieval call binding the contract method 0x7bbd164b.
//
// Solidity: function getTxCountForEon(uint64 eon) view returns(uint64)
func (_Sequencer *SequencerCallerSession) GetTxCountForEon(eon uint64) (uint64, error) {
	return _Sequencer.Contract.GetTxCountForEon(&_Sequencer.CallOpts, eon)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Sequencer *SequencerTransactor) SubmitDecryptionProgress(opts *bind.TransactOpts, message []byte) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "submitDecryptionProgress", message)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Sequencer *SequencerSession) SubmitDecryptionProgress(message []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.SubmitDecryptionProgress(&_Sequencer.TransactOpts, message)
}

// SubmitDecryptionProgress is a paid mutator transaction binding the contract method 0x2d32522e.
//
// Solidity: function submitDecryptionProgress(bytes message) returns()
func (_Sequencer *SequencerTransactorSession) SubmitDecryptionProgress(message []byte) (*types.Transaction, error) {
	return _Sequencer.Contract.SubmitDecryptionProgress(&_Sequencer.TransactOpts, message)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Sequencer *SequencerTransactor) SubmitEncryptedTransaction(opts *bind.TransactOpts, eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Sequencer.contract.Transact(opts, "submitEncryptedTransaction", eon, identityPrefix, encryptedTransaction, gasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Sequencer *SequencerSession) SubmitEncryptedTransaction(eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.SubmitEncryptedTransaction(&_Sequencer.TransactOpts, eon, identityPrefix, encryptedTransaction, gasLimit)
}

// SubmitEncryptedTransaction is a paid mutator transaction binding the contract method 0x6a69d2e1.
//
// Solidity: function submitEncryptedTransaction(uint64 eon, bytes32 identityPrefix, bytes encryptedTransaction, uint256 gasLimit) payable returns()
func (_Sequencer *SequencerTransactorSession) SubmitEncryptedTransaction(eon uint64, identityPrefix [32]byte, encryptedTransaction []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _Sequencer.Contract.SubmitEncryptedTransaction(&_Sequencer.TransactOpts, eon, identityPrefix, encryptedTransaction, gasLimit)
}

// SequencerDecryptionProgressSubmittedIterator is returned from FilterDecryptionProgressSubmitted and is used to iterate over the raw logs and unpacked data for DecryptionProgressSubmitted events raised by the Sequencer contract.
type SequencerDecryptionProgressSubmittedIterator struct {
	Event *SequencerDecryptionProgressSubmitted // Event containing the contract specifics and raw log

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
func (it *SequencerDecryptionProgressSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerDecryptionProgressSubmitted)
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
		it.Event = new(SequencerDecryptionProgressSubmitted)
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
func (it *SequencerDecryptionProgressSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerDecryptionProgressSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerDecryptionProgressSubmitted represents a DecryptionProgressSubmitted event raised by the Sequencer contract.
type SequencerDecryptionProgressSubmitted struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDecryptionProgressSubmitted is a free log retrieval operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Sequencer *SequencerFilterer) FilterDecryptionProgressSubmitted(opts *bind.FilterOpts) (*SequencerDecryptionProgressSubmittedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "DecryptionProgressSubmitted")
	if err != nil {
		return nil, err
	}
	return &SequencerDecryptionProgressSubmittedIterator{contract: _Sequencer.contract, event: "DecryptionProgressSubmitted", logs: logs, sub: sub}, nil
}

// WatchDecryptionProgressSubmitted is a free log subscription operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Sequencer *SequencerFilterer) WatchDecryptionProgressSubmitted(opts *bind.WatchOpts, sink chan<- *SequencerDecryptionProgressSubmitted) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "DecryptionProgressSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerDecryptionProgressSubmitted)
				if err := _Sequencer.contract.UnpackLog(event, "DecryptionProgressSubmitted", log); err != nil {
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

// ParseDecryptionProgressSubmitted is a log parse operation binding the contract event 0xa9a0645b33a70f18b8d490681d637cb46a859ec51707787e6f46b942f90e8f59.
//
// Solidity: event DecryptionProgressSubmitted(bytes message)
func (_Sequencer *SequencerFilterer) ParseDecryptionProgressSubmitted(log types.Log) (*SequencerDecryptionProgressSubmitted, error) {
	event := new(SequencerDecryptionProgressSubmitted)
	if err := _Sequencer.contract.UnpackLog(event, "DecryptionProgressSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerTransactionSubmittedIterator is returned from FilterTransactionSubmitted and is used to iterate over the raw logs and unpacked data for TransactionSubmitted events raised by the Sequencer contract.
type SequencerTransactionSubmittedIterator struct {
	Event *SequencerTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *SequencerTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerTransactionSubmitted)
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
		it.Event = new(SequencerTransactionSubmitted)
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
func (it *SequencerTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerTransactionSubmitted represents a TransactionSubmitted event raised by the Sequencer contract.
type SequencerTransactionSubmitted struct {
	Eon                  uint64
	TxIndex              uint64
	IdentityPrefix       [32]byte
	Sender               common.Address
	EncryptedTransaction []byte
	GasLimit             *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTransactionSubmitted is a free log retrieval operation binding the contract event 0xa7f1b5467be46c45249fb93063cceef96c63ddad03819246bc7770e32d4f5b7d.
//
// Solidity: event TransactionSubmitted(uint64 eon, uint64 txIndex, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Sequencer *SequencerFilterer) FilterTransactionSubmitted(opts *bind.FilterOpts) (*SequencerTransactionSubmittedIterator, error) {

	logs, sub, err := _Sequencer.contract.FilterLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return &SequencerTransactionSubmittedIterator{contract: _Sequencer.contract, event: "TransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchTransactionSubmitted is a free log subscription operation binding the contract event 0xa7f1b5467be46c45249fb93063cceef96c63ddad03819246bc7770e32d4f5b7d.
//
// Solidity: event TransactionSubmitted(uint64 eon, uint64 txIndex, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Sequencer *SequencerFilterer) WatchTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *SequencerTransactionSubmitted) (event.Subscription, error) {

	logs, sub, err := _Sequencer.contract.WatchLogs(opts, "TransactionSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerTransactionSubmitted)
				if err := _Sequencer.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
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

// ParseTransactionSubmitted is a log parse operation binding the contract event 0xa7f1b5467be46c45249fb93063cceef96c63ddad03819246bc7770e32d4f5b7d.
//
// Solidity: event TransactionSubmitted(uint64 eon, uint64 txIndex, bytes32 identityPrefix, address sender, bytes encryptedTransaction, uint256 gasLimit)
func (_Sequencer *SequencerFilterer) ParseTransactionSubmitted(log types.Log) (*SequencerTransactionSubmitted, error) {
	event := new(SequencerTransactionSubmitted)
	if err := _Sequencer.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
