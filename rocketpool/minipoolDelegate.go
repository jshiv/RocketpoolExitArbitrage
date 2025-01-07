package rocketpool

import (
	"context"
	"log/slog"
	"math/big"
	"rocketpoolArbitrage/rocketpool/contracts/minipoolDelegate"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	MinipoolStatusInitialised  = uint8(0)
	MinipoolStatusPrelaunch    = uint8(1)
	MinipoolStatusStaking      = uint8(2)
	MinipoolStatusWithdrawable = uint8(3)
	MinipoolStatusDissolved    = uint8(4)
)

type MinipoolDelegateInstance struct {
	logger *slog.Logger

	contractAddress common.Address
	contractAbi     *abi.ABI

	contractInstance *minipoolDelegate.MinipoolDelegate
}

func NewMinipoolDelegate(ctx context.Context, minipoolAddress common.Address) (*MinipoolDelegateInstance, error) {
	rpc := ctx.Value("RPC_LOCAL").(*ethclient.Client)
	logger := ctx.Value("logger").(*slog.Logger)

	abi, err := minipoolDelegate.MinipoolDelegateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	instance, err := minipoolDelegate.NewMinipoolDelegate(minipoolAddress, rpc)
	if err != nil {
		return nil, err
	}

	return &MinipoolDelegateInstance{
		logger:           logger.With(slog.String("module", "MinipoolDelegate")),
		contractAddress:  minipoolAddress,
		contractAbi:      abi,
		contractInstance: instance,
	}, nil
}

func (md *MinipoolDelegateInstance) GetBalance(ctx context.Context) (*big.Int, error) {
	rpc := ctx.Value("RPC_LOCAL").(*ethclient.Client)
	return rpc.BalanceAt(ctx, md.contractAddress, nil)
}

func (md *MinipoolDelegateInstance) GetStatus(ctx context.Context) (uint8, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetStatus()
}

func (md *MinipoolDelegateInstance) GetDelegateVersion(ctx context.Context) (uint8, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.Version()
}

func (md *MinipoolDelegateInstance) GetRefundBalance(ctx context.Context) (*big.Int, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetNodeRefundBalance()
}

func (md *MinipoolDelegateInstance) GetFinalised(ctx context.Context) (bool, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.GetFinalised()
}

func (md *MinipoolDelegateInstance) CalculateUserShare(ctx context.Context, fullAmount *big.Int) (*big.Int, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.CalculateUserShare(fullAmount)
}

func (md *MinipoolDelegateInstance) CalculateNodeShare(ctx context.Context, fullAmount *big.Int) (*big.Int, error) {
	session := &minipoolDelegate.MinipoolDelegateSession{
		Contract: md.contractInstance,
		CallOpts: bind.CallOpts{
			Context: ctx,
		},
	}

	return session.CalculateNodeShare(fullAmount)
}
