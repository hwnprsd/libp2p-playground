package utils

import (
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
)

const MAX_MESSAGE_LEN = 8000

func ReadStream(stream network.Stream) ([]byte, error) {
	var result []byte
	buf := make([]byte, 1024)

	for {
		n, err := stream.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		result = append(result, buf[:n]...)

		if MAX_MESSAGE_LEN > 0 && len(result) > MAX_MESSAGE_LEN {
			return nil, fmt.Errorf("received data exceeds maximum size")
		}
	}

	return result, nil
}
