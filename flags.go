package main

import (
	"encoding/base64"
	"flag"
	"strings"

	"github.com/libp2p/go-libp2p/core/crypto"
	maddr "github.com/multiformats/go-multiaddr"
)

// A new type we need for writing a custom flag parser
type addrList []maddr.Multiaddr

func (al *addrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) Set(value string) error {
	addr, err := maddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

func ParseFlags() (int, addrList, crypto.PrivKey) {
	var port int
	flag.IntVar(&port, "port", 3210, "Listen Port")
	bsPeers := addrList{}
	flag.Var(&bsPeers, "peer", "Bootstrap Peers")
	b64PrivKey := ""
	flag.StringVar(&b64PrivKey, "priv", "nil", "Private Key")
	flag.Parse()
	keyBytes, err := base64.StdEncoding.DecodeString(b64PrivKey)
	if err != nil {
		panic(err)
	}
	priv, err := crypto.UnmarshalPrivateKey(keyBytes)
	if err != nil {
		panic(err)
	}
	return port, bsPeers, priv
}
