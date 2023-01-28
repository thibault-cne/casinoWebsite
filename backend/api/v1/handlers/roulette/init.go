package roulette

import (
	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

var (
	Roulette *models.Roulette
)

func LoadRoulette(rg *gin.RouterGroup) {
	Roulette = models.NewRoulette()

	// routerGroup := rg.Group("/roulette")

	Roulette.Start()
}

func ListenEndpoint() {
	gosf.Listen("bet", bet)
}
