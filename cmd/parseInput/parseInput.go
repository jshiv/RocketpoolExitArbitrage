package parseInput

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"rocketpoolArbitrage/arbitrage"
	"rocketpoolArbitrage/rocketpoolContracts/rETH"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xtrooper/flashbots_client"
)

const (
	rETHAddressStr = "0xae78736Cd615f374D3085123A210448E74Fc6393"
)

func Input(ctx context.Context, logger *slog.Logger) (data *arbitrage.DataIn, err error) {
	logger.With(slog.String("function", "input"))

	data = &arbitrage.DataIn{}

	debugFlag := flag.Bool("debug", false, "Enable detailed debug logs")
	commandFlag := flag.String(
		"command",
		"docker exec rocketpool_node /go/bin/rocketpool",
		"Override the default command used to run the Rocket Pool smartnode daemon. Adjust if your container or binary path differs.",
	)
	flag.BoolVar(&data.LocalReth, "local-reth", false, "Use existing local rETH instead of taking a flashloan. If false, the CLI attempts a flashloan")
	minipoolFlag := flag.String("minipool", "", "Single minipool address to distribute. Use --minipools for multiple.")
	minipoolsFlag := flag.String("minipools", "", "Comma-separated list of minipool addresses to distribute.")
	SercherPrivateKeyFlag := flag.String("searcher-private-key", "", "Private key for the searcher used in Flashbots transactions. If not set, a random key is generated.")
	rpcFlag := flag.String("rpc", "http://localhost:8545", "Ethereum RPC endpoint for all on-chain calls. (default: http://localhost:8545)")
	rpcPortFlag := flag.String("rpc-port", "8545", "If using localhost but on a non-default port, override the port here.")
	flag.BoolVar(&data.SkipConfirmation, "skip-confirmation", false, "Skip confirmation prompt before executing")
	flag.BoolVar(&data.SkipConfirmation, "y", false, "Short flag for --skip-confirmation")
	flag.BoolVar(&data.CheckProfit, "check-profit", true, "If enabled, reverts when the profit is too low. (Default: true)")
	flag.BoolVar(&data.CheckProfitIgnoreDistributeCost, "ignoreDistributeCost", false, "Reverts when the profit is too low, but does not considering the distribute call(s). Best used if you want to distribute either way.")
	flag.BoolVar(&data.DryRun, "dry-run", false, "Perform a dry run without sending the bundle to Flashbots; only print the transaction bundle.")
	nodeAddressFlag := flag.String("node-address", "", "Node address used as caller. If not set, the first minipool's node address is used.")
	protocolFlag := flag.String("protocol", "best", "Protocol to use for arbitrage. Options: best, uniswap, paraswap")
	receiverFlag := flag.String("receiver", "", "Receiver address for the arbitrage. If not set, the node address is used.")

	flag.Parse()

	if *debugFlag {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	data.Command = *commandFlag
	logger.Debug("command", slog.String("command", data.Command))

	if *minipoolFlag == "" && *minipoolsFlag == "" {
		return nil, errors.New("\"--minipool\" or \"--minipools\" is required")
	}

	data.MinipoolAddresses = []common.Address{}
	if *minipoolFlag != "" {
		if !common.IsHexAddress(*minipoolFlag) {
			return nil, errors.New("minipool address is invalid")
		}

		data.MinipoolAddresses = append(data.MinipoolAddresses, common.HexToAddress(*minipoolFlag))
		logger.Debug("minipool", slog.String("minipool", common.HexToAddress(*minipoolFlag).Hex()))
	}

	if *minipoolsFlag != "" {
		minipools := strings.Split(*minipoolsFlag, ",")
		for _, minipool := range minipools {
			minipool = strings.TrimSpace(minipool) // handle whitespace
			if !common.IsHexAddress(minipool) {
				return nil, fmt.Errorf("minipool address %s is invalid", minipool)
			}

			data.MinipoolAddresses = append(data.MinipoolAddresses, common.HexToAddress(minipool))
			logger.Debug("minipool", slog.String("minipool", common.HexToAddress(minipool).Hex()))
		}
	}

	var url string
	if *rpcFlag == "http://localhost:8545" {
		if *rpcPortFlag != "8545" {
			url = "http://localhost:" + *rpcPortFlag
		} else {
			url = *rpcFlag
		}
	} else {
		// user should set the full url in the --rpc flag, check that they didn't set --rpcPort
		if *rpcPortFlag != "8545" {
			return nil, errors.New("only use --rpcPort without setting --rpc")
		}

		url = *rpcFlag
	}

	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, errors.Join(errors.New("failed to connect to rpc"), err)
	}

	networkID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Join(errors.New("failed to verify client connection"), err)
	}

	if networkID.Uint64() != 1 && networkID.Uint64() != 17000 {
		return nil, errors.New("only mainnet and holesky are supported")
	}

	data.Client = client
	logger.Debug("rpc connected and verified")

	data.RETHInstance, err = rETH.NewRETH(common.HexToAddress(rETHAddressStr), data.Client)
	if err != nil {
		return nil, err
	}

	var privateKey *ecdsa.PrivateKey
	if *SercherPrivateKeyFlag != "" {
		privateKey, err = crypto.HexToECDSA(*SercherPrivateKeyFlag)
		if err != nil {
			return nil, errors.Join(errors.New("failed to parse ECDSA private key for flashbots searcher"), err)
		}

		fmt.Printf("Using provided ECDSA private key for flashbots searcher (Address: %s)\n", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	} else {
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			return nil, errors.Join(errors.New("failed to generate ECDSA private key for flashbots searcher"), err)
		}
		data.RandomPrivateKey = true
	}

	logger.Debug("parsed ECDSA private key")
	data.FbClient, err = flashbots_client.NewClient(logger, data.Client, privateKey)
	if err != nil {
		return nil, errors.Join(errors.New("failed to create flashbots client"), err)
	}

	if *nodeAddressFlag != "" {
		if !common.IsHexAddress(*nodeAddressFlag) {
			return nil, errors.New("node address is invalid")
		}

		nodeAddress := common.HexToAddress(*nodeAddressFlag)
		data.NodeAddress = &nodeAddress
		logger.Debug("nodeAddress", slog.String("nodeAddress", nodeAddress.Hex()))
	}

	switch *protocolFlag {
	case "best", "b":
		data.Protocol = arbitrage.BestProtocol
	case "uniswap", "u":
		data.Protocol = arbitrage.UniswapProtocol
	case "paraswap", "p":
		data.Protocol = arbitrage.ParaswapProtocol
	default:
		return nil, errors.New("invalid protocol - Options: best, uniswap, paraswap")
	}

	if *receiverFlag != "" {
		if !common.IsHexAddress(*receiverFlag) {
			return nil, errors.New("receiver address is invalid")
		}

		receiverAddress := common.HexToAddress(*receiverFlag)
		data.ReceiverAddress = &receiverAddress
		logger.Debug("receiverAddress", slog.String("receiverAddress", receiverAddress.Hex()))
	}

	// overwrite this as local reth does not generate profit
	if data.LocalReth {
		data.CheckProfit = false
	}

	logger.Debug("localReth", slog.Bool("localReth", data.LocalReth))
	logger.Debug("dryRunFlag", slog.Bool("dryRunFlag", data.DryRun))
	logger.Debug("skipConfirmation", slog.Bool("skipConfirmation", data.SkipConfirmation))
	logger.Debug("checkProfitFlag", slog.Bool("checkProfitFlag", data.CheckProfit))
	logger.Debug("ignoreDistributeCostFlag",
		slog.Bool("checkProfitFlag", data.CheckProfit),
		slog.Bool("ignoreDistributeCostFlag", data.CheckProfitIgnoreDistributeCost),
	)

	return data, nil
}
