package claims

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func deleteClaim(ctx *gin.Context) {
	code := ctx.Query("code")

	claim, err := models.GetClaim(code)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not find claim"})
		return
	}

	err = claim.Delete()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not delete claim"})
		return
	}

	ctx.JSON(http.StatusOK, claim)
}
