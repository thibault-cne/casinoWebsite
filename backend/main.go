package main

import (
	"math/rand"
	"time"

	"casino.website/pkg/db"
	"casino.website/pkg/models"
	"casino.website/pkg/server"
)

func main() {
	// Init rand seed
	rand.Seed(time.Now().UnixNano())

	// Init the database
	db.InitDatabase()

	// Migrate the scheme
	models.Init(true)

	// Init the htpp server
	server.InitServer()
}
