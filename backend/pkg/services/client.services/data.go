package clientservices

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveAllClients() []ShortClient {
	var clients []Client

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while openning the database : %s", err.Error())
		return nil
	}

	r := db.Find(&clients)

	if r.Error != nil {
		panic(r.Error)
	}

	shortClients := make([]ShortClient, 0)

	for _, c := range clients {
		shortClients = append(shortClients, ShortClient{UserId: c.ID, Username: c.Username, AccessType: c.AccessType, Wallet: c.Wallet})
	}

	return shortClients
}
