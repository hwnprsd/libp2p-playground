package smartcontract

import (
	"github.com/solace-labs/skeyn/common"
)

type SmartContractWallet interface {
	validateSetup(signature []byte) (bool, error)
	validateRuleAddition(rule []byte, signature []byte, sender common.Addr) (bool, error)
	validateRuleDeletion(rule []byte, signature []byte, sender common.Addr) (bool, error)
}
