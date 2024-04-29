package main

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/feynmaz/go-ether/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetDefault()
	log.Debug(cfg.String())

	// TODO: make context
	ctx := context.TODO()

	client, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get block: %w", err)
	}

	fmt.Println("The latest block number:", block.Number())

	addr := "0xB61Ff46584078A85246B4520173b3640Db871666"
	address := common.HexToAddress(addr)

	weiBalance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}
	// 1 ether = 10^18 wei
	
	ethBalance := new(big.Float)
	ethBalance.SetString(weiBalance.String())
	ethBalance = new(big.Float).Quo(ethBalance, big.NewFloat(math.Pow10(18)))
	
	fmt.Printf("The %s address balance: %v ETH \n", addr, ethBalance)

	return nil
}
