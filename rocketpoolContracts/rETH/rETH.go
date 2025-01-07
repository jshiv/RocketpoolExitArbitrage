// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rETH

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

// RETHMetaData contains all meta data concerning the RETH contract.
var RETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokensBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokensMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rethAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositExcess\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositExcessCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rethAmount\",\"type\":\"uint256\"}],\"name\":\"getEthValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getRethValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RETHABI is the input ABI used to generate the binding from.
// Deprecated: Use RETHMetaData.ABI instead.
var RETHABI = RETHMetaData.ABI

// RETH is an auto generated Go binding around an Ethereum contract.
type RETH struct {
	RETHCaller     // Read-only binding to the contract
	RETHTransactor // Write-only binding to the contract
	RETHFilterer   // Log filterer for contract events
}

// RETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type RETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RETHSession struct {
	Contract     *RETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RETHCallerSession struct {
	Contract *RETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RETHTransactorSession struct {
	Contract     *RETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type RETHRaw struct {
	Contract *RETH // Generic contract binding to access the raw methods on
}

// RETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RETHCallerRaw struct {
	Contract *RETHCaller // Generic read-only contract binding to access the raw methods on
}

// RETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RETHTransactorRaw struct {
	Contract *RETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRETH creates a new instance of RETH, bound to a specific deployed contract.
func NewRETH(address common.Address, backend bind.ContractBackend) (*RETH, error) {
	contract, err := bindRETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RETH{RETHCaller: RETHCaller{contract: contract}, RETHTransactor: RETHTransactor{contract: contract}, RETHFilterer: RETHFilterer{contract: contract}}, nil
}

// NewRETHCaller creates a new read-only instance of RETH, bound to a specific deployed contract.
func NewRETHCaller(address common.Address, caller bind.ContractCaller) (*RETHCaller, error) {
	contract, err := bindRETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RETHCaller{contract: contract}, nil
}

// NewRETHTransactor creates a new write-only instance of RETH, bound to a specific deployed contract.
func NewRETHTransactor(address common.Address, transactor bind.ContractTransactor) (*RETHTransactor, error) {
	contract, err := bindRETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RETHTransactor{contract: contract}, nil
}

// NewRETHFilterer creates a new log filterer instance of RETH, bound to a specific deployed contract.
func NewRETHFilterer(address common.Address, filterer bind.ContractFilterer) (*RETHFilterer, error) {
	contract, err := bindRETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RETHFilterer{contract: contract}, nil
}

// bindRETH binds a generic wrapper to an already deployed contract.
func bindRETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RETH *RETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RETH.Contract.RETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RETH *RETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RETH.Contract.RETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RETH *RETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RETH.Contract.RETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RETH *RETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RETH *RETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RETH *RETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RETH *RETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RETH *RETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RETH.Contract.Allowance(&_RETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RETH *RETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RETH.Contract.Allowance(&_RETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RETH *RETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RETH *RETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RETH.Contract.BalanceOf(&_RETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RETH *RETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RETH.Contract.BalanceOf(&_RETH.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RETH *RETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RETH *RETHSession) Decimals() (uint8, error) {
	return _RETH.Contract.Decimals(&_RETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RETH *RETHCallerSession) Decimals() (uint8, error) {
	return _RETH.Contract.Decimals(&_RETH.CallOpts)
}

// GetCollateralRate is a free data retrieval call binding the contract method 0x852185fc.
//
// Solidity: function getCollateralRate() view returns(uint256)
func (_RETH *RETHCaller) GetCollateralRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "getCollateralRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralRate is a free data retrieval call binding the contract method 0x852185fc.
//
// Solidity: function getCollateralRate() view returns(uint256)
func (_RETH *RETHSession) GetCollateralRate() (*big.Int, error) {
	return _RETH.Contract.GetCollateralRate(&_RETH.CallOpts)
}

// GetCollateralRate is a free data retrieval call binding the contract method 0x852185fc.
//
// Solidity: function getCollateralRate() view returns(uint256)
func (_RETH *RETHCallerSession) GetCollateralRate() (*big.Int, error) {
	return _RETH.Contract.GetCollateralRate(&_RETH.CallOpts)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_RETH *RETHCaller) GetEthValue(opts *bind.CallOpts, _rethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "getEthValue", _rethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_RETH *RETHSession) GetEthValue(_rethAmount *big.Int) (*big.Int, error) {
	return _RETH.Contract.GetEthValue(&_RETH.CallOpts, _rethAmount)
}

// GetEthValue is a free data retrieval call binding the contract method 0x8b32fa23.
//
// Solidity: function getEthValue(uint256 _rethAmount) view returns(uint256)
func (_RETH *RETHCallerSession) GetEthValue(_rethAmount *big.Int) (*big.Int, error) {
	return _RETH.Contract.GetEthValue(&_RETH.CallOpts, _rethAmount)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_RETH *RETHCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_RETH *RETHSession) GetExchangeRate() (*big.Int, error) {
	return _RETH.Contract.GetExchangeRate(&_RETH.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_RETH *RETHCallerSession) GetExchangeRate() (*big.Int, error) {
	return _RETH.Contract.GetExchangeRate(&_RETH.CallOpts)
}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_RETH *RETHCaller) GetRethValue(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "getRethValue", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_RETH *RETHSession) GetRethValue(_ethAmount *big.Int) (*big.Int, error) {
	return _RETH.Contract.GetRethValue(&_RETH.CallOpts, _ethAmount)
}

// GetRethValue is a free data retrieval call binding the contract method 0x4346f03e.
//
// Solidity: function getRethValue(uint256 _ethAmount) view returns(uint256)
func (_RETH *RETHCallerSession) GetRethValue(_ethAmount *big.Int) (*big.Int, error) {
	return _RETH.Contract.GetRethValue(&_RETH.CallOpts, _ethAmount)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_RETH *RETHCaller) GetTotalCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "getTotalCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_RETH *RETHSession) GetTotalCollateral() (*big.Int, error) {
	return _RETH.Contract.GetTotalCollateral(&_RETH.CallOpts)
}

// GetTotalCollateral is a free data retrieval call binding the contract method 0xd6eb5910.
//
// Solidity: function getTotalCollateral() view returns(uint256)
func (_RETH *RETHCallerSession) GetTotalCollateral() (*big.Int, error) {
	return _RETH.Contract.GetTotalCollateral(&_RETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RETH *RETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RETH *RETHSession) Name() (string, error) {
	return _RETH.Contract.Name(&_RETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RETH *RETHCallerSession) Name() (string, error) {
	return _RETH.Contract.Name(&_RETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RETH *RETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RETH *RETHSession) Symbol() (string, error) {
	return _RETH.Contract.Symbol(&_RETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RETH *RETHCallerSession) Symbol() (string, error) {
	return _RETH.Contract.Symbol(&_RETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RETH *RETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RETH *RETHSession) TotalSupply() (*big.Int, error) {
	return _RETH.Contract.TotalSupply(&_RETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RETH *RETHCallerSession) TotalSupply() (*big.Int, error) {
	return _RETH.Contract.TotalSupply(&_RETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RETH *RETHCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RETH.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RETH *RETHSession) Version() (uint8, error) {
	return _RETH.Contract.Version(&_RETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RETH *RETHCallerSession) Version() (uint8, error) {
	return _RETH.Contract.Version(&_RETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RETH *RETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RETH *RETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Approve(&_RETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RETH *RETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Approve(&_RETH.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_RETH *RETHTransactor) Burn(opts *bind.TransactOpts, _rethAmount *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "burn", _rethAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_RETH *RETHSession) Burn(_rethAmount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Burn(&_RETH.TransactOpts, _rethAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _rethAmount) returns()
func (_RETH *RETHTransactorSession) Burn(_rethAmount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Burn(&_RETH.TransactOpts, _rethAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RETH *RETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RETH *RETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.DecreaseAllowance(&_RETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_RETH *RETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.DecreaseAllowance(&_RETH.TransactOpts, spender, subtractedValue)
}

// DepositExcess is a paid mutator transaction binding the contract method 0x6c985a88.
//
// Solidity: function depositExcess() payable returns()
func (_RETH *RETHTransactor) DepositExcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "depositExcess")
}

// DepositExcess is a paid mutator transaction binding the contract method 0x6c985a88.
//
// Solidity: function depositExcess() payable returns()
func (_RETH *RETHSession) DepositExcess() (*types.Transaction, error) {
	return _RETH.Contract.DepositExcess(&_RETH.TransactOpts)
}

// DepositExcess is a paid mutator transaction binding the contract method 0x6c985a88.
//
// Solidity: function depositExcess() payable returns()
func (_RETH *RETHTransactorSession) DepositExcess() (*types.Transaction, error) {
	return _RETH.Contract.DepositExcess(&_RETH.TransactOpts)
}

// DepositExcessCollateral is a paid mutator transaction binding the contract method 0x188e0dc6.
//
// Solidity: function depositExcessCollateral() returns()
func (_RETH *RETHTransactor) DepositExcessCollateral(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "depositExcessCollateral")
}

// DepositExcessCollateral is a paid mutator transaction binding the contract method 0x188e0dc6.
//
// Solidity: function depositExcessCollateral() returns()
func (_RETH *RETHSession) DepositExcessCollateral() (*types.Transaction, error) {
	return _RETH.Contract.DepositExcessCollateral(&_RETH.TransactOpts)
}

// DepositExcessCollateral is a paid mutator transaction binding the contract method 0x188e0dc6.
//
// Solidity: function depositExcessCollateral() returns()
func (_RETH *RETHTransactorSession) DepositExcessCollateral() (*types.Transaction, error) {
	return _RETH.Contract.DepositExcessCollateral(&_RETH.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RETH *RETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RETH *RETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.IncreaseAllowance(&_RETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_RETH *RETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.IncreaseAllowance(&_RETH.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_RETH *RETHTransactor) Mint(opts *bind.TransactOpts, _ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "mint", _ethAmount, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_RETH *RETHSession) Mint(_ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _RETH.Contract.Mint(&_RETH.TransactOpts, _ethAmount, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _ethAmount, address _to) returns()
func (_RETH *RETHTransactorSession) Mint(_ethAmount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _RETH.Contract.Mint(&_RETH.TransactOpts, _ethAmount, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_RETH *RETHTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_RETH *RETHSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Transfer(&_RETH.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_RETH *RETHTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.Transfer(&_RETH.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_RETH *RETHTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_RETH *RETHSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.TransferFrom(&_RETH.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_RETH *RETHTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RETH.Contract.TransferFrom(&_RETH.TransactOpts, sender, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RETH *RETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RETH *RETHSession) Receive() (*types.Transaction, error) {
	return _RETH.Contract.Receive(&_RETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RETH *RETHTransactorSession) Receive() (*types.Transaction, error) {
	return _RETH.Contract.Receive(&_RETH.TransactOpts)
}

// RETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RETH contract.
type RETHApprovalIterator struct {
	Event *RETHApproval // Event containing the contract specifics and raw log

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
func (it *RETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RETHApproval)
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
		it.Event = new(RETHApproval)
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
func (it *RETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RETHApproval represents a Approval event raised by the RETH contract.
type RETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RETH *RETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RETHApprovalIterator{contract: _RETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RETH *RETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RETHApproval)
				if err := _RETH.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RETH *RETHFilterer) ParseApproval(log types.Log) (*RETHApproval, error) {
	event := new(RETHApproval)
	if err := _RETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RETHEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the RETH contract.
type RETHEtherDepositedIterator struct {
	Event *RETHEtherDeposited // Event containing the contract specifics and raw log

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
func (it *RETHEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RETHEtherDeposited)
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
		it.Event = new(RETHEtherDeposited)
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
func (it *RETHEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RETHEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RETHEtherDeposited represents a EtherDeposited event raised by the RETH contract.
type RETHEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RETH *RETHFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*RETHEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RETH.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &RETHEtherDepositedIterator{contract: _RETH.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RETH *RETHFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *RETHEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RETH.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RETHEtherDeposited)
				if err := _RETH.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RETH *RETHFilterer) ParseEtherDeposited(log types.Log) (*RETHEtherDeposited, error) {
	event := new(RETHEtherDeposited)
	if err := _RETH.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RETHTokensBurnedIterator is returned from FilterTokensBurned and is used to iterate over the raw logs and unpacked data for TokensBurned events raised by the RETH contract.
type RETHTokensBurnedIterator struct {
	Event *RETHTokensBurned // Event containing the contract specifics and raw log

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
func (it *RETHTokensBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RETHTokensBurned)
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
		it.Event = new(RETHTokensBurned)
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
func (it *RETHTokensBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RETHTokensBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RETHTokensBurned represents a TokensBurned event raised by the RETH contract.
type RETHTokensBurned struct {
	From      common.Address
	Amount    *big.Int
	EthAmount *big.Int
	Time      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensBurned is a free log retrieval operation binding the contract event 0x19783b34589160c168487dc7f9c51ae0bcefe67a47d6708fba90f6ce0366d3d1.
//
// Solidity: event TokensBurned(address indexed from, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) FilterTokensBurned(opts *bind.FilterOpts, from []common.Address) (*RETHTokensBurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RETH.contract.FilterLogs(opts, "TokensBurned", fromRule)
	if err != nil {
		return nil, err
	}
	return &RETHTokensBurnedIterator{contract: _RETH.contract, event: "TokensBurned", logs: logs, sub: sub}, nil
}

// WatchTokensBurned is a free log subscription operation binding the contract event 0x19783b34589160c168487dc7f9c51ae0bcefe67a47d6708fba90f6ce0366d3d1.
//
// Solidity: event TokensBurned(address indexed from, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) WatchTokensBurned(opts *bind.WatchOpts, sink chan<- *RETHTokensBurned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RETH.contract.WatchLogs(opts, "TokensBurned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RETHTokensBurned)
				if err := _RETH.contract.UnpackLog(event, "TokensBurned", log); err != nil {
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

// ParseTokensBurned is a log parse operation binding the contract event 0x19783b34589160c168487dc7f9c51ae0bcefe67a47d6708fba90f6ce0366d3d1.
//
// Solidity: event TokensBurned(address indexed from, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) ParseTokensBurned(log types.Log) (*RETHTokensBurned, error) {
	event := new(RETHTokensBurned)
	if err := _RETH.contract.UnpackLog(event, "TokensBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RETHTokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the RETH contract.
type RETHTokensMintedIterator struct {
	Event *RETHTokensMinted // Event containing the contract specifics and raw log

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
func (it *RETHTokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RETHTokensMinted)
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
		it.Event = new(RETHTokensMinted)
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
func (it *RETHTokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RETHTokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RETHTokensMinted represents a TokensMinted event raised by the RETH contract.
type RETHTokensMinted struct {
	To        common.Address
	Amount    *big.Int
	EthAmount *big.Int
	Time      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x6155cfd0fd028b0ca77e8495a60cbe563e8bce8611f0aad6fedbdaafc05d44a2.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) FilterTokensMinted(opts *bind.FilterOpts, to []common.Address) (*RETHTokensMintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RETH.contract.FilterLogs(opts, "TokensMinted", toRule)
	if err != nil {
		return nil, err
	}
	return &RETHTokensMintedIterator{contract: _RETH.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x6155cfd0fd028b0ca77e8495a60cbe563e8bce8611f0aad6fedbdaafc05d44a2.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *RETHTokensMinted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RETH.contract.WatchLogs(opts, "TokensMinted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RETHTokensMinted)
				if err := _RETH.contract.UnpackLog(event, "TokensMinted", log); err != nil {
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

// ParseTokensMinted is a log parse operation binding the contract event 0x6155cfd0fd028b0ca77e8495a60cbe563e8bce8611f0aad6fedbdaafc05d44a2.
//
// Solidity: event TokensMinted(address indexed to, uint256 amount, uint256 ethAmount, uint256 time)
func (_RETH *RETHFilterer) ParseTokensMinted(log types.Log) (*RETHTokensMinted, error) {
	event := new(RETHTokensMinted)
	if err := _RETH.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RETH contract.
type RETHTransferIterator struct {
	Event *RETHTransfer // Event containing the contract specifics and raw log

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
func (it *RETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RETHTransfer)
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
		it.Event = new(RETHTransfer)
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
func (it *RETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RETHTransfer represents a Transfer event raised by the RETH contract.
type RETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RETH *RETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RETHTransferIterator{contract: _RETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RETH *RETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RETHTransfer)
				if err := _RETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RETH *RETHFilterer) ParseTransfer(log types.Log) (*RETHTransfer, error) {
	event := new(RETHTransfer)
	if err := _RETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
