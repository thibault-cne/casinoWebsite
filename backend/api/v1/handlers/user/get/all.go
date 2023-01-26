package get

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func all(ctx *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
