package main

import (
	"math/rand"
	"time"

	"casino.website/pkg/config/db"
	"casino.website/pkg/config/server"
)

func main() {
	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Init the database
	db.InitDatabase()

	// Init the htpp server
	server.InitServer()
}
