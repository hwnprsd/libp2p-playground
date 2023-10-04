package proto

import (
	"encoding/hex"
	"fmt"

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
	return []byte(fmt.Sprintf("%s%s%s%d", sc.Namespace, sc.Sender, sc.TokenAddress, sc.Cap))
}
