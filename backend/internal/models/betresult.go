package models

import "time"

type BetResult struct {
	ID        string `gorm:"primaryKey" json:"id"`
	GameId    string `json:"game_id"`
	Color     string `json:"color"`
	Amount    int    `json:"amount"`
	Win       bool   `json:"win"`
	CreatedAt time.Time
}
