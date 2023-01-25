package db

import (
	"log"

	"casino.website/internal/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Open up our database connection.
	db, err := gorm.Open(mysql.Open(env.Config.MysqlDSN()), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// Connect and check the server version
	var version string
	err = db.Raw("SELECT VERSION()").Scan(&version).Error
	if err != nil {
		panic(err.Error())
	}
	log.Println("[DB] Connected to MySQL:", version)

	DB = db
}
