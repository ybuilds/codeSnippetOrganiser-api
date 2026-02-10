package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/models"
)

func AddUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "db-error: user not created"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": user, "message": "user created"})
}

func GetUsers(context *gin.Context) {
	users, err := models.GetUsers()

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users from database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}
