package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/solace-labs/skeyn/common"
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

func HexToBytes(h string) []byte {
	return ethcommon.FromHex(h)
}

func VerifyEthSignature(message []byte, sig []byte, sender common.Addr) error {
	if len(sig) == 0 {
		// TODO: For testing
		return nil
	}
	// ETH_SIGNED_MESSAGE
	messageHash := accounts.TextHash(message)

	if sig[ethcrypto.RecoveryIDOffset] == 27 || sig[ethcrypto.RecoveryIDOffset] == 28 {
		sig[ethcrypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	pubKeyRaw, err := ethcrypto.Ecrecover(messageHash, sig)
	if err != nil {
		fmt.Printf("[1] Error recovering pubkey")
		fmt.Println(err)
		return err
	}

	pubKey, _ := ethcrypto.UnmarshalPubkey(pubKeyRaw)
	recovered := ethcrypto.PubkeyToAddress(*pubKey)
	expected := ethcommon.BytesToAddress(sender.Bytes())

	fmt.Printf("E: %s / R %s", expected, recovered)
	if expected == recovered {
		return nil
	} else {
		return fmt.Errorf("Signature verification failed")
	}
}

func EcdsaBytesToAddress(b []byte) string {
	hash := ethcrypto.Keccak256(b[1:])
	address := ethcommon.BytesToAddress(hash[12:]).Hex()
	return address
}
