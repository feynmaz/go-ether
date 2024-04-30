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
