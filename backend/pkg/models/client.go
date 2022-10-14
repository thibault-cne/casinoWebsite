package models

import (
	"fmt"
	"math/rand"
	"strings"

	"casino.website/pkg/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	AccessType int    `json:"accessType"`
	Wallet     int    `json:"wallet"`
}

func CheckUserPassword(username string, password string) bool {
	var user Client

	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}

	return checkPasswordHash(password, user.Password)
}

func generatePassword() string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetClientByUsername(username string) *Client {
	var user Client

	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func (c *Client) AddMoneyToClient(amount int) {
	c.Wallet += amount

	c.saveClient()
}

func (c *Client) RemoveMoneyToClient(amount int) bool {
	if c.checkClientAmount(amount) {
		c.Wallet -= amount

		c.saveClient()
		return true
	}

	return false
}

func (c *Client) checkClientAmount(amount int) bool {
	return c.Wallet >= amount
}

func (c *Client) saveClient() {
	db.DB.Save(c)
}

func NewClient(username string, accessType int) *Client {
	return &Client{Username: username, Password: generatePassword(), AccessType: accessType, Wallet: 0}
}

func CreateNewClient(username string, accessType int) (string, bool) {
	c := GetClientByUsername(username)

	if c != nil {
		return "", false
	}

	client := NewClient(username, accessType)

	password := client.Password
	hash, err := hashPassword(password)

	if err != nil {
		fmt.Printf("An error occured while hashing the password : %s", err.Error())
		return "", false
	}

	client.Password = hash
	client.saveClient()
	return password, true
}
