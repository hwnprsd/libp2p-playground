package common

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type Message interface {
	GetProtocol() protocol.ID
	GetData() []byte
	String() string
	GetPeerID() peer.ID
}

type IncomingMessage = Message

type OutgoingMessage = Message

type NodeMessage struct {
	Protocol protocol.ID
	PeerID   peer.ID
	Data     []byte
}

func (n NodeMessage) GetProtocol() protocol.ID {
	return n.Protocol
}

func (n NodeMessage) GetPeerID() peer.ID {
	return n.PeerID
}

func (n NodeMessage) GetData() []byte {
	return n.Data
}

func (n NodeMessage) String() string {
	return fmt.Sprintf("[%s] to %s // %s", n.GetProtocol(), n.GetPeerID(), string(n.GetData()))
}
