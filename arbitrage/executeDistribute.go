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
	"strings"
	"time"

	"log/slog"

	"github.com/0xtrooper/flashbots_client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

func ExecuteDistribute(ctx context.Context, logger *slog.Logger, dataIn *DataIn) error {
	logger.With(slog.String("function", "Simulate"))

	err := VerifyInputData(ctx, logger, dataIn)
	if err != nil {
		return errors.Join(errors.New("failed to verify input data"), err)
	}

	logger.Debug("verified input data")

	if dataIn.RefundAddress != nil {
		err = dataIn.FbClient.UpdateFeeRefundRecipient(*dataIn.RefundAddress)
		if err != nil {
			return errors.Join(errors.New("failed to update flashbots fee refund recipient"), err)
		}

		fmt.Printf("Updated flashbots fee refund recipient to supplied recipient (%s)\n", (*dataIn.RefundAddress).Hex())
	} else if dataIn.RandomPrivateKey {
		err := dataIn.FbClient.UpdateFeeRefundRecipient(*dataIn.NodeAddress)
		if err != nil {
			return errors.Join(errors.New("failed to update flashbots fee refund recipient to node address"), err)
		}

		fmt.Printf("Updated flashbots fee refund recipient to node address (%s)\n", (*dataIn.NodeAddress).Hex())
	}

	bundle, expectedProfit, err := BuildCall(ctx, logger, *dataIn)
	if err != nil {
		return errors.Join(errors.New("failed to build call"), err)
	}

	logger.Debug("created flashbots client")
	res, success, err := dataIn.FbClient.SimulateBundle(bundle, 0)
	if err != nil {
		return errors.Join(errors.New("failed to simulate bundle"), err)
	}

	for _, tx := range res.Results {
		if tx.Error != "" {
			fmt.Println("error: ", tx.Error)
			fmt.Println("revert reason: ", tx.RevertReason)
			return fmt.Errorf("bundle error: %s", tx.Error)
		}
	}

	maxBundleFees, maxArbitrageFees := evalGasPrices(bundle)

	maxBundleFeesFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(maxBundleFees), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	maxArbitrageFeesFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(maxArbitrageFees), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	expectedProfitFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(expectedProfit), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	fmt.Printf("Simulated bundle (Success: %t):\n", success)
	fmt.Printf("    Expected profit after fees: %.6f, with a tx fee of %.6f\n", expectedProfitFloat-maxBundleFeesFloat, maxBundleFeesFloat)
	fmt.Printf("    Expected profit after arbitrage fees: %.6f, with a tx fee of %.6f (interesting if you want to distribute regardless)\n\n", expectedProfitFloat-maxArbitrageFeesFloat, maxArbitrageFeesFloat)

	if dataIn.DryRun {
		txs := bundle.Transactions()
		fmt.Println("Dry run. Would have sent the following bundle:\n[")
		for i, tx := range txs {
			binary, err := tx.MarshalBinary()
			if err != nil {
				return errors.Join(errors.New("error marshalling transaction"), err)
			}

			txHex := "0x" + hex.EncodeToString(binary)
			fmt.Print("    \"" + txHex + "\"")
			if i != len(txs)-1 {
				fmt.Println(",")
			} else {
				fmt.Println()
			}
		}
		fmt.Println("]")
		return nil
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

	if !dataIn.SkipConfirmation && !waitForUserConfirmation() {
		return errors.New("user did not confirm to proceed")
	}

	// add more builders to improve chance to be included
	networkID, err := dataIn.Client.NetworkID(ctx)
	if err != nil {
		return errors.Join(errors.New("failed to get network id"), err)
	}
	bundle.UseAllBuilders(networkID.Uint64())

	// set target block number
	blockNumber, err := dataIn.Client.BlockNumber(ctx)
	if err != nil {
		return errors.Join(errors.New("failed to get block number"), err)
	}
	bundle.SetTargetBlockNumber(blockNumber + 1)

	// bundle is valid for one minutes
	bundle.SetMaxTimestamp(uint64((time.Now().Add(time.Minute)).Unix()))

	// set replacementUuid
	uuid := uuid.New().String()
	bundle.SetReplacementUuid(uuid)

	// send 3 bundles
	nextBundles, err := bundle.GetBundelsForNextNBlocks(2)
	if err != nil {
		return errors.Join(errors.New("failed to duplicate bundles for the next blocks"), err)
	}

	bundles := []flashbots_client.Bundle{*bundle}
	bundles = append(bundles, nextBundles...)

	var bundleHash common.Hash
	for _, b := range bundles {
		bundleHash, err = dataIn.FbClient.SendBundle(&b)
		if err != nil {
			return errors.Join(errors.New("failed to send bundle"), err)
		}
	}
	fmt.Printf("\nSent bundle with hash: %s. Waiting for up to one minute to see if the transaction is included...\n\n", bundleHash.Hex())

	var successfullyIncluded bool
	timeoutContext, cancel := context.WithTimeout(ctx, time.Second*60)
	for index, b := range bundles {
		successfullyIncluded = waitForBundle(timeoutContext, logger, dataIn.Client, dataIn.FbClient, index, b)
		if successfullyIncluded {
			break
		}		
	}

	if !successfullyIncluded {
		successfullyIncluded, err = dataIn.FbClient.WaitForBundleInclusion(timeoutContext, bundle)
	}
	cancel()
	if err != nil {
		return errors.Join(errors.New("failed to wait for bundle inclusion"), err)
	}

	if !successfullyIncluded {
		return errors.New("bundle was not included in the mempool")
	}

	err = dataIn.FbClient.CancelBundle(uuid)
	if err != nil {
		logger.Debug("failed to cancel bundle", slog.String("error", err.Error()))
	}

	if len(res.Results) == 2 {
		arbTx := res.Results[1]
		fmt.Printf("Distributed minipool! Arbitrage tx: https://etherscan.io/tx/%s\n\n", arbTx.TxHash.Hex())
	} else {
		arbTx := res.Results[len(res.Results)-1]
		fmt.Printf("Distributed minipools! Arbitrage tx: https://etherscan.io/tx/%s\n\n", arbTx.TxHash.Hex())
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

func waitForBundle(ctx context.Context, logger *slog.Logger, client *ethclient.Client, fbClient *flashbots_client.FlashbotsClient, number int, bundle flashbots_client.Bundle) bool {
	targetBlock := bundle.TargetBlockNumber()
	logger.Debug("waiting for bundle inclusion", slog.Uint64("targetBlock", targetBlock))

	var firstTime bool = true
	for {
		// 1. Check if the context has been canceled or timed out
		select {
		case <-ctx.Done():
			logger.Debug("timed out waiting for bundle inclusion")
			return false
		default:
			// continue
		}

		// 2. Check if the bundle has been included
		success, err := fbClient.CheckBundleIncusion(ctx, &bundle)
		if err != nil {
			logger.Debug("failed to check bundle inclusion", slog.String("error", err.Error()))
			return false
		}

		if success {
			return true
		}

		// 3. Check bundle stats
		stats, err := fbClient.GetBundleStatsV2(&bundle)
		if err != nil {
			logger.Debug("failed to get bundle stats", slog.String("error", err.Error()))
			return false
		}

		if logger.Enabled(ctx, slog.LevelInfo) {
			if !stats.IsSimulated {
				fmt.Printf("Bundle Nr. %d: Not yet seen by relay\n", number + 1)		
			} else {
				if firstTime {
					fmt.Printf("Bundle Nr. %d: Received at %s and simulated at %s\n", number + 1, stats.ReceivedAt, stats.SimulatedAt)
					firstTime = false
				}

				fmt.Printf("Bundle Nr. %d: Considered by %d builders and sealed by %d builders\n",
					number + 1,
					len(stats.ConsideredByBuilders),
					len(stats.SealedByBuilders),
				)
			}
		}

		// 4. Check if the target block has been reached
		blockNumber, err := client.BlockNumber(ctx)
		if err != nil {
			logger.Debug("failed to get block number", slog.String("error", err.Error()))
			return false
		}

		if blockNumber >= targetBlock {
			fmt.Printf("Bundle Nr. %d: Target block reached\n", number + 1)
			logger.Debug("target block reached")
			return false
		}

		time.Sleep(time.Second * 3)
	}
}
