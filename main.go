package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func main() {
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // if there is an error, it will return it.
}
