package internal

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() error {

	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id blob not null primary key, name text, completed boolean default false)")
	if err != nil {
		return err
	}
	return db.Close()
}

func ConnectDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "todos.db")
}
