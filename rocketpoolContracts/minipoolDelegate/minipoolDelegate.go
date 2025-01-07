// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package minipoolDelegate

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

// MinipoolDelegateMetaData contains all meta data concerning the MinipoolDelegate contract.
var MinipoolDelegateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBondAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBondAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BondReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawalProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolPrestaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolPromoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolScrubbed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bondAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolVacancyPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ScrubVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beginUserDistribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"calculateNodeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"calculateUserShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canPromote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dissolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_rewardsOnly\",\"type\":\"bool\"}],\"name\":\"distributeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositType\",\"outputs\":[{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeDepositAssigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRefundBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeTopUpValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreLaunchValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreMigrationBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"getScrubVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumMinipoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatusBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatusTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalScrubVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositAssigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositAssignedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDistributed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVacant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bondValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_validatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"preDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bondAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentBalance\",\"type\":\"uint256\"}],\"name\":\"prepareVacancy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"promote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reduceBondAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDistributeAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteScrub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinipoolDelegateABI is the input ABI used to generate the binding from.
// Deprecated: Use MinipoolDelegateMetaData.ABI instead.
var MinipoolDelegateABI = MinipoolDelegateMetaData.ABI

// MinipoolDelegate is an auto generated Go binding around an Ethereum contract.
type MinipoolDelegate struct {
	MinipoolDelegateCaller     // Read-only binding to the contract
	MinipoolDelegateTransactor // Write-only binding to the contract
	MinipoolDelegateFilterer   // Log filterer for contract events
}

// MinipoolDelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinipoolDelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinipoolDelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinipoolDelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinipoolDelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinipoolDelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinipoolDelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinipoolDelegateSession struct {
	Contract     *MinipoolDelegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinipoolDelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinipoolDelegateCallerSession struct {
	Contract *MinipoolDelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MinipoolDelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinipoolDelegateTransactorSession struct {
	Contract     *MinipoolDelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MinipoolDelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinipoolDelegateRaw struct {
	Contract *MinipoolDelegate // Generic contract binding to access the raw methods on
}

// MinipoolDelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinipoolDelegateCallerRaw struct {
	Contract *MinipoolDelegateCaller // Generic read-only contract binding to access the raw methods on
}

// MinipoolDelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinipoolDelegateTransactorRaw struct {
	Contract *MinipoolDelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinipoolDelegate creates a new instance of MinipoolDelegate, bound to a specific deployed contract.
func NewMinipoolDelegate(address common.Address, backend bind.ContractBackend) (*MinipoolDelegate, error) {
	contract, err := bindMinipoolDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegate{MinipoolDelegateCaller: MinipoolDelegateCaller{contract: contract}, MinipoolDelegateTransactor: MinipoolDelegateTransactor{contract: contract}, MinipoolDelegateFilterer: MinipoolDelegateFilterer{contract: contract}}, nil
}

// NewMinipoolDelegateCaller creates a new read-only instance of MinipoolDelegate, bound to a specific deployed contract.
func NewMinipoolDelegateCaller(address common.Address, caller bind.ContractCaller) (*MinipoolDelegateCaller, error) {
	contract, err := bindMinipoolDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateCaller{contract: contract}, nil
}

// NewMinipoolDelegateTransactor creates a new write-only instance of MinipoolDelegate, bound to a specific deployed contract.
func NewMinipoolDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*MinipoolDelegateTransactor, error) {
	contract, err := bindMinipoolDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateTransactor{contract: contract}, nil
}

// NewMinipoolDelegateFilterer creates a new log filterer instance of MinipoolDelegate, bound to a specific deployed contract.
func NewMinipoolDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*MinipoolDelegateFilterer, error) {
	contract, err := bindMinipoolDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateFilterer{contract: contract}, nil
}

// bindMinipoolDelegate binds a generic wrapper to an already deployed contract.
func bindMinipoolDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinipoolDelegateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinipoolDelegate *MinipoolDelegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinipoolDelegate.Contract.MinipoolDelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinipoolDelegate *MinipoolDelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.MinipoolDelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinipoolDelegate *MinipoolDelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.MinipoolDelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinipoolDelegate *MinipoolDelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinipoolDelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinipoolDelegate *MinipoolDelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinipoolDelegate *MinipoolDelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.contract.Transact(opts, method, params...)
}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) CalculateNodeShare(opts *bind.CallOpts, _balance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "calculateNodeShare", _balance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) CalculateNodeShare(_balance *big.Int) (*big.Int, error) {
	return _MinipoolDelegate.Contract.CalculateNodeShare(&_MinipoolDelegate.CallOpts, _balance)
}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) CalculateNodeShare(_balance *big.Int) (*big.Int, error) {
	return _MinipoolDelegate.Contract.CalculateNodeShare(&_MinipoolDelegate.CallOpts, _balance)
}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) CalculateUserShare(opts *bind.CallOpts, _balance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "calculateUserShare", _balance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) CalculateUserShare(_balance *big.Int) (*big.Int, error) {
	return _MinipoolDelegate.Contract.CalculateUserShare(&_MinipoolDelegate.CallOpts, _balance)
}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) CalculateUserShare(_balance *big.Int) (*big.Int, error) {
	return _MinipoolDelegate.Contract.CalculateUserShare(&_MinipoolDelegate.CallOpts, _balance)
}

