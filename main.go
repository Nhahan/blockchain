package main

import "crypto/sha256"

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	genesisBlock.hash = sha256.Sum256([]byte(genesisBlock.data + genesisBlock.hash))
}
