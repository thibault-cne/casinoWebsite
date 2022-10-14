package routes

import (
	"casino.website/pkg/routes/auth"
	"casino.website/pkg/routes/roulette"
	"github.com/gin-gonic/gin"
)

func Register(path *gin.RouterGroup) {
	auth.AuthRoutes(path)
	roulette.RouletteGameRoutes(path)
}
