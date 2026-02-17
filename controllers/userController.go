package controllers

import (
	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/handlers"
)

func UserController(server *gin.Engine) {
	server.GET("/users", handlers.GetUsers)
	server.GET("/users/:userid", handlers.GetUser)

	server.POST("/users/signup", handlers.AddUser)
	server.POST("/users/login", handlers.ValidateUser)

	server.PUT("/users/:userid", handlers.UpdateUser)

	server.DELETE("/users/:userid", handlers.DeleteUser)
}