// CanPromote is a free data retrieval call binding the contract method 0xc9c36b27.
//
// Solidity: function canPromote() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) CanPromote(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "canPromote")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPromote is a free data retrieval call binding the contract method 0xc9c36b27.
//
// Solidity: function canPromote() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) CanPromote() (bool, error) {
	return _MinipoolDelegate.Contract.CanPromote(&_MinipoolDelegate.CallOpts)
}

// CanPromote is a free data retrieval call binding the contract method 0xc9c36b27.
//
// Solidity: function canPromote() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) CanPromote() (bool, error) {
	return _MinipoolDelegate.Contract.CanPromote(&_MinipoolDelegate.CallOpts)
}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) CanStake(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "canStake")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) CanStake() (bool, error) {
	return _MinipoolDelegate.Contract.CanStake(&_MinipoolDelegate.CallOpts)
}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) CanStake() (bool, error) {
	return _MinipoolDelegate.Contract.CanStake(&_MinipoolDelegate.CallOpts)
}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetDepositType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getDepositType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateSession) GetDepositType() (uint8, error) {
	return _MinipoolDelegate.Contract.GetDepositType(&_MinipoolDelegate.CallOpts)
}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetDepositType() (uint8, error) {
	return _MinipoolDelegate.Contract.GetDepositType(&_MinipoolDelegate.CallOpts)
}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetFinalised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getFinalised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetFinalised() (bool, error) {
	return _MinipoolDelegate.Contract.GetFinalised(&_MinipoolDelegate.CallOpts)
}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetFinalised() (bool, error) {
	return _MinipoolDelegate.Contract.GetFinalised(&_MinipoolDelegate.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeAddress() (common.Address, error) {
	return _MinipoolDelegate.Contract.GetNodeAddress(&_MinipoolDelegate.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeAddress() (common.Address, error) {
	return _MinipoolDelegate.Contract.GetNodeAddress(&_MinipoolDelegate.CallOpts)
}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeDepositAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeDepositAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeDepositAssigned() (bool, error) {
	return _MinipoolDelegate.Contract.GetNodeDepositAssigned(&_MinipoolDelegate.CallOpts)
}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeDepositAssigned() (bool, error) {
	return _MinipoolDelegate.Contract.GetNodeDepositAssigned(&_MinipoolDelegate.CallOpts)
}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeDepositBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeDepositBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeDepositBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeDepositBalance(&_MinipoolDelegate.CallOpts)
}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeDepositBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeDepositBalance(&_MinipoolDelegate.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeFee() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeFee(&_MinipoolDelegate.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeFee() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeFee(&_MinipoolDelegate.CallOpts)
}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeRefundBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeRefundBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeRefundBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeRefundBalance(&_MinipoolDelegate.CallOpts)
}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeRefundBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeRefundBalance(&_MinipoolDelegate.CallOpts)
}

// GetNodeTopUpValue is a free data retrieval call binding the contract method 0xd2ceebd1.
//
// Solidity: function getNodeTopUpValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetNodeTopUpValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getNodeTopUpValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeTopUpValue is a free data retrieval call binding the contract method 0xd2ceebd1.
//
// Solidity: function getNodeTopUpValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetNodeTopUpValue() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeTopUpValue(&_MinipoolDelegate.CallOpts)
}

// GetNodeTopUpValue is a free data retrieval call binding the contract method 0xd2ceebd1.
//
// Solidity: function getNodeTopUpValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetNodeTopUpValue() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetNodeTopUpValue(&_MinipoolDelegate.CallOpts)
}

// GetPreLaunchValue is a free data retrieval call binding the contract method 0xd6047def.
//
// Solidity: function getPreLaunchValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetPreLaunchValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getPreLaunchValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreLaunchValue is a free data retrieval call binding the contract method 0xd6047def.
//
// Solidity: function getPreLaunchValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetPreLaunchValue() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetPreLaunchValue(&_MinipoolDelegate.CallOpts)
}

// GetPreLaunchValue is a free data retrieval call binding the contract method 0xd6047def.
//
// Solidity: function getPreLaunchValue() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetPreLaunchValue() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetPreLaunchValue(&_MinipoolDelegate.CallOpts)
}

