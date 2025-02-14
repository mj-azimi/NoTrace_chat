package config

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {


	db, err := sql.Open("sqlite3", GetConfig("db_name"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
