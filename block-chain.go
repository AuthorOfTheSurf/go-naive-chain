package go_naive_chain

import "time"
import (
	"crypto/sha256"
)

type Block struct {
	index        int
	previousHash [sha256.Size]byte
	timestamp    int64
	data        string
	hash         [sha256.Size]byte
}

type BlockChain struct {
	blocks          []Block
}

func (b *BlockChain) Blocks() []Block {
	return b.blocks
}

var genesisBlock = Block{
	index:        0,
	previousHash: sha256.Sum256([]byte("0")),
	timestamp:    time.Now().Unix(),
	data:         "The freedom to create coins",
	hash:         sha256.Sum256([]byte("Twitter 28/Oct/2017 Dodgers break it open in the ninth inning of Game 4")),
}
var goNaiveChain = BlockChain{
	blocks: []Block{genesisBlock},
}
