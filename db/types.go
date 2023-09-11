package db

// We defensively turn nil keys or values into []byte{} for
// most operations.
func checkNilBytes(bz []byte) []byte {
	if bz == nil {
		return []byte{}
	}
	return bz
}
