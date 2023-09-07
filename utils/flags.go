package utils

import (
	"flag"
	"strings"

	maddr "github.com/multiformats/go-multiaddr"
)

// A new type we need for writing a custom flag parser
type AddrList []maddr.Multiaddr

func (al *AddrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *AddrList) Set(value string) error {
	addr, err := maddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

var isCliFlagsParsed = false

func GetCliConfig() Config {
	if !isCliFlagsParsed {
		ParseFlags()
	}
	return cliConfig
}

func ParseFlags() {
	// Port
	flag.IntVar(&cliConfig.port, "port", 3210, "Listen Port")

	// Peers
	flag.Var(&cliConfig.peers, "peer", "Bootstrap Peers")

	// Should Run ExternalRPC Server

	// Private Key
	b64PrivKey := ""
	flag.StringVar(&b64PrivKey, "priv", "nil", "Private Key")

	// External RPC Server
	flag.BoolVar(&cliConfig.externalRpcServer, "rpc", false, "Should run an RPC Server for external calls")

	flag.Parse()

	priv, err := ParseB64Key(b64PrivKey)
	cliConfig.privKey = priv
	if err != nil {
		panic(err)
	}
	isCliFlagsParsed = true
}
