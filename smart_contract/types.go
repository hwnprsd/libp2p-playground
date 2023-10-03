package smartcontract

import (
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
)

type SmartContractWallet interface {
	validateSetup(signature []byte) (bool, error)
	validateRuleAddition(rule *proto.Rule, signature []byte, sender common.Addr) (bool, error)
	validateRuleDeletion(rule *proto.Rule, signature []byte, sender common.Addr) (bool, error)
}
