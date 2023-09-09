package squad

import (
	"context"
	"log"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/smart_contract"
)

type SquadPeers map[peer.ID]bool

type Squad struct {
	isInitialized bool
	peers         SquadPeers
	peerId        peer.ID
	sc            smartcontract.NetworkState
	ctx           context.Context
	writeCh       chan<- common.Message
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

func (s SquadPeers) List() []peer.ID {
	keys := make([]peer.ID, 0, len(s))
	for i := range s {
		keys = append(keys, i)
	}
	return keys
}

func (s *Squad) Init(ctx context.Context, sc smartcontract.NetworkState, squadId string, writeCh chan<- common.Message) {
	s.sc = sc
	peers, err := s.sc.GetPeerList(squadId)
	if err != nil {
		panic(err)
	}
	for _, peer := range peers {
		s.peers[peer] = true
	}

	log.Println(squadId, "- Squad Initialized")

	s.isInitialized = true
	s.ctx = ctx
	s.writeCh = writeCh
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
