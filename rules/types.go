package rules

import (
	"fmt"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"golang.org/x/exp/slices"
)

type ACL []*proto.AccessControlRule

func GetRulesForSender(tx *proto.SolaceTx, sender common.Addr, rules ACL) (ACL, ACL, ACL) {
	// This funciton assumes that the sender is already verified
	senderRules := utils.Filter(rules, func(rule *proto.AccessControlRule) bool {
		addrs, err := common.NewEthAddrSlice(rule.SenderGroup.Addresses)
		fmt.Printf("%#v\n", addrs)
		if err != nil {
			return false
		}
		return slices.Contains(addrs, sender)
	})

	// Split the available  rules into [RECIPIENT_LOCKED_RULES] and [VALUE RANGE LOCKED RULES] or ones with [BOTH]
	var (
		rcl  = make(ACL, 0)
		vrcl = make(ACL, 0)
		both = make(ACL, 0)
	)

	fmt.Printf("%#v\n", senderRules)
	for _, rule := range senderRules {
		if rule.TokenAddress != tx.TokenAddr {
			continue
		}

		if rule.RecipientAddress != "" && rule.ValueRangeClause == nil {
			rcl = append(rcl, rule)
		} else if rule.RecipientAddress == "" && (rule.ValueRangeClause != nil && (rule.ValueRangeClause.MaxVal != 0 || rule.ValueRangeClause.MinVal != 0)) {
			vrcl = append(vrcl, rule)
		} else {
			both = append(both, rule)
		}
	}

	if len(both) != 0 {
		// Apply Rule
		// Check if the tx fits the rules
	} else if len(rcl) != 0 && len(vrcl) == 0 {
		// Apply Rule
		// Check if the tx fit the Recipient based rules
	} else if len(rcl) == 0 && len(vrcl) != 0 {
		// Apply Rule
		// Check if the tx fit the value range based rules
	} else {
		// No rules apply - Reject Tx
	}

	// Check which sender rule applies
	return rcl, vrcl, both
}
