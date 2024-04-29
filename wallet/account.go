package wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func CreateAccount(key *keystore.KeyStore, passphrase string) (accounts.Account, error) {
	account, err := key.NewAccount(passphrase)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("failed to create new account: %w", err)
	}
	fmt.Println("new account address:", account.Address)

	return account, nil
}

func GetAddress(hex string) common.Address {
	return common.HexToAddress(hex)
}