// GetPreMigrationBalance is a free data retrieval call binding the contract method 0x49b42321.
//
// Solidity: function getPreMigrationBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetPreMigrationBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getPreMigrationBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreMigrationBalance is a free data retrieval call binding the contract method 0x49b42321.
//
// Solidity: function getPreMigrationBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetPreMigrationBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetPreMigrationBalance(&_MinipoolDelegate.CallOpts)
}

// GetPreMigrationBalance is a free data retrieval call binding the contract method 0x49b42321.
//
// Solidity: function getPreMigrationBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetPreMigrationBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetPreMigrationBalance(&_MinipoolDelegate.CallOpts)
}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetScrubVoted(opts *bind.CallOpts, _member common.Address) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getScrubVoted", _member)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetScrubVoted(_member common.Address) (bool, error) {
	return _MinipoolDelegate.Contract.GetScrubVoted(&_MinipoolDelegate.CallOpts, _member)
}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetScrubVoted(_member common.Address) (bool, error) {
	return _MinipoolDelegate.Contract.GetScrubVoted(&_MinipoolDelegate.CallOpts, _member)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateSession) GetStatus() (uint8, error) {
	return _MinipoolDelegate.Contract.GetStatus(&_MinipoolDelegate.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetStatus() (uint8, error) {
	return _MinipoolDelegate.Contract.GetStatus(&_MinipoolDelegate.CallOpts)
}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetStatusBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getStatusBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetStatusBlock() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetStatusBlock(&_MinipoolDelegate.CallOpts)
}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetStatusBlock() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetStatusBlock(&_MinipoolDelegate.CallOpts)
}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetStatusTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getStatusTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetStatusTime() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetStatusTime(&_MinipoolDelegate.CallOpts)
}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetStatusTime() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetStatusTime(&_MinipoolDelegate.CallOpts)
}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetTotalScrubVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getTotalScrubVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetTotalScrubVotes() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetTotalScrubVotes(&_MinipoolDelegate.CallOpts)
}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetTotalScrubVotes() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetTotalScrubVotes(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetUserDepositAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getUserDepositAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetUserDepositAssigned() (bool, error) {
	return _MinipoolDelegate.Contract.GetUserDepositAssigned(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetUserDepositAssigned() (bool, error) {
	return _MinipoolDelegate.Contract.GetUserDepositAssigned(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetUserDepositAssignedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getUserDepositAssignedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetUserDepositAssignedTime() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetUserDepositAssignedTime(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetUserDepositAssignedTime() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetUserDepositAssignedTime(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetUserDepositBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getUserDepositBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateSession) GetUserDepositBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetUserDepositBalance(&_MinipoolDelegate.CallOpts)
}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetUserDepositBalance() (*big.Int, error) {
	return _MinipoolDelegate.Contract.GetUserDepositBalance(&_MinipoolDelegate.CallOpts)
}

// GetUserDistributed is a free data retrieval call binding the contract method 0x7bfaef7d.
//
// Solidity: function getUserDistributed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetUserDistributed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getUserDistributed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetUserDistributed is a free data retrieval call binding the contract method 0x7bfaef7d.
//
// Solidity: function getUserDistributed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetUserDistributed() (bool, error) {
	return _MinipoolDelegate.Contract.GetUserDistributed(&_MinipoolDelegate.CallOpts)
}

// GetUserDistributed is a free data retrieval call binding the contract method 0x7bfaef7d.
//
// Solidity: function getUserDistributed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetUserDistributed() (bool, error) {
	return _MinipoolDelegate.Contract.GetUserDistributed(&_MinipoolDelegate.CallOpts)
}

// GetVacant is a free data retrieval call binding the contract method 0xbbe38fe1.
//
// Solidity: function getVacant() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) GetVacant(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "getVacant")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVacant is a free data retrieval call binding the contract method 0xbbe38fe1.
//
// Solidity: function getVacant() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) GetVacant() (bool, error) {
	return _MinipoolDelegate.Contract.GetVacant(&_MinipoolDelegate.CallOpts)
}

// GetVacant is a free data retrieval call binding the contract method 0xbbe38fe1.
//
// Solidity: function getVacant() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) GetVacant() (bool, error) {
	return _MinipoolDelegate.Contract.GetVacant(&_MinipoolDelegate.CallOpts)
}

// UserDistributeAllowed is a free data retrieval call binding the contract method 0x23e4e3e4.
//
// Solidity: function userDistributeAllowed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCaller) UserDistributeAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "userDistributeAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDistributeAllowed is a free data retrieval call binding the contract method 0x23e4e3e4.
//
// Solidity: function userDistributeAllowed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateSession) UserDistributeAllowed() (bool, error) {
	return _MinipoolDelegate.Contract.UserDistributeAllowed(&_MinipoolDelegate.CallOpts)
}

