package rules

import "github.com/solace-labs/skeyn/proto"

type ACL []*proto.AccessControlRule
type ACLRule *proto.AccessControlRule

const (
	RECIPIENT_WILD_CARD       = "0x00"
	errRecipientAddrViolation = "[1] Tx Recipient Addr doesn't match the rule"
	errValueRangeViolation    = "[2] Tx Value doesn't match the Value Range in the rule"
	errEscalationViolation    = "[3] Tx doesn't have signatures required by the Escalation Clause"
	errNoRulesFound           = "[4] No rules found for <Sender, Token> <%s, %s>"
)
