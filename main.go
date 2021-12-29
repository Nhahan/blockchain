package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func documentation(rw http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
