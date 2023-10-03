package squad

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
)

const TX_PREFIX = "\x19SOLACE_TX\n"

func (s *Squad) ValidateSolaceTx(tx *proto.SolaceTx) error {
	// 1. Verify tx signature
	message := tx.Namespace + tx.WalletAddr + tx.Sender.Addr + tx.ToAddr + tx.TokenAddr + fmt.Sprintf("%d", tx.Value) + fmt.Sprintf("%d", tx.Sender.Nonce)
	for _, sig := range tx.Signatures {
		message += sig
	}

	messageBytes := []byte(message)
	sig, err := hexutil.Decode(tx.TxSignature)
	if err != nil {
		return err
	}

	err = verifySignature(messageBytes, sig, tx.Sender.Addr)
	if err != nil {
		return err
	}

	// 2. Verify sender nonce
	nonce := s.GetCurrentNonce()
	if nonce.Int() != int(tx.Sender.Nonce) {
		return fmt.Errorf("Incorrect Nonce")
	}

	return nil
}

func HashSolaceTx(tx *proto.SolaceTx) ([]byte, error) {
	walletAddr, err := common.NewEthWalletAddressString(tx.WalletAddr)
	if err != nil {
		return nil, err
	}

	nonceBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceBytes, uint32(tx.Sender.Nonce))

	var hash []byte
	hash = append(hash, []byte(TX_PREFIX)...)
	hash = append(hash, walletAddr.Bytes()...)
	hash = append(hash, nonceBytes...)
	return hash, nil
}

func ParseSolaceTxHash(hashString string) ([]byte, common.Addr, error) {
	hash, err := hexutil.Decode(hashString)
	if err != nil {
		return nil, common.Addr(""), err
	}
	walletAddrB := hash[len(TX_PREFIX) : len(TX_PREFIX)+20]
	return hash, common.NewWalletAddress(walletAddrB), nil
}

func verifySignature(message []byte, sig []byte, sender string) error {
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
	expected := ethcommon.HexToAddress(sender)

	fmt.Printf("E: %s / R %s", expected, recovered)
	if expected == recovered {
		return nil
	} else {
		return fmt.Errorf("Signature verification failed")
	}
}
