package models

import (
	"log"

	"ybuilds.in/codesnippet-api/database"
)

type User struct {
	UserId       int64
	UserName     string `binding:"required"`
	UserEmail    string `binding:"required"`
	UserPassword string `binding:"required"`
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

func GetUserByUserid(userid int) (*User, error) {
	user := &User{}
	db := database.DB
	query := `SELECT userid, username, userpassword, useremail FROM users WHERE userid=$1`

	err := db.QueryRow(query, userid).Scan(&user.UserId, &user.UserName, &user.UserPassword, &user.UserEmail)
	if err != nil {
		log.Println("user with userid", userid, "not found")
		return nil, err
	}

	return user, nil
}

func (user *User) UpdateUserByUserid(userid int) (*User, error) {
	updatedUser := &User{}
	db := database.DB

	query := `UPDATE users SET (username, userpassword, useremail) = ($1, $2, $3) WHERE userid = $4 RETURNING userid, username, userpassword, useremail`

	oldUser, err := GetUserByUserid(userid)
	if err != nil {
		log.Println("user with userid", userid, "not found")
		return nil, err
	}

	//Validations
	user.UserId = int64(userid)
	if user.UserName == "" {
		user.UserName = oldUser.UserName
	}
	if user.UserPassword == "" {
		user.UserPassword = oldUser.UserPassword
	}
	if user.UserEmail == "" {
		user.UserEmail = oldUser.UserEmail
	}

	err = db.QueryRow(query, user.UserName, user.UserPassword, user.UserEmail, userid).Scan(&updatedUser.UserId, &updatedUser.UserName, &updatedUser.UserPassword, &updatedUser.UserEmail)
	if err != nil {
		log.Println("user with userid", userid, "not found")
		return nil, err
	}

	return updatedUser, nil
}

func DeleteUserByUserid(userid int) (int, error) {
	db := database.DB

	query := `DELETE FROM users WHERE userid=$1`

	res, err := db.Exec(query, userid)

	if err != nil {
		log.Println("user with userid", userid, "not deleted")
		return 0, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		log.Println("error getting number of affected rows")
		return 0, err
	}

	return int(rows), nil
}
