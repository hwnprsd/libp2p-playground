package squad

import (
	"context"
	"log"

	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/smart_contract"
)

type Squad struct {
	isInitialized bool
	peers         SquadPeers
	peerId        peer.ID
	sc            smartcontract.NetworkState
	ctx           context.Context
	writeCh       chan<- common.OutgoingMessage

	peerStore peerstore.Peerstore

	preParams   *keygen.LocalPreParams
	keyGenParty *tss.Party
	keyGenData  *keygen.LocalPartySaveData

	sigParty *tss.Party
}

func (cps *Squad) VerifyPeer(peerID peer.ID) bool {
	// Your custom logic here
	return cps.peers[peerID]
}

func NewSquad(peerId peer.ID) *Squad {
	return &Squad{
		peerId:        peerId,
		peers:         make(SquadPeers),
		isInitialized: false,
		// DKG and Key Info
	}
}

func (s *Squad) Init(ctx context.Context,
	sc smartcontract.NetworkState,
	squadId string,
	writeCh chan<- common.OutgoingMessage,
	peerStore peerstore.Peerstore,
) {
	s.sc = sc
	peers, err := s.sc.GetPeerList(squadId)
	if err != nil {
		panic(err)
	}
	// How to check if these peers are connected or not
	// Some communication is required from the node
	for _, peer := range peers {
		s.peers[peer] = true
	}

	log.Println(squadId, "- Squad Initialized")

	s.isInitialized = true
	s.ctx = ctx
	s.writeCh = writeCh
	s.peerStore = peerStore
}

func (s Squad) RefreshACL(ctx context.Context) {
	// Track last refresh?
}

func (s Squad) MakeACLVerificationMessage() {

}

func (s Squad) VerifyACL() {}

func (s Squad) Publish(message []byte) error {
	log.Println("Message Published")
	return nil
}
