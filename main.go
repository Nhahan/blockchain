package main

import (
	"github.com/Nhahan/blockchain//explorer"
	"github.com/Nhahan/blockchain/rest"
)

func main() {
	explorer.Start(3000)
	rest.Start(4000)
}
