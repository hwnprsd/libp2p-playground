package utils

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

const (
	MAX_MESSAGE_LEN = 500_000 // bytes
)

// Data = Len + Actual Data
func ReadStream(stream network.Stream) (walletAddr []byte, data []byte, err error) {
	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(stream, lenBuf); err != nil {
		return nil, nil, err
	}

	msgLen := binary.BigEndian.Uint32(lenBuf)

	if MAX_MESSAGE_LEN > 0 && int(msgLen) > MAX_MESSAGE_LEN {
		return nil, nil, fmt.Errorf("message too large - %d", int(msgLen))
	}

	buf := make([]byte, msgLen)

	if _, err := io.ReadFull(stream, buf); err != nil {
		return nil, nil, err
	}

	return buf[:60], buf[60:], nil
}

func WriteStream(stream network.Stream, data []byte, walletAddr []byte) error {
	lenBuf := make([]byte, 4)

	walletAddrPadded := make([]byte, 60)
	copy(walletAddrPadded, walletAddr)

	finalData := append(walletAddrPadded, data...)

	binary.BigEndian.PutUint32(lenBuf, uint32(len(finalData)))

	if _, err := stream.Write(lenBuf); err != nil {
		return err
	}

	if _, err := stream.Write(finalData); err != nil {
		return err
	}

	return nil
}
