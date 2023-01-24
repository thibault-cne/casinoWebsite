package auth

import (
	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func Connected(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}
