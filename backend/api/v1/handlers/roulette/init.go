package roulette

import (
	"casino.website/api/v1/middlewares"
	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func HandleRouletteGame(rg *gin.RouterGroup) {
	rGame := models.NewRouletteGame()

	routerGroup := rg.Group("/roulette")
	routerGroup.GET("/connect", middlewares.AuthRequired(), func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})

	go rGame.Start()
	go rGame.End()
}
