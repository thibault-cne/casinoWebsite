package models

import (
	"time"

	"casino.website/internal/db"
	"casino.website/internal/utils"
	"github.com/yyewolf/gosf"
)

type Bet struct {
	ID         string `gorm:"primaryKey"`
	UserId     string
	RouletteId string
	User       *User  `json:"user" gorm:"-:all"`
	Color      string `json:"color"`
	Amount     int    `json:"amount"`
	CreatedAt  time.Time
}

func NewBet() *Bet {
	return &Bet{
		ID:        utils.GenerateBetId(),
		CreatedAt: time.Now(),
	}
}

func (b *Bet) Save() error {
	b.UserId = b.User.ID
	return db.DB.Save(b).Error
}

func (b *Bet) SetUser(u *User) {
	b.User = u
}

func (b *Bet) Message() *gosf.Message {
	msg := gosf.NewSuccessMessage()
	msg.Body = gosf.StructToMap(b)

	return msg
}
