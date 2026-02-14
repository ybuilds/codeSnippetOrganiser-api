package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func init() {
	var err error

	connString := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable", "localhost", "postgres", "kjm40438", 5432, "codesnippetApi")

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		panic("unable to connect to database")
	}

	err = DB.Ping()
	if err != nil {
		panic("connection not stable to database")
	}
}