// UserDistributeAllowed is a free data retrieval call binding the contract method 0x23e4e3e4.
//
// Solidity: function userDistributeAllowed() view returns(bool)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) UserDistributeAllowed() (bool, error) {
	return _MinipoolDelegate.Contract.UserDistributeAllowed(&_MinipoolDelegate.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MinipoolDelegate.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateSession) Version() (uint8, error) {
	return _MinipoolDelegate.Contract.Version(&_MinipoolDelegate.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_MinipoolDelegate *MinipoolDelegateCallerSession) Version() (uint8, error) {
	return _MinipoolDelegate.Contract.Version(&_MinipoolDelegate.CallOpts)
}

// BeginUserDistribute is a paid mutator transaction binding the contract method 0xf09fa332.
//
// Solidity: function beginUserDistribute() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) BeginUserDistribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "beginUserDistribute")
}

// BeginUserDistribute is a paid mutator transaction binding the contract method 0xf09fa332.
//
// Solidity: function beginUserDistribute() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) BeginUserDistribute() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.BeginUserDistribute(&_MinipoolDelegate.TransactOpts)
}

// BeginUserDistribute is a paid mutator transaction binding the contract method 0xf09fa332.
//
// Solidity: function beginUserDistribute() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) BeginUserDistribute() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.BeginUserDistribute(&_MinipoolDelegate.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Close() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Close(&_MinipoolDelegate.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Close() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Close(&_MinipoolDelegate.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Deposit() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Deposit(&_MinipoolDelegate.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Deposit() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Deposit(&_MinipoolDelegate.TransactOpts)
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Dissolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "dissolve")
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Dissolve() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Dissolve(&_MinipoolDelegate.TransactOpts)
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Dissolve() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Dissolve(&_MinipoolDelegate.TransactOpts)
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x54efc6e5.
//
// Solidity: function distributeBalance(bool _rewardsOnly) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) DistributeBalance(opts *bind.TransactOpts, _rewardsOnly bool) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "distributeBalance", _rewardsOnly)
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x54efc6e5.
//
// Solidity: function distributeBalance(bool _rewardsOnly) returns()
func (_MinipoolDelegate *MinipoolDelegateSession) DistributeBalance(_rewardsOnly bool) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.DistributeBalance(&_MinipoolDelegate.TransactOpts, _rewardsOnly)
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x54efc6e5.
//
// Solidity: function distributeBalance(bool _rewardsOnly) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) DistributeBalance(_rewardsOnly bool) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.DistributeBalance(&_MinipoolDelegate.TransactOpts, _rewardsOnly)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Finalise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "finalise")
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Finalise() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Finalise(&_MinipoolDelegate.TransactOpts)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Finalise() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Finalise(&_MinipoolDelegate.TransactOpts)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _nodeAddress) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Initialise(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "initialise", _nodeAddress)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _nodeAddress) returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Initialise(_nodeAddress common.Address) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Initialise(&_MinipoolDelegate.TransactOpts, _nodeAddress)
}

// Initialise is a paid mutator transaction binding the contract method 0x9d6a890f.
//
// Solidity: function initialise(address _nodeAddress) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Initialise(_nodeAddress common.Address) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Initialise(&_MinipoolDelegate.TransactOpts, _nodeAddress)
}

// PreDeposit is a paid mutator transaction binding the contract method 0x3ca742e9.
//
// Solidity: function preDeposit(uint256 _bondValue, bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) PreDeposit(opts *bind.TransactOpts, _bondValue *big.Int, _validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "preDeposit", _bondValue, _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// PreDeposit is a paid mutator transaction binding the contract method 0x3ca742e9.
//
// Solidity: function preDeposit(uint256 _bondValue, bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_MinipoolDelegate *MinipoolDelegateSession) PreDeposit(_bondValue *big.Int, _validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.PreDeposit(&_MinipoolDelegate.TransactOpts, _bondValue, _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// PreDeposit is a paid mutator transaction binding the contract method 0x3ca742e9.
//
// Solidity: function preDeposit(uint256 _bondValue, bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) PreDeposit(_bondValue *big.Int, _validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.PreDeposit(&_MinipoolDelegate.TransactOpts, _bondValue, _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// PrepareVacancy is a paid mutator transaction binding the contract method 0x0871ffef.
//
// Solidity: function prepareVacancy(uint256 _bondAmount, uint256 _currentBalance) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) PrepareVacancy(opts *bind.TransactOpts, _bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "prepareVacancy", _bondAmount, _currentBalance)
}

// PrepareVacancy is a paid mutator transaction binding the contract method 0x0871ffef.
//
// Solidity: function prepareVacancy(uint256 _bondAmount, uint256 _currentBalance) returns()
func (_MinipoolDelegate *MinipoolDelegateSession) PrepareVacancy(_bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.PrepareVacancy(&_MinipoolDelegate.TransactOpts, _bondAmount, _currentBalance)
}

