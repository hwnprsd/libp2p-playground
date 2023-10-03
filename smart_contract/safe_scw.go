package smartcontract

import "github.com/solace-labs/skeyn/common"

type SafeScw struct {
	chainID    int
	walletAddr common.Addr
}

func NewSafeScw() SafeScw { return SafeScw{} }
