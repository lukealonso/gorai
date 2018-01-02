package gorai

import (
	"bytes"
	"encoding/binary"
)

const WorkThreshold uint64 = 0xffffffc000000000

type Work struct {
	Value [8]byte
}

func (w *Work) MarshalText() ([]byte, error) {
	return encodeHex(w.Value[:]), nil
}

func (w *Work) UnmarshalText(b []byte) error {
	return decodeHex(w.Value[:], b)
}

func (w *Work) String() string {
	return string(encodeHex(w.Value[:]))
}

func (w *Work) Verify(bh *BlockHash) bool {
	var checkBytes [8]byte
	// The reference implementation has many endianness bugs, this one is now part of the protocol.
	var littleEndianBytes [8]byte
	for i := 0; i < 8; i++ {
		littleEndianBytes[i] = w.Value[7-i]
	}
	hashBlake2b(checkBytes[:], littleEndianBytes[:], bh.Bytes())
	var checkValue uint64
	if binary.Read(bytes.NewReader(checkBytes[:]), binary.LittleEndian, &checkValue) != nil {
		return false
	}
	return checkValue >= WorkThreshold
}
