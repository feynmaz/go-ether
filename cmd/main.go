package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/feynmaz/go-ether/account"
	"github.com/feynmaz/go-ether/balance"
	"github.com/feynmaz/go-ether/config"
	"github.com/feynmaz/go-ether/todo"
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
		KeyPath:    "/root/go/src/github.com/feynmaz/go-ether/data/wallet/UTC--2024-05-02T20-22-14.699851585Z--02579e833024f2287612fba0b3767c1cd79a6007",
		Passphrase: "password0",
		Address:    "0x02579E833024f2287612FbA0b3767c1Cd79A6007",
	}
	key0, err := acc0.GetAccountKey()
	if err != nil {
		return fmt.Errorf("failed to get account 0 key: %w", err)
	}
	// acc1 := account.Account{
	// 	KeyPath:    "/root/go/src/github.com/feynmaz/go-ether/data/wallet/UTC--2024-05-02T20-22-16.535421374Z--9e23c27fa3912c6dc54e5dd8233918422966ba51",
	// 	Passphrase: "password1",
	// 	Address:    "0x9e23c27fA3912C6Dc54E5DD8233918422966BA51",
	// }

	cfg := config.GetDefault()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}
	defer client.Close()

	fmt.Println("---- Accouts ----")
	err = balance.PrintEthBalance(ctx, client, key0)
	if err != nil {
		return fmt.Errorf("failed to print balance of account 0: %w", err)
	}

	//

	fmt.Println("---- Chain ----")

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}
	fmt.Println("chainID:", chainID.String())

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to suggest gas price ID: %w", err)
	}
	fmt.Println("gasPrice:", gasPrice.String())

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to fetch the latest block: %w", err)
	}
	gasLimit := block.GasLimit()
	fmt.Println("gasLimit:", gasLimit)

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(acc0.Address))
	if err != nil {
		return fmt.Errorf("failed to get pending nonce: %w", err)
	}
	fmt.Println("nonce:", nonce)

	//

	fmt.Println("---- Smart contract ----")

	auth, err := bind.NewKeyedTransactorWithChainID(key0.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create new keyed transactor: %w", err)
	}
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(1))
	auth.Nonce = big.NewInt(int64(nonce))

	// td, err := todo.NewTodo(common.HexToAddress("0xA4ec0541d420f1526eb79B20DBE5f7Cc40dc8bF2"), client)
	// if err != nil {
	// 	return fmt.Errorf("failed to create new todo: %w", err)
	// }

	// tasks, err := td.List(&bind.CallOpts{
	// 	From: common.HexToAddress(acc0.Address),
	// })
	// if err != nil {
	// 	return fmt.Errorf("failed to list todo: %w", err)
	// }
	// fmt.Println(tasks)

	address, tx, td, err := todo.DeployTodo(auth, client)
	if err != nil {
		return fmt.Errorf("failed to deploy todo: %w", err)
	}
	fmt.Println("address:", address.Hex())
	fmt.Println("tx:", tx.Hash().Hex())

	tx, err = td.Add(auth, "task1")
	if err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}
	fmt.Println("todo.Add() tx:", tx.Hash().Hex())

	return nil
}
