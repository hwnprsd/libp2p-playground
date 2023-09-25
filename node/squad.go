package node

import (
	"context"
	"log"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/squad"
)

// Create the squad if not exists
func (n *Node) SetupSquad(ctx context.Context, walletAddress common.WalletAddress) {
	_, exists := n.squad[walletAddress]

	if exists {
		return
	}

	log.Println("Setting up Squad")

	sqd := squad.NewSquad(n.PeerID())

	outChan := n.setupOutgoingMessageHandler(ctx, walletAddress)
	peerStore := n.h().Peerstore()
	sqd.Init(
		ctx,
		n.smartContract,
		walletAddress.String(),
		outChan,
		peerStore,
	)

	n.squad[walletAddress] = sqd
}
