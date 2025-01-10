package beaconchain

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Minipool struct {
	PublicKey             string
	WithdrawalCredentials common.Hash
	Index                 int
	Status                string
	TimeToExitStart       time.Duration
	TimeToExit            time.Duration
	TimeToWithdrawEst     time.Duration
	ExitEpoch             uint64
	WithdrawableEpoch     uint64
	ActivationEpoch       uint64
	WithdrawalAddress     common.Address
	Balance               *big.Int
}

type Data struct {
	CurrentEpoch       uint64
	CurrentSlot        int
	Minipools          []Minipool
	ValidatorsEntering int
	ValidatorsExiting  int
	ValidatorsCount    int
	LastValidatorIndex int
	WithdrawalPerSlot  int
	NetworkId          int
}
