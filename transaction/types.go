package transaction

type Tx struct {
	From []byte
	To   []byte
}

type TxParser interface {
	Parse([]byte) Tx
	Rules()
}
