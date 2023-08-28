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
	http.HandleFunc("/clicked", ClickHandler)
	http.HandleFunc("/create-todo", ToDoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/public/index.html"))
	tmpl.Execute(w, nil)
	Connect()
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

func ToDoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	db := Connect()
	Create(db, r.FormValue("title"), r.FormValue("detail"))
	Read(db)
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

func Read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Read: ")
	for rows.Next() {
		var uid int
		var title string
		var detail string
		var created time.Time
		err = rows.Scan(&uid, &title, &detail, &created)
		if err != nil {
		 log.Fatal(err)
		}
		fmt.Println(uid)
		fmt.Println(title)
		fmt.Println(detail)
		fmt.Println(created)
	 }
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
