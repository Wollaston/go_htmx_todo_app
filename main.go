package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Title  string
	Detail string
}

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/static/output.css", ServeStyleSheet)
	http.HandleFunc("/clicked", ClickHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/public/index.html"))
	tmpl.Execute(w, nil)
	ConnectToDB()
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddItemHandler")
}

func ServeStyleSheet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/output.css")
}

func ClickHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClickHandler")
}

func ConnectToDB() {
	db, _ := sql.Open("sqlite3", "./todos.db")
	rows, _ := db.Query("SELECT * FROM todos")

	fmt.Println(rows)

	db.Close()
}