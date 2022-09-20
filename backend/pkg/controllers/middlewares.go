package controllers

import (
	"net/http"

	clientservices "casino.website/pkg/services/client.services"
	oauthservices "casino.website/pkg/services/oauth.services"
	"github.com/gin-gonic/gin"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loggedInInterface, _ := ctx.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
		}
	}
}

func ensureIsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdInterface, _ := ctx.Get("user_id")
		userId := userIdInterface.(int)

		if !clientservices.CheckClientModerationType(userId, false) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not an administrator."})
		}
	}
}

func ensureIsSuperAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdInterface, _ := ctx.Get("user_id")
		userId := userIdInterface.(int)

		if !clientservices.CheckClientModerationType(userId, true) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not an administrator."})
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.Request.Header.Get("Authorization")
		claims, err := oauthservices.RetrieveUserClaims(reqToken)

		if err != nil {
			ctx.Set("is_logged_in", false)
			return
		}

		ctx.Set("is_logged_in", true)
		ctx.Set("user_id", claims.User_id)
	}
}
