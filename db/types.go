package db

type Database interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, value []byte) error
	Delete(key []byte) error
}

// We defensively turn nil keys or values into []byte{} for
// most operations.
func checkNilBytes(bz []byte) []byte {
	if bz == nil {
		return []byte{}
	}
	return bz
}
