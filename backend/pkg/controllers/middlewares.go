package controllers

import (
	"net/http"

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
