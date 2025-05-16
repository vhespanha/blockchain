package blockchain

type BlockChain struct {
	Blocks []*MinedBlock // keep it public for kow
}

func NewBlockChain() (*BlockChain, error) {
	genesis, err := NewBlock("Genesis Block", []byte{})

	if err != nil {
		return nil, err
	}

	return &BlockChain{
		Blocks: []*MinedBlock{genesis},
	}, nil
}

func (bc *BlockChain) Add(data string) error {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock, err := NewBlock(data, prevBlock.Hash)

	if err != nil {
		return err
	}

	bc.Blocks = append(bc.Blocks, newBlock)

	return nil
}
