package arbitrage

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"os/exec"
	"rocketpoolArbitrage/arbitrage/contract"
	"rocketpoolArbitrage/rocketpoolContracts/minipoolDelegate"
	"rocketpoolArbitrage/rocketpoolContracts/rETH"
	uniswap "rocketpoolArbitrage/uniswapContracts"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"

	"github.com/0xtrooper/flashbots_client"
)

const (
	BURN_CALL_MAX_GAS = 200000 // roughly 140k to burn
	ARBITRAGE_CALL_MAX_GAS = 400000 // roughly uniswap 180k + burn 140k
	DISTRIBUTE_CALL_MAX_GAS = 500000 // roughly 400k

	arbitrageContractAddressStr = "0x241bfb6e47d478456bf20aad81ecb512aed223c8"
	rEthContractAddressStr      = "0xae78736Cd615f374D3085123A210448E74Fc6393"
)

func BuildCall(ctx context.Context, logger *slog.Logger, dataIn DataIn) (*flashbots_client.Bundle, *big.Int, error) {
	logger.With(slog.String("function", "BuildCall"))

	rethToBurn, expectedProfit, swapAmount, poolAddress, err := calcualteArbitrageData(ctx, logger, dataIn)
	if err != nil {
		return nil, nil, errors.Join(errors.New("failed to calculate arbitrage data"), err)
	}

	if logger.Enabled(ctx, slog.LevelInfo) {
		expectedProfitFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(expectedProfit), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
		fmt.Printf("Calculated Arbitrage: expected profit (before fees) = %.6f ETH\n\n", expectedProfitFloat)
	}

	baseGas, tipGas, err := getCurrentGasSettings(ctx, dataIn.Client)
	if err != nil {
		return nil, nil, errors.Join(errors.New("failed to get current gas settings"), err)
	}

	baseGasBoosted := new(big.Int).Div(new(big.Int).Mul(baseGas, big.NewInt(150)), big.NewInt(100))

	logger.Debug("fetched current gas settings",
		slog.String("baseGas", baseGas.String()),
		slog.String("baseGasBoosted", baseGasBoosted.String()),
		slog.String("tipGas", tipGas.String()),
	)

	nonce, err := getCurrentNonce(ctx, dataIn.Client, *dataIn.NodeAddress)
	if err != nil {
		return nil, nil, errors.Join(errors.New("failed to get current nonce"), err)
	}

	txs, err := generateAndBuildDistributeCalls(nonce, dataIn.MinipoolAddresses, baseGasBoosted, tipGas, logger, dataIn.Command)
	if err != nil {
		return nil, nil, errors.Join(errors.New("failed to generate distribute calls"), err)
	}

	count := len(txs)
	logger.Debug("signed distribute txs", slog.Int("count", count))

	if dataIn.LocalReth {
		rawBurnTx, err := generateBurnCall(nonce+uint64(count), rethToBurn, baseGasBoosted, tipGas)
		if err != nil {
			return nil, nil, errors.Join(errors.New("failed to generate burn call"), err)
		}

		signedBurnTx, err := signTransaction(logger, dataIn.Command, rawBurnTx)
		if err != nil {
			return nil, nil, errors.Join(errors.New("failed to sign burn tx"), err)
		}
	
		logger.Debug("signed burn tx", slog.String("txHash", signedBurnTx.Hash().Hex()))
		txs = append(txs, signedBurnTx)
	} else {
		maxGasUsed := len(dataIn.MinipoolAddresses) * DISTRIBUTE_CALL_MAX_GAS + ARBITRAGE_CALL_MAX_GAS
		minProfit := new(big.Int).Mul(baseGasBoosted, big.NewInt(int64(maxGasUsed)))
		if logger.Enabled(ctx, slog.LevelDebug) {
			minProfitFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(minProfit), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
			logger.Debug("calculated min profit", slog.Float64("minProfit", minProfitFloat))
		}

		rawArbitrageTx, err := generateArbitrageCall(nonce+uint64(count), poolAddress, swapAmount, minProfit, baseGasBoosted, tipGas)
		if err != nil {
			return nil, nil, errors.Join(errors.New("failed to generate arbitrage call"), err)
		}

		signedArbitrageTx, err := signTransaction(logger, dataIn.Command, rawArbitrageTx)
		if err != nil {
			return nil, nil, errors.Join(errors.New("failed to sign arbitrage tx"), err)
		}
	
		logger.Debug("signed arbitrage tx", slog.String("txHash", signedArbitrageTx.Hash().Hex()))
		txs = append(txs, signedArbitrageTx)
	}


	bundle := flashbots_client.NewBundleWithTransactions(txs)

	return bundle, expectedProfit, nil
}

