package get

import (
	"fmt"
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func bet(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	gameType := ctx.Query("game")

	if gameType == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing game querystring"})
		return
	}

	res := user.GetUserStat(gameType)

	fmt.Printf("%+v\n", res)

	ctx.JSON(http.StatusOK, res)
}
