package squad

import (
	"libp2p-playground/common"
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

// INCOMING MESSAGES
func (s *Squad) HandleIncomingMessages(ch <-chan []byte) error {
	go func() {
		for data := range ch {
			log.Println(string(data))
		}
	}()
	return nil
}

func (s *Squad) Broadcast(protocol protocol.ID, data []byte) {
	for peer := range s.acl {
		s.writeCh <- common.NodeMessage{PeerID: peer, Data: data, Protocol: protocol}
	}
	log.Println("Broadcast data complete")
}
