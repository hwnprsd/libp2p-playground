package node

import (
	"context"
	smartcontract "github.com/solace-labs/skeyn/smart_contract"
	"github.com/solace-labs/skeyn/squad"
)

func (n *Node) SetupSquads(ctx context.Context) {
	// TODO: Make it live in prod
	n.smartContract = &smartcontract.TestContract{}

	squadId, err := n.smartContract.GetSquadID(n.PeerID())
	if err != nil {
		panic(err)
	}

	sqd := squad.NewSquad(n.PeerID())

	ps := n.h().Peerstore()

	sqd.Init(ctx, n.smartContract, squadId, n.setupOutgoingMessageHandler(ctx), ps)

	n.squad = sqd
}
