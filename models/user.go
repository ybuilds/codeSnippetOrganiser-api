package models

import (
	"log"

	"ybuilds.in/codesnippet-api/database"
)

type User struct {
	UserId    int64
	UserName  string
	UserEmail string
}

func (user *User) Save() error {
	db := database.DB

	query := `INSERT INTO USERS(username, useremail) VALUES($1, $2) RETURNING userid`

	err := db.QueryRow(query, user.UserName, user.UserEmail).Scan(&user.UserId)

	if err != nil {
		log.Println("error saving user to database", err)
		return err
	}

	return nil
}
