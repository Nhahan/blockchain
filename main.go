package main

import (
	"fmt"
	"os"
)

func main() {
	// rest.Start(4000)
	if len(os.Args) < 2 {
		fmt.Printf("Welcome. Please use the following commands:\n")
		fmt.Printf("explorer:   Start the HTML Explorer\n")
		fmt.Printf("rest: 		Start the REST API\n")
	}
}
