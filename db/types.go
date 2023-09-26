package db

type Database interface {
	Get([]byte) ([]byte, error)
	Set([]byte, []byte) error
	Delete([]byte) error
	GetAll(string) [][]byte
}

// We defensively turn nil keys or values into []byte{} for
// most operations.
func checkNilBytes(bz []byte) []byte {
	if bz == nil {
		return []byte{}
	}
	return bz
}
