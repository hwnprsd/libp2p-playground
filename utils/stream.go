package utils

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

const MAX_MESSAGE_LEN = 8000

func ReadStream(stream network.Stream) ([]byte, error) {
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(stream, lenBuf); err != nil {
		return nil, err
	}

	msgLen := binary.BigEndian.Uint32(lenBuf)

	if MAX_MESSAGE_LEN > 0 && int(msgLen) > MAX_MESSAGE_LEN {
		return nil, fmt.Errorf("message too large")
	}

	buf := make([]byte, msgLen)

	if _, err := io.ReadFull(stream, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

// Data = Len + Actual Data
func WriteStream(stream network.Stream, data []byte) error {
	lenBuf := make([]byte, 4)

	binary.BigEndian.PutUint32(lenBuf, uint32(len(data)))
	if _, err := stream.Write(lenBuf); err != nil {
		return err
	}

	if _, err := stream.Write(data); err != nil {
		return err
	}

	return nil
}
