package arbitrage

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_fetchParaswapData(t *testing.T) {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	randomAddress := common.HexToAddress("0xfA82e08c42E6F62f95623F9ee8f2b15716F02aA6")
	type args struct {
		ctx                 context.Context
		logger              *slog.Logger
		amount              *big.Int
		senderAddress       *common.Address
		rEthContractAddress common.Address
		networkId           uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *ParaswapArbitrage
		wantErr bool
	}{}

	amounts := []int64{1, 1*24, 2*24, 3*24, 4*24, 5*24, 6*24, 7*24, 8*24, 9*24, 10*24}
	for _, amt := range amounts {
		newTest :=struct {
			name    string
			args    args
			want    *ParaswapArbitrage
			wantErr bool
		}{
			name: fmt.Sprintf("Swap %d ETH", amt),
			args: args{
				ctx:                 context.Background(),
				logger:              slog.Default(),
				amount:              new(big.Int).Mul(big.NewInt(1e18), big.NewInt(amt)),
				senderAddress:       &randomAddress,
				rEthContractAddress: common.HexToAddress(Mainnet_rEthContractAddressStr),
				networkId:           1,
			},
			want:    &ParaswapArbitrage{},
			wantErr: false,
		}	
		tests = append(tests, newTest)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchParaswapData(tt.args.ctx, tt.args.logger, tt.args.amount, tt.args.senderAddress, tt.args.rEthContractAddress, tt.args.networkId)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchParaswapData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.swapInAmountWeth.Cmp(big.NewInt(0)) == 0 {
				t.Errorf("fetchParaswapData() swapInAmountWeth is zero")
			}
			if got.swapOutAmountReth.Cmp(big.NewInt(0)) == 0 {
				t.Errorf("fetchParaswapData() swapOutAmountReth is zero")
			}
			if len(got.calldata) == 0 {
				t.Errorf("fetchParaswapData() calldata is empty")
			}
		})
	}
}
