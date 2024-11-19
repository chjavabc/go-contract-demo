// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package part5

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

// CrowdFundMetaData contains all meta data concerning the CrowdFund contract.
var CrowdFundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_goalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundingReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmountRaised\",\"type\":\"uint256\"}],\"name\":\"ProjectClosed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"closed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"}],\"name\":\"getContributionAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"goalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raisedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalContributorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CrowdFundABI is the input ABI used to generate the binding from.
// Deprecated: Use CrowdFundMetaData.ABI instead.
var CrowdFundABI = CrowdFundMetaData.ABI

// CrowdFund is an auto generated Go binding around an Ethereum contract.
type CrowdFund struct {
	CrowdFundCaller     // Read-only binding to the contract
	CrowdFundTransactor // Write-only binding to the contract
	CrowdFundFilterer   // Log filterer for contract events
}

// CrowdFundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdFundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdFundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdFundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdFundSession struct {
	Contract     *CrowdFund        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdFundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdFundCallerSession struct {
	Contract *CrowdFundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdFundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdFundTransactorSession struct {
	Contract     *CrowdFundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdFundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdFundRaw struct {
	Contract *CrowdFund // Generic contract binding to access the raw methods on
}

// CrowdFundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdFundCallerRaw struct {
	Contract *CrowdFundCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdFundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdFundTransactorRaw struct {
	Contract *CrowdFundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdFund creates a new instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFund(address common.Address, backend bind.ContractBackend) (*CrowdFund, error) {
	contract, err := bindCrowdFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdFund{CrowdFundCaller: CrowdFundCaller{contract: contract}, CrowdFundTransactor: CrowdFundTransactor{contract: contract}, CrowdFundFilterer: CrowdFundFilterer{contract: contract}}, nil
}

// NewCrowdFundCaller creates a new read-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundCaller(address common.Address, caller bind.ContractCaller) (*CrowdFundCaller, error) {
	contract, err := bindCrowdFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundCaller{contract: contract}, nil
}

// NewCrowdFundTransactor creates a new write-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdFundTransactor, error) {
	contract, err := bindCrowdFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundTransactor{contract: contract}, nil
}

// NewCrowdFundFilterer creates a new log filterer instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdFundFilterer, error) {
	contract, err := bindCrowdFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdFundFilterer{contract: contract}, nil
}

// bindCrowdFund binds a generic wrapper to an already deployed contract.
func bindCrowdFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrowdFundMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.CrowdFundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transact(opts, method, params...)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_CrowdFund *CrowdFundCaller) Closed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "closed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_CrowdFund *CrowdFundSession) Closed() (bool, error) {
	return _CrowdFund.Contract.Closed(&_CrowdFund.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_CrowdFund *CrowdFundCallerSession) Closed() (bool, error) {
	return _CrowdFund.Contract.Closed(&_CrowdFund.CallOpts)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_CrowdFund *CrowdFundCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "contributions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_CrowdFund *CrowdFundSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.Contributions(&_CrowdFund.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.Contributions(&_CrowdFund.CallOpts, arg0)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrowdFund *CrowdFundCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrowdFund *CrowdFundSession) Creator() (common.Address, error) {
	return _CrowdFund.Contract.Creator(&_CrowdFund.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_CrowdFund *CrowdFundCallerSession) Creator() (common.Address, error) {
	return _CrowdFund.Contract.Creator(&_CrowdFund.CallOpts)
}

// GetContributionAmount is a free data retrieval call binding the contract method 0x32c1f245.
//
// Solidity: function getContributionAmount(address contributor) view returns(uint256)
func (_CrowdFund *CrowdFundCaller) GetContributionAmount(opts *bind.CallOpts, contributor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "getContributionAmount", contributor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContributionAmount is a free data retrieval call binding the contract method 0x32c1f245.
//
// Solidity: function getContributionAmount(address contributor) view returns(uint256)
func (_CrowdFund *CrowdFundSession) GetContributionAmount(contributor common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.GetContributionAmount(&_CrowdFund.CallOpts, contributor)
}

// GetContributionAmount is a free data retrieval call binding the contract method 0x32c1f245.
//
// Solidity: function getContributionAmount(address contributor) view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) GetContributionAmount(contributor common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.GetContributionAmount(&_CrowdFund.CallOpts, contributor)
}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_CrowdFund *CrowdFundCaller) GoalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "goalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_CrowdFund *CrowdFundSession) GoalAmount() (*big.Int, error) {
	return _CrowdFund.Contract.GoalAmount(&_CrowdFund.CallOpts)
}

// GoalAmount is a free data retrieval call binding the contract method 0x2636b945.
//
// Solidity: function goalAmount() view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) GoalAmount() (*big.Int, error) {
	return _CrowdFund.Contract.GoalAmount(&_CrowdFund.CallOpts)
}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() view returns(uint256)
func (_CrowdFund *CrowdFundCaller) RaisedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "raisedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() view returns(uint256)
func (_CrowdFund *CrowdFundSession) RaisedAmount() (*big.Int, error) {
	return _CrowdFund.Contract.RaisedAmount(&_CrowdFund.CallOpts)
}

// RaisedAmount is a free data retrieval call binding the contract method 0xc59ee1dc.
//
// Solidity: function raisedAmount() view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) RaisedAmount() (*big.Int, error) {
	return _CrowdFund.Contract.RaisedAmount(&_CrowdFund.CallOpts)
}

// TotalContributorsCount is a free data retrieval call binding the contract method 0x72c24d92.
//
// Solidity: function totalContributorsCount() view returns(uint256)
func (_CrowdFund *CrowdFundCaller) TotalContributorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "totalContributorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalContributorsCount is a free data retrieval call binding the contract method 0x72c24d92.
//
// Solidity: function totalContributorsCount() view returns(uint256)
func (_CrowdFund *CrowdFundSession) TotalContributorsCount() (*big.Int, error) {
	return _CrowdFund.Contract.TotalContributorsCount(&_CrowdFund.CallOpts)
}

