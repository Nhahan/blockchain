package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	newTemplate, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	newTemplate.Execute()
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
