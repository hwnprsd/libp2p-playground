package squad

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	protoc "google.golang.org/protobuf/proto"
)

func (s *Squad) InitSigning(ctx context.Context, message []byte) chan error {
	log.Println("Initing Signing")
	shouldContinueInit, errChan := s.setupSigningParty(ctx, message)
	if !shouldContinueInit {
		return nil
	}
	go func() {
		err := (*s.sigParty).Start()
		log.Println("Starting to Sign")
		if err != nil {
			log.Println("SIG_ERROR", err)
		}
	}()
	return errChan
}

func (s *Squad) setupSigningParty(ctx context.Context, message []byte) (shouldContinueInit bool, errChan chan error) {
	// KeyGen is not completed
	if s.keyGenData.LocalPartySaveData == nil {
		// Check if it exists in the DB
		saveDataB, err := s.db.Get([]byte(s.LP_SAVE_DATA_KEY()))
		if err != nil {
			log.Println("Error reading SaveData from DB")
			panic(err)
		}
		if saveDataB == nil {
			log.Println("KeyGen SaveData does not exist")
			return false, nil
		}
		saveData := StoredSaveDataFromBytes(saveDataB)
		s.keyGenData = saveData
		log.Println("Stored Savedata fetched")
	}

	// In an ongoing session. No need to init
	// Or node is in a broken state
	if s.sigParty != nil {
		return false, nil
	}

	msg := new(big.Int).SetBytes(message)
	parties := s.SortedPartyIDs()
	peerCtx := tss.NewPeerContext(parties)
	params := tss.NewParameters(tss.S256(), peerCtx, s.PartyID(), len(parties), len(parties)-1)

	errChan = make(chan error)
	outChan := make(chan tss.Message)
	endChan := make(chan tsscommon.SignatureData)

	go func() {
		for {
			select {
			// case <-ctx.Done():
			// 	return
			case outData := <-outChan:
				s.handleSigningMessage(outData, message)
			case endData := <-endChan:
				s.handleSessionEnd(&endData)
			}
		}
	}()

	party := signing.NewLocalParty(msg, params, *s.keyGenData.LocalPartySaveData, outChan, endChan)
	s.sigParty = &party
	return true, errChan
}

func (s *Squad) handleSessionEnd(data *tsscommon.SignatureData) {
	log.Println(hex.EncodeToString(data.Signature))
	s.sigParty = nil
}

func (s *Squad) UpdateSigningParty(
	ctx context.Context,
	message UpdateMessage,
	peerId peer.ID,
) (chan error, error) {
	errChan := s.InitSigning(ctx, message.GetPayload())
	fromPartyId := s.GetSortedPartyID(&peerId)

	_, err := (*s.sigParty).UpdateFromBytes(message.GetWireMessage(), fromPartyId, message.GetIsBroadcast())
	if err != nil {
		return nil, err
	}
	return errChan, nil
}

func (s *Squad) handleSigningMessage(message tss.Message, signingData []byte) {
	log.Printf("[SIGNING] Received a message from outChan: %+v\n", message)
	dest := message.GetTo()
	wireBytes, _, _ := message.WireBytes()
	outMsg, err := protoc.Marshal(&proto.UpdateMessage{
		WireMessage: wireBytes,
		IsBroadcast: dest == nil,
		Payload:     signingData,
	})
	if err != nil {
		log.Println("{ERR} - Error serializing message to protobuf obj")
	}

	if dest == nil {
		s.Broadcast(common.SIGNING_PROTOCOL, outMsg)
	} else {
		toPeerId := s.ToPeerID(dest[0])
		if toPeerId == nil {
			panic("Unable to reconstruct Peer ID from TSS Party ID")
		}
		s.SendTo(*toPeerId, common.SIGNING_PROTOCOL, outMsg)
	}
}
