package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users from database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByUserid(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userid"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user id"})
		return
	}

	user, err := models.GetUserByUserid(int(userId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserByUserid(context *gin.Context) {
	var user models.User
	userId, err := strconv.ParseInt(context.Param("userid"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user id"})
		return
	}

	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body for update"})
		return
	}

	updatedUser, err := user.UpdateUserByUserid(int(userId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error updating user"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": updatedUser, "message": "user updated"})
}

func DeleteUserByUserid(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userid"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user id"})
		return
	}

	rows, err := models.DeleteUserByUserid(int(userId))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "error deleting user from database"})
		return
	}

	if rows == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with id %d not found", userId)})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user with id %d deleted", userId)})
}
