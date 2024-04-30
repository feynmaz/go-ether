package balance

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PrintEthBalance(ctx context.Context, client *ethclient.Client, key *keystore.Key) error {
	balanceWei, err := client.BalanceAt(ctx, key.Address, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	// 1 ether = 10^18 wei

	balanceEth := new(big.Float)
	balanceEth.SetString(balanceWei.String())
	balanceEth = new(big.Float).Quo(balanceEth, big.NewFloat(math.Pow10(18)))

	fmt.Printf("%s balance: %v ETH \n", key.Address, balanceEth)
	return nil
}
