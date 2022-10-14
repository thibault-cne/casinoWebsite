package roulette

import (
	"fmt"

	"casino.website/pkg/models"
	"casino.website/pkg/server/websocket"
	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *models.RouletteGame) {
	conn, err := websocket.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	loggedIn := ctx.MustGet("Logged").(bool)

	if !loggedIn {
		fmt.Printf("We are here\n")
		rGame.RegisterPlayerConn <- conn

	} else {
		client := ctx.MustGet("User").(*models.Client)

		rClient := &models.RouletteClient{
			Client: client,
			WsConn: conn,
		}

		rGame.RegisterPlayerConn <- rClient.WsConn

		go rClient.ReadRouletteClient(rGame)
	}
}
