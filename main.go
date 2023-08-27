package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title	string
	Detail	string
}

func main() {
	http.HandleFunc("/", DefaultHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./public/index.html"))
		tmpl.Execute(w, nil)
	}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddItemHandler")
}