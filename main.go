package main

import (
	"fmt"
	"github.com/AuthorOfTheSurf/go-naive-chain/blockchain"
)

func main() {
	goNaiveChain := blockchain.GetBlockChain()
	fmt.Printf("%s %#v", goNaiveChain.ChainName(), goNaiveChain.Blocks())
}