package squad

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/solace-labs/skeyn/proto"
	protob "google.golang.org/protobuf/proto"
)

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
	return nil
}

func (s *Squad) HashSolaceTx(tx *proto.SolaceTx) ([]byte, error) {
	b, err := protob.Marshal(tx)
	if err != nil {
		return nil, err
	}
	return append([]byte(TX_PREFIX), b...), nil
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
