package smartcontract

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/solace-labs/skeyn/common"
)

type WalletData struct {
	WalletAddress []byte
	Signer        []byte // the signer who is allowed to send messages on behalf of the wallet
}

type NetworkState interface {
	GetSquadID(ID peer.ID) (string, error)
	GetPeerList(squadID string) ([]peer.ID, error)
	GetWalletsUnderManagement(squadID string) ([]WalletData, error)
}

type TestContract struct{}

type SmartContractWallet interface {
	ValidateSetup(signature []byte) (bool, error)
	ValidateRuleAddition(rule []byte, signature []byte, sender common.Addr) (bool, error)
	ValidateRuleDeletion(rule []byte, signature []byte, sender common.Addr) (bool, error)
}
