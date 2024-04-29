package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	privateKeyData := crypto.FromECDSA(privateKey)
	fmt.Println("private key:", hexutil.Encode(privateKeyData))

	publicKeyData := crypto.FromECDSAPub(&privateKey.PublicKey)
	fmt.Println("public key:", hexutil.Encode(publicKeyData))

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("address:", address)
	
	return nil
}
