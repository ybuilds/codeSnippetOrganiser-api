package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error loading .env file", err)
		return
	}

	hostname := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, dbName)
	DB, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatalln("error creating database connection", err)
		return
	}

	if err = DB.Ping(); err != nil {
		log.Fatalln("error pinging database", err)
		return
	}

	log.Println("database ping successful")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
