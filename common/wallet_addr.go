package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const AddrLen = 20

// Keeping this as a string, as it's used as a Key in [SquadMap]
type Addr string

func (addr Addr) Bytes() []byte {
	return hexutil.MustDecode(string(addr))
}

func (addr Addr) String() string {
	return hexutil.Encode(addr.Bytes())
}

func ZeroAddr() Addr {
	return Addr("0x00")
}

func NewEthAddrSlice(addresses []string) ([]Addr, error) {
	addrs := make([]Addr, 0)
	for _, addr := range addresses {
		_addr, err := NewEthWalletAddressString(addr)
		addrs = append(addrs, _addr)
		if err != nil {
			return nil, err
		}
	}
	return addrs, nil
}

func NewWalletAddress(addr []byte) Addr {
	return Addr(hexutil.Encode(addr))
}

func NewEthWalletAddress(addr common.Address) Addr {
	// TODO: For now, store the hex string - later, implement chain ID, etc
	return Addr(hexutil.Encode(addr.Bytes()))
}

func NewEthWalletAddressString(addr string) (Addr, error) {
	b, err := hexutil.Decode(addr)
	if err != nil {
		return Addr(""), err
	}
	return Addr(hexutil.Encode(b)), nil
}
