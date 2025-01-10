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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"uniswapPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"}],\"name\":\"Arbitrage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RETH\",\"outputs\":[{\"internalType\":\"contractIRETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uniswapPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"arb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"arbParaswap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountWethBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMorphoFlashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_amountRETHDelta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_amountWETHDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e0604052731d8f8f00cfa6758d7be78336684788fb0ee0fa4673ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525073bbbbbbbbbb9cc5e90e3b3af64bdaf62c37eeffcb73ffffffffffffffffffffffffffffffffffffffff1660a09073ffffffffffffffffffffffffffffffffffffffff16815250736a000f20005980200259b80c510200304000106873ffffffffffffffffffffffffffffffffffffffff1660c09073ffffffffffffffffffffffffffffffffffffffff168152503480156100e4575f5ffd5b5060805160a05160c0516114846101135f395f61015a01525f81816105f501526107d301525f50506114845ff3fe608060405260043610610058575f3560e01c806331f57072146100635780638265b1211461008b578063a06a59f3146100b5578063ad5c4648146100dd578063c170bfa014610107578063fa461e331461012f5761005f565b3661005f57005b5f5ffd5b34801561006e575f5ffd5b5061008960048036038101906100849190610abb565b610157565b005b348015610096575f5ffd5b5061009f6103a2565b6040516100ac9190610b92565b60405180910390f35b3480156100c0575f5ffd5b506100db60048036038101906100d69190610be6565b6103ba565b005b3480156100e8575f5ffd5b506100f16105db565b6040516100fe9190610c6a565b60405180910390f35b348015610112575f5ffd5b5061012d60048036038101906101289190610c83565b6105f3565b005b34801561013a575f5ffd5b5061015560048036038101906101509190610d3a565b610814565b005b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16838360405161019f929190610de7565b5f604051808303815f865af19150503d805f81146101d8576040519150601f19603f3d011682016040523d82523d5f602084013e6101dd565b606091505b5050905080610221576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021890610e59565b60405180910390fd5b5f73ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161026f9190610e86565b602060405180830381865afa15801561028a573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102ae9190610eb3565b905073ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166342966c68826040518263ffffffff1660e01b81526004016102fd9190610eed565b5f604051808303815f87803b158015610314575f5ffd5b505af1158015610326573d5f5f3e3d5ffd5b5050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0866040518263ffffffff1660e01b81526004015f604051808303818588803b158015610384575f5ffd5b505af1158015610396573d5f5f3e3d5ffd5b50505050505050505050565b73ae78736cd615f374d3085123a210448e74fc639381565b8373ffffffffffffffffffffffffffffffffffffffff1663128acb08305f86600173fffd8963efd1fc6a506488495d951d5263988d266103fa9190610f33565b8960405160200161040b9190610e86565b6040516020818303038152906040526040518663ffffffff1660e01b815260040161043a959493929190611022565b60408051808303815f875af1158015610455573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610479919061108e565b50505f479050828110156104c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b99061113c565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516104e79061117d565b5f6040518083038185875af1925050503d805f8114610521576040519150601f19603f3d011682016040523d82523d5f602084013e610526565b606091505b505090508061056a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610561906111db565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a8888866040516105cb939291906111f9565b60405180910390a3505050505050565b73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc281565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e0232b4273c02aaa39b223fe8d0a0e5c4f27ead9083c756cc28787876040518563ffffffff1660e01b8152600401610666949392919061125a565b5f604051808303815f87803b15801561067d575f5ffd5b505af115801561068f573d5f5f3e3d5ffd5b505050505f479050828110156106da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d19061113c565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff16826040516106ff9061117d565b5f6040518083038185875af1925050503d805f8114610739576040519150601f19603f3d011682016040523d82523d5f602084013e61073e565b606091505b5050905080610782576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610779906111db565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a7f00000000000000000000000000000000000000000000000000000000000000008a86604051610803939291906111f9565b60405180910390a350505050505050565b5f828281019061082491906112d3565b90508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610894576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088b9061136e565b60405180910390fd5b73ae78736cd615f374d3085123a210448e74fc639373ffffffffffffffffffffffffffffffffffffffff166342966c68866108ce9061138c565b6040518263ffffffff1660e01b81526004016108ea9190610eed565b5f604051808303815f87803b158015610901575f5ffd5b505af1158015610913573d5f5f3e3d5ffd5b5050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004015f604051808303818588803b158015610971575f5ffd5b505af1158015610983573d5f5f3e3d5ffd5b505050505073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb82866040518363ffffffff1660e01b81526004016109d79291906113d2565b6020604051808303815f875af11580156109f3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a179190611423565b505050505050565b5f5ffd5b5f5ffd5b5f819050919050565b610a3981610a27565b8114610a43575f5ffd5b50565b5f81359050610a5481610a30565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610a7b57610a7a610a5a565b5b8235905067ffffffffffffffff811115610a9857610a97610a5e565b5b602083019150836001820283011115610ab457610ab3610a62565b5b9250929050565b5f5f5f60408486031215610ad257610ad1610a1f565b5b5f610adf86828701610a46565b935050602084013567ffffffffffffffff811115610b0057610aff610a23565b5b610b0c86828701610a66565b92509250509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610b5a610b55610b5084610b18565b610b37565b610b18565b9050919050565b5f610b6b82610b40565b9050919050565b5f610b7c82610b61565b9050919050565b610b8c81610b72565b82525050565b5f602082019050610ba55f830184610b83565b92915050565b5f610bb582610b18565b9050919050565b610bc581610bab565b8114610bcf575f5ffd5b50565b5f81359050610be081610bbc565b92915050565b5f5f5f5f60808587031215610bfe57610bfd610a1f565b5b5f610c0b87828801610bd2565b9450506020610c1c87828801610a46565b9350506040610c2d87828801610a46565b9250506060610c3e87828801610bd2565b91505092959194509250565b5f610c5482610b61565b9050919050565b610c6481610c4a565b82525050565b5f602082019050610c7d5f830184610c5b565b92915050565b5f5f5f5f5f60808688031215610c9c57610c9b610a1f565b5b5f610ca988828901610a46565b955050602086013567ffffffffffffffff811115610cca57610cc9610a23565b5b610cd688828901610a66565b94509450506040610ce988828901610a46565b9250506060610cfa88828901610bd2565b9150509295509295909350565b5f819050919050565b610d1981610d07565b8114610d23575f5ffd5b50565b5f81359050610d3481610d10565b92915050565b5f5f5f5f60608587031215610d5257610d51610a1f565b5b5f610d5f87828801610d26565b9450506020610d7087828801610d26565b935050604085013567ffffffffffffffff811115610d9157610d90610a23565b5b610d9d87828801610a66565b925092505092959194509250565b5f81905092915050565b828183375f83830152505050565b5f610dce8385610dab565b9350610ddb838584610db5565b82840190509392505050565b5f610df3828486610dc3565b91508190509392505050565b5f82825260208201905092915050565b7f5061726173776170206661696c656400000000000000000000000000000000005f82015250565b5f610e43600f83610dff565b9150610e4e82610e0f565b602082019050919050565b5f6020820190508181035f830152610e7081610e37565b9050919050565b610e8081610bab565b82525050565b5f602082019050610e995f830184610e77565b92915050565b5f81519050610ead81610a30565b92915050565b5f60208284031215610ec857610ec7610a1f565b5b5f610ed584828501610e9f565b91505092915050565b610ee781610a27565b82525050565b5f602082019050610f005f830184610ede565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f3d82610b18565b9150610f4883610b18565b9250828203905073ffffffffffffffffffffffffffffffffffffffff811115610f7457610f73610f06565b5b92915050565b5f8115159050919050565b610f8e81610f7a565b82525050565b610f9d81610d07565b82525050565b610fac81610b18565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610ff482610fb2565b610ffe8185610fbc565b935061100e818560208601610fcc565b61101781610fda565b840191505092915050565b5f60a0820190506110355f830188610e77565b6110426020830187610f85565b61104f6040830186610f94565b61105c6060830185610fa3565b818103608083015261106e8184610fea565b90509695505050505050565b5f8151905061108881610d10565b92915050565b5f5f604083850312156110a4576110a3610a1f565b5b5f6110b18582860161107a565b92505060206110c28582860161107a565b9150509250929050565b7f6172626974726167655769746864726177556e69737761703a2050726f6669745f8201527f20746f6f206c6f77000000000000000000000000000000000000000000000000602082015250565b5f611126602883610dff565b9150611131826110cc565b604082019050919050565b5f6020820190508181035f8301526111538161111a565b9050919050565b50565b5f6111685f83610dab565b91506111738261115a565b5f82019050919050565b5f6111878261115d565b9150819050919050565b7f5472616e73666572206661696c65642e000000000000000000000000000000005f82015250565b5f6111c5601083610dff565b91506111d082611191565b602082019050919050565b5f6020820190508181035f8301526111f2816111b9565b9050919050565b5f60608201905061120c5f830186610e77565b6112196020830185610ede565b6112266040830184610ede565b949350505050565b5f6112398385610fbc565b9350611246838584610db5565b61124f83610fda565b840190509392505050565b5f60608201905061126d5f830187610e77565b61127a6020830186610ede565b818103604083015261128d81848661122e565b905095945050505050565b5f6112a282610b18565b9050919050565b6112b281611298565b81146112bc575f5ffd5b50565b5f813590506112cd816112a9565b92915050565b5f602082840312156112e8576112e7610a1f565b5b5f6112f5848285016112bf565b91505092915050565b7f756e697377617056335377617043616c6c6261636b3a20556e617574686f72695f8201527f7a65640000000000000000000000000000000000000000000000000000000000602082015250565b5f611358602383610dff565b9150611363826112fe565b604082019050919050565b5f6020820190508181035f8301526113858161134c565b9050919050565b5f61139682610d07565b91507f800000000000000000000000000000000000000000000000000000000000000082036113c8576113c7610f06565b5b815f039050919050565b5f6040820190506113e55f830185610e77565b6113f26020830184610ede565b9392505050565b61140281610f7a565b811461140c575f5ffd5b50565b5f8151905061141d816113f9565b92915050565b5f6020828403121561143857611437610a1f565b5b5f6114458482850161140f565b9150509291505056fea264697066735822122054fadd42bb638eec198a013e48b9e8505793b9f543f3310fd3558b46fb0c79a664736f6c634300081c0033",
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

