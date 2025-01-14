package arbitrage

import (
	"bufio"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"rocketpoolArbitrage/rocketpoolContracts/storage"
	"strings"
	"time"

	"log/slog"

	"github.com/0xtrooper/flashbots_client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ExecuteDistribute(ctx context.Context, logger *slog.Logger, dataIn *DataIn) error {
	logger.With(slog.String("function", "Simulate"))

	err := VerifyInputData(ctx, logger, dataIn)
	if err != nil {
		return errors.Join(errors.New("failed to verify input data"), err)
	}

	logger.Debug("verified input data")

	// get node withdraw address
	isWithdrawalAddress := false
	isNodeAddress := false
	if dataIn.ReceiverAddress == nil {
		withdrawalAddress, err := getWithdrawalAddress(ctx, dataIn.Client, dataIn.NetworkId, *dataIn.NodeAddress)
		if err != nil {
			return errors.Join(errors.New("failed to get withdrawal address"), err)
		}

		logger.Debug("receiver address not set, updated to withdraw address", slog.String("receiverAddress", withdrawalAddress.Hex()))
		dataIn.ReceiverAddress = &withdrawalAddress
		if withdrawalAddress.Cmp(*dataIn.NodeAddress) == 0 {
			isNodeAddress = true
		} else {
			isWithdrawalAddress = true
		}
	}

	// update user on receiver address
	if logger.Enabled(ctx, slog.LevelInfo) {
		fmt.Print("Any profit will be sent to ")
		fmt.Print(colorOrange, (*dataIn.ReceiverAddress).Hex(), colorReset)
		if isWithdrawalAddress {
			fmt.Println(". This should be your withdrawal address.")
		} else if isNodeAddress {
			fmt.Println(". This should be your node address.")
		} else {
			fmt.Println(". Set with the --receiver flag.")
		}
	}

	// if using a random private key, update the fee refund recipient to the receiver address
	if dataIn.RandomPrivateKey && dataIn.NetworkId != 17000 {
		err = dataIn.FbClient.UpdateFeeRefundRecipient(*dataIn.ReceiverAddress)
		if err != nil {
			return errors.Join(errors.New("failed to update flashbots fee refund recipient"), err)
		}
		if logger.Enabled(ctx, slog.LevelInfo) {
			fmt.Print("Updated flashbots fee refund recipient to ")
			fmt.Print(colorOrange, (*dataIn.ReceiverAddress).Hex(), colorReset)
			fmt.Println(".")
		}
	}

	// build bundle
	var bundle *flashbots_client.Bundle
	var rethToBurn, rETHShare, expectedProfit *big.Int
	if dataIn.LocalReth {
		bundle, rethToBurn, rETHShare, err = BuildCallLocalReth(ctx, logger, *dataIn)
		if err != nil {
			return errors.Join(errors.New("failed to build call"), err)
		}
	} else {
		bundle, expectedProfit, err = BuildCall(ctx, logger, *dataIn)
		if err != nil {
			return errors.Join(errors.New("failed to build call"), err)
		}
	}

	logger.Debug("created flashbots client")
	res, success, err := dataIn.FbClient.SimulateBundle(bundle, 0)
	if err != nil {
		return errors.Join(errors.New("failed to simulate bundle"), err)
	}

	// udpate success based on tx results
	for index, tx := range res.Results {
		if tx.Error != "" {
			logger.Warn("tx failed", slog.Int("index", index), slog.String("error", tx.Error), slog.String("revertReason", hex.EncodeToString([]byte(tx.RevertReason))))
			success = false
		}
	}

	maxBundleFees, maxArbitrageFees := evalGasPrices(bundle)

	// print update based on user selection
	if logger.Enabled(ctx, slog.LevelInfo) {
		if dataIn.LocalReth {
			rEthBurnedFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(rethToBurn), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			ethReceivedFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(rETHShare), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			expectedFeeFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(maxBundleFees), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			fmt.Print("Simulated bundle (")
			if success {
				fmt.Print(string(colorGreen), "success", string(colorReset))
			} else {
				fmt.Print(string(colorRed), "failed", string(colorReset))
			}
			fmt.Println("):")
			fmt.Printf("    Expected to burn %.6f rETH for %.6f ETH, with a tx fee of %.6f\n", rEthBurnedFloat, ethReceivedFloat, expectedFeeFloat)
		} else {
			maxBundleFeesFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(maxBundleFees), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			maxArbitrageFeesFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(maxArbitrageFees), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			expectedProfitFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(expectedProfit), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			fmt.Print("Simulated bundle (")
			if success {
				fmt.Print(string(colorGreen), "success", string(colorReset))
			} else {
				fmt.Print(string(colorRed), "failed", string(colorReset))
			}
			fmt.Println("):")
			fmt.Printf("    Expected profit after fees: %.6f, with a tx fee of %.6f\n", expectedProfitFloat-maxBundleFeesFloat, maxBundleFeesFloat)
			fmt.Printf("    Expected profit after arbitrage fees: %.6f, with a tx fee of %.6f (interesting if you want to distribute regardless)\n\n", expectedProfitFloat-maxArbitrageFeesFloat, maxArbitrageFeesFloat)
		}
	}

	// print dry run and return if requested
	if dataIn.DryRun {
		txs := bundle.Transactions()
		fmt.Println("Dry run. Would have sent the following bundle:")
		for i, tx := range txs {
			baseGwei, _ := new(big.Float).Quo(new(big.Float).SetInt(tx.GasFeeCap()), new(big.Float).SetInt(big.NewInt(1e9))).Float64()
			tipGwei, _ := new(big.Float).Quo(new(big.Float).SetInt(tx.GasTipCap()), new(big.Float).SetInt(big.NewInt(1e9))).Float64()

			fmt.Printf("Transaction %d:\n", i+1)
			fmt.Printf("    From: %s\n", dataIn.NodeAddress.Hex())
			fmt.Printf("    To: %s\n", tx.To().Hex())
			fmt.Printf("    Value: %s\n", tx.Value().String())
			fmt.Printf("    Gas Limit: %d\n", tx.Gas())
			fmt.Printf("    Base Fee: %s (%.2f Gwei)\n", tx.GasFeeCap().String(), baseGwei)
			fmt.Printf("    Priority Fee: %s (%.4f Gwei)\n", tx.GasTipCap().String(), tipGwei)
			fmt.Printf("    Nonce: %d\n", tx.Nonce())
			fmt.Printf("    Data: %s\n", hex.EncodeToString(tx.Data()))
		}
		return nil
	}

	// end the attempt if the simulation failed - after printing the dryrun if it was requested!
	if !success {
		return errors.New("bundle simulation failed")
	}

	// this checks if a bundle makes sense to make arbitrage profits
	if dataIn.CheckProfit && !dataIn.CheckProfitIgnoreDistributeCost && expectedProfit.Cmp(maxBundleFees) < 0 {
		return errors.New("expected profit is less than max bundle fees")
	}

	// this checks if a bundle makes sense if the user wants to distribute
	// aka. distribute gas needs to be paid one way or another
	if dataIn.CheckProfit && dataIn.CheckProfitIgnoreDistributeCost && expectedProfit.Cmp(maxArbitrageFees) < 0 {
		return errors.New("expected profit is less than max arbitrage fees")
	}

	// ask for user confirmation
	if !dataIn.SkipConfirmation && !waitForUserConfirmation() {
		return errors.New("user did not confirm to proceed")
	}

	// add more builders to improve chance to be included
	bundle.UseAllBuilders(dataIn.NetworkId)

	// set target block number
	blockNumber, err := dataIn.Client.BlockNumber(ctx)
	if err != nil {
		return errors.Join(errors.New("failed to get block number"), err)
	}
	bundle.SetTargetBlockNumber(blockNumber + 1)

	fmt.Printf("\nSent bundle with hash: %s. Waiting for up to one minute to see if the transaction is included...\n\n", res.BundleHash)

	timeoutContext, cancel := context.WithTimeout(ctx, time.Second*60)
	successfullyIncluded, err := dataIn.FbClient.SendNBundleAndWait(timeoutContext, bundle, 4)
	cancel()
	if err != nil {
		return errors.Join(errors.New("failed to wait for bundle inclusion"), err)
	}

	if !successfullyIncluded {
		return errors.New("bundle was not included in the mempool")
	}

	// print successful inclusion and tx link
	var txType string
	if dataIn.LocalReth {
		txType = "Burn"
	} else {
		txType = "Arbitrage"
	}

	explorer := ""
	if dataIn.NetworkId == 1 {
		explorer = "https://etherscan.io/tx/"
	} else if dataIn.NetworkId == 17000 {
		explorer = "https://explorer.holesky.io/tx/"
	}
	if len(res.Results) == 2 {
		arbTx := res.Results[1]
		fmt.Printf("Distributed minipool! %s tx: %s%s\n\n", txType, explorer, arbTx.TxHash.Hex())
	} else {
		arbTx := res.Results[len(res.Results)-1]
		fmt.Printf("Distributed minipools! %s tx: %s%s\n\n", txType, explorer, arbTx.TxHash.Hex())
	}

	return nil
}

