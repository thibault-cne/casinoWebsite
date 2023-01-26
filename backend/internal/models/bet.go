package models

type Bet struct {
	User   *User  `json:"-"`
	Color  string `json:"color"`
	Amount int    `json:"amount"`
}

func (b *Bet) SetUser(u *User) {
	b.User = u
}
