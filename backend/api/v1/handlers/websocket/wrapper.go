package websocket

import "github.com/gin-gonic/gin"

func Wrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Ws.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
