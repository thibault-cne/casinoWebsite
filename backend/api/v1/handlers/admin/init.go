package admin

import (
	"casino.website/api/v1/handlers/admin/user"
	"github.com/gin-gonic/gin"
)

func LoadAdmin(path *gin.RouterGroup) {
	// TODO: add middlewares
	subpath := path.Group("/admin")

	user.LoadUser(subpath)
}
