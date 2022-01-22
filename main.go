package main

import (
	"flag"
	"fmt"
	"os"

	explorer "github.com/Nhahan/blockchain/explorer/templates"
	"github.com/Nhahan/blockchain/rest"
)

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "reset", "Choose between  'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)

}

func usage() {
	fmt.Printf("")
	fmt.Printf("Please use the following commands:\n")
	fmt.Printf("-port: Set the PORT of the server\n")
	fmt.Printf("-mode: Choose between 'html' and 'rest'\n")
	os.Exit(0)
}
