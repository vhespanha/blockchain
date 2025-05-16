package main

import (
	"fmt"
	"strconv"

	"github.com/vhespanha/blockchain/internal/blockchain"
)

func main() {
	bc, err := blockchain.NewBlockChain()
	if err != nil {
		panic(fmt.Errorf("could not initialize blockchain: %w", err))
	}

	if err := bc.Add("Send 1 Coin to Vinicius"); err != nil {
		panic(err)
	}

	if err := bc.Add("Send 4 more Coins to Vinicius"); err != nil {
		panic(err)
	}

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := blockchain.NewProofOfWork(&block.MetaBlock)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate(block.Nonce)))
		fmt.Println()
	}
}
