package squad

import (
	"math/big"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/libp2p/go-libp2p/core/peer"
)

type SquadPeers map[peer.ID]bool

func (s SquadPeers) List() []peer.ID {
	keys := make([]peer.ID, 0, len(s))
	for i := range s {
		keys = append(keys, i)
	}
	return keys
}

func (s SquadPeers) Match(moniker string) *peer.ID {
	for peer := range s {
		if peer.String() == moniker {
			return &peer
		}
	}
	return nil
}

func (s *Squad) SortedPartyIDs() []*tss.PartyID {
	parties := make([]*tss.PartyID, 0)
	for p := range s.peers {
		parties = append(parties, s.ToPartyID(&p))
	}
	parties = tss.SortPartyIDs(parties)
	return parties
}

// TODO: Better error handling
func (s *Squad) ToPartyID(p *peer.ID) *tss.PartyID {
	pubKeyBytes, _ := s.peerStore.PubKey(*p).Raw()
	return tss.NewPartyID(
		p.String(),
		p.String(), // So, things can be reproducable
		new(big.Int).SetBytes(pubKeyBytes),
	)
}

func (s *Squad) GetSortedPartyID(p *peer.ID) *tss.PartyID {
	targetId := s.ToPartyID(p)
	for _, id := range s.SortedPartyIDs() {
		if id.Id == targetId.Id {
			return id
		}
	}
	return nil
}

func (s *Squad) PartyID() *tss.PartyID {
	selfId := s.ToPartyID(&s.peerId)
	for _, id := range s.SortedPartyIDs() {
		if id.Id == selfId.Id {
			return id
		}
	}
	return nil
}

func (s *Squad) ToPeerID(p *tss.PartyID) *peer.ID {
	return s.peers.Match(p.GetMoniker())
}
