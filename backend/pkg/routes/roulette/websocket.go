package roulette

import (
	"fmt"
	"net/http"

	"casino.website/pkg/routes/auth"
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

var h http.Handler = gosf.GetHandler(map[string]interface{}{"path": "/api/v1/roulette/connect", "port": 8080})

func Wrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("%+v\n", auth.GetUser(ctx))

		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func InitWebsocket() {
	// Connect to the roulette room on conection
	gosf.OnConnect(func(c *gosf.Client, r *gosf.Request) {
		c.Join("roulette")
	})

	// Leave the roulette room on disconnect
	gosf.OnDisconnect(func(c *gosf.Client, request *gosf.Request) {
		c.Leave("roulette")
	})
}
