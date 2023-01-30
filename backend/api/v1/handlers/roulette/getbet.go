package roulette

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBet(ctx *gin.Context) {
	bets := Roulette.GetBets()

	ctx.JSON(http.StatusOK, bets)
}
