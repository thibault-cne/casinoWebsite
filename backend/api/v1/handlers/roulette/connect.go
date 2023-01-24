package roulette

import (
	"fmt"

	"casino.website/internal/models"
	"casino.website/internal/server"

	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *models.RouletteGame) {
	conn, err := server.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	user := ctx.MustGet("user").(*models.User)

	rClient := &models.RouletteClient{
		User:   user,
		WsConn: conn,
	}

	rGame.RegisterPlayerConn <- rClient.WsConn

	go rClient.ReadRouletteClient(rGame)
}
