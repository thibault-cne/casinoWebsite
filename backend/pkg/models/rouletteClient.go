package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type RouletteClient struct {
	Client *Client
	WsConn *websocket.Conn
}

func (rClient *RouletteClient) ReadRouletteClient(rGame *RouletteGame) {
	defer func() {
		fmt.Printf("Defering player conn\n")
		rGame.UnregisterPlayerConn <- rClient.WsConn
		rClient.WsConn.Close()
	}()

	for {
		_, p, err := rClient.WsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var frontBet FrontBet

		if err = json.Unmarshal(p, &frontBet); err != nil {
			fmt.Printf("An error occured while unmarshalling data : %s\n", err.Error())
			fmt.Printf("Here is what where sent : %q\n", p)
			return
		}

		if rClient.Client.RemoveMoneyToClient(frontBet.Amount) {
			rBet := &RouletteBets{
				Client:      rClient.Client,
				BetPosition: frontBet.Color,
				BetAmount:   frontBet.Amount,
			}

			broadcast := &Broadcast{
				DataType: "newBet",
				Data: map[string]interface{}{
					"clientName":  rClient.Client.Username,
					"betPosition": frontBet.Color,
					"betAmount":   frontBet.Amount,
				},
			}

			rGame.RegisterBet <- rBet
			rGame.Broadcast <- broadcast
		}
	}
}
