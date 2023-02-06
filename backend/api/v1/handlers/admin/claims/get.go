package claims

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func getAllClaims(ctx *gin.Context) {
	claims, err := models.GetAllClaims()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, claims)
}
