package rouletteservices

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"casino.website/pkg/models"
	clientservices "casino.website/pkg/services/client.services"
	"github.com/gorilla/websocket"
)

type RouletteClient struct {
	ClientId   string `json:"id"`
	ClientName string `json:"clientName"`
	WsConn     *websocket.Conn
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

		var frontBet models.FrontBet

		if err = json.Unmarshal(p, &frontBet); err != nil {
			fmt.Printf("An error occured while unmarshalling data : %s\n", err.Error())
			fmt.Printf("Here is what where sent : %q\n", p)
			return
		}

		id, _ := strconv.Atoi(rClient.ClientId)

		c := clientservices.GetClientById(id)

		if c.RemoveMoneyToClient(frontBet.Amount) {
			rBet := &models.RouletteBets{
				ClientId:    rClient.ClientId,
				BetPosition: frontBet.Color,
				BetAmount:   frontBet.Amount,
			}

			broadcast := &models.Broadcast{
				DataType: "newBet",
				Data: map[string]interface{}{
					"clientName":  rClient.ClientName,
					"betPosition": frontBet.Color,
					"betAmount":   frontBet.Amount,
				},
			}

			rGame.RegisterBet <- rBet
			rGame.Broadcast <- broadcast
		}
	}
}
