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
		KeyPath:    "../data/wallet/UTC--2024-05-01T15-39-49.425514661Z--a6fcba9bb03cd730e40984360ab15093f85520d3",
		Passphrase: "password0",
		Address:    "0xA6fCbA9bb03cd730e40984360aB15093f85520d3",
	}
	// acc1 := account.Account{
	// 	KeyPath:    "../data/wallet/UTC--2024-05-01T15-39-51.163922459Z--14655f7c9783aaa74990ab1704504032f2bd38ed",
	// 	Passphrase: "password1",
	// 	Address:    "0x14655F7c9783aAA74990aB1704504032f2bd38eD",
	// }

	key0, err := acc0.GetAccountKey()
	if err != nil {
		return fmt.Errorf("failed to get account 0 key: %w", err)
	}
	// key1, err := acc1.GetAccountKey()
	// if err != nil {
	// 	return fmt.Errorf("failed to get account 1 key: %w", err)
	// }

	cfg := config.GetDefault()
	log.Debug(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.EthClientURL)
	if err != nil {
		return fmt.Errorf("failed to dial eth client: %w", err)
	}
	defer client.Close()

	fmt.Println("---- Accouts ----")
	if err := balance.PrintEthBalance(ctx, client, key0); err != nil {
		return fmt.Errorf("failed to print balance of account 0: %w", err)
	}
	// if err := balance.PrintEthBalance(ctx, client, key1); err != nil {
	// 	return fmt.Errorf("failed to print balance of account 1: %w", err)
	// }

	chainID := big.NewInt(5777)
	var gasLimit uint64 = 6721975
	gasPrice := big.NewInt(20000000000)
	// gasPrice, err := client.SuggestGasPrice(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to suggest gas price: %w", err)
	// }

	nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(acc0.Address))
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	fmt.Println("---- Chain ----")
	fmt.Println("chainID", chainID)
	fmt.Println("gasLimit:", gasLimit, "wei")
	fmt.Println("gasPrice:", gasPrice, "wei")
	fmt.Println("nonce:", nonce)

	// #region Smart contract
	fmt.Println("---- Smart contract ----")
	auth, err := bind.NewKeyedTransactorWithChainID(key0.PrivateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create new keyed transactor: %w", err)
	}
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	address, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		return fmt.Errorf("failed to deploy todo: %w", err)
	}
	fmt.Println("address:", address.Hex())
	fmt.Println("tx:", tx.Hash().Hex())

	// contractHexAddress := "0xdB842622d5404D5c93ECdc2E9c2E4f1796CD173f"
	// contractAddress := common.HexToAddress(contractHexAddress)

	// t, err := todo.NewTodo(contractAddress, client)
	// if err != nil {
	// 	return fmt.Errorf("failed to create new todo: %w", err)
	// }

	// tx, err := t.Add(auth, "First task")
	// if err != nil {
	// 	return fmt.Errorf("failed to add task: %w", err)
	// }
	// fmt.Println(tx)

	// addressFrom := crypto.PubkeyToAddress(key0.PrivateKey.PublicKey)
	// tasks, err := t.List(&bind.CallOpts{
	// 	From: addressFrom,
	// })
	// if err != nil {
	// 	return fmt.Errorf("failed to list tasks: %w", err)
	// }
	// fmt.Println("tasks:", tasks)

	// #endregion

	// // #region Make transaction
	// var gasLimit uint64 = 30000000
	// gasPrice, err := client.SuggestGasPrice(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to suggest gas price: %w", err)
	// }

	// nonce, err := client.PendingNonceAt(ctx, common.HexToAddress(acc0.Address))
	// if err != nil {
	// 	return fmt.Errorf("failed to get nonce: %w", err)
	// }

	// fmt.Println("gasLimit:", gasLimit, "wei")
	// fmt.Println("gasPrice:", gasPrice, "wei")
	// fmt.Println("nonce:", nonce)

	// gasLimitBigInt := big.NewInt(int64(gasLimit))
	// totalCostWei := new(big.Int).Mul(gasPrice, gasLimitBigInt)
	// etherValue := new(big.Float).Quo(new(big.Float).SetInt(totalCostWei), big.NewFloat(1e18))
	// etherValueF, _ := etherValue.Float64()
	// fmt.Printf("transaction cost = gasPrice*gasLimit: %v ETH \n", etherValueF)

	// amount := big.NewInt(1e17) // 0.1 ETH
	// tx := types.NewTransaction(
	// 	nonce,
	// 	common.HexToAddress(acc1.Address),
	// 	amount,
	// 	gasLimit,
	// 	gasPrice,
	// 	nil,
	// )

	// // Sign transaction
	// chainID := big.NewInt(1337)
	// // chainID, err := client.NetworkID(ctx)
	// // if err != nil {
	// // 	return fmt.Errorf("failed to get network id: %w", err)
	// // }

	// tx, err = types.SignTx(
	// 	tx,
	// 	types.LatestSignerForChainID(chainID),
	// 	key0.PrivateKey,
	// )
	// if err != nil {
	// 	return fmt.Errorf("failed to sign transaction: %w", err)
	// }

	// // Send transaction
	// err = client.SendTransaction(ctx, tx)
	// if err != nil {
	// 	return fmt.Errorf("failed to send transaction: %w", err)
	// }
	// fmt.Printf("transaction hash: %s", tx.Hash().Hex())

	// // #endregion

	return nil
}
