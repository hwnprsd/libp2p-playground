package squad

import (
	"log"
	"math/big"

	"github.com/bnb-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p/core/crypto"
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

func (s SquadPeers) SortedPartyIDs() []*tss.PartyID {
	parties := make([]*tss.PartyID, 0)
	for _, p := range s.List() {
		parties = append(parties, ToPartyID(&p))
	}
	parties = tss.SortPartyIDs(parties)
	return parties
}

// TODO: Better error handling
func ToPartyID(p *peer.ID) *tss.PartyID {
	pubKey, err := p.ExtractPublicKey()
	if err != nil {
		log.Println("ERR extracting pub key", p)
		return nil
	}
	pubKeyBytes, err := pubKey.Raw()
	if err != nil {
		log.Println("ERR extracting bytes from pubkey", p)
		return nil
	}
	return tss.NewPartyID(
		string(*p),
		p.ShortString(),
		new(big.Int).SetBytes(pubKeyBytes),
	)
}

func (s Squad) PartyID() *tss.PartyID {
	return ToPartyID(&s.peerId)
}

func ToPeerID(p *tss.PartyID) *peer.ID {
	pubKeyBytes := p.GetKey()
	pubkey, err := crypto.UnmarshalPublicKey(pubKeyBytes)
	if err != nil {
		log.Println("Error marshalling pubkey", err)
	}
	id, err := peer.IDFromPublicKey(pubkey)
	if err != nil {
		log.Println("Error creating ID from pubkey", err)
	}
	return &id
}
