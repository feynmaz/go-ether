package account

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type Account struct {
	KeyPath    string
	Passphrase string
}

func GetAccountKey(account Account) (*keystore.Key, error) {
	b, err := os.ReadFile(account.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	key, err := keystore.DecryptKey(b, account.Passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %w", err)
	}

	return key, nil
}
