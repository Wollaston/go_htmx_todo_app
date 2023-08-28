package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
