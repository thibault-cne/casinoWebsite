package models

import "casino.website/pkg/db"

func Init(t bool) {
	if t {
		db.DB.Migrator().DropTable("clients")

		// Migrate the scheme
		db.DB.AutoMigrate(&Client{})
	} else {
		// Migrate the scheme
		db.DB.AutoMigrate(&Client{})
	}
}
