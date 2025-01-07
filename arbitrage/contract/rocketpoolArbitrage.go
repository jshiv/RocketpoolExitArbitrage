// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uniswapPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"name\":\"Arbitrage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RETH\",\"outputs\":[{\"internalType\":\"contractIRETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uniswapPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"}],\"name\":\"arb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_amountRETHDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_amountWETHDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052731d8f8f00cfa6758d7be78336684788fb0ee0fa4673ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156055575f5ffd5b50608051610d4461006b5f395f5050610d445ff3fe608060405260043610610042575f3560e01c80635ac6a53b1461004d5780638265b12114610075578063ad5c46481461009f578063fa461e33146100c957610049565b3661004957005b5f5ffd5b348015610058575f5ffd5b50610073600480360381019061006e91906105d3565b6100f1565b005b348015610080575f5ffd5b50610089610303565b604051610096919061067e565b60405180910390f35b3480156100aa575f5ffd5b506100b361031b565b6040516100c091906106b7565b60405180910390f35b3480156100d4575f5ffd5b506100ef60048036038101906100ea9190610764565b610333565b005b8273ffffffffffffffffffffffffffffffffffffffff1663128acb08305f8573fffd8963efd1fc6a506488495d951d5263988d2688604051602001610136919061081a565b6040516020818303038152906040526040518663ffffffff1660e01b81526004016101659594939291906108eb565b60408051808303815f875af1158015610180573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101a49190610957565b50505f479050818110156101ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e490610a15565b60405180910390fd5b5f3373ffffffffffffffffffffffffffffffffffffffff168260405161021290610a60565b5f6040518083038185875af1925050503d805f811461024c576040519150601f19603f3d011682016040523d82523d5f602084013e610251565b606091505b5050905080610295576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028c90610abe565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f2a447723197b74967a6126c418f254abf2f6e9dc7b9c4f3a09321daaba444d1f86856040516102f4929190610aeb565b60405180910390a35050505050565b73ae78736cd615f374d3085123a210448e74fc639381565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b5f82828101906103439190610b4d565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103aa90610be8565b60405180910390fd5b73ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166342966c68866103ed90610c33565b6040518263ffffffff1660e01b81526004016104099190610c79565b5f604051808303815f87803b158015610420575f5ffd5b505af1158015610432573d5f5f3e3d5ffd5b5050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004015f604051808303818588803b158015610490575f5ffd5b505af11580156104a2573d5f5f3e3d5ffd5b505050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb82866040518363ffffffff1660e01b81526004016104f6929190610c92565b6020604051808303815f875af1158015610512573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105369190610ce3565b505050505050565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61056f82610546565b9050919050565b61057f81610565565b8114610589575f5ffd5b50565b5f8135905061059a81610576565b92915050565b5f819050919050565b6105b2816105a0565b81146105bc575f5ffd5b50565b5f813590506105cd816105a9565b92915050565b5f5f5f606084860312156105ea576105e961053e565b5b5f6105f78682870161058c565b9350506020610608868287016105bf565b9250506040610619868287016105bf565b9150509250925092565b5f819050919050565b5f61064661064161063c84610546565b610623565b610546565b9050919050565b5f6106578261062c565b9050919050565b5f6106688261064d565b9050919050565b6106788161065e565b82525050565b5f6020820190506106915f83018461066f565b92915050565b5f6106a18261064d565b9050919050565b6106b181610697565b82525050565b5f6020820190506106ca5f8301846106a8565b92915050565b5f819050919050565b6106e2816106d0565b81146106ec575f5ffd5b50565b5f813590506106fd816106d9565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261072457610723610703565b5b8235905067ffffffffffffffff81111561074157610740610707565b5b60208301915083600182028301111561075d5761075c61070b565b5b9250929050565b5f5f5f5f6060858703121561077c5761077b61053e565b5b5f610789878288016106ef565b945050602061079a878288016106ef565b935050604085013567ffffffffffffffff8111156107bb576107ba610542565b5b6107c78782880161070f565b925092505092959194509250565b5f8160601b9050919050565b5f6107eb826107d5565b9050919050565b5f6107fc826107e1565b9050919050565b61081461080f82610565565b6107f2565b82525050565b5f6108258284610803565b60148201915081905092915050565b61083d81610565565b82525050565b5f8115159050919050565b61085781610843565b82525050565b610866816106d0565b82525050565b61087581610546565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6108bd8261087b565b6108c78185610885565b93506108d7818560208601610895565b6108e0816108a3565b840191505092915050565b5f60a0820190506108fe5f830188610834565b61090b602083018761084e565b610918604083018661085d565b610925606083018561086c565b818103608083015261093781846108b3565b90509695505050505050565b5f81519050610951816106d9565b92915050565b5f5f6040838503121561096d5761096c61053e565b5b5f61097a85828601610943565b925050602061098b85828601610943565b9150509250929050565b5f82825260208201905092915050565b7f6172626974726167655769746864726177556e69737761703a2050726f6669745f8201527f20746f6f206c6f77000000000000000000000000000000000000000000000000602082015250565b5f6109ff602883610995565b9150610a0a826109a5565b604082019050919050565b5f6020820190508181035f830152610a2c816109f3565b9050919050565b5f81905092915050565b50565b5f610a4b5f83610a33565b9150610a5682610a3d565b5f82019050919050565b5f610a6a82610a40565b9150819050919050565b7f5472616e73666572206661696c65642e000000000000000000000000000000005f82015250565b5f610aa8601083610995565b9150610ab382610a74565b602082019050919050565b5f6020820190508181035f830152610ad581610a9c565b9050919050565b610ae5816105a0565b82525050565b5f604082019050610afe5f830185610adc565b610b0b6020830184610adc565b9392505050565b5f610b1c82610546565b9050919050565b610b2c81610b12565b8114610b36575f5ffd5b50565b5f81359050610b4781610b23565b92915050565b5f60208284031215610b6257610b6161053e565b5b5f610b6f84828501610b39565b91505092915050565b7f756e697377617056335377617043616c6c6261636b3a20556e617574686f72695f8201527f7a65640000000000000000000000000000000000000000000000000000000000602082015250565b5f610bd2602383610995565b9150610bdd82610b78565b604082019050919050565b5f6020820190508181035f830152610bff81610bc6565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610c3d826106d0565b91507f80000000000000000000000000000000000000000000000000000000000000008203610c6f57610c6e610c06565b5b815f039050919050565b5f602082019050610c8c5f830184610adc565b92915050565b5f604082019050610ca55f830185610834565b610cb26020830184610adc565b9392505050565b610cc281610843565b8114610ccc575f5ffd5b50565b5f81519050610cdd81610cb9565b92915050565b5f60208284031215610cf857610cf761053e565b5b5f610d0584828501610ccf565b9150509291505056fea2646970667358221220c50dbdff5886df2819661a39ae8648a636c55d402861d28f83ddd8a52ea6d36664736f6c634300081c0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// RETH is a free data retrieval call binding the contract method 0x8265b121.
//
// Solidity: function RETH() view returns(address)
func (_Contract *ContractCaller) RETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "RETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RETH is a free data retrieval call binding the contract method 0x8265b121.
//
// Solidity: function RETH() view returns(address)
func (_Contract *ContractSession) RETH() (common.Address, error) {
	return _Contract.Contract.RETH(&_Contract.CallOpts)
}

