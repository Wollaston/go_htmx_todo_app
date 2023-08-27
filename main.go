package main

import (
	"fmt"
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
		fmt.Println("DefaultHandler")
	}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddItemHandler")
}