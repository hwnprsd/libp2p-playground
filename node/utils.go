package node

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/solace-labs/skeyn/common"
)

func (n *Node) verifyWalletAddr(walletAddrString string) (common.WalletAddress, error) {
	walletAddressEth := ethcommon.HexToAddress(walletAddrString)
	if walletAddressEth.Bytes() == nil {
		return common.WalletAddress(""), fmt.Errorf("[1] Invalid wallet address - %s", walletAddrString)
	}

	walletAddr := common.NewEthWalletAddress(walletAddressEth)

	if _, exists := n.squad[walletAddr]; !exists {
		return walletAddr, fmt.Errorf("[2] WalletAddr not being managed - %s", walletAddrString)
	}

	return walletAddr, nil
}
