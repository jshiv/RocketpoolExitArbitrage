package main

import (
	"errors"
	"flag"
	"fmt"
	"rocketpoolArbitrage/beaconchain"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EntryChurnLimit = 8
	ExitChurnLimit = 16
	EpochsAfterExitStart = 256
)

func main() {
	nodeAddress, eth2Url, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := beaconchain.GetBeaconchainData(nodeAddress, eth2Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	estimateTimeToExit(data)

	printTimeToExit(data)	
}

func estimateTimeToExit(data *beaconchain.Data) {
	for i, minipool := range data.Minipools {
		switch minipool.Status {
		case "unknown":
			minipool.TimeToExit = 0
		case "active_exiting":
			// 1st check when exit will happen:
			epochsBeforeExitStart := uint64(0)
			if data.CurrentEpoch < minipool.ExitEpoch {
				epochsBeforeExitStart = minipool.ExitEpoch - data.CurrentEpoch
			}

			// 2nd check withdrawal queue
			withdrawalQueueLength := data.ValidatorsExiting / ExitChurnLimit

			var epochsToGo uint64
			// does the validator need to queue?
			if withdrawalQueueLength > EpochsAfterExitStart {
				epochsToGo = epochsBeforeExitStart + uint64(withdrawalQueueLength)
			} else {
				epochsToGo = epochsBeforeExitStart + uint64(EpochsAfterExitStart)
			}

			minipool.TimeToExitStart = time.Duration(epochsBeforeExitStart) * 32 * 12 * time.Second // 32 slots per epoch @ 12s per slot
			minipool.TimeToExit = time.Duration(epochsToGo) * 32 * 12 * time.Second // 32 slots per epoch @ 12s per slot

			// estimate time to withdraw
			var slotsToNextWithdraw int
			if minipool.Index > data.LastValidatorIndex {
				slotsToNextWithdraw = minipool.Index - data.LastValidatorIndex
			} else {
				slotsToNextWithdraw = data.ValidatorsCount - minipool.Index
			}
			minipool.TimeToWithdrawEst = time.Duration(slotsToNextWithdraw / data.WithdrawalPerSlot) * 12 * time.Second
			timeForOneSweep := time.Duration(data.ValidatorsCount / data.WithdrawalPerSlot) * 12 * time.Second
			for minipool.TimeToWithdrawEst < minipool.TimeToExit {
				minipool.TimeToWithdrawEst += timeForOneSweep
			}
		case "active_ongoing":
			// 1st check if old enough:
			epochsBeforeExitStart := uint64(0)
			if data.CurrentEpoch - 256 < minipool.ActivationEpoch {
				epochsBeforeExitStart = 256 - (data.CurrentEpoch - minipool.ActivationEpoch)
			}

			// 2nd check withdrawal queue
			withdrawalQueueLength := data.ValidatorsExiting / ExitChurnLimit

			var epochsToGo uint64
			// does the validator need to queue?
			if withdrawalQueueLength > EpochsAfterExitStart {
				epochsToGo = epochsBeforeExitStart + uint64(withdrawalQueueLength)
			} else {
				epochsToGo = epochsBeforeExitStart + uint64(EpochsAfterExitStart)
			}

			minipool.TimeToExitStart = time.Duration(epochsBeforeExitStart) * 32 * 12 * time.Second // 32 slots per epoch @ 12s per slot
			minipool.TimeToExit = time.Duration(epochsToGo) * 32 * 12 * time.Second // 32 slots per epoch @ 12s per slot

			// estimate time to withdraw
			var slotsToNextWithdraw int
			if minipool.Index > data.LastValidatorIndex {
				slotsToNextWithdraw = minipool.Index - data.LastValidatorIndex
			} else {
				slotsToNextWithdraw = data.ValidatorsCount - minipool.Index
			}
			minipool.TimeToWithdrawEst = time.Duration(slotsToNextWithdraw / data.WithdrawalPerSlot) * 12 * time.Second
			timeForOneSweep := time.Duration(data.ValidatorsCount / data.WithdrawalPerSlot) * 12 * time.Second
			for minipool.TimeToWithdrawEst < minipool.TimeToExit {
				minipool.TimeToWithdrawEst += timeForOneSweep
			}
		case "active_slashed":
			var epochsToGo uint64
			withdrawalQueueLength := data.ValidatorsExiting / ExitChurnLimit
			// does the validator need to queue?
			if withdrawalQueueLength > EpochsAfterExitStart {
				epochsToGo = uint64(withdrawalQueueLength)
			} else {
				epochsToGo = EpochsAfterExitStart
			}
			minipool.TimeToExit = time.Duration(epochsToGo) * 32 * 12 * time.Second // 32 slots per epoch @ 12s per slot
		case "exited_slashed", "exited_unslashed", "withdrawal_possible":
			var slotsToNextWithdraw int
			if minipool.Index > data.LastValidatorIndex {
				slotsToNextWithdraw = minipool.Index - data.LastValidatorIndex
			} else {
				slotsToNextWithdraw = data.ValidatorsCount - minipool.Index
			}
			minipool.TimeToWithdrawEst = time.Duration(slotsToNextWithdraw / data.WithdrawalPerSlot) * 12 * time.Second // 32 slots per epoch @ 12s per slot
		case "pending_initialized", "pending_queued", "withdrawal_done":
			minipool.TimeToExit = 0
		}

		data.Minipools[i] = minipool
	}
}

func printTimeToExit(data *beaconchain.Data) {
	var states = make(map[string]int)
	for _, minipool := range data.Minipools {
		switch minipool.Status {
		case "withdrawal_done", "pending_initialized", "pending_queued":
			//skip
		case "unknown":
			fmt.Printf("Unknown status for validator %s\n", minipool.PublicKey)
		case "exited_slashed", "exited_unslashed":
			fmt.Printf("Validator %s is exiting, will be withdrawn in less then %s\n", minipool.PublicKey, minipool.TimeToWithdrawEst)
		case "withdrawal_possible":
			fmt.Printf("Validator %s is exited, will be withdrawn in less then %s\n", minipool.PublicKey, minipool.TimeToWithdrawEst)

		case "active_exiting":

		case "active_ongoing":
			if minipool.TimeToExitStart > 0 {
				fmt.Printf("Validator %s is active. You can start the exit it in %s, it will be exitable in %s and withdrawn in roughly %s.\n", 
					minipool.PublicKey,
					minipool.TimeToExitStart,
					minipool.TimeToExit,
					minipool.TimeToWithdrawEst,
				)
			} else {
				fmt.Printf("Validator %s is active. If you exit now it will be exitable in %s and withdrawn in roughly %s.\n", 
					minipool.PublicKey,
					minipool.TimeToExit,
					minipool.TimeToWithdrawEst,
				)
			}

		case "active_slashed":
			fmt.Printf("Validator %s was slashed, will be exited in %s\n", 
				minipool.PublicKey,
				minipool.TimeToExit,
			)
		}
		states[minipool.Status]++
	}
}

func parseInput() (common.Address, string, error) {
	nodeAddressFlag := flag.String("nodeAddress", "", "Node address")
	rpcFlag := flag.String("rpc", "http://localhost:5052", "Ethereum Beacon Node RPC endpoint for all on-chain calls. (default: http://localhost:5052)")
	rpcPortFlag := flag.String("rpc-port", "5052", "If using localhost but on a non-default port, override the port here.")

	flag.Parse()

	var nodeAddress common.Address
	if *nodeAddressFlag != "" {
		if !common.IsHexAddress(*nodeAddressFlag) {
			return common.Address{}, "", errors.New("invalid node address")
		}

		nodeAddress = common.HexToAddress(*nodeAddressFlag)
	}	

	var eth2Url string
	if *rpcFlag == "http://localhost:5052" {
		if *rpcPortFlag != "5052" {
			eth2Url = "http://localhost:" + *rpcPortFlag
		} else {
			eth2Url = *rpcFlag
		}
	} else {
		// user should set the full url in the --rpc flag, check that they didn't set --rpcPort
		if *rpcPortFlag != "5052" {
			return common.Address{}, "", errors.New("only use --rpc-port without setting --rpc")
		}

		eth2Url = *rpcFlag
	}

	eth2Url = strings.TrimRight(eth2Url, "/")
	
	return nodeAddress, eth2Url, nil
}
