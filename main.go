package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	pageTitle string
}

func home(rw http.ResponseWriter, r *http.Request) {
	newTemplate, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	data := homeData{"Home"}
	newTemplate.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
