package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	data := homeData{"Home"}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
