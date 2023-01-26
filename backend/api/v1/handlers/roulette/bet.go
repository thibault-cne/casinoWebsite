package roulette

import (
	"casino.website/internal/models"
	"github.com/yyewolf/gosf"
)

func bet(c *gosf.Client, r *gosf.Request) *gosf.Message {
	// msg := r.Message.Body
	b := models.Bet{}

	err := gosf.MapToStruct(r.Message.Body, &b)

	if err != nil {
		return gosf.NewFailureMessage()
	}

	u, err := getUser(c)

	if err != nil {
		return gosf.NewFailureMessage()
	}

	b.SetUser(u)

	Roulette.RegisterBet(&b)

	return nil
}
