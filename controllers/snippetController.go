package controllers

import (
	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/handlers"
)

func SnippetController(server *gin.Engine) {
	server.GET("/snippet", handlers.GetSnippets)
	server.GET("/snippet/:snippetid", handlers.GetSnippet)
	server.POST("/snippet", handlers.AddSnippet)
	server.PUT("/snippet/:snippetid", handlers.UpdateSnippet)
	server.DELETE("/snippet/:snippetid", handlers.DeleteSnippet)
}
