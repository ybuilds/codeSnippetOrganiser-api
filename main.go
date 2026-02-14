package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"ybuilds.in/codesnippet-api/controllers"
	"ybuilds.in/codesnippet-api/handlers"
)

func main() {
	server := gin.Default()

	//health check only
	server.GET("/health-check", handlers.HealthCheck)

	//user controller
	controllers.UserController(server)

	err := http.ListenAndServe(":8000", server)
	if err != nil {
		panic("unable to start server")
	}
}
