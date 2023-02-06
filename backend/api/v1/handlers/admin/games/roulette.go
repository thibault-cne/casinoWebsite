package games

import (
	"net/http"

	"casino.website/api/v1/handlers/roulette"
	"github.com/gin-gonic/gin"
)

func toggleRoulette(ctx *gin.Context) {
	roulette.Roulette.Toggle()

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
