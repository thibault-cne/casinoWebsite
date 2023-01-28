package roulette

import (
	"casino.website/internal/models"
	"github.com/yyewolf/gosf"
)

func bet(c *gosf.Client, r *gosf.Request) *gosf.Message {
	b := models.NewBet()
	msg := gosf.NewFailureMessage()

	err := gosf.MapToStruct(r.Message.Body, b)

	if err != nil {
		msg.Text = "Invalid bet format"
		return msg
	}

	if b.Amount <= 0 {
		msg.Text = "Cannot bet a value <= 0"
		return msg
	}

	u, err := getUser(c)

	if err != nil {
		msg.Text = "Invalid user"
		return msg
	}

	b.SetUser(u)

	Roulette.RegisterBet(b)

	msg = gosf.NewSuccessMessage()
	msg.Body = r.Message.Body

	return msg
}
