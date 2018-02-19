package blockchain

type BlockChain struct {
	chainName string
	blocks    []Block
}

func (bc *BlockChain) ChainName() string {
	return bc.chainName
}

func (bc *BlockChain) GetLatestBlock() *Block {
	return &bc.blocks[len(bc.blocks) - 1]
}

func NewBlockChain(chainName string, genesisBlock Block) BlockChain {
	return BlockChain{
		chainName,
		[]Block{genesisBlock},
	}
}
