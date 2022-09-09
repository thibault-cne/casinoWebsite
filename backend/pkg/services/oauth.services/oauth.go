package oauthservices

import (
	"fmt"

	"casino.website/pkg/models"
	clientservices "casino.website/pkg/services/client.services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CheckUserPassword(username string, password string) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return false
	}

	var user models.Client

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}

	return clientservices.CheckPasswordHash(password, user.Password)
}
