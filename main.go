package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("hashtest"))
	fmt.Println("%x\n", hash)
}
