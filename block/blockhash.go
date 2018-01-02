package gorai

type BlockHash struct {
	Digest [32]byte
}

func (h *BlockHash) MarshalText() ([]byte, error) {
	return encodeHex(h.Digest[:]), nil
}

func (h *BlockHash) UnmarshalText(b []byte) error {
	return decodeHex(h.Digest[:], b)
}

func (h *BlockHash) String() string {
	return string(encodeHex(h.Digest[:]))
}

func (h *BlockHash) Bytes() []byte {
	return h.Digest[:]
}

func (h *BlockHash) Hash(parts ...[]byte) {
	hashBlake2b(h.Digest[:], parts...)
}

func NewBlockHash(parts ...[]byte) *BlockHash {
	bh := &BlockHash{}
	bh.Hash(parts...)
	return bh
}

func DecodeBlockHash(text string) (*BlockHash, error) {
	bh := &BlockHash{}
	err := bh.UnmarshalText([]byte(text))
	return bh, err
}
