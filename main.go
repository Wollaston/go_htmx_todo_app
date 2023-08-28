package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Title  string
	Detail string
}

func main() {
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/static/output.css", ServeStyleSheet)
	http.HandleFunc("/create-todo", ToDoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/public/templates/index.html", "./src/public/templates/todo.html"))
	db := Connect()
	records := Read(db)
	tmpl.Execute(w, records)
	Close(db)
}

func ServeStyleSheet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/output.css")
}

func ToDoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db := Connect()
	Create(db, r.FormValue("title"), r.FormValue("detail"))

	records := Read(db)
	fmt.Println(records)

	tmpl := template.Must(template.ParseFiles("./src/public/templates/todo.html"))
	tmpl.ExecuteTemplate(w, "todo", records)

	Close(db)
}
