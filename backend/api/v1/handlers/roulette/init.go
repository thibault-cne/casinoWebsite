package roulette

import (
	"casino.website/api/v1/middlewares"
	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

var (
	Engine   *gin.Engine
	Roulette *models.Roulette
)

func LoadRoulette(rg *gin.RouterGroup, engine *gin.Engine) {
	Roulette = models.NewRoulette()

	Engine = engine

	routerGroup := rg.Group("/roulette")
	routerGroup.Any("/ws/", middlewares.AuthRequired(), Wrapper())

	Roulette.Start()
}
