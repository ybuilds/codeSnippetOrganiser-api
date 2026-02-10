package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"ybuilds.in/codesnippet-api/handlers"
)

func main() {
	server := gin.Default()

	//Health check end point
	server.GET("/health-check", handlers.HealthCheck)

	//User end point
	server.GET("/users", handlers.GetUsers)
	server.GET("/users/:userid", handlers.GetUserByUserid)
	server.POST("/users", handlers.AddUser)
	server.PUT("/users/:userid", handlers.UpdateUserByUserid)
	server.DELETE("/users/:userid", handlers.DeleteUserByUserid)

	//Snippet end point

	err := server.Run(":8000")

	if err != nil {
		log.Fatalln("error starting server", err)
		return
	}
}
