package main

import (
	"context"
	"fmt"

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

	fmt.Println(block.Number())

	return nil
}
