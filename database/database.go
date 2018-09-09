package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func NewConnection() *sql.DB {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func Migration() {

	db := NewConnection()
	defer db.Close()

	var q = ""

    q = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT (DATETIME('now','localtime')),
		update_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))
	)
	`
	_, err := db.Exec(q)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
