package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/Nhahan/blockchain/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// Must is the function that if there's an error returns error, otherwise returns value
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	newTemplate.Execute(rw, data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
