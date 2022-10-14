package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	// Génération d'un UUID pour la clé du cookie
	token, err := uuid.NewV4()
	if err != nil {
		panic("impossible de générer une clé pour les cookies")
	}
	key := token.Bytes()
	maxAge := 86400 * 30 // 30 days cookie

	store = sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Path = "/"
}

func AuthRoutes(rg *gin.RouterGroup) {
	subpath := rg.Group("/auth")

	subpath.POST("/login", login)
	subpath.GET("/logout", logout)
}
