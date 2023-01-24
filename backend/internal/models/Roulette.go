package models

import (
	"fmt"
	"time"

	"casino.website/internal/utils"
	"github.com/gorilla/websocket"
)

type RouletteBets struct {
	ClientId    string `json:"id"`
	BetPosition string `json:"position"`
	BetAmount   int    `json:"amount"`
}

type Broadcast struct {
	DataType string                 `json:"dataType"`
	Data     map[string]interface{} `json:"data"`
}

type FrontBet struct {
	Color  string `json:"color"`
	Amount int    `json:"amount"`
}

type RouletteGame struct {
	Id                   string
	RegisterBet          chan *RouletteBets
	RegisterPlayerConn   chan *websocket.Conn
	UnregisterPlayerConn chan *websocket.Conn
	Bets                 map[*RouletteBets]bool
	ConnectedPlayers     map[*websocket.Conn]bool
	Broadcast            chan *Broadcast
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   utils.GenerateRandomId(),
		RegisterBet:          make(chan *RouletteBets),
		RegisterPlayerConn:   make(chan *websocket.Conn),
		UnregisterPlayerConn: make(chan *websocket.Conn),
		Bets:                 make(map[*RouletteBets]bool),
		ConnectedPlayers:     make(map[*websocket.Conn]bool),
		Broadcast:            make(chan *Broadcast),
	}
}

func (rGame *RouletteGame) cleanRouletteGame() {
	rGame.Bets = make(map[*RouletteBets]bool)
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

	result, color := utils.GenerateRouletteColor()

	fmt.Printf("The good color is : %s\n", color)

	for bet := range rGame.Bets {
		if bet.BetPosition == color {
			fmt.Printf("Client %s won %d\n", bet.ClientId, bet.BetAmount)
		}
	}

	endBroadcast := Broadcast{
		DataType: "endGame",
		Data: map[string]interface{}{
			"color":  color,
			"number": result,
		},
	}

	rGame.Broadcast <- &endBroadcast
}
