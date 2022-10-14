package auth

import "github.com/gin-gonic/gin"

func logout(ctx *gin.Context) {
	session, _ := store.Get(ctx.Request, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(ctx.Request, ctx.Writer)
}
