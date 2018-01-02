package gorai

import (
	"bytes"
	"testing"

	"github.com/golang/crypto/blake2b"
	"github.com/lukealonso/ed25519"
	"github.com/stretchr/testify/assert"
)

func TestSignatureVerify(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(sendBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	var bh BlockHash
	b.Hash(&bh)
	a, err := DecodeAccount("xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuohr3")
	assert.Nil(t, err)
	assert.True(t, b.Signature().Verify(a.PublicKey(), &bh))
	bh2 := BlockHash{}
	assert.Nil(t, err)
	assert.False(t, b.Signature().Verify(a.PublicKey(), &bh2))
	a, err = DecodeAccount("xrb_13ezf4od79h1tgj9aiu4djzcmmguendtjfuhwfukhuucboua8cpoihmh8byo")
	assert.Nil(t, err)
	assert.False(t, b.Signature().Verify(a.PublicKey(), &bh))
}

func TestSignatureSign(t *testing.T) {

	b, err := DecodeBlockJSON(bytes.NewReader([]byte(sendBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)

	var bh BlockHash
	b.Hash(&bh)
	sig := b.Signature()

	h, _ := blake2b.New512(nil)
	pub, prv, _ := ed25519.GenerateKey(nil, h)

	assert.Nil(t, sig.Sign(prv, &bh))
	assert.Nil(t, err)
	assert.True(t, sig.Verify(pub, &bh))
}
