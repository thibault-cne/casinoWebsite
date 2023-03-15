package models

import (
	"fmt"
	"os"
	"time"

	"casino.website/internal/db"
	"golang.org/x/crypto/bcrypt"
)

var (
	statusMap = map[string]int{"user": 1, "admin": 2, "superAdmin": 3}
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Status    string    `json:"status"`
	Wallet    int       `json:"wallet"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPasswordHash(attempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(attempt))
	return err == nil
}

func (u *User) Create() error {
	err := db.DB.Create(u).Error
	return err
}

func (u *User) Save() error {
	u.UpdatedAt = time.Now()
	return db.DB.Save(u).Error
}

func CheckClientModerationType(id string) bool {
	user, err := GetUserByID(id)

	if err != nil {
		return false
	}

	if statusMap[user.Status] >= 2 {
		return true
	}

	return false
}

func GetUserByID(id string) (*User, error) {
	user := &User{
		ID: id,
	}

	err := db.DB.Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User

	err := db.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) AddMoney(amount int) error {
	u.Wallet += amount

	return u.Save()
}

func (u *User) RemoveMoney(amount int) error {
	u.Wallet -= amount

	return u.Save()
}

func GetAllUsers() ([]*User, error) {
	var users []*User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetUserStat(gameType string) []*BetResult {
	var res []*BetResult

	if gameType == "roulette" {
		var bets []*RouletteBet

		db.DB.Where("user_id = ?", u.ID).Find(&bets)

		for _, bet := range bets {
			res = append(res, bet.ToResult())
		}
	}

	return res
}

func (u *User) GetPicture() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/images/webp/%s.webp", path, u.ID)

	if _, err := os.Stat(filePath); err != nil {
		filePath = fmt.Sprintf("%s/images/webp/default.webp", path)
	}

	return filePath, nil
}
