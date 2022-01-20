package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server (default 4000)")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
		fmt.Println("start REST API")
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start Server")
	}
}

func usage() {
	fmt.Printf("Welcome. Please use the following commands:\n")
	fmt.Printf("explorer:   Start the HTML Explorer\n")
	fmt.Printf("rest: 		Start the REST API\n")
	os.Exit(1)
}
