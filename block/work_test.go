package gorai

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkVerify(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(sendBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.True(t, b.Work().Verify(b.Previous()))
	b, err = DecodeBlockJSON(bytes.NewReader([]byte(receiveBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.True(t, b.Work().Verify(b.Previous()))
	w := b.Work()
	w.UnmarshalText([]byte("0000000000000000"))
	assert.False(t, w.Verify(b.Previous()))
}
