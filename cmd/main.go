package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/feynmaz/go-ether/account"
	"github.com/feynmaz/go-ether/balance"
	"github.com/feynmaz/go-ether/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	acc0 := account.Account{
		KeyPath:    "../data/wallet/UTC--2024-04-30T18-16-11.115357251Z--c0e3cad8caf06c53588efe33d160316a7b9e7ce8",
		Passphrase: "password0",
		Address:    "0xc0E3Cad8CAF06C53588EFE33D160316a7B9e7cE8",
	}
	acc1 := account.Account{
		KeyPath:    "../data/wallet/UTC--2024-04-30T18-16-12.792380414Z--057728f9f90c3bef651c3f33430d72089d0b8fb6",
		Passphrase: "password1",
		Address:    "0x057728F9f90c3BEF651c3f33430D72089d0b8fB6",
	}

	key0, err := acc0.GetAccountKey()
	if err != nil {
		return fmt.Errorf("failed to get account 0 key: %w", err)
	}
	key1, err := acc1.GetAccountKey()
	if err != nil {
		return fmt.Errorf("failed to get account 1 key: %w", err)
	}

	cfg := config.GetDefault()
	log.Debug(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}
	defer client.Close()

	if err := balance.PrintEthBalance(ctx, client, key0); err != nil {
		return fmt.Errorf("failed to print balance of account 0: %w", err)
	}
	if err := balance.PrintEthBalance(ctx, client, key1); err != nil {
		return fmt.Errorf("failed to print balance of account 1: %w", err)
	}

	// Make transaction
	var gasLimit uint64 = 30_000_000
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(acc0.Address))
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	fmt.Println("gasLimit:", gasLimit, "wei")
	fmt.Println("gasPrice:", gasPrice, "wei")
	fmt.Println("nonce:", nonce, "wei")

	gasLimitBigInt := big.NewInt(int64(gasLimit))
	totalCostWei := new(big.Int).Mul(gasPrice, gasLimitBigInt)
	etherValue := new(big.Float).Quo(new(big.Float).SetInt(totalCostWei), big.NewFloat(1e18))
	etherValueF, _ := etherValue.Float64()
	fmt.Printf("transaction cost = gasPrice*gasLimit: %v ETH \n", etherValueF)

	amount := big.NewInt(1e17) // 0.1 ETH
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(acc1.Address),
		amount,
		gasLimit,
		gasPrice,
		nil,
	)

	// Sign transaction
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get network id: %w", err)
	}

	tx, err = types.SignTx(
		tx,
		types.LatestSignerForChainID(chainID),
		key0.PrivateKey,
	)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}
	fmt.Printf("transaction hash: %s", tx.Hash().Hex())

	return nil
}
