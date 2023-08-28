package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

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

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	return db
}

func Read(db *sql.DB) []Todo {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	var records []Todo

	for rows.Next() {
		var uid int
		var title string
		var detail string
		var created time.Time
		err = rows.Scan(&uid, &title, &detail, &created)
		if err != nil {
			log.Fatal(err)
		}
		record := Todo{Title: title, Detail: detail}
		records = append(records, record)
	}
	return records
}

func Create(db *sql.DB, title string, detail string) {
	stmt, err := db.Prepare("INSERT INTO todos(title, detail, created) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(title, detail, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Create: ")
	fmt.Println(res)
}

func Close(db *sql.DB) {
	db.Close()
	fmt.Println("DB Connection Closed.")
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
