package gorai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeHex(t *testing.T) {
	var b [2]byte
	assert.Nil(t, decodeHex(b[:], []byte("AABB")))
	assert.NotNil(t, decodeHex(b[:], []byte("AA")))
	assert.NotNil(t, decodeHex(b[:], []byte("AABBCC")))
	assert.NotNil(t, decodeHex(b[:], []byte("AAXX")))
	assert.NotNil(t, decodeHex(b[:], []byte("AABB.")))
}
