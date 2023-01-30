package models

import (
	"time"

	"casino.website/internal/db"
	"casino.website/internal/utils"
	"github.com/yyewolf/gosf"
)

type RouletteBet struct {
	ID        string `gorm:"primaryKey"`
	UserId    string `json:"-"`
	GameId    string `json:"-"`
	User      *User  `json:"user" gorm:"-:all"`
	Color     string `json:"color"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
}

func NewBet() *RouletteBet {
	return &RouletteBet{
		ID:        utils.GenerateBetId(),
		CreatedAt: time.Now(),
	}
}

func (b *RouletteBet) Save() error {
	b.UserId = b.User.ID
	return db.DB.Save(b).Error
}

func (b *RouletteBet) SetUser(u *User) {
	b.User = u
}

func (b *RouletteBet) Message() *gosf.Message {
	msg := gosf.NewSuccessMessage()
	msg.Body = gosf.StructToMap(b)

	return msg
}

func (b *RouletteBet) ToResult() *BetResult {
	var game Roulette
	game.ID = b.GameId

	db.DB.Find(&game)

	return &BetResult{
		ID:        b.ID,
		GameId:    b.GameId,
		Color:     b.Color,
		Amount:    b.Amount,
		Win:       game.Color == b.Color,
		CreatedAt: b.CreatedAt,
	}
}
