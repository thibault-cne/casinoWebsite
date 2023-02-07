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
	colorMap      = map[string]int{"green": 0, "red": 1, "black": 2}
	toColorMap    = map[int]string{0: "green", 1: "red", 2: "black"}
	multiplierMap = map[string]int{"green": 14, "red": 2, "black": 2}
	isActive      = false
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

func (r *Roulette) RegisterBet(b *RouletteBet) {
	if b.User.Wallet < b.Amount {
		return
	}

	_, ok := r.Wager[b.User.ID]

	if !ok {
		r.Wager[b.User.ID] = &[3]int{0, 0, 0}
	}

	b.GameId = r.ID
	b.Save()
	b.User.RemoveMoney(b.Amount)
	r.Wager[b.User.ID][colorMap[b.Color]] += b.Amount
	gosf.Broadcast("roulette", "newbet", b.Message())
}

func (r *Roulette) Start() {
	isActive = true
	go r.Roll()
}

func (r *Roulette) Stop() {
	isActive = false
}

func (r *Roulette) Toggle() {
	if isActive {
		r.Stop()
	} else {
		r.Start()
	}
}

func (r *Roulette) Roll() {
	defer func() {
		r.reset()

		if !isActive {
			return
		}

		go r.Roll()
	}()

	gosf.Broadcast("roulette", "start", r.startTimeMessage())
	time.Sleep(30 * time.Second)

	result, color := generateRouletteColor()
	r.Color = color
	r.Number = result
	r.FinishedAt = time.Now()
	r.Save()

	for userId, wager := range r.Wager {
		if wager[colorMap[color]] > 0 {
			user, _ := GetUserByID(userId)
			user.AddMoney(wager[colorMap[color]] * multiplierMap[color])
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
	r.CreatedAt = time.Now()
	r.Wager = make(map[string]*[3]int)
}

func (r *Roulette) GetBets() []map[string]interface{} {
	var res []map[string]interface{}

	for userId, wager := range r.Wager {
		u, _ := GetUserByID(userId)
		for i, amount := range wager {
			if amount == 0 {
				continue
			}
			b := RouletteBet{
				User:   u,
				Amount: amount,
				Color:  toColorMap[i],
			}
			res = append(res, gosf.StructToMap(b))
		}
	}

	return res
}

func (r *Roulette) startTimeMessage() *gosf.Message {
	msg := gosf.NewSuccessMessage()
	msg.Body = map[string]interface{}{
		"time": r.CreatedAt.Add(time.Second * 30).UnixMilli(),
	}

	return msg
}
