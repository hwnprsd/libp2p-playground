package smartcontract

import "github.com/solace-labs/skeyn/common"

type SmartContractWallet interface {
	GetOwner() (common.WalletAddress, error)
}
