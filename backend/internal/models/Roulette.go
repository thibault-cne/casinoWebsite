package models

import (
	"math/rand"
	"time"

	"casino.website/internal/db"
	"casino.website/internal/utils"
	"github.com/yyewolf/gosf"
)

type Roulette struct {
	ID         string             `gorm:"primaryKey"`
	Users      map[string]int     `gorm:"-:all"`
	Wager      map[string]*[3]int `gorm:"-:all"`
	Color      string
	Number     int
	CreatedAt  time.Time
	FinishedAt time.Time
}

var (
	colorMap = map[string]int{"green": 0, "red": 1, "black": 2}
)

func NewRoulette() *Roulette {
	return &Roulette{
		ID:        utils.GenerateRouletteId(),
		Users:     make(map[string]int),
		Wager:     make(map[string]*[3]int),
		CreatedAt: time.Now(),
	}
}

func (r *Roulette) Save() error {
	return db.DB.Save(r).Error
}

func (r *Roulette) AddUser(u *User) {
	r.Users[u.ID] += 1
}

func (r *Roulette) RemoveUser(u *User) {
	r.Users[u.ID] -= 1
}

func (r *Roulette) RegisterBet(b *Bet) {
	if b.User.Wallet < b.Amount {
		return
	}

	_, ok := r.Wager[b.User.ID]

	if !ok {
		r.Wager[b.User.ID] = &[3]int{0, 0, 0}
	}

	b.RouletteId = r.ID
	b.Save()
	b.User.RemoveMoney(b.Amount)
	r.Wager[b.User.ID][colorMap[b.Color]] += b.Amount
	gosf.Broadcast("roulette", "newbet", b.Message())
}

func (r *Roulette) Roll() {
	defer func() {
		r.reset()

		go r.Roll()
	}()

	time.Sleep(30 * time.Second)

	result, color := generateRouletteColor()
	r.Color = color
	r.Number = result
	r.FinishedAt = time.Now()
	r.Save()

	for userId, wager := range r.Wager {
		if wager[colorMap[color]] > 0 {
			user, _ := GetUserByID(userId)
			user.AddMoney(wager[colorMap[color]] * 2)
		}
	}

	res := &Result{
		Number: result,
		Color:  color,
	}

	gosf.Broadcast("roulette", "endgame", res.Message())
}

func generateRouletteColor() (int, string) {
	result := randInt(0, 14)
	var color string

	if result == 0 {
		color = "green"
	} else if result < 8 {
		color = "red"
	} else {
		color = "black"
	}

	return result, color
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (r *Roulette) reset() {
	r.ID = utils.GenerateRouletteId()
	r.Wager = make(map[string]*[3]int)
}
