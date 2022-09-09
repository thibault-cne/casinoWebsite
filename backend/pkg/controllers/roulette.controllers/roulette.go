package roulettecontrollers

import (
	"fmt"
	"strconv"

	"casino.website/pkg/config/server/websocket"
	clientservices "casino.website/pkg/services/client.services"
	oauthservices "casino.website/pkg/services/oauth.services"
	rouletteservices "casino.website/pkg/services/roulette.services"
	"github.com/gin-gonic/gin"
)

func connectRoulette(ctx *gin.Context, rGame *rouletteservices.RouletteGame) {
	conn, err := websocket.WsUpgrader(ctx.Writer, ctx.Request)

	if err != nil {
		fmt.Printf("An error occured while upgrading connection : %s.\n", err.Error())
		return
	}

	loggedInInterface, _ := ctx.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)

	if !loggedIn {
		fmt.Printf("We are here\n")
		rGame.RegisterPlayerConn <- conn

	} else {
		userIdInterface, _ := ctx.Get("user_id")
		userId := userIdInterface.(int)

		fmt.Printf("Here !\n")
		client := clientservices.GetClientById(userId)

		rClient := &rouletteservices.RouletteClient{
			ClientId:   strconv.Itoa(userId),
			ClientName: client.Username,
			WsConn:     conn,
		}

		rGame.RegisterPlayerConn <- rClient.WsConn

		go rClient.ReadRouletteClient(rGame)
	}
}

func HandleRouletteGame(rg *gin.RouterGroup) {
	rGame := rouletteservices.NewRouletteGame()

	routerGroup := rg.Group("/roulette")
	routerGroup.GET("/connect", setUserStatus(), func(ctx *gin.Context) {
		connectRoulette(ctx, rGame)
	})

	go rGame.Start()
	go rGame.End()
}

func setUserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.Query("accessToken")

		if reqToken == "" {
			ctx.Set("is_logged_in", false)
			return
		}

		claims, err := oauthservices.DecodeToken(reqToken)

		if err != nil {
			ctx.Set("is_logged_in", false)
			return
		}

		ctx.Set("is_logged_in", true)
		ctx.Set("user_id", claims.User_id)
	}
}
