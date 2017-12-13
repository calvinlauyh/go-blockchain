package core

type BlockChain struct {
	Blocks []*Block
}

// NewBlockChain creates and returns the pointer to a newly created BlockChain
// with the Genesis Block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// AppendBlock appends a new Block with the speicifed data to the BlockChain
func (chain *BlockChain) AppendBlock(data string) {
	prevBlk := chain.Blocks[len(chain.Blocks)-1]
	chain.Blocks = append(chain.Blocks, NewBlock(data, prevBlk.Hash))
}
