package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/feynmaz/go-ether/account"
	"github.com/feynmaz/go-ether/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{ PrettyPrint: true})
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	accounts := []account.Account{
		{
			KeyPath:    "../data/wallet/UTC--2024-04-30T18-16-11.115357251Z--c0e3cad8caf06c53588efe33d160316a7b9e7ce8",
			Passphrase: "password0",
		},
		{
			KeyPath:    "../data/wallet/UTC--2024-04-30T18-16-12.792380414Z--057728f9f90c3bef651c3f33430d72089d0b8fb6",
			Passphrase: "password1",
		},
	}

	key0, err := account.GetAccountKey(accounts[0])
	if err != nil {
		return fmt.Errorf("failed to get account 0 key: %w", err)
	}
	key1, err := account.GetAccountKey(accounts[1])
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

	balance0, err := client.BalanceAt(ctx, key0.Address, nil)
	if err != nil {
		return fmt.Errorf("failed to get 0 balance: %w", err)
	}
	log.Infof("%s balance: %d", key0.Address, balance0)

	balance1, err := client.BalanceAt(ctx, key1.Address, nil)
	if err != nil {
		return fmt.Errorf("failed to get 1 balance: %w", err)
	}
	log.Infof("%s balance: %d", key1.Address, balance1)


	// gasPrice, err := client.SuggestGasPrice(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to suggest gas price: %w", err)
	// }

	// nonce1, err := client.PendingNonceAt(ctx, hexAddress1)
	// if err != nil {
	// 	return fmt.Errorf("failed to get nonce: %w", err)
	// }

	// amount := big.NewInt(10_000_000_000_000_000) // 0.01 ETH

	// networkId, err := client.NetworkID(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to get network id: %w", err)
	// }

	// privateKey, err := wallet.GetPrivateKey(walletPath, hexAddress1.String(), passphrase1)
	// if err != nil {
	// 	return fmt.Errorf("failed to get private key: %w", err)
	// }

	// tx := types.NewTransaction(
	// 	nonce1,
	// 	hexAddress2,
	// 	amount,
	// 	defaultGasLimit,
	// 	gasPrice,
	// 	nil,
	// )

	// tx, err = types.SignTx(tx, types.NewEIP155Signer(networkId), privateKey)
	// if err != nil {
	// 	return fmt.Errorf("failed to sign transaction: %w", err)
	// }

	// err = client.SendTransaction(ctx, tx)
	// if err != nil {
	// 	return fmt.Errorf("failed to send transaction: %w", err)
	// }

	// log.Infof("transaction sent: %s", tx.Hash().Hex())

	return nil
}