// PrepareVacancy is a paid mutator transaction binding the contract method 0x0871ffef.
//
// Solidity: function prepareVacancy(uint256 _bondAmount, uint256 _currentBalance) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) PrepareVacancy(_bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.PrepareVacancy(&_MinipoolDelegate.TransactOpts, _bondAmount, _currentBalance)
}

// Promote is a paid mutator transaction binding the contract method 0x13dc01dc.
//
// Solidity: function promote() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Promote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "promote")
}

// Promote is a paid mutator transaction binding the contract method 0x13dc01dc.
//
// Solidity: function promote() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Promote() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Promote(&_MinipoolDelegate.TransactOpts)
}

// Promote is a paid mutator transaction binding the contract method 0x13dc01dc.
//
// Solidity: function promote() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Promote() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Promote(&_MinipoolDelegate.TransactOpts)
}

// ReduceBondAmount is a paid mutator transaction binding the contract method 0xd191ea9c.
//
// Solidity: function reduceBondAmount() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) ReduceBondAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "reduceBondAmount")
}

// ReduceBondAmount is a paid mutator transaction binding the contract method 0xd191ea9c.
//
// Solidity: function reduceBondAmount() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) ReduceBondAmount() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.ReduceBondAmount(&_MinipoolDelegate.TransactOpts)
}

// ReduceBondAmount is a paid mutator transaction binding the contract method 0xd191ea9c.
//
// Solidity: function reduceBondAmount() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) ReduceBondAmount() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.ReduceBondAmount(&_MinipoolDelegate.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Refund() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Refund(&_MinipoolDelegate.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Refund() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Refund(&_MinipoolDelegate.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Slash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "slash")
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Slash() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Slash(&_MinipoolDelegate.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Slash() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Slash(&_MinipoolDelegate.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) Stake(opts *bind.TransactOpts, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "stake", _validatorSignature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_MinipoolDelegate *MinipoolDelegateSession) Stake(_validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Stake(&_MinipoolDelegate.TransactOpts, _validatorSignature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) Stake(_validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.Stake(&_MinipoolDelegate.TransactOpts, _validatorSignature, _depositDataRoot)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) UserDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "userDeposit")
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateSession) UserDeposit() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.UserDeposit(&_MinipoolDelegate.TransactOpts)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) UserDeposit() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.UserDeposit(&_MinipoolDelegate.TransactOpts)
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactor) VoteScrub(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinipoolDelegate.contract.Transact(opts, "voteScrub")
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_MinipoolDelegate *MinipoolDelegateSession) VoteScrub() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.VoteScrub(&_MinipoolDelegate.TransactOpts)
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_MinipoolDelegate *MinipoolDelegateTransactorSession) VoteScrub() (*types.Transaction, error) {
	return _MinipoolDelegate.Contract.VoteScrub(&_MinipoolDelegate.TransactOpts)
}

