package models

import (
	"log"

	"ybuilds.in/codesnippet-api/database"
)

type User struct {
	UserId       int64
	UserName     string
	UserEmail    string
	UserPassword string
}

func (user *User) Save() error {
	db := database.DB

	query := `INSERT INTO USERS(username, userpassword, useremail) VALUES($1, $2, $3) RETURNING userid`

	err := db.QueryRow(query, user.UserName, user.UserPassword, user.UserEmail).Scan(&user.UserId)

	if err != nil {
		log.Println("error saving user to database", err)
		return err
	}

	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	db := database.DB

	query := `SELECT userid, username, userpassword, useremail FROM users`

	res, err := db.Query(query)

	if err != nil {
		log.Println("error fetching users from database")
		return nil, err
	}

	for res.Next() {
		var user User
		err := res.Scan(&user.UserId, &user.UserName, &user.UserPassword, &user.UserEmail)

		if err != nil {
			log.Println("error parsing users from database")
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
