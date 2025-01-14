package arbitrage

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	Mainnet_rocketpoolStorageAddressStr = "0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46"
	Mainnet_arbitrageContractAddressStr = "0x2631618408497d27D455aBA9c99A6f61eF305559"
	Mainnet_rEthContractAddressStr      = "0xae78736Cd615f374D3085123A210448E74Fc6393"
	Mainnet_WETHContractAddressStr      = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	Mainnet_ParaswapV6_2AddressStr      = "0x6a000f20005980200259b80c5102003040001068"

	Holesky_rocketpoolStorageAddressStr = "0x594Fb75D3dc2DFa0150Ad03F99F97817747dd4E1"
	Holesky_rEthContractAddressStr      = "0x7322c24752f79c05ffd1e2a6fcb97020c1c264f1"
	Holesky_WETHContractAddressStr      = "0x94373a4919B3240D86eA41593D5eBa789FEF3848"
)

func GetRocketpoolStorageAddress(networkId uint64) (common.Address, error) {
	switch networkId {
	case 1:
		return common.HexToAddress(Mainnet_rocketpoolStorageAddressStr), nil
	case 17000:
		return common.HexToAddress(Holesky_rocketpoolStorageAddressStr), nil
	default:
		return common.Address{}, errors.New("unsupported network")
	}
}

func GetArbitrageContractAddress(networkId uint64) (common.Address, error) {
	switch networkId {
	case 1:
		return common.HexToAddress(Mainnet_arbitrageContractAddressStr), nil
	default:
		return common.Address{}, errors.New("unsupported network")
	}
}

func GetREthContractAddress(networkId uint64) (common.Address, error) {
	switch networkId {
	case 1:
		return common.HexToAddress(Mainnet_rEthContractAddressStr), nil
	case 17000:
		return common.HexToAddress(Holesky_rEthContractAddressStr), nil
	default:
		return common.Address{}, errors.New("unsupported network")
	}
}

func GetWETHContractAddress(networkId uint64) (common.Address, error) {
	switch networkId {
	case 1:
		return common.HexToAddress(Mainnet_WETHContractAddressStr), nil
	case 17000:
		return common.HexToAddress(Holesky_WETHContractAddressStr), nil
	default:
		return common.Address{}, errors.New("unsupported network")
	}
}

func GetParaswapV6_2Address(networkId uint64) (common.Address, error) {
	switch networkId {
	case 1:
		return common.HexToAddress(Mainnet_ParaswapV6_2AddressStr), nil
	default:
		return common.Address{}, errors.New("unsupported network")
	}
}
