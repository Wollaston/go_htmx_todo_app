package main

import (
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
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
	tmpl := template.Must(template.ParseFiles("./src/public/index.html"))
	tmpl.Execute(w, nil)
	CreateTodos(w, r)
	Connect()
}

func ServeStyleSheet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/output.css")
}

func ToDoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db := Connect()
	Create(db, r.FormValue("title"), r.FormValue("detail"))

	tmpl := template.Must(template.ParseFiles("./src/public/index.html"))
	tmpl.ExecuteTemplate(w, "todo-list-element", Todo{Title: r.FormValue("title"), Detail: r.FormValue("detail")})

	Close(db)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	db := Connect()

	records := Read(db)
	tmpl := template.Must(template.ParseFiles("./src/public/index.html"))

	for i := 0; i < len(records); i++ {
		tmpl.ExecuteTemplate(w, "todo-list-element", Todo{Title: records[i].Title, Detail: records[i].Detail})
	}

	Close(db)
}
