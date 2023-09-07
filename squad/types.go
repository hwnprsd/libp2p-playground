package squad

import (
	"context"
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
)

type Squad struct {
	*pubsub.PubSub
	isInitialized bool
	acl           map[peer.ID]bool
	peerId        peer.ID
	id            string
	sc            SCInterface
	ctx           context.Context
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

func NewSquad(peerId peer.ID, pb *pubsub.PubSub) *Squad {
	return &Squad{
		PubSub:        pb,
		peerId:        peerId,
		acl:           make(map[peer.ID]bool),
		isInitialized: false,
		id:            "",
	}
}

func (s *Squad) Init(sc SCInterface) {
	s.sc = sc
	squadId := s.sc.GetSquadID(s.peerId)
	peers := s.sc.GetPeerList(squadId)
	for _, peer := range peers {
		s.acl[peer] = true
	}
	s.RegisterTopicValidator(squadId, s.validator)
	log.Println(squadId, "- Squad Initialized")
	s.isInitialized = true
	s.id = squadId
}

func (s Squad) RefreshACL(ctx context.Context) {
	// Track last refresh?
}

func (s Squad) MakeACLVerificationMessage() {

}

func (s Squad) VerifyACL() {}
