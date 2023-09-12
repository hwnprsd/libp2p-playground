package proto

import (
	"encoding/hex"

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
