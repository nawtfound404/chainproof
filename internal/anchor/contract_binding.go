// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package anchor

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

// AnchorMetaData contains all meta data concerning the Anchor contract.
var AnchorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"proofExists\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"storeProof\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ProofStored\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ProofAlreadyExists\",\"inputs\":[]}]",
}

// AnchorABI is the input ABI used to generate the binding from.
// Deprecated: Use AnchorMetaData.ABI instead.
var AnchorABI = AnchorMetaData.ABI

// Anchor is an auto generated Go binding around an Ethereum contract.
type Anchor struct {
	AnchorCaller     // Read-only binding to the contract
	AnchorTransactor // Write-only binding to the contract
	AnchorFilterer   // Log filterer for contract events
}

// AnchorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnchorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnchorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnchorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnchorSession struct {
	Contract     *Anchor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnchorCallerSession struct {
	Contract *AnchorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AnchorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnchorTransactorSession struct {
	Contract     *AnchorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnchorRaw struct {
	Contract *Anchor // Generic contract binding to access the raw methods on
}

// AnchorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnchorCallerRaw struct {
	Contract *AnchorCaller // Generic read-only contract binding to access the raw methods on
}

// AnchorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnchorTransactorRaw struct {
	Contract *AnchorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnchor creates a new instance of Anchor, bound to a specific deployed contract.
func NewAnchor(address common.Address, backend bind.ContractBackend) (*Anchor, error) {
	contract, err := bindAnchor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Anchor{AnchorCaller: AnchorCaller{contract: contract}, AnchorTransactor: AnchorTransactor{contract: contract}, AnchorFilterer: AnchorFilterer{contract: contract}}, nil
}

// NewAnchorCaller creates a new read-only instance of Anchor, bound to a specific deployed contract.
func NewAnchorCaller(address common.Address, caller bind.ContractCaller) (*AnchorCaller, error) {
	contract, err := bindAnchor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorCaller{contract: contract}, nil
}

// NewAnchorTransactor creates a new write-only instance of Anchor, bound to a specific deployed contract.
func NewAnchorTransactor(address common.Address, transactor bind.ContractTransactor) (*AnchorTransactor, error) {
	contract, err := bindAnchor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorTransactor{contract: contract}, nil
}

// NewAnchorFilterer creates a new log filterer instance of Anchor, bound to a specific deployed contract.
func NewAnchorFilterer(address common.Address, filterer bind.ContractFilterer) (*AnchorFilterer, error) {
	contract, err := bindAnchor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnchorFilterer{contract: contract}, nil
}

// bindAnchor binds a generic wrapper to an already deployed contract.
func bindAnchor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnchorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anchor *AnchorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anchor.Contract.AnchorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anchor *AnchorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anchor.Contract.AnchorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anchor *AnchorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anchor.Contract.AnchorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Anchor *AnchorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Anchor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Anchor *AnchorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Anchor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Anchor *AnchorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Anchor.Contract.contract.Transact(opts, method, params...)
}

// ProofExists is a free data retrieval call binding the contract method 0xd938d202.
//
// Solidity: function proofExists(bytes32 hash) view returns(bool)
func (_Anchor *AnchorCaller) ProofExists(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Anchor.contract.Call(opts, &out, "proofExists", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProofExists is a free data retrieval call binding the contract method 0xd938d202.
//
// Solidity: function proofExists(bytes32 hash) view returns(bool)
func (_Anchor *AnchorSession) ProofExists(hash [32]byte) (bool, error) {
	return _Anchor.Contract.ProofExists(&_Anchor.CallOpts, hash)
}

// ProofExists is a free data retrieval call binding the contract method 0xd938d202.
//
// Solidity: function proofExists(bytes32 hash) view returns(bool)
func (_Anchor *AnchorCallerSession) ProofExists(hash [32]byte) (bool, error) {
	return _Anchor.Contract.ProofExists(&_Anchor.CallOpts, hash)
}

// StoreProof is a paid mutator transaction binding the contract method 0x8952877b.
//
// Solidity: function storeProof(bytes32 hash) returns()
func (_Anchor *AnchorTransactor) StoreProof(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Anchor.contract.Transact(opts, "storeProof", hash)
}

// StoreProof is a paid mutator transaction binding the contract method 0x8952877b.
//
// Solidity: function storeProof(bytes32 hash) returns()
func (_Anchor *AnchorSession) StoreProof(hash [32]byte) (*types.Transaction, error) {
	return _Anchor.Contract.StoreProof(&_Anchor.TransactOpts, hash)
}

// StoreProof is a paid mutator transaction binding the contract method 0x8952877b.
//
// Solidity: function storeProof(bytes32 hash) returns()
func (_Anchor *AnchorTransactorSession) StoreProof(hash [32]byte) (*types.Transaction, error) {
	return _Anchor.Contract.StoreProof(&_Anchor.TransactOpts, hash)
}

// AnchorProofStoredIterator is returned from FilterProofStored and is used to iterate over the raw logs and unpacked data for ProofStored events raised by the Anchor contract.
type AnchorProofStoredIterator struct {
	Event *AnchorProofStored // Event containing the contract specifics and raw log

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
func (it *AnchorProofStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorProofStored)
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
		it.Event = new(AnchorProofStored)
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
func (it *AnchorProofStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorProofStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorProofStored represents a ProofStored event raised by the Anchor contract.
type AnchorProofStored struct {
	Hash   [32]byte
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProofStored is a free log retrieval operation binding the contract event 0xfda99acf408a6ae16a17a1706ea07ffd3d505b55b914399e9b238d050e2a5edb.
//
// Solidity: event ProofStored(bytes32 indexed hash, address indexed sender)
func (_Anchor *AnchorFilterer) FilterProofStored(opts *bind.FilterOpts, hash [][32]byte, sender []common.Address) (*AnchorProofStoredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Anchor.contract.FilterLogs(opts, "ProofStored", hashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AnchorProofStoredIterator{contract: _Anchor.contract, event: "ProofStored", logs: logs, sub: sub}, nil
}

// WatchProofStored is a free log subscription operation binding the contract event 0xfda99acf408a6ae16a17a1706ea07ffd3d505b55b914399e9b238d050e2a5edb.
//
// Solidity: event ProofStored(bytes32 indexed hash, address indexed sender)
func (_Anchor *AnchorFilterer) WatchProofStored(opts *bind.WatchOpts, sink chan<- *AnchorProofStored, hash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Anchor.contract.WatchLogs(opts, "ProofStored", hashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorProofStored)
				if err := _Anchor.contract.UnpackLog(event, "ProofStored", log); err != nil {
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

// ParseProofStored is a log parse operation binding the contract event 0xfda99acf408a6ae16a17a1706ea07ffd3d505b55b914399e9b238d050e2a5edb.
//
// Solidity: event ProofStored(bytes32 indexed hash, address indexed sender)
func (_Anchor *AnchorFilterer) ParseProofStored(log types.Log) (*AnchorProofStored, error) {
	event := new(AnchorProofStored)
	if err := _Anchor.contract.UnpackLog(event, "ProofStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
