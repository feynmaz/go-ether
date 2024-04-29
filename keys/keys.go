package keys

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrintPrivateKeyData(privateKey *ecdsa.PrivateKey) {
	privateKeyData := crypto.FromECDSA(privateKey)
	fmt.Println("private key:", hexutil.Encode(privateKeyData))

	publicKeyData := crypto.FromECDSAPub(&privateKey.PublicKey)
	fmt.Println("public key:", hexutil.Encode(publicKeyData))

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("address:", address)
}
