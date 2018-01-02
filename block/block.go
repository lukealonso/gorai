package gorai

import (
	"encoding/json"
	"fmt"
	"io"
)

type BlockType byte

const (
	TypeOpen BlockType = iota
	TypeReceive
	TypeSend
	TypeChange
)

type Block interface {
	Hash(*BlockHash)
	Previous() *BlockHash
	Signature() *Signature
	Work() *Work
}

type BlockWorkSig struct {
	Wrk Work      `json:"work"`
	Sig Signature `json:"signature"`
}

func (b *BlockWorkSig) Signature() *Signature {
	return &b.Sig
}

func (b *BlockWorkSig) Work() *Work {
	return &b.Wrk
}

type OpenBlock struct {
	Source         BlockHash `json:"source"`
	Representative Account   `json:"representative"`
	Account        Account   `json:"account"`
	BlockWorkSig
}

func (b *OpenBlock) Hash(bh *BlockHash) {
	bh.Hash(b.Source.Bytes(), b.Representative.Bytes(), b.Account.Bytes())
}

func (b *OpenBlock) Previous() *BlockHash {
	return nil
}

func (b *OpenBlock) MarshalJSON() ([]byte, error) {
	type OpenBlockMarshal OpenBlock
	return json.Marshal(&struct {
		BlockType string `json:"type"`
		OpenBlockMarshal
	}{
		BlockType:        "open",
		OpenBlockMarshal: OpenBlockMarshal(*b),
	})
}

type ReceiveBlock struct {
	Prev BlockHash `json:"previous"`
	Src  BlockHash `json:"source"`
	BlockWorkSig
}

func (b *ReceiveBlock) Previous() *BlockHash {
	return &b.Prev
}

func (b *ReceiveBlock) Hash(bh *BlockHash) {
	bh.Hash(b.Prev.Bytes(), b.Src.Bytes())
}

func (b *ReceiveBlock) MarshalJSON() ([]byte, error) {
	type ReceiveBlockMarshal ReceiveBlock
	return json.Marshal(&struct {
		BlockType string `json:"type"`
		ReceiveBlockMarshal
	}{
		BlockType:           "receive",
		ReceiveBlockMarshal: ReceiveBlockMarshal(*b),
	})
}

type SendBlock struct {
	Prev    BlockHash `json:"previous"`
	Dest    Account   `json:"destination"`
	Balance Amount    `json:"balance"`
	BlockWorkSig
}

func (b *SendBlock) Previous() *BlockHash {
	return &b.Prev
}

func (b *SendBlock) Hash(bh *BlockHash) {
	bh.Hash(b.Prev.Bytes(), b.Dest.Bytes(), b.Balance.Bytes())
}

func (b *SendBlock) MarshalJSON() ([]byte, error) {
	type SendBlockMarshal SendBlock
	return json.Marshal(&struct {
		BlockType string `json:"type"`
		SendBlockMarshal
	}{
		BlockType:        "send",
		SendBlockMarshal: SendBlockMarshal(*b),
	})
}

type ChangeBlock struct {
	Prev BlockHash `json:"previous"`
	Rep  Account   `json:"representative"`
	BlockWorkSig
}

func (b *ChangeBlock) Previous() *BlockHash {
	return &b.Prev
}

func (b *ChangeBlock) Hash(bh *BlockHash) {
	bh.Hash(b.Prev.Bytes(), b.Rep.Bytes())
}

func (b *ChangeBlock) MarshalJSON() ([]byte, error) {
	type ChangeBlockMarshal ChangeBlock
	return json.Marshal(&struct {
		BlockType string `json:"type"`
		ChangeBlockMarshal
	}{
		BlockType:          "change",
		ChangeBlockMarshal: ChangeBlockMarshal(*b),
	})
}

func newBlockFromType(t string) (Block, error) {
	switch t {
	case "open":
		return &OpenBlock{}, nil
	case "send":
		return &SendBlock{}, nil
	case "change":
		return &ChangeBlock{}, nil
	case "receive":
		return &ReceiveBlock{}, nil
	default:
		return nil, fmt.Errorf("unknown block type '%s'", t)
	}
}

type blockIdent struct {
	Type string `json:"type"`
}

func DecodeBlockJSON(r io.Reader) (Block, error) {
	jd := json.NewDecoder(r)

	var raw json.RawMessage
	if err := jd.Decode(&raw); err != nil {
		return nil, fmt.Errorf("error parsing block JSON: %s", err)
	}

	var bi blockIdent
	if err := json.Unmarshal(raw, &bi); err != nil {
		return nil, fmt.Errorf("error parsing block type: %s", err)
	}

	block, err := newBlockFromType(bi.Type)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(raw, block); err != nil {
		return nil, fmt.Errorf("error parsing block type: %s", err)
	}

	return block, nil
}
