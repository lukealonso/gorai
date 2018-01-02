package gorai

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sendBlock = `{
    "type": "send",
    "previous": "4270f4fb3a820fe81827065f967a9589df5ca860443f812d21ece964ac359e05",
    "destination": "xrb_1111111111111111111111111111111111111111111111111111hifc8npp",
    "balance": "000000041c06df91d202b70a40000011",
    "work": "7202df8a7c380578",
    "signature": "047115cb577ac78f5c66ad79bbf47540de97a441456004190f22025fe4255285f57010d962601ae64c266c98fa22973dd95ac62309634940b727ac69f0c86d03"
}`

const openBlock = `{
    "type": "open",
    "source": "4270f4fb3a820fe81827065f967a9589df5ca860443f812d21ece964ac359e05",
    "representative": "xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
    "account": "xrb_1ipx847tk8o46pwxt5qjdbncjqcbwcc1rrmqnkztrfjy5k7z4imsrata9est",
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

const receiveBlock = `{
    "type": "receive",
    "previous": "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8",
    "source": "3b09d9743b8e7d0e43b80f177a10bb8d272806e98bc5948999a175508c5d92f5",
    "work": "ed873ff0c0203175",
    "signature": "380b1cf8198fd7fe2e997406177166129a5084b908696dae516e8a138a1f9ad2f9d47c121e7db9680c95b2d658ed580049cb0987c38c65e6c1926faf1040e30e"
}`

// Fake, need to find a real change block
const changeBlock = `{
    "type": "change",
    "previous": "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8",
    "representative": "xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

const unknownBlock = `{
    "type": "unknown",
    "previous": "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8",
    "representative": "xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

const badBlock1 = `{
    type": "unknown",
    "previous": "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8",
    "representative": "xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

const badBlock2 = `{
    "type": 2,
    "previous": "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8",
    "representative": "xrb_1awsn43we17c1oshdru4azeqjz9wii41dy8npubm4rg11so7dx3jtqgoeahy",
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

const badBlock3 = `{
    "type": "open",
    "previous": 20,
    "representative": 30,
    "work": "12d10d44912c9085",
    "signature": "712df7c4af0bd92446ec64d3f61d54510a93a591a638db9de812d9cb1b0b47edad31e7b23ecf3fd8ab9a948162a0cb5c1b8ab29e0c672029f53135f6b933b804"
}`

func TestBlockCodecJSONSend(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(sendBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.EqualValues(t, "4270f4fb3a820fe81827065f967a9589df5ca860443f812d21ece964ac359e05", b.Previous().String())
	assert.EqualValues(t, "7202df8a7c380578", b.Work().String())
	jb, err := json.MarshalIndent(b, "", "    ")
	assert.Nil(t, err)
	assert.EqualValues(t, sendBlock, string(jb))
}

func TestBlockCodecJSONOpen(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(openBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.Nil(t, b.Previous())
	assert.EqualValues(t, "12d10d44912c9085", b.Work().String())
	jb, err := json.MarshalIndent(b, "", "    ")
	assert.Nil(t, err)
	assert.EqualValues(t, openBlock, string(jb))
}

func TestBlockCodecJSONReceive(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(receiveBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.EqualValues(t, "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8", b.Previous().String())
	assert.EqualValues(t, "ed873ff0c0203175", b.Work().String())
	jb, err := json.MarshalIndent(b, "", "    ")
	assert.Nil(t, err)
	assert.EqualValues(t, receiveBlock, string(jb))
}

func TestBlockCodecJSONChange(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(changeBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	assert.EqualValues(t, "849df42e3e550d1a385e40615acd630124216b4f40d710547ed9bb6ba8b6feb8", b.Previous().String())
	assert.EqualValues(t, "12d10d44912c9085", b.Work().String())
	jb, err := json.MarshalIndent(b, "", "    ")
	assert.Nil(t, err)
	assert.EqualValues(t, changeBlock, string(jb))
}

func TestBlockDecodeUnknown(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(unknownBlock)))
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestBlockDecodeBad(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(badBlock1)))
	assert.Nil(t, b)
	assert.NotNil(t, err)
}
func TestBlockDecodeBad2(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(badBlock2)))
	assert.Nil(t, b)
	assert.NotNil(t, err)
}
func TestBlockDecodeBad3(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(badBlock3)))
	assert.Nil(t, b)
	assert.NotNil(t, err)
}

func TestBlockHashSend(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(sendBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	var h BlockHash
	b.Hash(&h)
	assert.EqualValues(t, "eccb8cb65cd3106eda8ce9aa893fead497a91bca903890cbd7a5c59f06ab9113", h.String())
}

func TestBlockHashOpen(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(openBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	var h BlockHash
	b.Hash(&h)
	assert.EqualValues(t, "8f02d66117cac96ad0c66db2dd583f8452d1cce979faea5c72e4937f33f4ada4", h.String())
}

func TestBlockHashReceive(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(receiveBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	var h BlockHash
	b.Hash(&h)
	assert.EqualValues(t, "3a753e4b6ab82e5dba277eb1a38166e80c0c001f02f0aae30c228fd39fe21609", h.String())
}

func TestBlockHashChange(t *testing.T) {
	b, err := DecodeBlockJSON(bytes.NewReader([]byte(changeBlock)))
	assert.NotNil(t, b)
	assert.Nil(t, err)
	var h BlockHash
	b.Hash(&h)
	assert.EqualValues(t, "dab6181b7d9813d4c387dc2cdff80a27aac56f8c6e14b5a7a84cc816449640f1", h.String())
}
