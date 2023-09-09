package node

import (
	"context"
	smartcontract "libp2p-playground/smart_contract"
	"libp2p-playground/squad"
)

func (n *Node) SetupSquads(ctx context.Context) {
	// TODO: Make it live in prod
	n.smartContract = &smartcontract.TestContract{}

	squadId, err := n.smartContract.GetSquadID(n.PeerID())
	if err != nil {
		panic(err)
	}

	sqd := squad.NewSquad(n.PeerID())

	sqd.Init(ctx, n.smartContract, squadId, n.setupOutgoingMessageHandler(ctx))

	n.squad = sqd
}
