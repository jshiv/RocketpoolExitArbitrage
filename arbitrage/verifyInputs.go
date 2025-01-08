package arbitrage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"rocketpoolArbitrage/rocketpoolContracts/minipoolDelegate"
)

func VerifyInputData(ctx context.Context, logger *slog.Logger, dataIn *DataIn) error {
	logger.With(slog.String("function", "VerifyInputData"))

	var verifyAllCallsFromNO bool
	if dataIn.NodeAddress == nil {
		verifyAllCallsFromNO = true
	}

	for _, minipoolAddress := range dataIn.MinipoolAddresses {
		minipoolInstance, err := minipoolDelegate.NewMinipoolDelegate(minipoolAddress, dataIn.Client)
		if err != nil {
			return errors.Join(fmt.Errorf("%s: failed to create minipool instance", minipoolAddress), err)
		}

		// check minipool is V3
		version, err := GetMinipoolDelegateVersion(ctx, minipoolInstance)
		if err != nil {
			return errors.Join(fmt.Errorf("%s: failed to get minipool version", minipoolAddress), err)
		}

		logger.Debug("minipool version", slog.Uint64("version", uint64(version)))

		if version != 3 {
			return fmt.Errorf("%s: only minipool V3 is supported", minipoolAddress)
		}

		status, err := GetMinipoolStatus(ctx, minipoolInstance)
		if err != nil {
			return errors.Join(fmt.Errorf("%s: failed to get minipool status", minipoolAddress), err)
		}

		logger.Debug("minipool status", slog.Uint64("status", uint64(status)))

		// check if status is staking
		if status != uint8(2) {
			return fmt.Errorf("%s: minipool is not staking", minipoolAddress)
		}

		if verifyAllCallsFromNO {
			nodeAddress, err := GetMinipoolNodeAddress(ctx, minipoolInstance)
			if err != nil {
				return errors.Join(fmt.Errorf("%s: failed to get node address", minipoolAddress), err)
			}

			// first time we set the node address here
			if dataIn.NodeAddress == nil {
				dataIn.NodeAddress = &nodeAddress
			} else if *dataIn.NodeAddress != nodeAddress {
				return fmt.Errorf("%s: node address does not match", minipoolAddress)
			}
		}

		minipoolBalance, err := dataIn.Client.BalanceAt(ctx, minipoolAddress, nil)
		if err != nil {
			logger.Warn("failed to get minipool balance", slog.String("minipool", minipoolAddress.Hex()))
			continue
		}

		if minipoolBalance.Cmp(big.NewInt(8e18)) > 0 {
			// get node address
			nodeAddress, err := GetMinipoolNodeAddress(ctx, minipoolInstance)
			if err != nil {
				return errors.Join(fmt.Errorf("%s: failed to get node address", minipoolAddress), err)
			}

			if *dataIn.NodeAddress != nodeAddress {
				return fmt.Errorf("%s: node address does not match. Minipools with over 8 ETH need to be finalized from the NO address", minipoolAddress)
			}
		}
	}

	return nil
}
