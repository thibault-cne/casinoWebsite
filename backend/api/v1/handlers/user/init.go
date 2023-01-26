package user

import (
	"casino.website/api/v1/handlers/user/get"
	"casino.website/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func LoadUser(path *gin.RouterGroup) {
	subpath := path.Group("/user", middlewares.AuthRequired())

	get.LoadGet(subpath)
}
