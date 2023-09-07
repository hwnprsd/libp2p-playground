package utils

import (
	"github.com/libp2p/go-libp2p/core/crypto"
)

type Config interface {
	GetPrivKey() crypto.PrivKey
	GetPeers() AddrList
	GetPort() int
	GetShouldRunExternalRPCServer() bool
}

type config struct {
	port              int
	peers             AddrList
	privKey           crypto.PrivKey
	externalRpcServer bool
}

var cliConfig = config{
	peers: make(AddrList, 0),
}

func (c config) GetPrivKey() crypto.PrivKey {
	return c.privKey
}

func (c config) GetPeers() AddrList {
	return c.peers
}

func (c config) GetPort() int {
	return c.port
}

func (c config) GetShouldRunExternalRPCServer() bool {
	return c.externalRpcServer
}