func evalGasPrices(bundle *flashbots_client.Bundle) (bundleGasPrice, arbitrageGasPrice *big.Int) {
	bundleGasPrice = bundle.MaximumGasFeePaid()

	txs := bundle.Transactions()
	arbTx := txs[len(txs)-1]

	return bundleGasPrice, arbTx.Cost()
}

func waitForUserConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to proceed? (y/n): ")
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSpace(response)
	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("Invalid input. Please type 'y' or 'n'.")
		return waitForUserConfirmation()
	}
}

func getWithdrawalAddress(ctx context.Context, client *ethclient.Client, networkId uint64, nodeAddress common.Address) (common.Address, error) {
	rocketpoolStorageAddress, err := GetRocketpoolStorageAddress(networkId)
	if err != nil {
		return common.Address{}, errors.Join(errors.New("failed to get rocketpool storage address"), err)
	}

	storageInterface, err := storage.NewStorage(rocketpoolStorageAddress, client)
	if err != nil {
		return common.Address{}, errors.Join(errors.New("failed to create storage contract instance"), err)
	}

	timoutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	opts := &bind.CallOpts{
		Pending: false,
		Context: timoutCtx,
	}

	address, err := storageInterface.GetNodeWithdrawalAddress(opts, nodeAddress)
	if err != nil {
		return common.Address{}, errors.Join(errors.New("failed to get node withdrawal address"), err)
	}

	return address, nil
}
