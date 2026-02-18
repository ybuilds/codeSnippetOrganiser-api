package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ybuilds.in/codesnippet-api/util"
)

func Authenticate(ctx *gin.Context) {
	jwt := ctx.Request.Header.Get("Authorization")
	if jwt == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no jwt provided"})
		return
	}

	userId, err := util.VerifyToken(jwt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
