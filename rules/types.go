package rules

import (
	"fmt"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"golang.org/x/exp/slices"
)

type ACL []*proto.AccessControlRule
type ACLRule *proto.AccessControlRule

const (
	RECIPIENT_WILD_CARD       = "0x00"
	errRecipientAddrViolation = "[1] Tx Recipient Addr doesn't match the rule"
	errValueRangeViolation    = "[2] Tx Value doesn't match the Value Range in the rule"
	errEscalationViolation    = "[3] Tx doesn't have signatures required by the Escalation Clause"
	errNoRulesFound           = "[4] No rules found for <Sender, Token> combo"
)

func ValidateTx(tx *proto.SolaceTx, sender common.Addr, rules ACL) error {
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
		rule := both[len(both)-1]

		if rule.RecipientAddress != RECIPIENT_WILD_CARD && rule.RecipientAddress != tx.ToAddr {
			return fmt.Errorf(errRecipientAddrViolation)
		}

		isValid, rule := applyValueRangeClause(both, tx)
		if isValid {
			isValid = applyEscalationClause(rule, tx)
			if !isValid {
				return fmt.Errorf(errEscalationViolation)
			}
			return nil
		}
		return fmt.Errorf(errValueRangeViolation)

	} else if len(rcl) != 0 && len(vrcl) == 0 {
		// Only apply the last rule
		rule := rcl[len(rcl)-1]
		if rule.RecipientAddress != RECIPIENT_WILD_CARD && rule.RecipientAddress != tx.ToAddr {
			return fmt.Errorf(errRecipientAddrViolation)
		}

		if rule.EscalationClause == nil {
			// Rule passes
			return nil
		} else {
			isValid := applyEscalationClause(rule, tx)
			if !isValid {
				return fmt.Errorf(errEscalationViolation)
			}
		}
		// Check if escalation rules exist
	} else if len(rcl) == 0 && len(vrcl) != 0 {
		// go over every rule which is applicable, and check if the value matches
		isValid, _ := applyValueRangeClause(vrcl, tx)
		if !isValid {
			return fmt.Errorf(errValueRangeViolation)
		} else {
			return nil
		}
	} else {
		return fmt.Errorf(errNoRulesFound)
	}

	// Check which sender rule applies
	return nil
}

func applyValueRangeClause(rules ACL, tx *proto.SolaceTx) (bool, ACLRule) {
	for _, rule := range rules {
		if tx.Value < int32(rule.ValueRangeClause.MaxVal) && tx.Value > int32(rule.ValueRangeClause.MinVal) {
			if rule.EscalationClause != nil {
				isPassing := applyEscalationClause(rule, tx)
				fmt.Println(errEscalationViolation)
				return isPassing, rule
			} else {
				return true, rule
			}
		}
	}
	return false, nil
}

func applyEscalationClause(rules ACLRule, tx *proto.SolaceTx) bool {
	return true
}
