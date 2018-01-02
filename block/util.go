package gorai

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/crypto/blake2b"
)

func decodeHex(dst, src []byte) error {
	if len(src)/2 != len(dst) {
		return fmt.Errorf("invalid hex string size")
	}
	nr, err := hex.Decode(dst, src)
	if err != nil {
		return err
	}
	if nr != len(dst) {
		// Shouldn't happen, checked above.
		return fmt.Errorf("invalid hex string size")
	}
	return nil
}

func encodeHex(src []byte) []byte {
	b := make([]byte, len(src)*2)
	hex.Encode(b, src)
	return b
}

func hashBlake2b(dest []byte, parts ...[]byte) {
	size := len(dest)
	h, err := blake2b.New(size, nil)
	if err != nil {
		// blake2b doesn't return any errors.
		panic(err)
	}
	for _, part := range parts {
		h.Write(part)
	}
	bw := h.Sum(dest[:0])
	if len(bw) != size {
		// Shouldn't happen.
		panic("hash sum error")
	}
}
