package main

import (
	"context"
	"fmt"
	"log/slog"
	"rocketpoolArbitrage/arbitrage"
	"rocketpoolArbitrage/cmd/parseInput"
)

func main() {
	ctx := context.Background()
	logger := slog.Default()

	logger = logger.With(slog.String("module", "distribute"))

	dataIn, err := parseInput.Input(ctx, logger)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = arbitrage.ExecuteDistribute(ctx, logger, dataIn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
