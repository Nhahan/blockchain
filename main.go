package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

}

func usage() {
	fmt.Printf("Welcome. Please use the following commands:\n")
	fmt.Printf("explorer:   Start the HTML Explorer\n")
	fmt.Printf("rest: 		Start the REST API\n")
	os.Exit(1)
}
