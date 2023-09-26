package common

import "github.com/libp2p/go-libp2p/core/protocol"

var (
	DKG_PROTOCOL       = protocol.ID("/dkg")
	SIGNING_PROTOCOL   = protocol.ID("/sign")
	RESHARING_PROTOCOL = protocol.ID("/resharing")

	CREATE_RULE = protocol.ID("/create-rule")
)
