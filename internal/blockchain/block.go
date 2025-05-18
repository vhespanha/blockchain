package blockchain

import (
	"errors"
	"time"
)

type Block struct {
	Data      []byte
	PrevHash  []byte
	Hash      []byte
	Nonce     int
	Timestamp int64
}

func NewBlock(data string, prevHash []byte) (*Block, error) {
	b := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
	}

	return b.mineBlock()
}

func (b *Block) mineBlock() (*Block, error) {
	c := NewProofOfWork(b).Run()

	result, ok := <-c
	if !ok {
		return nil, errors.New("no valid proof found")
	}

	b.Nonce = result.Nonce
	b.Hash = result.Hash

	return b, nil
}
