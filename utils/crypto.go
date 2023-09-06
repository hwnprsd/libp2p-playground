package utils

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func ParseB64Key(b64 string) (crypto.PrivKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.UnmarshalPrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func HexToPubkey(h string) (crypto.PubKey, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalECDSAPublicKey(b)
}
