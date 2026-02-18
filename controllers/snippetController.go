package controllers

import (
	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/handlers"
	"ybuilds.in/codesnippet-api/middleware"
)

func SnippetController(server *gin.Engine) {
	server.GET("/snippet", handlers.GetSnippets)
	server.GET("/snippet/:snippetid", handlers.GetSnippet)
	server.GET("/snippet/category/:category", handlers.GetSnippetByCategory)
	server.GET("/snippet/language/:language", handlers.GetSnippetByLanguage)

	server.POST("/snippet", middleware.Authenticate, handlers.AddSnippet)

	server.PUT("/snippet/:snippetid", middleware.Authenticate, handlers.UpdateSnippet)

	server.DELETE("/snippet/:snippetid", middleware.Authenticate, handlers.DeleteSnippet)
}
