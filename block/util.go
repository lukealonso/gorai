package gorai

import (
	"encoding/hex"
	"fmt"
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
