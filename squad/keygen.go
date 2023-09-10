package squad

import (
	"context"
	"log"
	"time"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	protoc "google.golang.org/protobuf/proto"
)

func (s *Squad) InitKeygen(ctx context.Context) chan error {
	log.Println("Initing Keygen")
	shouldContinueInit, errChan := s.setupKeygenParty(ctx)
	if !shouldContinueInit {
		return nil
	}

	go s.startKeygen()
	return errChan
}

func (s *Squad) startKeygen() {
	err := (*s.keyGenParty).Start()
	if err != nil {
		log.Println("ERR", err)
	}
}

// Should continue init
func (s *Squad) setupKeygenParty(ctx context.Context) (shouldContinueInit bool, errChan chan error) {
	// Keygen Party exists for this session
	if s.keyGenParty != nil {
		return false, nil
	}

	parties := s.SortedPartyIDs()

	peerCtx := tss.NewPeerContext(parties)

	params := tss.NewParameters(tss.S256(), peerCtx, s.PartyID(), len(parties), len(parties)-1)

	errChan = make(chan error)
	outChan := make(chan tss.Message)
	endChan := make(chan keygen.LocalPartySaveData)
	preParams, err := keygen.GeneratePreParams(1 * time.Minute)
	if err != nil {
		log.Println("Error generating pre-params")
		panic(err)
	}
	party := keygen.NewLocalParty(params, outChan, endChan, *preParams)
	s.keyGenParty = &party
	s.preParams = preParams

	go func() {
		for {
			select {
			// case <-ctx.Done():
			// 	return
			case outData := <-outChan:
				s.handleKeygenMessage(outData)
			case endData := <-endChan:
				s.handleKeygenEnd(endData)
			}
		}

	}()

	return true, errChan
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
) (chan error, error) {
	errChan := s.InitKeygen(ctx)
	fromPartyId := s.GetSortedPartyID(&peerId)

	_, err := (*s.keyGenParty).UpdateFromBytes(message.GetWireMessage(), fromPartyId, message.GetIsBroadcast())
	if err != nil {
		return nil, err
	}
	return errChan, nil
}

// Messages coming in from the TSS-Lib channels
func (s *Squad) handleKeygenMessage(message tss.Message) {
	log.Printf("[KEYGEN] Received a message from outChan: %+v\n", message)
	dest := message.GetTo()
	wireBytes, _, _ := message.WireBytes()
	outMsg, err := protoc.Marshal(&proto.UpdateMessage{
		WireMessage: wireBytes,
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
		toPeerId := s.ToPeerID(dest[0])
		if toPeerId == nil {
			panic("Unable to reconstruct Peer ID from TSS Party ID")
		}
		s.SendTo(*toPeerId, common.DKG_PROTOCOL, outMsg)
	}
}
