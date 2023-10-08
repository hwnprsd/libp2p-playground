package proto

import (
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
)

func (r *Payload) ID() string {
	b, _ := proto.Marshal(r)
	return hex.EncodeToString(b)
}

func PayloadFromID(ID string) []byte {
	protoBytes, err := hex.DecodeString(ID)
	if err != nil {
		return nil
	}
	return protoBytes
}

func (sc *SpendingCap) Bytes() []byte {
	return []byte(fmt.Sprintf("%s%s", sc.Sender, sc.TokenAddress, sc.Cap))
}

// Use IDs to prevent rule collusions while storing new rules
func (rule *AccessControlRule) Ids() []string {
	ids := make([]string, 0)
	isRecipientLocked := rule.RecipientAddress != ""
	isValueRangeLocked := rule.ValueRangeClause != nil && rule.ValueRangeClause.MaxVal != 0 && rule.ValueRangeClause.MinVal != 0
	if isRecipientLocked {
		for _, sender := range rule.SenderGroup.Addresses {
			var id []byte
			id = append(id, []byte(sender)...)
			id = append(id, []byte(rule.TokenAddress)...)
			id = append(id, []byte(rule.RecipientAddress)...)
			ids = append(ids, string(id))
		}
	}
	if isValueRangeLocked {
		for _, sender := range rule.SenderGroup.Addresses {
			var id []byte
			id = append(id, []byte(sender)...)
			id = append(id, []byte(rule.TokenAddress)...)
			id = append(id, []byte(fmt.Sprintf("%d", rule.ValueRangeClause.MinVal))...)
			id = append(id, []byte(fmt.Sprintf("%d", rule.ValueRangeClause.MaxVal))...)
			ids = append(ids, string(id))
		}
	}
	return ids
}

func (rule *AccessControlRule) Bytes() []byte {
	var res []byte
	res = append(res, rule.SenderGroup.Bytes()...)
	res = append(res, rule.ValueRangeClause.Bytes()...)
	res = append(res, rule.TimeWindowClause.Bytes()...)
	res = append(res, rule.EscalationClause.Bytes()...)
	return res
}

func (clause *SenderGroup) Bytes() []byte {
	addresses := clause.GetAddresses()
	slices.Sort(addresses)
	return []byte(strings.Join(addresses, "+"))
}

func (clause *ValueRangeClause) Bytes() []byte {
	return []byte(fmt.Sprintf("%d->%d", clause.MinVal, clause.MaxVal))
}

func (clause *TimeWindowClause) Bytes() []byte {
	return []byte("TBD")
}

func (clause *EscalationClause) Bytes() []byte {
	return []byte(strings.Join(clause.Addresses, "+") + clause.Logic)
}
