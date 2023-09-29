package squad

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	tsscommon "github.com/bnb-chain/tss-lib/v2/common"
	"github.com/bnb-chain/tss-lib/v2/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	protob "google.golang.org/protobuf/proto"
	protoc "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const TX_PREFIX = "SOLACETX####"

func (s *Squad) InitSigning(tx *proto.SolaceTx) (chan error, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	ctx := context.Background()

	shouldContinueInit, errChan := s.setupSigningParty(ctx, tx)
	if !shouldContinueInit {
		return nil, nil
	}

	// Only run it if not already inited
	err := s.validateTx(tx)
	if err != nil {
		log.Println("error validating Tx", err)
		s.cleanupSigning()
		return nil, err
	}

	go func() {
		err := (*s.sigParty).Start()
		log.Println("Starting sig process")
		if err != nil {
			log.Println("SIG_ERROR", err)
		}
	}()
	return errChan, nil
}

func (s *Squad) setupSigningParty(ctx context.Context, tx *proto.SolaceTx) (shouldContinueInit bool, errChan chan error) {
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
		if !saveData.Validate() {
			panic("SaveData is corrupt")
			// TODO: Handle corrupt save data
		}
	}

	// In an ongoing session. No need to init
	// Or node is in a broken state
	if s.sigParty != nil {
		return false, nil
	}

	msg := new(big.Int).SetBytes([]byte("DATA"))
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
				s.handleSigningMessage(outData, tx)
			case endData := <-endChan:
				// Find a way to parse a unique ID for the transaction
				s.handleSessionEnd(&endData, tx)
			}
		}
	}()

	party := signing.NewLocalParty(msg, params, *s.keyGenData.LocalPartySaveData, outChan, endChan)
	s.sigParty = &party
	return true, errChan
}

func (s *Squad) cleanupSigning() {
	s.sigParty = nil
}

func (s *Squad) GetSig(key []byte) ([]byte, error) {
	return s.db.Get(key)
}

func (s *Squad) GetTransactions() []*proto.Signature {
	index := s.getDbIndex()
	sigs := make([]*proto.Signature, 0)
	count := 0

	for i := 0; i < index.Int(); i++ {
		txB, err := s.db.Get(IndexFromInt(i))
		if err != nil {
			log.Println("Error fetching tx at index", i)
			continue
		}
		sig := &proto.Signature{}
		err = protob.Unmarshal(txB, sig)
		if err != nil {
			log.Println("[WARN] Error unmarshalling tx", err)
			continue
		}
		sigs = append(sigs, sig)
		count++
	}
	log.Println("Sig Len", count)
	// txSlice := s.db.GetAll(TX_PREFIX)
	// for _, txB := range txSlice {
	// 	sig := &proto.Signature{}
	// 	err := protob.Unmarshal(txB, sig)
	// 	if err != nil {
	// 		log.Println("[WARN] Error unmarshalling tx", err)
	// 		continue
	// 	}
	// 	sigs = append(sigs, sig)
	// }
	return sigs
}

func (s *Squad) UpdateSigningParty(
	ctx context.Context,
	message UpdateMessage,
	peerId peer.ID,
) (chan error, error) {

	tx := &proto.SolaceTx{}
	err := protob.Unmarshal(message.GetPayload(), tx)
	if err != nil {
		log.Println("error unmarshalling tx")
		return nil, err
	}

	message.GetPayload()
	// Ignoring the error
	errChan, _ := s.InitSigning(tx)
	fromPartyId := s.GetSortedPartyID(&peerId)

	_, err2 := (*s.sigParty).UpdateFromBytes(message.GetWireMessage(), fromPartyId, message.GetIsBroadcast())

	if err2 != nil {
		return nil, err
	}
	return errChan, nil
}

func (s *Squad) handleSigningMessage(message tss.Message, tx *proto.SolaceTx) {
	log.Printf("[SIGNING] Received a message from outChan: %+v\n", message)
	dest := message.GetTo()
	wireBytes, _, _ := message.WireBytes()

	txB, err := protob.Marshal(tx)
	if err != nil {
		log.Println("{ERR} - Error marshalling payload")
	}

	outMsg, err := protoc.Marshal(&proto.UpdateMessage{
		WireMessage: wireBytes,
		IsBroadcast: dest == nil,
		Payload:     txB,
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

func (s *Squad) handleSessionEnd(data *tsscommon.SignatureData, tx *proto.SolaceTx) {
	key, err := s.HashSolaceTx(tx)
	if err != nil {
		log.Println("error marshalling tx", err)
		s.cleanupSigning()
		return
	}

	keyHex := hexutil.Encode(key)

	val := &proto.Signature{
		Signature: hex.EncodeToString(data.Signature),
		Timestamp: timestamppb.Now(),
		Id:        keyHex,
		Tx:        tx,
	}

	valB, err := protob.Marshal(val)
	if err != nil {
		log.Println("Error marshalling signature value")
	}

	err = s.db.Set(key, valB)

	if err != nil {
		log.Println("Error setting signature")
	} else {
		log.Println("Sig Saved")
		log.Println(hex.EncodeToString(data.Signature))
	}

	index := s.getDbIndex()
	_ = s.db.Set(index.Bytes(), valB)
	_ = s.updateIndex()

	s.cleanupSigning()
}

func (s *Squad) GetSignature() {

}
