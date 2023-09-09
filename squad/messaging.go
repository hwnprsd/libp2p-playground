package squad

import (
	"errors"
	"github.com/solace-labs/skeyn/common"
	"log"

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

func (s *Squad) VerifyMessage(msg common.Message) error {
	if !s.peers[msg.GetPeerID()] {
		return errors.New("Invalid sender")
	}
	return nil
}

// INCOMING MESSAGES
func (s *Squad) HandleIncomingMessages(ch <-chan common.Message) {
	go func() {
		for msg := range ch {
			err := s.VerifyMessage(msg)
			if err != nil {
				log.Println("Message from invalid sender!")
				continue
			}
			log.Println(string(msg.GetData()))
		}
	}()
}

func (s *Squad) Broadcast(protocol protocol.ID, data []byte) {
	for peer := range s.peers {
		s.writeCh <- common.NodeMessage{PeerID: peer, Data: data, Protocol: protocol}
	}
	log.Println("Broadcast data complete")
}
