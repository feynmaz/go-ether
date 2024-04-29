package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	passphrase := "password"
	account, err := key.NewAccount(passphrase)
	if err != nil {
		return fmt.Errorf("failed to create new account: %w", err)
	}
	fmt.Println("new account address:", account.Address)

	return nil
}
