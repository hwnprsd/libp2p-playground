package smartcontract

import (
	"libp2p-playground/utils"

	"github.com/libp2p/go-libp2p/core/peer"
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

func (*TestContract) GetSquadID(ID peer.ID) (string, error) {
	return "SQUAD_1", nil
}

func (*TestContract) GetPeerList(squadID string) ([]peer.ID, error) {
	pk1, _ := utils.HexToPubkey("3059301306072a8648ce3d020106082a8648ce3d03010703420004916991a9993dfed0e3f10949e1579a8cce6a9f7cec55bab8dc3d9b695be43f62ae074bb24622cc9ae81408ed43b6cf60c2eddfea08f6d40278183b4acd308cfc")
	peer1, _ := peer.IDFromPublicKey(pk1)
	_ = peer1

	pk2, _ := utils.HexToPubkey("3059301306072a8648ce3d020106082a8648ce3d03010703420004a077b88da4d409066bc7feb31c647464c8d18259238a1f3a60b0e484fde7b3d08846eebd518448684c4b502e6eb00e296df414d95e88c975136cd0c6b3d6e293")
	peer2, _ := peer.IDFromPublicKey(pk2)

	pk3, _ := utils.HexToPubkey("3059301306072a8648ce3d020106082a8648ce3d03010703420004e6427210ebc8c2b96f68a66b6de9e79afcfa99490d6b516db873491eaea2cc1c9389cfa55a835942345e7f35796289c4adaa2f02eb7681389d3416c48a029252")
	peer3, _ := peer.IDFromPublicKey(pk3)

	return []peer.ID{peer2, peer3}, nil
}

func (*TestContract) GetWalletsUnderManagement(squadID string) ([]WalletData, error) {
	priv, _ := utils.ParseB64Key("CAMSeTB3AgEBBCBEjggC6S8a6PX6spH0dCTW/9VGJV/f3nPVnZgkAdu35qAKBggqhkjOPQMBB6FEA0IABD2LHlnjffc10PVUVj3PIlahYdo/H4evtzp/j9ClnSYh1HHbjxyEkazedUSUe8prqqifwG3Ypmnr2xZ3gD9zwYI=")
	walletAddress, _ := priv.GetPublic().Raw()

	priv2, _ := utils.ParseB64Key("CAMSeTB3AgEBBCBaUjSijhsw2U3/fsQVQ2+HT11NlesyQr71RXpbh0zE76AKBggqhkjOPQMBB6FEA0IABASuKQIl4u3J1jho+29tYUNEOWgGiQyieoAHoyChd4wsqUe5OWHPanCG0pTY2NiKMNuUQUrKda8KxUXZf5dtwXk=")
	signerAddress, _ := priv2.GetPublic().Raw()

	return []WalletData{
		{
			WalletAddress: walletAddress,
			Signer:        signerAddress,
		},
	}, nil
}
