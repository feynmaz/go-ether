package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/feynmaz/go-ether/config"
	"github.com/feynmaz/go-ether/wallet"
	log "github.com/sirupsen/logrus"
)

var (
	walletPath        = "../data/wallet"
	createNewAccounts = false
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetDefault()
	log.Debug(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}
	defer client.Close()

	var hexAddress1, hexAddress2 common.Address
	if createNewAccounts {
		walletPath = "../data/wallet"
		key := wallet.CreateWallet(walletPath)

		passphrase1 := "password1"
		acc1, err := wallet.CreateAccount(key, passphrase1)
		if err != nil {
			return fmt.Errorf("failed to create account: %w", err)
		}
		hexAddress1 = acc1.Address

		passphrase2 := "password1"
		acc2, err := wallet.CreateAccount(key, passphrase2)
		if err != nil {
			return fmt.Errorf("failed to create account: %w", err)
		}
		hexAddress2 = acc2.Address

	} else {
		hexAddress1 = wallet.GetAddress("0xAc5fc0Da40229Da5b0Fa582d41354F9BF4f4ca8a")
		hexAddress2 = wallet.GetAddress("0x11Bc2AfD266FDa58C771630119dD5b3a7D0Fa6B7")
	}

	balance1, err := client.BalanceAt(ctx, hexAddress1, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	balance2, err := client.BalanceAt(ctx, hexAddress2, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	fmt.Println("balance 1:", balance1)
	fmt.Println("balance 2:", balance2)

	return nil
}
