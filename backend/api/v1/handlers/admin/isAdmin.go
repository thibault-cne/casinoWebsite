package admin

import (
	"github.com/gin-gonic/gin"
)

func isAdmin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": true,
	})
}
