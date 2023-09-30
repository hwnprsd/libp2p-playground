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

func (idx index) increment() error {
	if len(idx) != 4 {
		return fmt.Errorf("byte slice length must be 4 for int32")
	}

	carry := uint16(1)
	for i := 0; i < 4; i++ {
		sum := uint16(idx[i]) + carry
		idx[i] = uint8(sum & 0xFF)
		carry = sum >> 8
	}
	return nil
}

func (s *Squad) GetCurrentNonce() index {
	data, err := s.db.Get([]byte("NONCE"))
	if err != nil {
		log.Println("Error getting sender nonce", err)
		return []byte{0, 0, 0, 0}
	}
	if data == nil {
		data = []byte{0, 0, 0, 0}
	}
	err = index(data).increment()
	if err != nil {
		log.Println("Error incrementing sender nonce", err)
		return []byte{0, 0, 0, 0}
	}
	return data
}

func (s *Squad) updateNonce() error {
	senderNonce := s.GetCurrentNonce()
	return s.db.Set([]byte("NONCE"), senderNonce)
}

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
	err := index.increment()
	if err != nil {
		log.Println("err incrementing index")
		return err
	}
	return s.db.Set([]byte(db_index_key), index)
}