// Arb is a paid mutator transaction binding the contract method 0xa06a59f3.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactor) Arb(opts *bind.TransactOpts, _uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "arb", _uniswapPool, _amount, _minProfit, _receiver)
}

// Arb is a paid mutator transaction binding the contract method 0xa06a59f3.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractSession) Arb(_uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _amount, _minProfit, _receiver)
}

// Arb is a paid mutator transaction binding the contract method 0xa06a59f3.
//
// Solidity: function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit, address _receiver) returns()
func (_Contract *ContractTransactorSession) Arb(_uniswapPool common.Address, _amount *big.Int, _minProfit *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Arb(&_Contract.TransactOpts, _uniswapPool, _amount, _minProfit, _receiver)
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
	Receiver    common.Address
	UniswapPool common.Address
	Amount      *big.Int
	Profit      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterArbitrage is a free log retrieval operation binding the contract event 0x1b9af080202f245b0221dc0001aaa951553203bd8a00ae62d2fb3940c12bad1a.
//
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address uniswapPool, uint256 amount, uint256 profit)
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
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address uniswapPool, uint256 amount, uint256 profit)
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
// Solidity: event Arbitrage(address indexed caller, address indexed receiver, address uniswapPool, uint256 amount, uint256 profit)
func (_Contract *ContractFilterer) ParseArbitrage(log types.Log) (*ContractArbitrage, error) {
	event := new(ContractArbitrage)
	if err := _Contract.contract.UnpackLog(event, "Arbitrage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