// TotalContributorsCount is a free data retrieval call binding the contract method 0x72c24d92.
//
// Solidity: function totalContributorsCount() view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) TotalContributorsCount() (*big.Int, error) {
	return _CrowdFund.Contract.TotalContributorsCount(&_CrowdFund.CallOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CrowdFund *CrowdFundTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CrowdFund *CrowdFundSession) Contribute() (*types.Transaction, error) {
	return _CrowdFund.Contract.Contribute(&_CrowdFund.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_CrowdFund *CrowdFundTransactorSession) Contribute() (*types.Transaction, error) {
	return _CrowdFund.Contract.Contribute(&_CrowdFund.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CrowdFund *CrowdFundTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CrowdFund *CrowdFundSession) Withdraw() (*types.Transaction, error) {
	return _CrowdFund.Contract.Withdraw(&_CrowdFund.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CrowdFund *CrowdFundTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CrowdFund.Contract.Withdraw(&_CrowdFund.TransactOpts)
}

// CrowdFundFundingReceivedIterator is returned from FilterFundingReceived and is used to iterate over the raw logs and unpacked data for FundingReceived events raised by the CrowdFund contract.
type CrowdFundFundingReceivedIterator struct {
	Event *CrowdFundFundingReceived // Event containing the contract specifics and raw log

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
func (it *CrowdFundFundingReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundFundingReceived)
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
		it.Event = new(CrowdFundFundingReceived)
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
func (it *CrowdFundFundingReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundFundingReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundFundingReceived represents a FundingReceived event raised by the CrowdFund contract.
type CrowdFundFundingReceived struct {
	Contributor common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundingReceived is a free log retrieval operation binding the contract event 0x23bac35767dec08bf219a302474bbc6687dfaaf3776a62d45bb7811de040e26d.
//
// Solidity: event FundingReceived(address indexed contributor, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) FilterFundingReceived(opts *bind.FilterOpts, contributor []common.Address) (*CrowdFundFundingReceivedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "FundingReceived", contributorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundFundingReceivedIterator{contract: _CrowdFund.contract, event: "FundingReceived", logs: logs, sub: sub}, nil
}

// WatchFundingReceived is a free log subscription operation binding the contract event 0x23bac35767dec08bf219a302474bbc6687dfaaf3776a62d45bb7811de040e26d.
//
// Solidity: event FundingReceived(address indexed contributor, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) WatchFundingReceived(opts *bind.WatchOpts, sink chan<- *CrowdFundFundingReceived, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "FundingReceived", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundFundingReceived)
				if err := _CrowdFund.contract.UnpackLog(event, "FundingReceived", log); err != nil {
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

// ParseFundingReceived is a log parse operation binding the contract event 0x23bac35767dec08bf219a302474bbc6687dfaaf3776a62d45bb7811de040e26d.
//
// Solidity: event FundingReceived(address indexed contributor, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) ParseFundingReceived(log types.Log) (*CrowdFundFundingReceived, error) {
	event := new(CrowdFundFundingReceived)
	if err := _CrowdFund.contract.UnpackLog(event, "FundingReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundProjectClosedIterator is returned from FilterProjectClosed and is used to iterate over the raw logs and unpacked data for ProjectClosed events raised by the CrowdFund contract.
type CrowdFundProjectClosedIterator struct {
	Event *CrowdFundProjectClosed // Event containing the contract specifics and raw log

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
func (it *CrowdFundProjectClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundProjectClosed)
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
		it.Event = new(CrowdFundProjectClosed)
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
func (it *CrowdFundProjectClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundProjectClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundProjectClosed represents a ProjectClosed event raised by the CrowdFund contract.
type CrowdFundProjectClosed struct {
	TotalAmountRaised *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterProjectClosed is a free log retrieval operation binding the contract event 0xa10da3bd85a8a7569325abd94177d0bbf25fbded5a3f6622a5588ce2cf1153ba.
//
// Solidity: event ProjectClosed(uint256 totalAmountRaised)
func (_CrowdFund *CrowdFundFilterer) FilterProjectClosed(opts *bind.FilterOpts) (*CrowdFundProjectClosedIterator, error) {

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "ProjectClosed")
	if err != nil {
		return nil, err
	}
	return &CrowdFundProjectClosedIterator{contract: _CrowdFund.contract, event: "ProjectClosed", logs: logs, sub: sub}, nil
}

// WatchProjectClosed is a free log subscription operation binding the contract event 0xa10da3bd85a8a7569325abd94177d0bbf25fbded5a3f6622a5588ce2cf1153ba.
//
// Solidity: event ProjectClosed(uint256 totalAmountRaised)
func (_CrowdFund *CrowdFundFilterer) WatchProjectClosed(opts *bind.WatchOpts, sink chan<- *CrowdFundProjectClosed) (event.Subscription, error) {

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "ProjectClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundProjectClosed)
				if err := _CrowdFund.contract.UnpackLog(event, "ProjectClosed", log); err != nil {
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

// ParseProjectClosed is a log parse operation binding the contract event 0xa10da3bd85a8a7569325abd94177d0bbf25fbded5a3f6622a5588ce2cf1153ba.
//
// Solidity: event ProjectClosed(uint256 totalAmountRaised)
func (_CrowdFund *CrowdFundFilterer) ParseProjectClosed(log types.Log) (*CrowdFundProjectClosed, error) {
	event := new(CrowdFundProjectClosed)
	if err := _CrowdFund.contract.UnpackLog(event, "ProjectClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
