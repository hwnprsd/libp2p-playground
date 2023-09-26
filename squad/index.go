package squad

import (
	"encoding/binary"
	"fmt"
	"log"
)

type index []byte

func (i index) Int() int {
	value := int(binary.LittleEndian.Uint32(i))
	return value
}

func (i index) Bytes() []byte {
	return i
}

func IndexFromInt(val int) index {
	// Convert int32 to byte slice
	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, uint32(val))
	return index(valueBytes)
}

const db_index_key = "DB_INDEX"

func (s *Squad) getDbIndex() index {
	indexB, err := s.db.Get([]byte(db_index_key))
	if err != nil {
		log.Println("error fetching db index")
		return []byte{0, 0, 0, 0}
	}
	if indexB == nil {
		return []byte{0, 0, 0, 0}
	}
	return indexB
}

func (s *Squad) updateIndex() error {
	index := s.getDbIndex()
	err := increment(index)
	if err != nil {
		log.Println("err incrementing index")
		return err
	}
	return s.db.Set([]byte(db_index_key), index)
}

func increment(valueBytes []byte) error {
	if len(valueBytes) != 4 {
		return fmt.Errorf("byte slice length must be 4 for int32")
	}

	carry := uint16(1)
	for i := 0; i < 4; i++ {
		sum := uint16(valueBytes[i]) + carry
		valueBytes[i] = uint8(sum & 0xFF)
		carry = sum >> 8
	}
	return nil
}
