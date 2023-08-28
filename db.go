package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Record struct {
	Title  string
	Detail string
	Uid int
	Created time.Time
}

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	return db
}

func Read(db *sql.DB) []Record {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	var records []Record

	for rows.Next() {
		var uid int
		var title string
		var detail string
		var created time.Time
		err = rows.Scan(&uid, &title, &detail, &created)
		if err != nil {
			log.Fatal(err)
		}
		record := Record{Title: title, Detail: detail, Uid: uid, Created: created}
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

func DeleteOne(uid string, db *sql.DB) {
	stmt, err := db.Prepare("DELETE FROM todos WHERE uid=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(uid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Delete: ")
	fmt.Println(res)
}
