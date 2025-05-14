package blockchain

type BlockChain struct {
	Blocks []*Block // keep it public for kow
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{NewBlock("Genesis Block", []byte{})},
	}
}

func (bc *BlockChain) Add(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}
