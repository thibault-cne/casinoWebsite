package db

import (
	"fmt"
	"os"

	"casino.website/pkg/config"
	"casino.website/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("An error occured while creating the db : %s", err.Error())
		return
	}

	// Init env variables
	config.InitEnv()

	POPULATE_TEST_DATABASE := os.Getenv("POPULATE_TEST_DATABASE")

	if POPULATE_TEST_DATABASE == "true" {
		db.Migrator().DropTable(&models.Client{})
		db.AutoMigrate(&models.Client{})
		populateDatabase()
	} else {
		db.AutoMigrate(&models.Client{})
	}

}