func calcaulteDistributedBalance(ctx context.Context, logger *slog.Logger, dataIn DataIn) (*big.Int, error) {
	totalNodeShare := new(big.Int)
	totalDistributeAmount := new(big.Int)
	for _, minipoolAddress := range dataIn.MinipoolAddresses {
		balance, err := dataIn.Client.BalanceAt(ctx, minipoolAddress, nil)
		if err != nil {
			return nil, errors.Join(errors.New("failed to get minipool balance"), err)
		}

		logger.Debug("fetched minipool balance",
			slog.String("minipool", minipoolAddress.Hex()),
			slog.String("balance", balance.String()),
		)

		minipoolInstance, err := minipoolDelegate.NewMinipoolDelegate(minipoolAddress, dataIn.Client)
		if err != nil {
			return nil, err
		}

		refundBalance, err := GetMinipoolRefundBalance(ctx, minipoolInstance)
		if err != nil {
			return nil, errors.Join(errors.New("failed to get refund balance"), err)
		}

		distributeAmount := new(big.Int).Sub(balance, refundBalance)

		logger.Debug("fetched refund balance",
			slog.String("minipool", minipoolAddress.Hex()),
			slog.String("refundBalance", refundBalance.String()),
			slog.String("distributeAmount", distributeAmount.String()),
		)

		rETHShare, err := CalculateMinipoolUserShare(ctx, minipoolInstance, distributeAmount)
		if err != nil {
			return nil, errors.Join(errors.New("failed to calculate rETH share"), err)
		}

		nodeShare := new(big.Int).Sub(balance, rETHShare)

		totalDistributeAmount = new(big.Int).Add(totalDistributeAmount, rETHShare)
		totalNodeShare = new(big.Int).Add(totalNodeShare, nodeShare)
	}

	totalNodeShareFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(totalNodeShare), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	rETHShareFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(totalDistributeAmount), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	fmt.Printf("Calculated distribution amounts: %.6f ETH send to NO, %.6f ETH send to RP\n\n", totalNodeShareFloat, rETHShareFloat)

	return totalDistributeAmount, nil
}