// MinipoolDelegateBondReducedIterator is returned from FilterBondReduced and is used to iterate over the raw logs and unpacked data for BondReduced events raised by the MinipoolDelegate contract.
type MinipoolDelegateBondReducedIterator struct {
	Event *MinipoolDelegateBondReduced // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateBondReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateBondReduced)
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
		it.Event = new(MinipoolDelegateBondReduced)
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
func (it *MinipoolDelegateBondReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateBondReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateBondReduced represents a BondReduced event raised by the MinipoolDelegate contract.
type MinipoolDelegateBondReduced struct {
	PreviousBondAmount *big.Int
	NewBondAmount      *big.Int
	Time               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBondReduced is a free log retrieval operation binding the contract event 0x90e131460b9acb17565f1719b9ebc49998aec6b07a4743a09b1b700545769eb6.
//
// Solidity: event BondReduced(uint256 previousBondAmount, uint256 newBondAmount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterBondReduced(opts *bind.FilterOpts) (*MinipoolDelegateBondReducedIterator, error) {

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "BondReduced")
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateBondReducedIterator{contract: _MinipoolDelegate.contract, event: "BondReduced", logs: logs, sub: sub}, nil
}

// WatchBondReduced is a free log subscription operation binding the contract event 0x90e131460b9acb17565f1719b9ebc49998aec6b07a4743a09b1b700545769eb6.
//
// Solidity: event BondReduced(uint256 previousBondAmount, uint256 newBondAmount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchBondReduced(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateBondReduced) (event.Subscription, error) {

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "BondReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateBondReduced)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "BondReduced", log); err != nil {
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

// ParseBondReduced is a log parse operation binding the contract event 0x90e131460b9acb17565f1719b9ebc49998aec6b07a4743a09b1b700545769eb6.
//
// Solidity: event BondReduced(uint256 previousBondAmount, uint256 newBondAmount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseBondReduced(log types.Log) (*MinipoolDelegateBondReduced, error) {
	event := new(MinipoolDelegateBondReduced)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "BondReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherDepositedIterator struct {
	Event *MinipoolDelegateEtherDeposited // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateEtherDeposited)
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
		it.Event = new(MinipoolDelegateEtherDeposited)
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
func (it *MinipoolDelegateEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateEtherDeposited represents a EtherDeposited event raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*MinipoolDelegateEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateEtherDepositedIterator{contract: _MinipoolDelegate.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateEtherDeposited)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseEtherDeposited(log types.Log) (*MinipoolDelegateEtherDeposited, error) {
	event := new(MinipoolDelegateEtherDeposited)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateEtherWithdrawalProcessedIterator is returned from FilterEtherWithdrawalProcessed and is used to iterate over the raw logs and unpacked data for EtherWithdrawalProcessed events raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherWithdrawalProcessedIterator struct {
	Event *MinipoolDelegateEtherWithdrawalProcessed // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateEtherWithdrawalProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateEtherWithdrawalProcessed)
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
		it.Event = new(MinipoolDelegateEtherWithdrawalProcessed)
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
func (it *MinipoolDelegateEtherWithdrawalProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateEtherWithdrawalProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateEtherWithdrawalProcessed represents a EtherWithdrawalProcessed event raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherWithdrawalProcessed struct {
	Executed     common.Address
	NodeAmount   *big.Int
	UserAmount   *big.Int
	TotalBalance *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawalProcessed is a free log retrieval operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterEtherWithdrawalProcessed(opts *bind.FilterOpts, executed []common.Address) (*MinipoolDelegateEtherWithdrawalProcessedIterator, error) {

	var executedRule []interface{}
	for _, executedItem := range executed {
		executedRule = append(executedRule, executedItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "EtherWithdrawalProcessed", executedRule)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateEtherWithdrawalProcessedIterator{contract: _MinipoolDelegate.contract, event: "EtherWithdrawalProcessed", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawalProcessed is a free log subscription operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchEtherWithdrawalProcessed(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateEtherWithdrawalProcessed, executed []common.Address) (event.Subscription, error) {

	var executedRule []interface{}
	for _, executedItem := range executed {
		executedRule = append(executedRule, executedItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "EtherWithdrawalProcessed", executedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateEtherWithdrawalProcessed)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherWithdrawalProcessed", log); err != nil {
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

// ParseEtherWithdrawalProcessed is a log parse operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseEtherWithdrawalProcessed(log types.Log) (*MinipoolDelegateEtherWithdrawalProcessed, error) {
	event := new(MinipoolDelegateEtherWithdrawalProcessed)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherWithdrawalProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherWithdrawnIterator struct {
	Event *MinipoolDelegateEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateEtherWithdrawn)
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
		it.Event = new(MinipoolDelegateEtherWithdrawn)
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
func (it *MinipoolDelegateEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateEtherWithdrawn represents a EtherWithdrawn event raised by the MinipoolDelegate contract.
type MinipoolDelegateEtherWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, to []common.Address) (*MinipoolDelegateEtherWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "EtherWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateEtherWithdrawnIterator{contract: _MinipoolDelegate.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateEtherWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "EtherWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateEtherWithdrawn)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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

// ParseEtherWithdrawn is a log parse operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseEtherWithdrawn(log types.Log) (*MinipoolDelegateEtherWithdrawn, error) {
	event := new(MinipoolDelegateEtherWithdrawn)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateMinipoolPrestakedIterator is returned from FilterMinipoolPrestaked and is used to iterate over the raw logs and unpacked data for MinipoolPrestaked events raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolPrestakedIterator struct {
	Event *MinipoolDelegateMinipoolPrestaked // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateMinipoolPrestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateMinipoolPrestaked)
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
		it.Event = new(MinipoolDelegateMinipoolPrestaked)
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
func (it *MinipoolDelegateMinipoolPrestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateMinipoolPrestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateMinipoolPrestaked represents a MinipoolPrestaked event raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolPrestaked struct {
	ValidatorPubkey       []byte
	ValidatorSignature    []byte
	DepositDataRoot       [32]byte
	Amount                *big.Int
	WithdrawalCredentials []byte
	Time                  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMinipoolPrestaked is a free log retrieval operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterMinipoolPrestaked(opts *bind.FilterOpts) (*MinipoolDelegateMinipoolPrestakedIterator, error) {

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "MinipoolPrestaked")
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateMinipoolPrestakedIterator{contract: _MinipoolDelegate.contract, event: "MinipoolPrestaked", logs: logs, sub: sub}, nil
}

// WatchMinipoolPrestaked is a free log subscription operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchMinipoolPrestaked(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateMinipoolPrestaked) (event.Subscription, error) {

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "MinipoolPrestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateMinipoolPrestaked)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolPrestaked", log); err != nil {
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

// ParseMinipoolPrestaked is a log parse operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseMinipoolPrestaked(log types.Log) (*MinipoolDelegateMinipoolPrestaked, error) {
	event := new(MinipoolDelegateMinipoolPrestaked)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolPrestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateMinipoolPromotedIterator is returned from FilterMinipoolPromoted and is used to iterate over the raw logs and unpacked data for MinipoolPromoted events raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolPromotedIterator struct {
	Event *MinipoolDelegateMinipoolPromoted // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateMinipoolPromotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateMinipoolPromoted)
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
		it.Event = new(MinipoolDelegateMinipoolPromoted)
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
func (it *MinipoolDelegateMinipoolPromotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateMinipoolPromotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateMinipoolPromoted represents a MinipoolPromoted event raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolPromoted struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMinipoolPromoted is a free log retrieval operation binding the contract event 0xa5c869f853c40dbf5557240b202402a69e253565e0eb171fa239d8e95b1b1c2e.
//
// Solidity: event MinipoolPromoted(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterMinipoolPromoted(opts *bind.FilterOpts) (*MinipoolDelegateMinipoolPromotedIterator, error) {

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "MinipoolPromoted")
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateMinipoolPromotedIterator{contract: _MinipoolDelegate.contract, event: "MinipoolPromoted", logs: logs, sub: sub}, nil
}

// WatchMinipoolPromoted is a free log subscription operation binding the contract event 0xa5c869f853c40dbf5557240b202402a69e253565e0eb171fa239d8e95b1b1c2e.
//
// Solidity: event MinipoolPromoted(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchMinipoolPromoted(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateMinipoolPromoted) (event.Subscription, error) {

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "MinipoolPromoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateMinipoolPromoted)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolPromoted", log); err != nil {
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

// ParseMinipoolPromoted is a log parse operation binding the contract event 0xa5c869f853c40dbf5557240b202402a69e253565e0eb171fa239d8e95b1b1c2e.
//
// Solidity: event MinipoolPromoted(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseMinipoolPromoted(log types.Log) (*MinipoolDelegateMinipoolPromoted, error) {
	event := new(MinipoolDelegateMinipoolPromoted)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolPromoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateMinipoolScrubbedIterator is returned from FilterMinipoolScrubbed and is used to iterate over the raw logs and unpacked data for MinipoolScrubbed events raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolScrubbedIterator struct {
	Event *MinipoolDelegateMinipoolScrubbed // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateMinipoolScrubbedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateMinipoolScrubbed)
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
		it.Event = new(MinipoolDelegateMinipoolScrubbed)
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
func (it *MinipoolDelegateMinipoolScrubbedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateMinipoolScrubbedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateMinipoolScrubbed represents a MinipoolScrubbed event raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolScrubbed struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMinipoolScrubbed is a free log retrieval operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterMinipoolScrubbed(opts *bind.FilterOpts) (*MinipoolDelegateMinipoolScrubbedIterator, error) {

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "MinipoolScrubbed")
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateMinipoolScrubbedIterator{contract: _MinipoolDelegate.contract, event: "MinipoolScrubbed", logs: logs, sub: sub}, nil
}

// WatchMinipoolScrubbed is a free log subscription operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchMinipoolScrubbed(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateMinipoolScrubbed) (event.Subscription, error) {

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "MinipoolScrubbed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateMinipoolScrubbed)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolScrubbed", log); err != nil {
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

// ParseMinipoolScrubbed is a log parse operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseMinipoolScrubbed(log types.Log) (*MinipoolDelegateMinipoolScrubbed, error) {
	event := new(MinipoolDelegateMinipoolScrubbed)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolScrubbed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateMinipoolVacancyPreparedIterator is returned from FilterMinipoolVacancyPrepared and is used to iterate over the raw logs and unpacked data for MinipoolVacancyPrepared events raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolVacancyPreparedIterator struct {
	Event *MinipoolDelegateMinipoolVacancyPrepared // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateMinipoolVacancyPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateMinipoolVacancyPrepared)
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
		it.Event = new(MinipoolDelegateMinipoolVacancyPrepared)
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
func (it *MinipoolDelegateMinipoolVacancyPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateMinipoolVacancyPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateMinipoolVacancyPrepared represents a MinipoolVacancyPrepared event raised by the MinipoolDelegate contract.
type MinipoolDelegateMinipoolVacancyPrepared struct {
	BondAmount     *big.Int
	CurrentBalance *big.Int
	Time           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinipoolVacancyPrepared is a free log retrieval operation binding the contract event 0xf7cb92de8d4b074aafcfa9bdb83842b1ef40f49087a9eb04996d68a012de105c.
//
// Solidity: event MinipoolVacancyPrepared(uint256 bondAmount, uint256 currentBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterMinipoolVacancyPrepared(opts *bind.FilterOpts) (*MinipoolDelegateMinipoolVacancyPreparedIterator, error) {

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "MinipoolVacancyPrepared")
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateMinipoolVacancyPreparedIterator{contract: _MinipoolDelegate.contract, event: "MinipoolVacancyPrepared", logs: logs, sub: sub}, nil
}

// WatchMinipoolVacancyPrepared is a free log subscription operation binding the contract event 0xf7cb92de8d4b074aafcfa9bdb83842b1ef40f49087a9eb04996d68a012de105c.
//
// Solidity: event MinipoolVacancyPrepared(uint256 bondAmount, uint256 currentBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchMinipoolVacancyPrepared(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateMinipoolVacancyPrepared) (event.Subscription, error) {

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "MinipoolVacancyPrepared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateMinipoolVacancyPrepared)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolVacancyPrepared", log); err != nil {
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

// ParseMinipoolVacancyPrepared is a log parse operation binding the contract event 0xf7cb92de8d4b074aafcfa9bdb83842b1ef40f49087a9eb04996d68a012de105c.
//
// Solidity: event MinipoolVacancyPrepared(uint256 bondAmount, uint256 currentBalance, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseMinipoolVacancyPrepared(log types.Log) (*MinipoolDelegateMinipoolVacancyPrepared, error) {
	event := new(MinipoolDelegateMinipoolVacancyPrepared)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "MinipoolVacancyPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateScrubVotedIterator is returned from FilterScrubVoted and is used to iterate over the raw logs and unpacked data for ScrubVoted events raised by the MinipoolDelegate contract.
type MinipoolDelegateScrubVotedIterator struct {
	Event *MinipoolDelegateScrubVoted // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateScrubVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateScrubVoted)
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
		it.Event = new(MinipoolDelegateScrubVoted)
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
func (it *MinipoolDelegateScrubVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateScrubVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateScrubVoted represents a ScrubVoted event raised by the MinipoolDelegate contract.
type MinipoolDelegateScrubVoted struct {
	Member common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterScrubVoted is a free log retrieval operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterScrubVoted(opts *bind.FilterOpts, member []common.Address) (*MinipoolDelegateScrubVotedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "ScrubVoted", memberRule)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateScrubVotedIterator{contract: _MinipoolDelegate.contract, event: "ScrubVoted", logs: logs, sub: sub}, nil
}

// WatchScrubVoted is a free log subscription operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchScrubVoted(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateScrubVoted, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "ScrubVoted", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateScrubVoted)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "ScrubVoted", log); err != nil {
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

// ParseScrubVoted is a log parse operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseScrubVoted(log types.Log) (*MinipoolDelegateScrubVoted, error) {
	event := new(MinipoolDelegateScrubVoted)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "ScrubVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinipoolDelegateStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the MinipoolDelegate contract.
type MinipoolDelegateStatusUpdatedIterator struct {
	Event *MinipoolDelegateStatusUpdated // Event containing the contract specifics and raw log

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
func (it *MinipoolDelegateStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinipoolDelegateStatusUpdated)
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
		it.Event = new(MinipoolDelegateStatusUpdated)
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
func (it *MinipoolDelegateStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinipoolDelegateStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinipoolDelegateStatusUpdated represents a StatusUpdated event raised by the MinipoolDelegate contract.
type MinipoolDelegateStatusUpdated struct {
	Status uint8
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) FilterStatusUpdated(opts *bind.FilterOpts, status []uint8) (*MinipoolDelegateStatusUpdatedIterator, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.FilterLogs(opts, "StatusUpdated", statusRule)
	if err != nil {
		return nil, err
	}
	return &MinipoolDelegateStatusUpdatedIterator{contract: _MinipoolDelegate.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *MinipoolDelegateStatusUpdated, status []uint8) (event.Subscription, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _MinipoolDelegate.contract.WatchLogs(opts, "StatusUpdated", statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinipoolDelegateStatusUpdated)
				if err := _MinipoolDelegate.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
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

// ParseStatusUpdated is a log parse operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_MinipoolDelegate *MinipoolDelegateFilterer) ParseStatusUpdated(log types.Log) (*MinipoolDelegateStatusUpdated, error) {
	event := new(MinipoolDelegateStatusUpdated)
	if err := _MinipoolDelegate.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
