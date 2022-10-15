package roulette

import (
	"casino.website/pkg/models"
	middlewares "casino.website/pkg/routes/middleware"
	"github.com/gin-gonic/gin"
)

func RouletteGameRoutes(rg *gin.RouterGroup) {
	rGame := models.NewRouletteGame()

	subpath := rg.Group("/roulette")
	subpath.GET("/connect", middlewares.UserStatus(), func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})
	subpath.GET("/ws/", middlewares.UserStatus(), Wrapper())

	go rGame.Start()
	go rGame.End()
}
