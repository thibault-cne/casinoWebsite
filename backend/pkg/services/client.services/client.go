package clientservices

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateNewClient(username string, accessType int) string {
	// TO DO : check if username is available
	client := NewClient(username, accessType)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return ""
	}

	password := client.Password
	hash, err := hashPassword(password)

	if err != nil {
		fmt.Printf("An error occured while hashing the password : %s", err.Error())
		return ""
	}

	client.Password = hash
	db.Save(client)
	return password
}

func CheckClientModerationType(userId int) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return false
	}

	var user Client

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return false
	}

	if user.AccessType >= 2 {
		return true
	}

	return false
}

func GetClientByUsername(username string) *Client {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return nil
	}

	var user Client

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func GetClientById(id int) *Client {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return nil
	}

	var user Client

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Save(c)
}
