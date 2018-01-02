package gorai

import (
	"fmt"

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

func (s *Signature) Verify(publicKey []byte, message []byte) bool {
	h, err := blake2b.New512(nil)
	if err != nil {
		// blake2b doesn't return errors from New512, but check anyways.
		panic(err)
	}
	return ed25519.Verify(publicKey, message, s.Bytes(), h)
}

func (s *Signature) Sign(privateKey []byte, message []byte) error {
	h, err := blake2b.New512(nil)
	if err != nil {
		// blake2b doesn't return errors from New512, but check anyways.
		panic(err)
	}
	if len(privateKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid private key size")
	}
	sig := ed25519.Sign(ed25519.PrivateKey(privateKey), bh.Bytes(), h)
	if sig == nil || len(sig) != ed25519.SignatureSize || len(sig) != len(s.Data) {
		return fmt.Errorf("invalid signature size")
	}
	copy(s.Data[:], sig)
	return nil
}
