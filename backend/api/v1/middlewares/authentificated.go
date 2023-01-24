package middlewares

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user")
		if userID == nil {
			// Abort the request with the appropriate error code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Vous devez être connecté pour faire cette action"})
			return
		}

		// Can cast
		switch userID.(type) {
		case string:
		default:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Vous devez être connecté pour faire cette action"})
			return
		}

		// Fill the user in the context
		u, err := models.GetUserByID(userID.(string))
		if err != nil {
			session.Clear()
			session.Save()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Vous devez être connecté pour faire cette action"})
			return
		}
		c.Set("user", u)

		// Continue
		c.Next()
	}
}
