package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		panic("error loading environment variables")
	}

	var DBHOST = os.Getenv("DB_HOST")
	var DBUSER = os.Getenv("DB_USER")
	var DBPASS = os.Getenv("DB_PASS")
	var DBNAME = os.Getenv("DB_NAME")
	var DBPORT = os.Getenv("DB_PORT")

	connString := fmt.Sprintf("host=%v user=%v password=%v port=%v dbname=%v sslmode=disable", DBHOST, DBUSER, DBPASS, DBPORT, DBNAME)

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		panic("unable to connect to database")
	}

	err = DB.Ping()
	if err != nil {
		panic("connection not stable to database")
	}
}
