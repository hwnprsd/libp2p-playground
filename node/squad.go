package node

import (
	"context"

	"github.com/solace-labs/skeyn/squad"
)

// Create the squad if not exists
func (n *Node) SetupSquad(ctx context.Context, walletAddress string) {
	_, exists := n.squad[walletAddress]
	squadId, err := n.smartContract.GetSquadID(n.PeerID())
	if err != nil {
		panic(err)
	}

	outChan := n.setupOutgoingMessageHandler(ctx)

	if !exists {
		sqd := squad.NewSquad(n.PeerID())
		peerStore := n.h().Peerstore()
		sqd.Init(
			ctx,
			n.smartContract,
			squadId,
			outChan,
			peerStore,
		)
	}
}
