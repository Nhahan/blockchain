package main

import (
	"fmt"

	"github.com/Nhahan/blockchain/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block.Data)
	}
}
