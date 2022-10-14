package middlewares

import (
	"net/http"

	"casino.website/pkg/models"
	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("User").(*models.Client)

		if user.AccessType < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not administrator."})
		}
	}
}
