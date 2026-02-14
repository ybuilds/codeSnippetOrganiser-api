package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/models"
)

func GetSnippet(ctx *gin.Context) {
	snippetid, err := strconv.ParseInt(ctx.Param("snippetid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing snippet ID"})
		return
	}

	snippet, err := models.GetSnippet(snippetid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching snippet from snippet model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"snippet": snippet})
}

func GetSnippets(ctx *gin.Context) {
	snippets, err := models.GetSnippets()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching snippets from user model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": snippets})
}

func AddSnippet(ctx *gin.Context) {
	var snippet models.Snippet

	err := ctx.ShouldBindJSON(&snippet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	err = snippet.AddSnippet()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error saving snippet to database from model"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("snippet with id %d created", snippet.Id)})
}

func UpdateSnippet(ctx *gin.Context) {
	snippetid, err := strconv.ParseInt(ctx.Param("snippetid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing snippet ID"})
		return
	}

	_, err = models.GetUser(snippetid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", snippetid)})
		return
	}

	var updatedSnippet models.Snippet
	err = ctx.ShouldBindJSON(&updatedSnippet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	updatedSnippet.Id = snippetid

	err = updatedSnippet.UpdateSnippet()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error updating snippet in database from model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": updatedSnippet})
}

func DeleteSnippet(ctx *gin.Context) {
	snippetid, err := strconv.ParseInt(ctx.Param("snippetid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing snippet ID"})
		return
	}

	snippet, err := models.GetSnippet(snippetid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("snippet with id %d not found", snippetid)})
		return
	}

	err = snippet.DeleteSnippet()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting snippet from database from model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("snippet with id %d deleted", snippet.Id)})
}
