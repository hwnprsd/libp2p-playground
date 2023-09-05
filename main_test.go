package main

import (
	"encoding/base64"
	"log"
	"testing"

	"github.com/libp2p/go-libp2p/core/crypto"
)

// Generate private keys to be used for stable Peer IDs
func Test_CreatePrivKKey(t *testing.T) {
	priv, _, _ := crypto.GenerateKeyPair(crypto.ECDSA, -1)
	keyBytes, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		panic(err)
	}
	base64Key := base64.StdEncoding.EncodeToString(keyBytes)
	log.Println(base64Key)
}
