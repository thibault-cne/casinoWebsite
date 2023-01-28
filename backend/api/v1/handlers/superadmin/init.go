package superadmin

import (
	"casino.website/api/v1/handlers/superadmin/user"
	"casino.website/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func LoadSuperAdmin(path *gin.RouterGroup) {
	subpath := path.Group("/superadmin", middlewares.AuthRequired(), middlewares.IsSuperAdmin())

	user.LoadUser(subpath)
}
