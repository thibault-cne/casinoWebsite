package middlewares

import (
	"casino.website/pkg/routes/auth"
	"github.com/gin-gonic/gin"
)

func UserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		u := auth.GetUser(ctx)

		if u == nil {
			ctx.Set("Logged", false)
			return
		}

		ctx.Set("Logged", true)
		ctx.Set("UserID", u.ID)
		ctx.Set("User", u)
	}
}
