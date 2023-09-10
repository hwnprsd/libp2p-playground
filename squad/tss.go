package squad

import "github.com/libp2p/go-libp2p/core/peer"

type UpdateMessage interface {
	GetWireMessage() []byte
	GetIsBroadcast() bool
	GetSigMessage() []byte
	GetPeerID() *peer.ID
}
