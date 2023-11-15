// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getItem\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200096c3803806200096c8339818101604052810190620000369190620001d3565b805f908162000046919062000459565b50506200053d565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620000af8262000067565b810181811067ffffffffffffffff82111715620000d157620000d062000077565b5b80604052505050565b5f620000e56200004e565b9050620000f38282620000a4565b919050565b5f67ffffffffffffffff82111562000115576200011462000077565b5b620001208262000067565b9050602081019050919050565b5f5b838110156200014c5780820151818401526020810190506200012f565b5f8484015250505050565b5f6200016d6200016784620000f8565b620000da565b9050828152602081018484840111156200018c576200018b62000063565b5b620001998482856200012d565b509392505050565b5f82601f830112620001b857620001b76200005f565b5b8151620001ca84826020860162000157565b91505092915050565b5f60208284031215620001eb57620001ea62000057565b5b5f82015167ffffffffffffffff8111156200020b576200020a6200005b565b5b6200021984828501620001a1565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200027157607f821691505b6020821081036200028757620002866200022c565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620002eb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002ae565b620002f78683620002ae565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620003416200033b62000335846200030f565b62000318565b6200030f565b9050919050565b5f819050919050565b6200035c8362000321565b620003746200036b8262000348565b848454620002ba565b825550505050565b5f90565b6200038a6200037c565b6200039781848462000351565b505050565b5b81811015620003be57620003b25f8262000380565b6001810190506200039d565b5050565b601f8211156200040d57620003d7816200028d565b620003e2846200029f565b81016020851015620003f2578190505b6200040a62000401856200029f565b8301826200039c565b50505b505050565b5f82821c905092915050565b5f6200042f5f198460080262000412565b1980831691505092915050565b5f6200044983836200041e565b9150826002028217905092915050565b620004648262000222565b67ffffffffffffffff81111562000480576200047f62000077565b5b6200048c825462000259565b62000499828285620003c2565b5f60209050601f831160018114620004cf575f8415620004ba578287015190505b620004c685826200043c565b86555062000535565b601f198416620004df866200028d565b5f5b828110156200050857848901518255600182019150602085019450602081019050620004e1565b8683101562000528578489015162000524601f8916826200041e565b8355505b6001600288020188555050505b505050505050565b610421806200054b5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806348f343f31461004e57806354fd4d501461007e578063aa0372e71461009c578063f56256c7146100cc575b5f80fd5b6100686004803603810190610063919061022c565b6100e8565b6040516100759190610266565b60405180910390f35b6100866100fd565b6040516100939190610309565b60405180910390f35b6100b660048036038101906100b1919061022c565b610188565b6040516100c39190610266565b60405180910390f35b6100e660048036038101906100e19190610329565b6101a2565b005b6001602052805f5260405f205f915090505481565b5f805461010990610394565b80601f016020809104026020016040519081016040528092919081815260200182805461013590610394565b80156101805780601f1061015757610100808354040283529160200191610180565b820191905f5260205f20905b81548152906001019060200180831161016357829003601f168201915b505050505081565b5f60015f8381526020019081526020015f20549050919050565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d482826040516101e99291906103c4565b60405180910390a15050565b5f80fd5b5f819050919050565b61020b816101f9565b8114610215575f80fd5b50565b5f8135905061022681610202565b92915050565b5f60208284031215610241576102406101f5565b5b5f61024e84828501610218565b91505092915050565b610260816101f9565b82525050565b5f6020820190506102795f830184610257565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156102b657808201518184015260208101905061029b565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6102db8261027f565b6102e58185610289565b93506102f5818560208601610299565b6102fe816102c1565b840191505092915050565b5f6020820190508181035f83015261032181846102d1565b905092915050565b5f806040838503121561033f5761033e6101f5565b5b5f61034c85828601610218565b925050602061035d85828601610218565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103ab57607f821691505b6020821081036103be576103bd610367565b5b50919050565b5f6040820190506103d75f830185610257565b6103e46020830184610257565b939250505056fea2646970667358221220ac0e73ed5717fcd54a285c4ee3ef5a73c5fb5efeb113415fd7fdff599d4d916d64736f6c63430008170033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _version string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// GetItem is a free data retrieval call binding the contract method 0xaa0372e7.
//
// Solidity: function getItem(bytes32 key) view returns(bytes32)
func (_Store *StoreCaller) GetItem(opts *bind.CallOpts, key [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getItem", key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetItem is a free data retrieval call binding the contract method 0xaa0372e7.
//
// Solidity: function getItem(bytes32 key) view returns(bytes32)
func (_Store *StoreSession) GetItem(key [32]byte) ([32]byte, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// GetItem is a free data retrieval call binding the contract method 0xaa0372e7.
//
// Solidity: function getItem(bytes32 key) view returns(bytes32)
func (_Store *StoreCallerSession) GetItem(key [32]byte) ([32]byte, error) {
	return _Store.Contract.GetItem(&_Store.CallOpts, key)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Store *StoreCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Store.Contract.Items(&_Store.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Store *StoreCallerSession) Version() (string, error) {
	return _Store.Contract.Version(&_Store.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Store *StoreTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Store.Contract.SetItem(&_Store.TransactOpts, key, value)
}

// StoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Store contract.
type StoreItemSetIterator struct {
	Event *StoreItemSet // Event containing the contract specifics and raw log

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
func (it *StoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreItemSet)
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
		it.Event = new(StoreItemSet)
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
func (it *StoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreItemSet represents a ItemSet event raised by the Store contract.
type StoreItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) FilterItemSet(opts *bind.FilterOpts) (*StoreItemSetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &StoreItemSetIterator{contract: _Store.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *StoreItemSet) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreItemSet)
				if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Store *StoreFilterer) ParseItemSet(log types.Log) (*StoreItemSet, error) {
	event := new(StoreItemSet)
	if err := _Store.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
