package auth

import (
	"casino.website/api/v1/middlewares"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.RouterGroup) {
	sg := r.Group("/auth")

	sg.POST("/register", middlewares.DeauthRequired(), Register)
	sg.POST("/login", middlewares.DeauthRequired(), Login)
	sg.GET("/logout", middlewares.AuthRequired(), Logout)
	sg.GET("/connected", middlewares.AuthRequired(), Connected)
}
