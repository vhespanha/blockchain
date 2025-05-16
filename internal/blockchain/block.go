package blockchain

import (
	"fmt"
	"time"
)

type MetaBlock struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
}

type MinedBlock struct {
	MetaBlock
	Nonce int
	Hash  []byte
}

var (
	ErrMaxNonce = fmt.Errorf("no valid proof found")
)

func NewBlock(data string, prevHash []byte) (*MinedBlock, error) {
	return MineBlock(NewMetaBlock(data, prevHash))
}

func NewMetaBlock(data string, prevHash []byte) *MetaBlock {
	return &MetaBlock{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
	}
}

func MineBlock(b *MetaBlock) (*MinedBlock, error) {
	p := NewProofOfWork(b)
	c := p.Run()

	result, ok := <-c
	if !ok {
		return nil, ErrMaxNonce
	}

	return &MinedBlock{
		MetaBlock: *b,
		Nonce:     result.Nonce,
		Hash:      result.Hash,
	}, nil
}
