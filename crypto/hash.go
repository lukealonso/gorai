package crypto

import (
	"hash"

	"github.com/golang/crypto/blake2b"
)

func newBlake2b() hash.Hash {
	h, err := blake2b.New512(nil)
	if err != nil {
		// blake2b doesn't return errors from New512, but check anyways.
		panic(err)
	}
	return h
}

func HashBlake2b(dest []byte, parts ...[]byte) {
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
