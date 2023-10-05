package rules

import (
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"golang.org/x/exp/slices"
)

func GetRulesForSender(tx *proto.SolaceTx, sender common.Addr, rules []*proto.AccessControlRule) []proto.AccessControlRule {
	// This funciton assumes that the sender is already verified
	senderRules := utils.Filter(rules, func(rule *proto.AccessControlRule) bool {
		addrs, err := common.NewEthAddrSlice(rule.SenderGroup.Addresses)
		if err != nil {
			return false
		}
		return slices.Contains(addrs, sender)
	})

	// Check which sender rule applies
	return nil
}
