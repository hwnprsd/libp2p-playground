package squad

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func (s *Squad) Broadcast(message []byte) {
	s.Publish(s.id, message)
}

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
