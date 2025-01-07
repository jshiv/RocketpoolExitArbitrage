package arbitrage

import (
	"rocketpoolArbitrage/rocketpoolContracts/rETH"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xtrooper/flashbots_client"
)

type DataIn struct {
	Command                         string
	LocalReth                       bool
	MinipoolAddresses               []common.Address
	NodeAddress                     *common.Address
	RefundAddress                   *common.Address
	RETHInstance                    *rETH.RETH
	Client                          *ethclient.Client
	FbClient                        *flashbots_client.FlashbotsClient
	RandomPrivateKey                bool
	SkipConfirmation                bool
	CheckProfit                     bool
	CheckProfitIgnoreDistributeCost bool
	DryRun                          bool
}
