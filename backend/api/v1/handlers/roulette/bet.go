package roulette

import (
	"casino.website/api/v1/handlers/websocket"
	"casino.website/internal/models"
	"github.com/yyewolf/gosf"
)

func bet(c *gosf.Client, r *gosf.Request) *gosf.Message {
	b := models.NewBet()

	err := gosf.MapToStruct(r.Message.Body, b)

	if err != nil {
		return nil
	}

	if b.Amount <= 0 {
		return nil
	}

	u, err := websocket.GetUser(c)

	if err != nil {
		return nil
	}

	b.SetUser(u)

	Roulette.RegisterBet(b)

	return b.Message()
}
