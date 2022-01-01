package main

import (
	explorer "github.com/Nhahan/blockchain/explorer/templates"
	"github.com/Nhahan/blockchain/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
