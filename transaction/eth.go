package transaction

type EthereumTx struct {
	From         []byte
	To           []byte
	TokenAddress []byte
	Value        []byte
	Signatures   [][]byte
}

func (e EthereumTx) GetFrom() []byte {
	return e.From
}

func (e EthereumTx) GetTo() []byte {
	return e.To
}

func (e EthereumTx) GetTokenAddress() []byte {
	return e.TokenAddress
}

func (e EthereumTx) GetValue() []byte {
	return e.Value
}

func (e EthereumTx) GetSignatures() [][]byte {
	return e.Signatures
}
