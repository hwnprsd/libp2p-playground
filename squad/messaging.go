package squad

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
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

// This will largely recieve TSS Messages
func (s *Squad) RecieveMessages(ch chan<- *pubsub.Message) error {
	sub, err := s.Subscribe(s.id)
	if err != nil {
		return err
	}
	go func() {
		for {
			msg, err := sub.Next(s.ctx)
			if err != nil {
				log.Println("Message Err: ", err)
			}
			ch <- msg
		}
	}()
	return nil
}