func calcualteArbitrageData(ctx context.Context, logger *slog.Logger, dataIn DataIn) (*big.Int, *big.Int, *big.Int, common.Address, error) {
	rETHShare, err := calcaulteDistributedBalance(ctx, logger, dataIn)
	if err != nil {
		return nil, nil, nil, common.Address{}, errors.Join(errors.New("failed to calculate distributed balance"), err)
	}

	// calcaulte amount rETH to burn
	rethToBurn, err := ConvertWethToReth(ctx, dataIn.RETHInstance, rETHShare)
	if err != nil {
		return nil, nil, nil, common.Address{}, errors.Join(errors.New("failed to convert rETH to WETH"), err)
	}

	// get best pool to swap rETH
	poolAddress, uniswapReturnAmountWeth, err := uniswap.GetBestPoolWithdrawArb(ctx, logger, dataIn.Client, rethToBurn)
	if err != nil {
		return nil, nil, nil, common.Address{}, errors.Join(errors.New("failed to get best pool"), err)
	}
	
	poolAmountInFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(uniswapReturnAmountWeth), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	poolAmountOutFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(rethToBurn), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	secondaryRatio := poolAmountInFloat / poolAmountOutFloat

	rethToBurnFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(rethToBurn), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	ethUnlockedFloat, _ := new(big.Float).Quo(new(big.Float).SetInt(rETHShare), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	primaryRatio := ethUnlockedFloat / rethToBurnFloat

	uniswapReturnAmountWeth = new(big.Int).Sub(uniswapReturnAmountWeth, big.NewInt(5))
	
	// early exit if we use the users rETH
	if dataIn.LocalReth {
		if logger.Enabled(ctx, slog.LevelInfo) {
			fmt.Printf("At the secondary ratio of %f, you would swap %.6f rETH to %.6f WETH\n\n", secondaryRatio, poolAmountOutFloat, poolAmountInFloat)
			fmt.Printf("At the primary ratio of %f, you will burn %.6f rETH for %.6f ETH\n\n", primaryRatio, rethToBurnFloat, ethUnlockedFloat)
		}
		return rethToBurn, new(big.Int).Sub(rETHShare, uniswapReturnAmountWeth), big.NewInt(0), common.Address{}, nil
	} else {
		if logger.Enabled(ctx, slog.LevelInfo) {
			// update user about the secondary ratio
			fmt.Printf("Selected uniswap pool - %s:\n", poolAddress.String())
			fmt.Printf("    Swapping %.6f WETH to %.6f rETH at a secondary ratio of %.5f\n\n",
				poolAmountInFloat,
				poolAmountOutFloat,
				secondaryRatio,
			)
	
			// update user about the primary ratio
			fmt.Printf("Calculated rETH to burn: Burning %.6f rETH for %.6f ETH at a primary ratio of %.5f.\n\n", 
				rethToBurnFloat,
				ethUnlockedFloat,
				primaryRatio,
			)
		}
	
		return rethToBurn, new(big.Int).Sub(rETHShare, uniswapReturnAmountWeth), uniswapReturnAmountWeth, poolAddress, nil
	}

}

func getCurrentGasSettings(ctx context.Context, client *ethclient.Client) (baseGas *big.Int, tipGas *big.Int, err error) {
	// Get gas price
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	var eg errgroup.Group
	eg.Go(func() error {
		var err error
		baseGas, err = client.SuggestGasPrice(timeoutCtx)
		if err != nil {
			return errors.Join(errors.New("failed to get suggested gas price"), err)
		} else {
			return nil
		}
	})

	eg.Go(func() error {
		var err error
		tipGas, err = client.SuggestGasTipCap(timeoutCtx)
		if err != nil {
			return errors.Join(errors.New("failed to get gas tip cap"), err)
		} else {
			return nil
		}
	})

	return baseGas, tipGas, eg.Wait()
}

func getCurrentNonce(ctx context.Context, client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return 0, errors.Join(errors.New("failed to get current nonce"), err)
	}

	return nonce, nil
}

func generateAndBuildDistributeCalls(nonce uint64, minipoolAddresses []common.Address, baseGas, tipGas *big.Int, logger *slog.Logger, apiCommand string) ([]*types.Transaction, error) {
	var txs []*types.Transaction

	for i, minipoolAddress := range minipoolAddresses {
		rawTx, err := generateDistributeCall(nonce+uint64(i), minipoolAddress, baseGas, tipGas)
		if err != nil {
			return nil, errors.Join(errors.New("failed to generate distribute call"), err)
		}

		signedTx, err := signTransaction(logger, apiCommand, rawTx)
		if err != nil {
			return nil, errors.Join(errors.New("failed to sign distribute tx"), err)
		}

		txs = append(txs, signedTx)
	}

	return txs, nil
}

