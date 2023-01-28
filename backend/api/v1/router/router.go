package router

import (
	"casino.website/api/v1/handlers/admin"
	"casino.website/api/v1/handlers/auth"
	"casino.website/api/v1/handlers/roulette"
	"casino.website/api/v1/handlers/superadmin"
	"casino.website/api/v1/handlers/user"
	"casino.website/api/v1/handlers/websocket"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	path := engine.Group("/api/v1")

	auth.LoadRoutes(path)
	websocket.LoadWebsocket(path, engine)
	roulette.LoadRoulette(path)
	user.LoadUser(path)
	admin.LoadAdmin(path)
	superadmin.LoadSuperAdmin(path)
}
