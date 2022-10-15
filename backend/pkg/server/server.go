package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"casino.website/pkg/config"
	"casino.website/pkg/models"
	"casino.website/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Init environment variables
	config.InitEnv()

	APP_DOMAIN := os.Getenv("APP_DOMAIN")
	APP_PORT := os.Getenv("APP_PORT")

	r := gin.Default()
	// Config CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	basepath := r.Group("/")

	// Add all controllers
	routes.Register(basepath)

	s := &http.Server{
		Addr:           APP_DOMAIN + ":" + APP_PORT,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	p, _ := models.CreateNewClient("test", 1)
	fmt.Printf("Username : %s Password : %s\n", "test", p)

	// Start server
	s.ListenAndServe()
}
