package smartcontract

import "github.com/solace-labs/skeyn/proto"

type SmartContractWallet interface {
	validateRule(proto.Rule) (bool, error)
}
