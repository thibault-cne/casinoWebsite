package db

import (
	"database/sql"
	"log"

	"casino.website/internal/env"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	// Open up our database connection.
	db, err := sql.Open("mysql", env.Config.MysqlDSN())
	if err != nil {
		panic(err.Error())
	}

	// Connect and check the server version
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		panic(err.Error())
	}
	log.Println("[DB] Connected to MySQL:", version)

	DB = db
}
