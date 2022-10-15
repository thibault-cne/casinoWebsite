package routes

import (
	"casino.website/pkg/routes/auth"
	"casino.website/pkg/routes/roulette"
	"github.com/gin-gonic/gin"
)

func Register(path *gin.RouterGroup) {
	subpath := path.Group("/api/v1")

	auth.AuthRoutes(subpath)
	roulette.RouletteGameRoutes(subpath)

	roulette.InitWebsocket()
}
