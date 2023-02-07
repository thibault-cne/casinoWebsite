package roulette

import (
	"casino.website/api/v1/handlers/websocket"
	"casino.website/api/v1/middlewares"
	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

var (
	Roulette *models.Roulette
)

func LoadRoulette(rg *gin.RouterGroup) {
	Roulette = models.NewRoulette()

	// Init the websocket endpoints
	initWebsocket()

	subpath := rg.Group("/roulette", middlewares.AuthRequired())

	subpath.GET("/bet", getBet)
	subpath.GET("/time", getTime)

	go Roulette.Start()
}

func initWebsocket() {
	// Listen ws endpoint
	listenEndpoint()

	// Join roulette room
	gosf.OnConnect(func(c *gosf.Client, r *gosf.Request) {
		user, err := websocket.GetUser(c)

		if err != nil || user == nil {
			c.Disconnect()
			return
		}

		c.Join("roulette")
		c.Join(user.ID)
	})

	// Leave the roulette room on disconnect
	gosf.OnDisconnect(func(c *gosf.Client, request *gosf.Request) {
		user, err := websocket.GetUser(c)

		if err != nil || user == nil {
			c.Disconnect()
			return
		}

		c.Leave("roulette")
		c.Leave(user.ID)
	})
}

func listenEndpoint() {
	gosf.Listen("bet", bet)
}
