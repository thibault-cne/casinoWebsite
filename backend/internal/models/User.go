package models

import (
	"time"

	"casino.website/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Status    int       `json:"status"`
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

func CheckClientModerationType(id string) bool {
	user, err := GetUserByID(id)

	if err != nil {
		return false
	}

	if user.Status >= 2 {
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
	user := &User{
		Username: username,
	}

	err := db.DB.Find(user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
