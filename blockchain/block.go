package blockchain

import "crypto/sha256"

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