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
	return []byte(fmt.Sprintf("%s%s%d", sc.Sender, sc.TokenAddress, sc.Cap))
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
