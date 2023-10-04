package squad

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/utils"
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

	senderAddr, err := common.NewEthWalletAddressString(tx.Sender.Addr)
	if err != nil {
		return err
	}

	err = utils.VerifyEthSignature(messageBytes, sig, senderAddr)
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
