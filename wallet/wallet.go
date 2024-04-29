package wallet

import "github.com/ethereum/go-ethereum/accounts/keystore"

func CreateWallet(walletPath string) *keystore.KeyStore {
	return keystore.NewKeyStore(
		walletPath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
}
