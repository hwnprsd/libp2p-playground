package squad

import (
	"context"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"libp2p-playground/common"
	"libp2p-playground/smart_contract"
)

type Squad struct {
	isInitialized bool
	acl           map[peer.ID]bool
	peerId        peer.ID
	sc            smartcontract.NetworkState
	ctx           context.Context
	writeCh       chan<- common.Message
}

func (cps *Squad) VerifyPeer(peerID peer.ID) bool {
	// Your custom logic here
	return cps.acl[peerID]
}

func (cps *Squad) validator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	if cps.VerifyPeer(peer) {
		return pubsub.ValidationAccept
	}
	return pubsub.ValidationReject
}

func NewSquad(peerId peer.ID) *Squad {
	return &Squad{
		peerId:        peerId,
		acl:           make(map[peer.ID]bool),
		isInitialized: false,
		// DKG and Key Info
	}
}

func (s *Squad) Init(ctx context.Context, sc smartcontract.NetworkState, squadId string, writeCh chan<- common.Message) {
	s.sc = sc
	peers, err := s.sc.GetPeerList(squadId)
	if err != nil {
		panic(err)
	}
	for _, peer := range peers {
		s.acl[peer] = true
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
