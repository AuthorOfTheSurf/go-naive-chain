package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	bc "github.com/AuthorOfTheSurf/go-naive-chain/blockchain"
)

var genesisBlock = bc.NewBlock(
	0,
	sha256.Sum256([]byte("0")),
	time.Now().Unix(),
	"The freedom to create coins",
	sha256.Sum256([]byte("Twitter 28/Oct/2017 Dodgers break it open in the ninth inning of Game 4")),
)

func main() {
	goNaiveChain := bc.NewBlockChain("Go Naive Chain", genesisBlock)
	fmt.Printf("%s %#v", goNaiveChain.ChainName(), goNaiveChain.GetLatestBlock())
}
