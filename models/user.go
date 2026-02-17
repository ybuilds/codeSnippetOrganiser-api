package models

import (
	"errors"
	"log"

	"ybuilds.in/codesnippet-api/util"
)

type User struct {
	Id       int64
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func GetUsers() ([]User, error) {
	var users []User
	query := `select id, name, email, password from users`

	res, err := db.Query(query)
	if err != nil {
		log.Println("error fetching users from database")
		return nil, err
	}

	for res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Println("error parsing user data from database")
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUser(userid int64) (*User, error) {
	var user User

	query := `select id, name, email, password from users where id = $1`

	err := db.QueryRow(query, userid).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("error fetching user from database")
		return nil, err
	}

	return &user, nil
}

func (user *User) AddUser() error {
	hashPwd, err := util.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing user password")
		return err
	}

	query := `insert into users(name, email, password) values($1, $2, $3) returning id`

	err = db.QueryRow(query, user.Name, user.Email, hashPwd).Scan(&user.Id)
	if err != nil {
		log.Println("error adding user to database")
		return err
	}

	return nil
}

func (user *User) UpdateUser() error {
	hashPwd, err := util.HashPassword(user.Password)
	if err != nil {
		log.Println("error hashing user password")
		return err
	}

	query := `update users set name=$1, email=$2, password=$3 where id=$4 returning id`

	err = db.QueryRow(query, user.Name, user.Email, hashPwd, user.Id).Scan(&user.Id)
	if err != nil {
		log.Println("error updating user in database")
		return err
	}

	return nil
}

func (user *User) DeleteUser() error {
	query := `delete from users where id=$1`

	_, err := db.Query(query, user.Id)
	if err != nil {
		log.Println("error deleting user from database")
		return err
	}

	return nil
}

func ValidateUser(email, password string) error {
	var dbPassword string
	query := `select password from users where email=$1`

	err := db.QueryRow(query, email).Scan(&dbPassword)
	if err != nil {
		log.Println("error fetching user from database")
		return err
	}

	validate := util.ValidatePassword(dbPassword, password)
	if !validate {
		return errors.New("passwords do not match")
	}

	return nil
}
