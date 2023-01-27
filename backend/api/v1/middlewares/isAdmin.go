package middlewares

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*models.User)

		if user.Status == "user" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Vous devez être administrateur pour faire cette action"})
			return
		}

		ctx.Next()
	}
}
