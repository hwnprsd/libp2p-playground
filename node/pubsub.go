package node

import (
	"context"
	"libp2p-playground/squad"

	"github.com/libp2p/go-libp2p-pubsub"
)

func (n *Node) SetupPubSub(ctx context.Context) {
	ps, err := pubsub.NewGossipSub(ctx, n.h())
	if err != nil {
		panic(err)
	}
	n.ps = ps
	sqd := squad.NewSquad(n.PeerID(), ps)
	sqd.Init(&squad.TestContract{})
}
