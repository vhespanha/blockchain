package blockchain

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{NewBlock("Genesis Block", []byte{})},
	}
}

func (bc *BlockChain) Add(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
