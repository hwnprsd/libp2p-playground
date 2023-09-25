package common

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type WalletAddress string

func (addr WalletAddress) Bytes() []byte {
	return []byte(addr)
}

func (addr WalletAddress) String() string {
	return strings.ToLower(strings.TrimSpace(string(bytes.Trim([]byte(addr), "\x00"))))
}
func NewWalletAddress(addr []byte) WalletAddress {
	return WalletAddress(strings.ToLower(strings.TrimSpace(string(bytes.Trim([]byte(addr), "\x00")))))
}

func NewEthWalletAddress(addr common.Address) WalletAddress {
	// TODO: For now, store the hex string - later, implement chain ID, etc
	return WalletAddress(strings.ToLower(strings.TrimSpace(string(bytes.Trim([]byte(addr.Hex()), "\x00")))))
}