func generateDistributeCall(nonce uint64, minipoolAddress common.Address, baseGas, tipGas *big.Int) (*types.Transaction, error) {
	minipoolAbi, err := abi.JSON(strings.NewReader(minipoolDelegate.MinipoolDelegateABI))
	if err != nil {
		return nil, errors.Join(errors.New("failed to get minipool ABI"), err)
	}

	// Pack the function call with parameters
	callData, err := minipoolAbi.Pack("distributeBalance", false)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function data: %v", err)
	}

	dynTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     nonce,
		GasFeeCap: baseGas,
		GasTipCap: tipGas,
		To:        &minipoolAddress,
		Value:     big.NewInt(0),
		Gas:       DISTRIBUTE_CALL_MAX_GAS,
		Data:      callData,
	}

	return types.NewTx(dynTx), nil
}

func generateArbitrageCall(nonce uint64, uniswapPool common.Address, amountWethToTrade, minProfit, baseGas, tipGas *big.Int) (*types.Transaction, error) {
	arbitrageAbi, err := abi.JSON(strings.NewReader(contract.ContractABI))
	if err != nil {
		return nil, errors.Join(errors.New("failed to get arbitrage ABI"), err)
	}

	callData, err := arbitrageAbi.Pack("arb", uniswapPool, amountWethToTrade, minProfit)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function data: %v", err)
	}

	arbitrageContractAddress := common.HexToAddress(arbitrageContractAddressStr)

	dynTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     nonce,
		GasFeeCap: baseGas,
		GasTipCap: tipGas,
		To:        &arbitrageContractAddress,
		Value:     big.NewInt(0),
		Gas:       ARBITRAGE_CALL_MAX_GAS,
		Data:      callData,
	}

	return types.NewTx(dynTx), nil
}

func generateBurnCall(nonce uint64, rethToBurn, baseGas, tipGas *big.Int) (*types.Transaction, error) {
	rEthAbi, err := abi.JSON(strings.NewReader(rETH.RETHABI))
	if err != nil {
		return nil, errors.Join(errors.New("failed to get arbitrage ABI"), err)
	}

	callData, err := rEthAbi.Pack("burn", rethToBurn)
	if err != nil {
		return nil, fmt.Errorf("failed to pack function data: %v", err)
	}

	rethContractAddress := common.HexToAddress(rEthContractAddressStr)

	dynTx := &types.DynamicFeeTx{
		ChainID:   big.NewInt(1),
		Nonce:     nonce,
		GasFeeCap: baseGas,
		GasTipCap: tipGas,
		To:        &rethContractAddress,
		Value:     big.NewInt(0),
		Gas:       BURN_CALL_MAX_GAS,
		Data:      callData,
	}

	return types.NewTx(dynTx), nil
}

func signTransaction(logger *slog.Logger, apiCommand string, tx *types.Transaction) (*types.Transaction, error) {
	serializedTx, err := tx.MarshalBinary()
	if err != nil {
		return nil, errors.Join(errors.New("failed to serialize tx"), err)
	}

	args := strings.Fields(apiCommand) // split on whitespace, not commas
	args = append(args, "api", "node", "sign", hex.EncodeToString(serializedTx))

	cmd := exec.Command(args[0], args[1:]...)

	logger.Debug("signing tx", slog.String("cmd", strings.Join(args, " ")))
	stdout, err := cmd.Output()

	if err != nil {
		return nil, errors.Join(errors.New("failed to sign tx"), err)
	}

	var result struct {
		Status     string `json:"status"`
		Error      string `json:"error"`
		SignedData string `json:"signedData"`
	}

	if err := json.Unmarshal(stdout, &result); err != nil {
		return nil, errors.Join(errors.New("failed to parse JSON response"), err)
	}

	if result.Status != "success" {
		return nil, fmt.Errorf("failed to sign tx: %s", result.Error)
	}

	parsedData, err := hex.DecodeString(result.SignedData[2:])
	if err != nil {
		return nil, errors.Join(errors.New("failed to decode signed tx"), err)
	}

	signedTx := types.Transaction{}
	if err := signedTx.UnmarshalBinary(parsedData); err != nil {
		return nil, errors.Join(errors.New("failed to unmarshal signed tx"), err)
	}

	return &signedTx, nil
}