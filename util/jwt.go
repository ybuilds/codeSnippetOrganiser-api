package util

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int64, email string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	unsignedString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := unsignedString.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("error generating JWT")
		return "", err
	}

	return token, nil
}

func VerifyToken(userToken string) (int64, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.Parse(userToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("signing method of JWT does not match")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return -1, err
	}

	if !parsedToken.Valid {
		return -1, errors.New("JWT is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("invalid claims")
	}

	userIdValue, ok := claims["userId"]
	if !ok {
		return -1, errors.New("userId not found in claims")
	}

	userId, ok := userIdValue.(float64)
	if !ok {
		return -1, errors.New("error converting userId from claims to int")
	}

	return int64(userId), nil
}
