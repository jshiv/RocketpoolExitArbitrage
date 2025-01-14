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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"flashloanProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"name\":\"Arbitrage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RETH\",\"outputs\":[{\"internalType\":\"contractIRETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uniswapPool\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"_sqrtPriceLimitX96\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"arb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"arbParaswap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountWethBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoFlashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_amountRETHDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_amountWETHDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061154d8061001c5f395ff3fe608060405260043610610058575f3560e01c806331f57072146100635780638265b1211461008b5780639631b2a0146100b5578063ad5c4648146100dd578063c170bfa014610107578063fa461e331461012f5761005f565b3661005f57005b5f5ffd5b34801561006e575f5ffd5b5061008960048036038101906100849190610bb2565b610157565b005b348015610096575f5ffd5b5061009f6104de565b6040516100ac9190610c89565b60405180910390f35b3480156100c0575f5ffd5b506100db60048036038101906100d69190610d07565b6104f6565b005b3480156100e8575f5ffd5b506100f16106e7565b6040516100fe9190610d9e565b60405180910390f35b348015610112575f5ffd5b5061012d60048036038101906101289190610db7565b6106ff565b005b34801561013a575f5ffd5b5061015560048036038101906101509190610e6e565b610908565b005b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663095ea7b3736a000f20005980200259b80c5102003040001068856040518363ffffffff1660e01b81526004016101ba929190610efd565b6020604051808303815f875af11580156101d6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101fa9190610f59565b505f736a000f20005980200259b80c510200304000106873ffffffffffffffffffffffffffffffffffffffff168383604051610237929190610fc0565b5f604051808303815f865af19150503d805f8114610270576040519150601f19603f3d011682016040523d82523d5f602084013e610275565b606091505b50509050806102b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b090611032565b60405180910390fd5b5f73ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103079190611050565b602060405180830381865afa158015610322573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610346919061107d565b905073ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166342966c68826040518263ffffffff1660e01b815260040161039591906110a8565b5f604051808303815f87803b1580156103ac575f5ffd5b505af11580156103be573d5f5f3e3d5ffd5b5050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0866040518263ffffffff1660e01b81526004015f604051808303818588803b15801561041c575f5ffd5b505af115801561042e573d5f5f3e3d5ffd5b505050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663095ea7b373bbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb876040518363ffffffff1660e01b8152600401610496929190610efd565b6020604051808303815f875af11580156104b2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104d69190610f59565b505050505050565b73ae78736cd615f374d3085123a210448e74fc639381565b8473ffffffffffffffffffffffffffffffffffffffff1663128acb08305f868860405180602001604052805f8152506040518663ffffffff1660e01b815260040161054595949392919061115e565b60408051808303815f875af1158015610560573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058491906111ca565b50505f479050828110156105cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c490611252565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516105f290611293565b5f6040518083038185875af1925050503d805f811461062c576040519150601f19603f3d011682016040523d82523d5f602084013e610631565b606091505b5050905080610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066c906112f1565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a8988866040516106d69392919061130f565b60405180910390a350505050505050565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b73bbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb73ffffffffffffffffffffffffffffffffffffffff1663e0232b4273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28787876040518563ffffffff1660e01b81526004016107669493929190611370565b5f604051808303815f87803b15801561077d575f5ffd5b505af115801561078f573d5f5f3e3d5ffd5b505050505f479050828110156107da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d190611252565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516107ff90611293565b5f6040518083038185875af1925050503d805f8114610839576040519150601f19603f3d011682016040523d82523d5f602084013e61083e565b606091505b5050905080610882576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610879906112f1565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a73bbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb8a866040516108f79392919061130f565b60405180910390a350505050505050565b5f841261094a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610941906113f8565b60405180910390fd5b5f831361098c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098390611486565b60405180910390fd5b73ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166342966c68856109c6906114d1565b6040518263ffffffff1660e01b81526004016109e291906110a8565b5f604051808303815f87803b1580156109f9575f5ffd5b505af1158015610a0b573d5f5f3e3d5ffd5b5050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0846040518263ffffffff1660e01b81526004015f604051808303818588803b158015610a69575f5ffd5b505af1158015610a7b573d5f5f3e3d5ffd5b505050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b8152600401610acf929190610efd565b6020604051808303815f875af1158015610aeb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b0f9190610f59565b5050505050565b5f5ffd5b5f5ffd5b5f819050919050565b610b3081610b1e565b8114610b3a575f5ffd5b50565b5f81359050610b4b81610b27565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610b7257610b71610b51565b5b8235905067ffffffffffffffff811115610b8f57610b8e610b55565b5b602083019150836001820283011115610bab57610baa610b59565b5b9250929050565b5f5f5f60408486031215610bc957610bc8610b16565b5b5f610bd686828701610b3d565b935050602084013567ffffffffffffffff811115610bf757610bf6610b1a565b5b610c0386828701610b5d565b92509250509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610c51610c4c610c4784610c0f565b610c2e565b610c0f565b9050919050565b5f610c6282610c37565b9050919050565b5f610c7382610c58565b9050919050565b610c8381610c69565b82525050565b5f602082019050610c9c5f830184610c7a565b92915050565b5f610cac82610c0f565b9050919050565b610cbc81610ca2565b8114610cc6575f5ffd5b50565b5f81359050610cd781610cb3565b92915050565b610ce681610c0f565b8114610cf0575f5ffd5b50565b5f81359050610d0181610cdd565b92915050565b5f5f5f5f5f60a08688031215610d2057610d1f610b16565b5b5f610d2d88828901610cc9565b9550506020610d3e88828901610cf3565b9450506040610d4f88828901610b3d565b9350506060610d6088828901610b3d565b9250506080610d7188828901610cc9565b9150509295509295909350565b5f610d8882610c58565b9050919050565b610d9881610d7e565b82525050565b5f602082019050610db15f830184610d8f565b92915050565b5f5f5f5f5f60808688031215610dd057610dcf610b16565b5b5f610ddd88828901610b3d565b955050602086013567ffffffffffffffff811115610dfe57610dfd610b1a565b5b610e0a88828901610b5d565b94509450506040610e1d88828901610b3d565b9250506060610e2e88828901610cc9565b9150509295509295909350565b5f819050919050565b610e4d81610e3b565b8114610e57575f5ffd5b50565b5f81359050610e6881610e44565b92915050565b5f5f5f5f60608587031215610e8657610e85610b16565b5b5f610e9387828801610e5a565b9450506020610ea487828801610e5a565b935050604085013567ffffffffffffffff811115610ec557610ec4610b1a565b5b610ed187828801610b5d565b925092505092959194509250565b610ee881610ca2565b82525050565b610ef781610b1e565b82525050565b5f604082019050610f105f830185610edf565b610f1d6020830184610eee565b9392505050565b5f8115159050919050565b610f3881610f24565b8114610f42575f5ffd5b50565b5f81519050610f5381610f2f565b92915050565b5f60208284031215610f6e57610f6d610b16565b5b5f610f7b84828501610f45565b91505092915050565b5f81905092915050565b828183375f83830152505050565b5f610fa78385610f84565b9350610fb4838584610f8e565b82840190509392505050565b5f610fcc828486610f9c565b91508190509392505050565b5f82825260208201905092915050565b7f5061726173776170206661696c656400000000000000000000000000000000005f82015250565b5f61101c600f83610fd8565b915061102782610fe8565b602082019050919050565b5f6020820190508181035f83015261104981611010565b9050919050565b5f6020820190506110635f830184610edf565b92915050565b5f8151905061107781610b27565b92915050565b5f6020828403121561109257611091610b16565b5b5f61109f84828501611069565b91505092915050565b5f6020820190506110bb5f830184610eee565b92915050565b6110ca81610f24565b82525050565b6110d981610e3b565b82525050565b6110e881610c0f565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611130826110ee565b61113a81856110f8565b935061114a818560208601611108565b61115381611116565b840191505092915050565b5f60a0820190506111715f830188610edf565b61117e60208301876110c1565b61118b60408301866110d0565b61119860608301856110df565b81810360808301526111aa8184611126565b90509695505050505050565b5f815190506111c481610e44565b92915050565b5f5f604083850312156111e0576111df610b16565b5b5f6111ed858286016111b6565b92505060206111fe858286016111b6565b9150509250929050565b7f50726f66697420746f6f206c6f770000000000000000000000000000000000005f82015250565b5f61123c600e83610fd8565b915061124782611208565b602082019050919050565b5f6020820190508181035f83015261126981611230565b9050919050565b50565b5f61127e5f83610f84565b915061128982611270565b5f82019050919050565b5f61129d82611273565b9150819050919050565b7f5472616e73666572206661696c65642e000000000000000000000000000000005f82015250565b5f6112db601083610fd8565b91506112e6826112a7565b602082019050919050565b5f6020820190508181035f830152611308816112cf565b9050919050565b5f6060820190506113225f830186610edf565b61132f6020830185610eee565b61133c6040830184610eee565b949350505050565b5f61134f83856110f8565b935061135c838584610f8e565b61136583611116565b840190509392505050565b5f6060820190506113835f830187610edf565b6113906020830186610eee565b81810360408301526113a3818486611344565b905095945050505050565b7f72455448206d7573742062652073656e7420746f2074686520706f6f6c0000005f82015250565b5f6113e2601d83610fd8565b91506113ed826113ae565b602082019050919050565b5f6020820190508181035f83015261140f816113d6565b9050919050565b7f57455448206d7573742062652072656365697665642066726f6d2074686520705f8201527f6f6f6c0000000000000000000000000000000000000000000000000000000000602082015250565b5f611470602383610fd8565b915061147b82611416565b604082019050919050565b5f6020820190508181035f83015261149d81611464565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6114db82610e3b565b91507f8000000000000000000000000000000000000000000000000000000000000000820361150d5761150c6114a4565b5b815f03905091905056fea26469706673582212200ced343d9a9fd4cf6dccfad03c485651e10670fe26deeafb78e376aa28bf99c064736f6c634300081c0033",
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

