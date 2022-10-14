package models

import (
	"crypto/rand"
	"fmt"
	r "math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type RouletteBets struct {
	Client      *Client `json:"id"`
	BetPosition string  `json:"position"`
	BetAmount   int     `json:"amount"`
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
	Color                string
	Number               int
}

func NewRouletteGame() *RouletteGame {
	return &RouletteGame{
		Id:                   generateRandomId(),
		RegisterBet:          make(chan *RouletteBets),
		RegisterPlayerConn:   make(chan *websocket.Conn),
		UnregisterPlayerConn: make(chan *websocket.Conn),
		Bets:                 make(map[*RouletteBets]bool),
		ConnectedPlayers:     make(map[*websocket.Conn]bool),
		Broadcast:            make(chan *Broadcast),
	}
}

func generateRandomId() (randId string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	randId = fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func (rGame *RouletteGame) cleanRouletteGame() {
	for bet := range rGame.Bets {
		if bet.BetPosition == rGame.Color {
			bet.Client.AddMoneyToClient(2 * bet.BetAmount)
		}
	}

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
	time.Sleep(30 * time.Second)

	rGame.calculateRouletteColor()

	fmt.Printf("The good color is : %s\n", rGame.Color)

	for bet := range rGame.Bets {
		if bet.BetPosition == rGame.Color {
			fmt.Printf("Client %s won %d\n", bet.Client.ID, bet.BetAmount)
		}
	}

	endBroadcast := Broadcast{
		DataType: "endGame",
		Data: map[string]interface{}{
			"color":  rGame.Color,
			"number": rGame.Number,
		},
	}

	rGame.Broadcast <- &endBroadcast

	defer func() {
		fmt.Printf("Ended a roulette game.\n")
		rGame.cleanRouletteGame()
		fmt.Printf("Game cleaned.\n")
		go rGame.End()
	}()
}

func (r *RouletteGame) calculateRouletteColor() {
	result := randInt(0, 14)
	color := ""

	if result == 0 {
		color = "green"
	} else if result < 8 {
		color = "red"
	} else {
		color = "black"
	}

	fmt.Printf("Rand number : %d\n", result)

	r.Color = color
	r.Number = result
}

func randInt(min int, max int) int {
	return min + r.Intn(max-min)
}
