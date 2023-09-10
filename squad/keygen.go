package squad

import (
	"context"
	"log"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	protoc "google.golang.org/protobuf/proto"
)

func (s *Squad) InitKeygen(ctx context.Context) (*chan tss.Message, *chan error) {
	shouldContinueInit, outChan, errChan := s.setupKeygenParty(ctx)
	if !shouldContinueInit {
		return nil, nil
	}
	go func() {
		err := (*s.keyGenParty).Start()
		if err != nil {
			log.Println("ERR", err)
		}
	}()
	return outChan, errChan
}

// Should continue init
func (s *Squad) setupKeygenParty(ctx context.Context) (shouldContinueInit bool, outChan *chan tss.Message, errChan *chan error) {
	// Keygen Party exists for this session
	if s.keyGenParty != nil {
		return false, nil, nil
	}

	parties := s.peers.SortedPartyIDs()

	peerCtx := tss.NewPeerContext(parties)

	params := tss.NewParameters(tss.S256(), peerCtx, s.PartyID(), len(parties), len(parties)-1)

	outChan, errChan = s.setupChannels()
	endChan := make(chan keygen.LocalPartySaveData)
	party := keygen.NewLocalParty(params, *outChan, endChan, *s.preParams)
	s.keyGenParty = &party

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case outData := <-*outChan:
				s.handleKeygenMessage(outData)
			case endData := <-endChan:
				s.handleKeygenEnd(endData)
			}
		}
	}()

	return true, outChan, errChan
}

func (s *Squad) handleKeygenEnd(data keygen.LocalPartySaveData) {
	s.keyGenData = &data
	log.Println("Keygen Complete")
	// x, y := data.ECDSAPub.X(), data.ECDSAPub.Y()
	// pk := ecdsa.PublicKey{
	// 	Curve: tss.EC(),
	// 	X:     x,
	// 	Y:     y,
	// }
	// pubKeyBytes := elliptic.Marshal(pk.Curve, pk.X, pk.Y)
	// n.logger.Sugar().Infof("Session - %s", sAddress)
	// n.logger.Sugar().Infof("Public Key - %s", hex.EncodeToString(pubKeyBytes))
}

func (s *Squad) UpdateKeygenParty(
	ctx context.Context,
	message UpdateMessage,
	peerId peer.ID,
) (*chan tss.Message, *chan error, error) {
	outChan, errChan := s.InitKeygen(ctx)
	fromPartyId := ToPartyID(&peerId)

	_, err := (*s.keyGenParty).UpdateFromBytes(message.GetWireMessage(), fromPartyId, message.GetIsBroadcast())
	if err != nil {
		return nil, nil, err
	}
	return outChan, errChan, nil
}

// Messages coming in from the TSS-Lib channels
func (s *Squad) handleKeygenMessage(message tss.Message) {
	// n.logger.Sugar().Infof("[KEYGEN] Received a message from outChan: %+v", message)
	dest := message.GetTo()
	outMsg, err := protoc.Marshal(&proto.UpdateMessage{
		WireMessage: message.WireMsg().Message.Value,
		IsBroadcast: dest == nil,
		Payload:     []byte(""),
	})
	if err != nil {
		log.Println("{ERR} - Error serializing message to protobuf obj")
	}

	if dest == nil {
		// Broadcast
		s.Broadcast(common.DKG_PROTOCOL, outMsg)
	} else {
		s.SendTo(*ToPeerID(dest[0]), common.DKG_PROTOCOL, outMsg)
	}
}

func (n *Squad) setupChannels() (*chan tss.Message, *chan error) {
	outChan := make(chan tss.Message)
	errChan := make(chan error)
	return &outChan, &errChan
}
