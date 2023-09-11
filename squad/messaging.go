package squad

import (
	"context"
	"errors"
	"log"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	protoc "google.golang.org/protobuf/proto"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// This will broadcast largely TSS Messages
// func (s *Squad) Broadcast(message *proto.Request) error {
// 	messageBytes, err := gproto.Marshal(message)
// 	if err != nil {
// 		return err
// 	}
// 	s.Publish(s.id, messageBytes)
// 	return nil
// }

func (s *Squad) VerifyMessage(msg common.IncomingMessage) error {
	if !s.peers[msg.GetPeerID()] {
		return errors.New("Invalid sender")
	}
	return nil
}

// Given a recieve-only channel, this funcion will recieve messages and handle them gracefully
func (s *Squad) HandleIncomingMessages(ctx context.Context, ch <-chan common.IncomingMessage) {
	log.Println("Handling incoming messages")
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-ch:
			err := s.VerifyMessage(msg)
			if err != nil {
				log.Println("Message from invalid sender!")
				continue
			}
			// Parse the data into a DKG Understandable message
			updateMsg := &proto.UpdateMessage{}
			if err := protoc.Unmarshal(msg.GetData(), updateMsg); err != nil {
				log.Println("Error unmarshalling Data from wire", err)
				continue
			}
			switch msg.GetProtocol() {
			case common.DKG_PROTOCOL:
				_, err = s.UpdateKeygenParty(ctx, updateMsg, msg.GetPeerID())
			case common.SIGNING_PROTOCOL:
				_, err = s.UpdateSigningParty(ctx, updateMsg, msg.GetPeerID())
			}
			if err != nil {
				log.Println("[ERR] Updating Keygen/Signing party", err)
			}
		}
	}
}

func (s *Squad) Broadcast(protocol protocol.ID, data []byte) {
	for peer := range s.peers {
		if peer == s.peerId {
			continue
		}
		s.writeCh <- common.NodeMessage{PeerID: peer, Data: data, Protocol: protocol}
	}
	log.Println("Broadcast data complete")
}

func (s *Squad) SendTo(peer peer.ID, protocol protocol.ID, data []byte) {
	s.writeCh <- common.NodeMessage{PeerID: peer, Data: data, Protocol: protocol}
	log.Println("Message sent to", peer)
}
