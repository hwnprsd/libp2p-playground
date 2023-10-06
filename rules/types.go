package rules

import (
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"golang.org/x/exp/slices"
)

func GetRulesForSender(tx *proto.SolaceTx, sender common.Addr, rules []*proto.AccessControlRule) error {
	// This funciton assumes that the sender is already verified
	senderRules := utils.Filter(rules, func(rule *proto.AccessControlRule) bool {
		addrs, err := common.NewEthAddrSlice(rule.SenderGroup.Addresses)
		if err != nil {
			return false
		}
		return slices.Contains(addrs, sender)
	})

	// Split the available  rules into [RECIPIENT_LOCKED_RULES] and [VALUE RANGE LOCKED RULES] or ones with [BOTH]
	var (
		rcl  = make([]*proto.AccessControlRule, 0)
		vrcl = make([]*proto.AccessControlRule, 0)
		both = make([]*proto.AccessControlRule, 0)
	)

	for _, rule := range senderRules {
		if rule.TokenAddress != tx.TokenAddr {
			continue
		}

		if rule.RecipientAddress != "" && (rule.ValueRangeClause.MaxVal == 0 && rule.ValueRangeClause.MinVal == 0) {
			rcl = append(rcl, rule)
		} else if rule.RecipientAddress == "" && (rule.ValueRangeClause.MaxVal != 0 || rule.ValueRangeClause.MinVal != 0) {
			vrcl = append(vrcl, rule)
		} else {
			both = append(both, rule)
		}
	}

	if len(both) != 0 {
		// Check if the tx fits the rules
	} else if len(rcl) != 0 && len(vrcl) == 0 {
		// Check if the tx fit the Recipient based rules
	} else if len(rcl) == 0 && len(vrcl) != 0 {
		// Check if the tx fit the value range based rules
	} else {
		// No rules apply - Reject Tx
	}

	// Check which sender rule applies
	return nil
}