// Arb is a paid mutator transaction binding the contract method 0x9631b2a0.
//
// Solidity: function arb(address _uniswapPool, uint160 _sqrtPriceLimitX96, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactor) Arb(opts *bind.TransactOpts, _uniswapPool common.Address, _sqrtPriceLimitX96 *big.Int, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "arb", _uniswapPool, _sqrtPriceLimitX96, _amount, _minProfit, _receiver)
}

// Arb is a paid mutator transaction binding the contract method 0x9631b2a0.
//
// Solidity: function arb(address _uniswapPool, uint160 _sqrtPriceLimitX96, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractSession) Arb(_uniswapPool common.Address, _sqrtPriceLimitX96 *big.Int, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _sqrtPriceLimitX96, _amount, _minProfit, _receiver)
}

// Arb is a paid mutator transaction binding the contract method 0x9631b2a0.
//
// Solidity: function arb(address _uniswapPool, uint160 _sqrtPriceLimitX96, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactorSession) Arb(_uniswapPool common.Address, _sqrtPriceLimitX96 *big.Int, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _sqrtPriceLimitX96, _amount, _minProfit, _receiver)
}

// ArbParaswap is a paid mutator transaction binding the contract method 0xc170bfa0.
//
// Solidity: function arbParaswap(uint256 _amount, bytes _data, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactor) ArbParaswap(opts *bind.TransactOpts, _amount *big.Int, _data []byte, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "arbParaswap", _amount, _data, _minProfit, _receiver)
}

