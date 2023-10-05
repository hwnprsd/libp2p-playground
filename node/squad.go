package node

import (
	"context"
	"log"

	"github.com/solace-labs/skeyn/common"
	smartcontract "github.com/solace-labs/skeyn/smart_contract"
	"github.com/solace-labs/skeyn/squad"
)

// Create the squad if not exists
func (n *Node) SetupSquad(ctx context.Context, walletAddress common.Addr) {
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
		// TODO:  Replace this with actual SCWs deployed on networks
		smartcontract.TestEvmScw,
	)

	n.squad[walletAddress] = sqd
}
