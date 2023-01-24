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
	Email     string    `json:"email"`
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
	_, err := db.DB.Exec("INSERT INTO users (id, username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", u.ID, u.Username, u.Password, u.Email, u.CreatedAt, u.UpdatedAt)
	return err
}

func CheckClientModerationType(id string) bool {
	var user User

	err := db.DB.QueryRow("SELECT status FROM users WHERE id = ?", id).Scan(&user.Status)
	if err != nil {
		return false
	}

	if user.Status >= 2 {
		return true
	}

	return false
}

func GetUserByID(id string) (*User, error) {
	var u User
	err := db.DB.QueryRow("SELECT id, username, password, email, created_at, updated_at FROM users WHERE id = ?", id).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByUsername(username string) (*User, error) {
	var u User
	err := db.DB.QueryRow("SELECT id, username, password, email, created_at, updated_at FROM users WHERE username = ?", username).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByEmail(email string) (*User, error) {
	var u User
	err := db.DB.QueryRow("SELECT id, username, password, email, created_at, updated_at FROM users WHERE email = ?", email).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
