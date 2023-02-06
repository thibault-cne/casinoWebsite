package claims

import (
	"net/http"

	"casino.website/internal/models"
	"casino.website/internal/utils"
	"github.com/gin-gonic/gin"
)

func newClaim(ctx *gin.Context) {
	claim := &models.Claims{
		Code: utils.GenerateClaimsCode(),
		Use:  5,
	}

	err := claim.Create()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, claim)
}
