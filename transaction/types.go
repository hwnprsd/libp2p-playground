package transaction

type SolaceTranasction interface {
	GetFrom() []byte
	GetTo() []byte
	GetTokenAddress() []byte
	GetValue() []byte
	GetSignatures() [][]byte
}
