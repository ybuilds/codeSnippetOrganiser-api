package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"ybuilds.in/codesnippet-api/models"
)

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "system-ok"})
}

func addUser(context *gin.Context) {
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

func main() {
	server := gin.Default()

	server.GET("/health-check", healthCheck)

	server.POST("/users", addUser)

	err := server.Run(":8000")

	if err != nil {
		log.Fatalln("error starting server", err)
		return
	}
}
