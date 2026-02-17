package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println("error hashing password")
		return "", err
	}

	return string(bytes), nil
}

func ValidatePassword(dbPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	return err == nil
}
