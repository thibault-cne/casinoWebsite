package admin

import (
	"casino.website/api/v1/handlers/admin/user"
	"casino.website/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func LoadAdmin(path *gin.RouterGroup) {
	// TODO: add middlewares
	subpath := path.Group("/admin", middlewares.AuthRequired(), middlewares.IsAdmin())

	subpath.GET("/isAdmin", isAdmin)

	user.LoadUser(subpath)
}
