package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

func main() {
	fmt.Printf("Listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
