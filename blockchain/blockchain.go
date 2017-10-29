package blockchain

import (
	"crypto/sha256"
)

type Block struct {
	index        int
	previousHash [sha256.Size]byte
	timestamp    int64
	data         string
	hash         [sha256.Size]byte
}

func NewBlock(
	index int,
	previousHash [sha256.Size]byte,
	timestamp int64,
	data string,
	hash [sha256.Size]byte,
) Block {
	return Block{
		index,
		previousHash,
		timestamp,
		data,
		hash,
	}
}

type BlockChain struct {
	chainName string
	blocks    []Block
}

func (bc *BlockChain) ChainName() string {
	return bc.chainName
}

func (bc *BlockChain) Blocks() []Block {
	return bc.blocks
}

func NewBlockChain(chainName string, genesisBlock Block) BlockChain {
	return BlockChain{
		chainName,
		[]Block{genesisBlock},
	}
}
