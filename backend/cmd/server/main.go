package main

import (
	"time"

	"casino.website/api/v1/router"
	"casino.website/internal/env"
	"casino.website/internal/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var secret = []byte(env.Config.CookieKey)

func main() {
	app := gin.Default()
	// Config CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.Config.FrontendURL()},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup the session
	store := cookie.NewStore(secret)
	store.Options(sessions.Options{
		MaxAge: 10 * 60 * 60 * 24,
		Path:   "/",
	})
	app.Use(sessions.Sessions(env.Config.CookieName, store))

	router.Route(app)
	models.Migrate()

	app.Run(":8000")
}
