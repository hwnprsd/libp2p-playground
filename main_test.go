package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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

type request struct {
	T       string `json:"type"`
	Payload struct {
		WalletAddress string `json:"walletAddress"`
		Signature     string `json:"signature"`
		Data          string `json:"data"`
	} `json:"payload"`
}

func Test_CreateWalletData(t *testing.T) {
	// Generate Ethereum-compatible wallet keys
	walletPriv, _ := ethcrypto.GenerateKey()
	walletPub := walletPriv.PublicKey
	walletAddr := ethcrypto.PubkeyToAddress(walletPub)

	// Generate Ethereum-compatible user keys
	userPriv, _ := ethcrypto.GenerateKey()
	userPub := userPriv.PublicKey
	// userPubBytes := ethcrypto.FromECDSAPub(&userPub)

	// Sign data
	signData := []byte("SIGN ME PLEASE")
	signDataHash := ethcrypto.Keccak256(signData)
	sig, _ := ethcrypto.Sign(signDataHash, userPriv)

	log.Println("UserAddr", ethcrypto.PubkeyToAddress(userPub).Hex())

	r := request{}
	r.T = "1"
	r.Payload.Data = hex.EncodeToString(signData)
	r.Payload.WalletAddress = walletAddr.Hex()
	r.Payload.Signature = hex.EncodeToString(sig)

	j, _ := json.Marshal(r)
	fmt.Println("\n\ncurl -X POST 'http://localhost:5050/v1/transaction' \\")
	fmt.Println("-H 'Content-Type: application/json' \\")
	fmt.Println("-H 'Content-Type: application/json' \\")
	fmt.Printf("-d '%s'\n\n", string(j))
	fmt.Println("")
}
