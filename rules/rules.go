package rules

import (
	"fmt"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
	"golang.org/x/exp/slices"
)

func ValidateTx(tx *proto.SolaceTx, verifiedSender common.Addr, rules ACL) (ACLRule, error) {
	// This funciton assumes that the sender is already verified
	senderRules := utils.Filter(rules, func(rule *proto.AccessControlRule) bool {
		addrs, err := common.NewEthAddrSlice(rule.SenderGroup.Addresses)
		if err != nil {
			return false
		}
		return slices.Contains(addrs, verifiedSender)
	})

	// Split the available  rules into [RECIPIENT_LOCKED_RULES] and [VALUE RANGE LOCKED RULES] or ones with [BOTH]
	var (
		rcl  = make(ACL, 0)
		vrcl = make(ACL, 0)
		both = make(ACL, 0)
	)

	for _, rule := range senderRules {
		if rule.TokenAddr != tx.TokenAddr {
			continue
		}

		if rule.RecipientAddr != "" && rule.ValueRangeClause == nil {
			rcl = append(rcl, rule)
		} else if rule.RecipientAddr == "" && (rule.ValueRangeClause != nil && (rule.ValueRangeClause.MaxVal != 0 || rule.ValueRangeClause.MinVal != 0)) {
			vrcl = append(vrcl, rule)
		} else {
			both = append(both, rule)
		}
	}

	// fmt.Println("RC LEN", len(rcl), len(vrcl), len(both))

	if len(both) != 0 {
		rule := both[len(both)-1]

		if rule.RecipientAddr != RECIPIENT_WILD_CARD && rule.RecipientAddr != tx.ToAddr {
			return nil, fmt.Errorf(errRecipientAddrViolation)
		}

		isValid, rule := applyValueRangeClause(both, tx)
		if isValid {
			isValid = applyEscalationClause(rule, tx)
			if !isValid {
				return nil, fmt.Errorf(errEscalationViolation)
			}
			return rule, nil
		}
		return nil, fmt.Errorf(errValueRangeViolation)

	} else if len(rcl) != 0 && len(vrcl) == 0 {
		fmt.Println("Applying RCL")
		// Only apply the last rule
		rule := rcl[len(rcl)-1]
		if rule.RecipientAddr != RECIPIENT_WILD_CARD && rule.RecipientAddr != tx.ToAddr {
			return nil, fmt.Errorf(errRecipientAddrViolation)
		}

		if rule.EscalationClause == nil {
			// Rule passes
			return rule, nil
		} else {
			isValid := applyEscalationClause(rule, tx)
			if !isValid {
				return nil, fmt.Errorf(errEscalationViolation)
			}
		}
		// Check if escalation rules exist
	} else if len(rcl) == 0 && len(vrcl) != 0 {
		// go over every rule which is applicable, and check if the value matches
		isValid, rule := applyValueRangeClause(vrcl, tx)
		if !isValid {
			return nil, fmt.Errorf(errValueRangeViolation)
		} else {
			return rule, nil
		}
	} else if len(rcl) != 0 && len(vrcl) != 0 {
		// A condition where both a value range clause is valid && recipient
		// If the recipient matches, apply the rule otherwise
		rclRule := rcl[len(rcl)-1]
		isValid := false
		if rclRule.RecipientAddr == tx.ToAddr {
			isValid = applyEscalationClause(rclRule, tx)
			if !isValid {
				return nil, fmt.Errorf(errEscalationViolation)
			} else {
				return rclRule, nil
			}
		} else {
			_isValid, rule := applyValueRangeClause(vrcl, tx)
			isValid = _isValid
			if isValid {
				isValid = applyEscalationClause(rule, tx)
				if isValid {
					return rule, nil
				} else {
					return nil, fmt.Errorf(errEscalationViolation)
				}
			} else {
				return nil, fmt.Errorf(errValueRangeViolation)
			}
		}
	} else {
		fmt.Println("HEER")
		return nil, fmt.Errorf(errNoRulesFound, verifiedSender.String(), tx.TokenAddr)
	}

	return nil, fmt.Errorf(errNoRulesFound, verifiedSender.String(), tx.ToAddr)
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

func applyEscalationClause(rule ACLRule, tx *proto.SolaceTx) bool {
	return true
}
