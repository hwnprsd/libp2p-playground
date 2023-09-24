package node

import (
	"context"
	"log"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/utils"

	"github.com/libp2p/go-libp2p/core/network"
)

// Message sink
func (n *Node) setupMessageRecieverHandler(ctx context.Context) <-chan common.IncomingMessage {
	ch := make(chan common.IncomingMessage)
	go func() {
		// Setup recievers
		n.h().SetStreamHandler(common.DKG_PROTOCOL, func(s network.Stream) {
			data, err := utils.ReadStream(s)
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.DKG_PROTOCOL,
			}
			ch <- nodeMessage
		})

		n.h().SetStreamHandler(common.SIGNING_PROTOCOL, func(s network.Stream) {
			data, err := utils.ReadStream(s)
			if err != nil {
				log.Println("Error Reading Stream Data")
				log.Println(err)
			}

			nodeMessage := common.NodeMessage{
				PeerID:   s.Conn().RemotePeer(),
				Data:     data,
				Protocol: common.SIGNING_PROTOCOL,
			}
			ch <- nodeMessage
		})

	}()
	return ch
}

// Message Export Hub for the node
func (n *Node) setupOutgoingMessageHandler(ctx context.Context) chan common.OutgoingMessage {
	ch := make(chan common.OutgoingMessage)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-ch:
				stream, err := n.h().NewStream(ctx, msg.GetPeerID(), msg.GetProtocol())
				if err != nil {
					log.Println("Error creating stream", err)
					continue
				}

				if err := utils.WriteStream(stream, msg.GetData()); err != nil {
					log.Println("Error writing to stream", err)
					continue
				}
			}
		}
	}()
	return ch
}
