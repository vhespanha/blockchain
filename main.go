package main

import (
	"fmt"
	"strconv"

	"github.com/vhespanha/blockchain/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()

	bc.Add("Send 1 Coin to Vinicius")
	bc.Add("Send 4 more Coins to Vinicius")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
