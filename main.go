package main

import (
	"github.com/Nhahan/blockchain/blockchain"
)

func main() {
	blockchain.Blockchain().AddBlock("1")
	blockchain.Blockchain().AddBlock("2")
}
