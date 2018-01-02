package gorai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockHash(t *testing.T) {
	bh := BlockHash{}
	bh.Hash([]byte("1234"), []byte("5678"))
	assert.EqualValues(t, "33756ca9d42894dd598b1e84a323c4c9850781cc0e6c687737997120e5ab71f4", bh.String())
	bh2 := NewBlockHash([]byte("12345678"))
	assert.EqualValues(t, "33756ca9d42894dd598b1e84a323c4c9850781cc0e6c687737997120e5ab71f4", bh2.String())
	dec, err := DecodeBlockHash("33756ca9d42894dd598b1e84a323c4c9850781cc0e6c687737997120e5ab71f4")
	assert.Nil(t, err)
	assert.EqualValues(t, "33756ca9d42894dd598b1e84a323c4c9850781cc0e6c687737997120e5ab71f4", dec.String())
}
