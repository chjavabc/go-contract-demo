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

// CalculatorMetaData contains all meta data concerning the Calculator contract.
var CalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"subtract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// CalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CalculatorMetaData.ABI instead.
var CalculatorABI = CalculatorMetaData.ABI

// Calculator is an auto generated Go binding around an Ethereum contract.
type Calculator struct {
	CalculatorCaller     // Read-only binding to the contract
	CalculatorTransactor // Write-only binding to the contract
	CalculatorFilterer   // Log filterer for contract events
}

// CalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalculatorSession struct {
	Contract     *Calculator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalculatorCallerSession struct {
	Contract *CalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalculatorTransactorSession struct {
	Contract     *CalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalculatorRaw struct {
	Contract *Calculator // Generic contract binding to access the raw methods on
}

// CalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalculatorCallerRaw struct {
	Contract *CalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// CalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalculatorTransactorRaw struct {
	Contract *CalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalculator creates a new instance of Calculator, bound to a specific deployed contract.
func NewCalculator(address common.Address, backend bind.ContractBackend) (*Calculator, error) {
	contract, err := bindCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calculator{CalculatorCaller: CalculatorCaller{contract: contract}, CalculatorTransactor: CalculatorTransactor{contract: contract}, CalculatorFilterer: CalculatorFilterer{contract: contract}}, nil
}

// NewCalculatorCaller creates a new read-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorCaller(address common.Address, caller bind.ContractCaller) (*CalculatorCaller, error) {
	contract, err := bindCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorCaller{contract: contract}, nil
}

// NewCalculatorTransactor creates a new write-only instance of Calculator, bound to a specific deployed contract.
func NewCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CalculatorTransactor, error) {
	contract, err := bindCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalculatorTransactor{contract: contract}, nil
}

// NewCalculatorFilterer creates a new log filterer instance of Calculator, bound to a specific deployed contract.
func NewCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CalculatorFilterer, error) {
	contract, err := bindCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalculatorFilterer{contract: contract}, nil
}

// bindCalculator binds a generic wrapper to an already deployed contract.
func bindCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.CalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calculator.Contract.CalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calculator *CalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calculator *CalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calculator *CalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calculator.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Calculator.contract.Call(opts, &out, "add", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Add(&_Calculator.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Add(&_Calculator.CallOpts, a, b)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorCaller) Subtract(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Calculator.contract.Call(opts, &out, "subtract", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Subtract(&_Calculator.CallOpts, a, b)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 a, uint256 b) pure returns(uint256)
func (_Calculator *CalculatorCallerSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Calculator.Contract.Subtract(&_Calculator.CallOpts, a, b)
}
