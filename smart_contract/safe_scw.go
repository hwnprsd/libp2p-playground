package smartcontract

import "github.com/solace-labs/skeyn/common"

type SafeScw struct {
	chainID    int
	walletAddr common.WalletAddress
}

func NewSafeScw() SafeScw { return SafeScw{} }
