package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

type Result struct {
	Nonce int
	Hash  []byte
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TargetBits))

	return &ProofOfWork{b, target}
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			p.block.PrevHash,
			p.block.Data,
			intToHex(p.block.Timestamp),
			intToHex(int64(TargetBits)),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)
}

func (p *ProofOfWork) Run() <-chan Result {
	c := make(chan Result, 1)

	go func() {
		defer close(c)

		var hashInt big.Int

		fmt.Printf("Mining the block containing \"%s\"\n", p.block.Data)

		for nonce := range MaxNonce {
			data := p.prepareData(nonce)
			hash := sha256.Sum256(data)

			hashInt.SetBytes(hash[:])

			if hashInt.Cmp(p.target) == -1 {
				c <- Result{Nonce: nonce, Hash: hash[:]}
				return
			}
		}

		fmt.Println()
	}()

	return c
}

func (p *ProofOfWork) Validate(nonce int) bool {
	var hashInt big.Int

	data := p.prepareData(nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(p.target) == -1
}

func intToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
