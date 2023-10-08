package rules

import (
	"fmt"

	"github.com/solace-labs/skeyn/proto"
)

type ACL []*proto.AccessControlRule
type ACLRule *proto.AccessControlRule

const (
	RECIPIENT_WILD_CARD       = "0x00"
	errRecipientAddrViolation = "[1] Tx Recipient Addr doesn't match the rule"
	errValueRangeViolation    = "[2] Tx Value doesn't match the Value Range in the rule"
	errEscalationViolation    = "[3] Tx doesn't have signatures required by the Escalation Clause"
	errNoRulesFound           = "[4] No rules found for <Sender, Token> <%s, %s>"
)

// Use IDs to prevent rule collusions while storing new rules
func getRuleIds(rule ACLRule) []string {
	ids := make([]string, 0)
	isRecipientLocked := rule.RecipientAddr != ""
	isValueRangeLocked := rule.ValueRangeClause != nil && rule.ValueRangeClause.MaxVal != 0 && rule.ValueRangeClause.MinVal != 0
	if isRecipientLocked {
		for _, sender := range rule.SenderGroup.Addresses {
			var id []byte
			id = append(id, []byte(sender)...)
			id = append(id, []byte(rule.TokenAddr)...)
			id = append(id, []byte(rule.RecipientAddr)...)
			ids = append(ids, string(id))
		}
	}
	if isValueRangeLocked {
		for _, sender := range rule.SenderGroup.Addresses {
			var id []byte
			id = append(id, []byte(sender)...)
			id = append(id, []byte(rule.TokenAddr)...)
			id = append(id, []byte(fmt.Sprintf("%d", rule.ValueRangeClause.MinVal))...)
			id = append(id, []byte(fmt.Sprintf("%d", rule.ValueRangeClause.MaxVal))...)
			ids = append(ids, string(id))
		}
	}
	return ids
}
