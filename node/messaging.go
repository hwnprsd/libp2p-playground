package node

import (
	"context"
	"log"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/utils"

	"github.com/libp2p/go-libp2p/core/network"
)

// Message sink
func (n *Node) setupMessageRecieverHandler() {
	go func() {
		// Setup recievers
		n.h().SetStreamHandler(common.DKG_PROTOCOL, func(s network.Stream) {
			log.Println("Message In From", s.Conn().RemotePeer())
			walletAddrBytes, data, err := utils.ReadStream(s)
			// TODO: Parse the wallet address out
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.DKG_PROTOCOL,
			}

			// Call the squad HandleIncomingMessage func
			walletAddress := common.NewWalletAddress(walletAddrBytes)
			n.SetupSquad(context.Background(), walletAddress)
			n.squad[walletAddress].HandleIncomingMessages(context.Background(), nodeMessage)
		})

		n.h().SetStreamHandler(common.SIGNING_PROTOCOL, func(s network.Stream) {
			walletAddrBytes, data, err := utils.ReadStream(s)
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.SIGNING_PROTOCOL,
			}

			walletAddress := common.NewWalletAddress(walletAddrBytes)
			n.SetupSquad(context.Background(), walletAddress)
			n.squad[walletAddress].HandleIncomingMessages(context.Background(), nodeMessage)
		})

	}()
}

// Message Export Hub for the node
func (n *Node) setupOutgoingMessageHandler(ctx context.Context, walletAddress common.WalletAddress) chan common.OutgoingMessage {
	ch := make(chan common.OutgoingMessage)
	go func() {
		for {
			select {
			case <-ctx.Done():
				// return
			case msg := <-ch:
				stream, err := n.h().NewStream(context.Background(), msg.GetPeerID() /* TO */, msg.GetProtocol())
				if err != nil {
					log.Println("Error creating stream", err)
					continue
				}

				if err := utils.WriteStream(stream, msg.GetData(), walletAddress.Bytes()); err != nil {
					log.Println("Error writing to stream", err)
					continue
				}
			}
		}
	}()
	return ch
}
