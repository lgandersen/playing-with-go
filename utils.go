package main

import (
	"database/sql"
)

func create_database(filename string) *sql.DB {
	createTable := `
	CREATE TABLE IF NOT EXISTS files(
		path TEXT,
		name TEXT
	);

	CREATE TABLE IF NOT EXISTS imdb_pages (
		file_id INTEGER,
		body BLOB,
		FOREIGN KEY(file_id) REFERENCES files(rowid)
	);
	`
	db, err := sql.Open("sqlite3", filename)
	checkError(err)

	_, err = db.Exec(createTable)
	checkError(err)
	return db
}

func open_database(filename string) *sql.DB {
	db, err := sql.Open("sqlite3", filename)
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
