package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

var passphrase = "password"

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var err error

	// err = createWallet()
	// if err != nil {
	// 	return fmt.Errorf("failed to create wallet: %w", err)
	// }

	err = readWalletFile()
	if err != nil {
		return fmt.Errorf("failed to read wallet file: %w", err)
	}

	return nil
}

func createWallet() error {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := key.NewAccount(passphrase)
	if err != nil {
		return fmt.Errorf("failed to create new account: %w", err)
	}
	fmt.Println("new account address:", account.Address)

	return nil
}

func readWalletFile() error {
	b, err := os.ReadFile("./wallet/UTC--2024-04-29T15-47-46.521383016Z--0bba35d0cc610cff57fb0d5d8c97427892f948e5")
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	key, err := keystore.DecryptKey(b, passphrase)
	if err != nil {
		return fmt.Errorf("failed to decrypt key: %w", err)
	}

	privateKey := key.PrivateKey

	privateKeyData := crypto.FromECDSA(privateKey)
	fmt.Println("private key:", hexutil.Encode(privateKeyData))

	publicKeyData := crypto.FromECDSAPub(&privateKey.PublicKey)
	fmt.Println("public key:", hexutil.Encode(publicKeyData))

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("address:", address)

	return nil
}
