package get

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func wallet(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	ctx.JSON(http.StatusOK, gin.H{"wallet": user.Wallet})
}
