package arbitrage

import (
	"context"
	"math/big"
	"rocketpoolArbitrage/rocketpoolContracts/minipoolDelegate"
	"rocketpoolArbitrage/rocketpoolContracts/rETH"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetMinipoolStatus(ctx context.Context, instance *minipoolDelegate.MinipoolDelegate) (uint8, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetStatus()
}

func GetMinipoolDelegateVersion(ctx context.Context, instance *minipoolDelegate.MinipoolDelegate) (uint8, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.Version()
}

func GetMinipoolRefundBalance(ctx context.Context, instance *minipoolDelegate.MinipoolDelegate) (*big.Int, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetNodeRefundBalance()
}

func CalculateMinipoolUserShare(ctx context.Context, instance *minipoolDelegate.MinipoolDelegate, amount *big.Int) (*big.Int, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.CalculateUserShare(amount)
}

func GetMinipoolNodeAddress(ctx context.Context, instance *minipoolDelegate.MinipoolDelegate) (common.Address, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetNodeAddress()
}

func ConvertRethToWeth(ctx context.Context, instance *rETH.RETH, rEthAmount *big.Int) (*big.Int, error) {
	session := &rETH.RETHSession{
		Contract: instance,
		TransactOpts: bind.TransactOpts{
			Context: ctx,
		},
	}

	return session.GetEthValue(rEthAmount)
}

func ConvertWethToReth(ctx context.Context, instance *rETH.RETH, wethAmount *big.Int) (*big.Int, error) {
	session := &rETH.RETHSession{
		Contract: instance,
		TransactOpts: bind.TransactOpts{
			Context: ctx,
		},
	}

	return session.GetRethValue(wethAmount)
}
