package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Account struct {
	KeyPath    string
	Passphrase string
	Address    string
}

func CreateAccount(keyStore *keystore.KeyStore, passphrase string) (*Account, error) {
	acc, err := keyStore.NewAccount(passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to create new account in keystore: %w", err)
	}

	return &Account{
		KeyPath:    acc.URL.Path,
		Passphrase: passphrase,
		Address:    acc.Address.Hex(),
	}, err

}

func (a *Account) GetAccountKey() (*keystore.Key, error) {
	b, err := os.ReadFile(a.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	key, err := keystore.DecryptKey(b, a.Passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %w", err)
	}

	return key, nil
}
