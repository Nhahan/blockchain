package main

import (
	"github.com/Nhahan/blockchain/cli"
	"github.com/Nhahan/blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
