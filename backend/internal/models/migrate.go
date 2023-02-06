package models

import "casino.website/internal/db"

func Migrate() {
	db.DB.AutoMigrate(&User{})
	db.DB.AutoMigrate(&Roulette{})
	db.DB.AutoMigrate(&RouletteBet{})
	db.DB.AutoMigrate(&Claims{})
}
