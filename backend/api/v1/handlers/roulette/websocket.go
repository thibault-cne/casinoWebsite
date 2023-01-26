package roulette

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

var (
	ws http.Handler = gosf.GetHandler(map[string]interface{}{})
)

func init() {
	// Connect to the roulette room on conection
	gosf.OnConnect(func(c *gosf.Client, r *gosf.Request) {
		user, err := getUser(c)

		if err != nil || user == nil {
			c.Disconnect()
			return
		}

		c.Join("roulette")
		c.Join(user.ID)
	})

	// Leave the roulette room on disconnect
	gosf.OnDisconnect(func(c *gosf.Client, request *gosf.Request) {
		user, err := getUser(c)

		if err != nil || user == nil {
			c.Disconnect()
			return
		}

		c.Leave("roulette")
		c.Leave(user.ID)
	})

	// Listen on endpoints
	listenEndpoints()
}

func listenEndpoints() {
	gosf.Listen("bet", bet)
}

func Wrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ws.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
