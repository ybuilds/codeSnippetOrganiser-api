package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/models"
)

func GetUser(ctx *gin.Context) {
	userid, err := strconv.ParseInt(ctx.Param("userid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user ID"})
		return
	}

	user, err := models.GetUser(userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user from user model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(ctx *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching users from user model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func AddUser(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	err = user.AddUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error saving user to database from model"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("user with id %d created", user.Id)})
}

func UpdateUser(ctx *gin.Context) {
	userid, err := strconv.ParseInt(ctx.Param("userid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user ID"})
		return
	}

	_, err = models.GetUser(userid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", userid)})
		return
	}

	var updatedUser models.User
	err = ctx.ShouldBindJSON(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	updatedUser.Id = userid

	err = updatedUser.UpdateUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error updating user in database from model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": updatedUser})
}

func DeleteUser(ctx *gin.Context) {
	userid, err := strconv.ParseInt(ctx.Param("userid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing user ID"})
		return
	}

	user, err := models.GetUser(userid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user with id %d not found", userid)})
		return
	}

	err = user.DeleteUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting user from database from model"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user with id %d deleted", user.Id)})
}

func ValidateUser(ctx *gin.Context) {
	type authUser struct {
		Email    string
		Password string
	}

	var user authUser

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing request body"})
		return
	}

	err = models.ValidateUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user authenticated successfully"})
}
