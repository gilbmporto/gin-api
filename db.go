package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		fmt.Printf("Error opening database: %s\n", err.Error())
	}

	if db == nil {
		log.Fatal("Could not connect to the database")
	}

	createTable(db)
	return db

}

func createTable(db *sql.DB) {
	sqlTable := `
	CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL
  );`

	_, err := db.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}
}
