package gorai

import (
	"github.com/golang/crypto/blake2b"
	"github.com/lukealonso/ed25519"
)

type Signature struct {
	Data [64]byte
}

func (s *Signature) MarshalText() ([]byte, error) {
	return encodeHex(s.Data[:]), nil
}

func (s *Signature) UnmarshalText(b []byte) error {
	return decodeHex(s.Data[:], b)
}

func (s *Signature) Bytes() []byte {
	return s.Data[:]
}

func (s *Signature) Verify(a *Account, bh *BlockHash) bool {
	h, err := blake2b.New512(nil)
	if err != nil {
		// blake2b doesn't return errors from New512, but check anyways.
		return false
	}
	return ed25519.Verify(a.PublicKey(), bh.Bytes(), s.Bytes(), h)
}
