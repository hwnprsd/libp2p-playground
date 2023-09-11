package squad

import (
	"github.com/bnb-chain/tss-lib/v2/ecdsa/keygen"
	"github.com/solace-labs/skeyn/proto"
	protob "google.golang.org/protobuf/proto"
)

var curve = []byte("256")

type StoredSaveData struct {
	*keygen.LocalPartySaveData
}

func NewStoredSaveData(saveData *keygen.LocalPartySaveData) StoredSaveData {
	return StoredSaveData{saveData}
}

func (lpsd StoredSaveData) Bytes() []byte {
	protoLPSD := &proto.LocalPartySaveData{
		LocalPreParams: &proto.LocalPreParams{
			PaillierSK: privateKeyToProto(lpsd.LocalPreParams.PaillierSK),
			Ntildei:    bigIntToBytes(lpsd.LocalPreParams.NTildei),
			H1I:        bigIntToBytes(lpsd.LocalPreParams.H1i),
			H2I:        bigIntToBytes(lpsd.LocalPreParams.H2i),
			Alpha:      bigIntToBytes(lpsd.LocalPreParams.Alpha),
			Beta:       bigIntToBytes(lpsd.LocalPreParams.Beta),
			P:          bigIntToBytes(lpsd.LocalPreParams.P),
			Q:          bigIntToBytes(lpsd.LocalPreParams.Q),
		},
		LocalSecrets: &proto.LocalSecrets{
			Xi:      bigIntToBytes(lpsd.LocalSecrets.Xi),
			ShareID: bigIntToBytes(lpsd.LocalSecrets.ShareID),
		},
		// Populate ECDSAPub assuming you have ECPoint to bytes conversion method
		EcdsaPub: ecPointToProto(lpsd.ECDSAPub),
	}

	// Convert Ks []*big.Int to repeated bytes field
	for _, k := range lpsd.Ks {
		protoLPSD.Ks = append(protoLPSD.Ks, bigIntToBytes(k))
	}

	// Convert NTildej, H1j, H2j []*big.Int to repeated bytes field
	for _, value := range lpsd.NTildej {
		protoLPSD.Ntildej = append(protoLPSD.Ntildej, bigIntToBytes(value))
	}
	for _, value := range lpsd.H1j {
		protoLPSD.H1J = append(protoLPSD.H1J, bigIntToBytes(value))
	}
	for _, value := range lpsd.H2j {
		protoLPSD.H2J = append(protoLPSD.H2J, bigIntToBytes(value))
	}

	// Convert BigXj []*crypto.ECPoint to repeated ECPoint messages
	for _, point := range lpsd.BigXj {
		protoLPSD.BigXj = append(protoLPSD.BigXj, &proto.ECPoint{
			Curve: curve,
			X:     bigIntToBytes(point.X()),
			Y:     bigIntToBytes(point.Y()),
		})
	}

	// Convert PaillierPKs []*paillier.PublicKey to repeated bytes
	for _, pk := range lpsd.PaillierPKs {
		protoLPSD.PaillierPKs = append(protoLPSD.PaillierPKs, publicKeyToProto(pk))
	}

	b, _ := protob.Marshal(protoLPSD)
	return b
}

func StoredSaveDataFromBytes(data []byte) StoredSaveData {
	protoLSD := &proto.LocalPartySaveData{}
	if err := protob.Unmarshal(data, protoLSD); err != nil {
		panic(err)
	}

	lpsd := &keygen.LocalPartySaveData{
		LocalPreParams: keygen.LocalPreParams{
			PaillierSK: protoToPrivateKey(protoLSD.LocalPreParams.PaillierSK),
			NTildei:    bytesToBigInt(protoLSD.LocalPreParams.Ntildei),
			H1i:        bytesToBigInt(protoLSD.LocalPreParams.H1I),
			H2i:        bytesToBigInt(protoLSD.LocalPreParams.H2I),
			Alpha:      bytesToBigInt(protoLSD.LocalPreParams.Alpha),
			Beta:       bytesToBigInt(protoLSD.LocalPreParams.Beta),
			P:          bytesToBigInt(protoLSD.LocalPreParams.P),
			Q:          bytesToBigInt(protoLSD.LocalPreParams.Q),
		},
		LocalSecrets: keygen.LocalSecrets{
			Xi:      bytesToBigInt(protoLSD.LocalSecrets.Xi),
			ShareID: bytesToBigInt(protoLSD.LocalSecrets.ShareID),
		},
		ECDSAPub: protoToECPoint(protoLSD.EcdsaPub),
	}
	for _, k := range protoLSD.Ks {
		lpsd.Ks = append(lpsd.Ks, bytesToBigInt(k))
	}

	for _, value := range protoLSD.Ntildej {
		lpsd.NTildej = append(lpsd.NTildej, bytesToBigInt(value))
	}
	for _, value := range protoLSD.H1J {
		lpsd.H1j = append(lpsd.H1j, bytesToBigInt(value))
	}
	for _, value := range protoLSD.H2J {
		lpsd.H2j = append(lpsd.H2j, bytesToBigInt(value))
	}

	for _, point := range protoLSD.BigXj {
		lpsd.BigXj = append(lpsd.BigXj, protoToECPoint(point))
	}

	for _, pk := range protoLSD.PaillierPKs {
		lpsd.PaillierPKs = append(lpsd.PaillierPKs, protoToPublicKey(pk))
	}
	return StoredSaveData{
		lpsd,
	}
}