// RETH is a free data retrieval call binding the contract method 0x8265b121.
//
// Solidity: function RETH() view returns(address)
func (_Contract *ContractCallerSession) RETH() (common.Address, error) {
	return _Contract.Contract.RETH(&_Contract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contract *ContractCallerSession) WETH() (common.Address, error) {
	return _Contract.Contract.WETH(&_Contract.CallOpts)
}

// Arb is a paid mutator transaction binding the contract method 0x5ac6a53b.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit) returns()
func (_Contract *ContractTransactor) Arb(opts *bind.TransactOpts, _uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "arb", _uniswapPool, _amount, _minProfit)
}

// Arb is a paid mutator transaction binding the contract method 0x5ac6a53b.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit) returns()
func (_Contract *ContractSession) Arb(_uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _amount, _minProfit)
}

// Arb is a paid mutator transaction binding the contract method 0x5ac6a53b.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit) returns()
func (_Contract *ContractTransactorSession) Arb(_uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _amount, _minProfit)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes _data) returns()
func (_Contract *ContractTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, _amountRETHDelta *big.Int, _amountWETHDelta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3SwapCallback", _amountRETHDelta, _amountWETHDelta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes _data) returns()
func (_Contract *ContractSession) UniswapV3SwapCallback(_amountRETHDelta *big.Int, _amountWETHDelta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, _amountRETHDelta, _amountWETHDelta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes _data) returns()
func (_Contract *ContractTransactorSession) UniswapV3SwapCallback(_amountRETHDelta *big.Int, _amountWETHDelta *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, _amountRETHDelta, _amountWETHDelta, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractArbitrageIterator is returned from FilterArbitrage and is used to iterate over the raw logs and unpacked data for Arbitrage events raised by the Contract contract.
type ContractArbitrageIterator struct {
	Event *ContractArbitrage // Event containing the contract specifics and raw log

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
func (it *ContractArbitrageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractArbitrage)
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
		it.Event = new(ContractArbitrage)
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
func (it *ContractArbitrageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractArbitrageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractArbitrage represents a Arbitrage event raised by the Contract contract.
type ContractArbitrage struct {
	Caller      common.Address
	UniswapPool common.Address
	Amount      *big.Int
	Profit      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterArbitrage is a free log retrieval operation binding the contract event 0x2a447723197b74967a6126c418f254abf2f6e9dc7b9c4f3a09321daaba444d1f.
//
// Solidity: event Arbitrage(address indexed caller, address indexed uniswapPool, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) FilterArbitrage(opts *bind.FilterOpts, caller []common.Address, uniswapPool []common.Address) (*ContractArbitrageIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Arbitrage", callerRule, uniswapPoolRule)
	if err != nil {
		return nil, err
	}
	return &ContractArbitrageIterator{contract: _Contract.contract, event: "Arbitrage", logs: logs, sub: sub}, nil
}

// WatchArbitrage is a free log subscription operation binding the contract event 0x2a447723197b74967a6126c418f254abf2f6e9dc7b9c4f3a09321daaba444d1f.
//
// Solidity: event Arbitrage(address indexed caller, address indexed uniswapPool, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) WatchArbitrage(opts *bind.WatchOpts, sink chan<- *ContractArbitrage, caller []common.Address, uniswapPool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var uniswapPoolRule []interface{}
	for _, uniswapPoolItem := range uniswapPool {
		uniswapPoolRule = append(uniswapPoolRule, uniswapPoolItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Arbitrage", callerRule, uniswapPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractArbitrage)
				if err := _Contract.contract.UnpackLog(event, "Arbitrage", log); err != nil {
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

// ParseArbitrage is a log parse operation binding the contract event 0x2a447723197b74967a6126c418f254abf2f6e9dc7b9c4f3a09321daaba444d1f.
//
// Solidity: event Arbitrage(address indexed caller, address indexed uniswapPool, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) ParseArbitrage(log types.Log) (*ContractArbitrage, error) {
	event := new(ContractArbitrage)
	if err := _Contract.contract.UnpackLog(event, "Arbitrage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
