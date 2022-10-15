package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	subpath := rg.Group("/auth")

	subpath.POST("/login", login)
	subpath.GET("/logout", logout)
	subpath.GET("/test", test)
}
