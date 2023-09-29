package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const WalletAddrLen = 20

type WalletAddress string

func (addr WalletAddress) Bytes() []byte {
	return hexutil.MustDecode(string(addr))
}

func (addr WalletAddress) String() string {
	return hexutil.Encode(addr.Bytes())
}
func NewWalletAddress(addr []byte) WalletAddress {
	return WalletAddress(hexutil.Encode(addr))
}

func NewEthWalletAddress(addr common.Address) WalletAddress {
	// TODO: For now, store the hex string - later, implement chain ID, etc
	return WalletAddress(hexutil.Encode(addr.Bytes()))
}

func NewEthWalletAddressString(addr string) (WalletAddress, error) {
	b, err := hexutil.Decode(addr)
	if err != nil {
		return WalletAddress(""), err
	}
	return WalletAddress(hexutil.Encode(b)), nil
}
