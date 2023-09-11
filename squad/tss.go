package squad

import (
	"crypto/elliptic"
	"math/big"

	"github.com/bnb-chain/tss-lib/v2/crypto"
	"github.com/bnb-chain/tss-lib/v2/crypto/paillier"
	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/solace-labs/skeyn/proto"
)

func (s Squad) LP_SAVE_DATA_KEY() string {
	return "SAVE_DATA_" + s.ID + s.peerId.String()
}

func bigIntToBytes(bi *big.Int) []byte {
	return bi.Bytes()
}

func bytesToBigInt(b []byte) *big.Int {
	return new(big.Int).SetBytes(b)
}

func publicKeyToProto(publicKey *paillier.PublicKey) *proto.PublicKey {
	return &proto.PublicKey{
		N: bigIntToBytes(publicKey.N),
	}
}

func protoToPublicKey(publicKey *proto.PublicKey) *paillier.PublicKey {
	return &paillier.PublicKey{
		N: bytesToBigInt(publicKey.N),
	}
}

func privateKeyToProto(pk *paillier.PrivateKey) *proto.PrivateKey {
	return &proto.PrivateKey{
		PublicKey: publicKeyToProto(&pk.PublicKey),
		LambdaN:   bigIntToBytes(pk.LambdaN),
		PhiN:      bigIntToBytes(pk.PhiN),
		P:         bigIntToBytes(pk.P),
		Q:         bigIntToBytes(pk.Q),
	}
}

func ecPointToProto(point *crypto.ECPoint) *proto.ECPoint {
	return &proto.ECPoint{
		X:     bigIntToBytes(point.X()),
		Y:     bigIntToBytes(point.Y()),
		Curve: []byte("256"),
	}
}

func ecPointToBytes(point *crypto.ECPoint) ([]byte, error) {
	xBytes := bigIntToBytes(point.X())
	yBytes := bigIntToBytes(point.Y())
	return append(xBytes, yBytes...), nil
}

func bytesToECPoint(data []byte, curve elliptic.Curve) (*crypto.ECPoint, error) {
	byteLen := len(data) / 2
	x := new(big.Int).SetBytes(data[:byteLen])
	y := new(big.Int).SetBytes(data[byteLen:])
	return crypto.NewECPoint(curve, x, y)
}

func protoToECPoint(point *proto.ECPoint) *crypto.ECPoint {
	x := new(big.Int).SetBytes(point.X)
	y := new(big.Int).SetBytes(point.Y)
	ecPoint, err := crypto.NewECPoint(tss.EC(), x, y)
	if err != nil {
		panic(err)
	}
	return ecPoint
}

func protoToPrivateKey(protoData *proto.PrivateKey) *paillier.PrivateKey {
	return &paillier.PrivateKey{
		PublicKey: paillier.PublicKey{
			N: bytesToBigInt(protoData.PublicKey.N),
		},
		LambdaN: bytesToBigInt(protoData.LambdaN),
		PhiN:    bytesToBigInt(protoData.PhiN),
		P:       bytesToBigInt(protoData.P),
		Q:       bytesToBigInt(protoData.Q),
	}
}

type UpdateMessage interface {
	GetWireMessage() []byte
	GetIsBroadcast() bool
	GetPayload() []byte
}

// What should the keys be?
// Wallet Address + Squad ID?
