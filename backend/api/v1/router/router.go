package router

import (
	"casino.website/api/v1/handlers/admin"
	"casino.website/api/v1/handlers/auth"
	"casino.website/api/v1/handlers/roulette"
	"casino.website/api/v1/handlers/user"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	path := engine.Group("/api/v1")

	auth.LoadRoutes(path)
	roulette.LoadRoulette(path, engine)
	user.LoadUser(path)
	admin.LoadAdmin(path)
}
