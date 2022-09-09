package rouletteservices

import (
	"fmt"
	"time"

	"casino.website/pkg/models"
	"github.com/gorilla/websocket"
)

type RouletteGame struct {
	Id                   string
	RegisterBet          chan *models.RouletteBets
	RegisterPlayerConn   chan *websocket.Conn
	UnregisterPlayerConn chan *websocket.Conn
	Bets                 map[*models.RouletteBets]bool
	ConnectedPlayers     map[*websocket.Conn]bool
	Broadcast            chan *models.Broadcast
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   generateRandomId(),
		RegisterBet:          make(chan *models.RouletteBets),
		RegisterPlayerConn:   make(chan *websocket.Conn),
		UnregisterPlayerConn: make(chan *websocket.Conn),
		Bets:                 make(map[*models.RouletteBets]bool),
		ConnectedPlayers:     make(map[*websocket.Conn]bool),
		Broadcast:            make(chan *models.Broadcast),
	}
}

func (rGame *RouletteGame) cleanRouletteGame() {
	rGame.Bets = make(map[*models.RouletteBets]bool)
}

func (rGame *RouletteGame) Start() {
	for {
		select {
		case bet := <-rGame.RegisterBet:
			fmt.Printf("Registering a new bet %+v\n", bet)
			rGame.Bets[bet] = true
			break
		case conn := <-rGame.RegisterPlayerConn:
			fmt.Printf("Registering a new player.\n")
			rGame.ConnectedPlayers[conn] = true
			break
		case conn := <-rGame.UnregisterPlayerConn:
			delete(rGame.ConnectedPlayers, conn)
			break
		case broadcast := <-rGame.Broadcast:
			for conn := range rGame.ConnectedPlayers {
				if err := conn.WriteJSON(broadcast); err != nil {
					fmt.Printf("An error occured while broadcasting : %s", err.Error())
					return
				}
			}
			break
		}
	}
}

func (rGame *RouletteGame) End() {
	defer func() {
		fmt.Printf("Ended a roulette game.\n")
		rGame.cleanRouletteGame()
		fmt.Printf("Game cleaned.\n")
		go rGame.End()
	}()

	time.Sleep(30 * time.Second)

	result, color := calculateRouletteColor()

	fmt.Printf("The good color is : %s\n", color)

	for bet := range rGame.Bets {
		if bet.BetPosition == color {
			fmt.Printf("Client %s won %d\n", bet.ClientId, bet.BetAmount)
		}
	}

	endBroadcast := models.Broadcast{
		DataType: "endGame",
		Data: map[string]interface{}{
			"color":  color,
			"number": result,
		},
	}

	rGame.Broadcast <- &endBroadcast
}