// ArbParaswap is a paid mutator transaction binding the contract method 0xc170bfa0.
//
// Solidity: function arbParaswap(uint256 _amount, bytes _data, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractSession) ArbParaswap(_amount *big.Int, _data []byte, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ArbParaswap(&_Contract.TransactOpts, _amount, _data, _minProfit, _receiver)
}

// ArbParaswap is a paid mutator transaction binding the contract method 0xc170bfa0.
//
// Solidity: function arbParaswap(uint256 _amount, bytes _data, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactorSession) ArbParaswap(_amount *big.Int, _data []byte, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ArbParaswap(&_Contract.TransactOpts, _amount, _data, _minProfit, _receiver)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 amountWethBorrowed, bytes data) returns()
func (_Contract *ContractTransactor) OnMorphoFlashLoan(opts *bind.TransactOpts, amountWethBorrowed *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "onMorphoFlashLoan", amountWethBorrowed, data)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 amountWethBorrowed, bytes data) returns()
func (_Contract *ContractSession) OnMorphoFlashLoan(amountWethBorrowed *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnMorphoFlashLoan(&_Contract.TransactOpts, amountWethBorrowed, data)
}

// OnMorphoFlashLoan is a paid mutator transaction binding the contract method 0x31f57072.
//
// Solidity: function onMorphoFlashLoan(uint256 amountWethBorrowed, bytes data) returns()
func (_Contract *ContractTransactorSession) OnMorphoFlashLoan(amountWethBorrowed *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnMorphoFlashLoan(&_Contract.TransactOpts, amountWethBorrowed, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes ) returns()
func (_Contract *ContractTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, _amountRETHDelta *big.Int, _amountWETHDelta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "uniswapV3SwapCallback", _amountRETHDelta, _amountWETHDelta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes ) returns()
func (_Contract *ContractSession) UniswapV3SwapCallback(_amountRETHDelta *big.Int, _amountWETHDelta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, _amountRETHDelta, _amountWETHDelta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 _amountRETHDelta, int256 _amountWETHDelta, bytes ) returns()
func (_Contract *ContractTransactorSession) UniswapV3SwapCallback(_amountRETHDelta *big.Int, _amountWETHDelta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Contract.Contract.UniswapV3SwapCallback(&_Contract.TransactOpts, _amountRETHDelta, _amountWETHDelta, arg2)
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
	Caller            common.Address
	Receiver          common.Address
	FlashloanProvider common.Address
	Amount            *big.Int
	Profit            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterArbitrage is a free log retrieval operation binding the contract event 0x1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a.
//
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address flashloanProvider, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) FilterArbitrage(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address) (*ContractArbitrageIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Arbitrage", callerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractArbitrageIterator{contract: _Contract.contract, event: "Arbitrage", logs: logs, sub: sub}, nil
}

// WatchArbitrage is a free log subscription operation binding the contract event 0x1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a.
//
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address flashloanProvider, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) WatchArbitrage(opts *bind.WatchOpts, sink chan<- *ContractArbitrage, caller []common.Address, receiver []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Arbitrage", callerRule, receiverRule)
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

// ParseArbitrage is a log parse operation binding the contract event 0x1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a.
//
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address flashloanProvider, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) ParseArbitrage(log types.Log) (*ContractArbitrage, error) {
	event := new(ContractArbitrage)
	if err := _Contract.contract.UnpackLog(event, "Arbitrage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
