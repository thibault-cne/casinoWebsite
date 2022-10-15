package auth

import (
	"net/http"

	"casino.website/pkg/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if !models.CheckUserPassword(username, password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		return
	}

	s := sessions.Default(ctx)

	s.Set("authentificated", true)
	s.Set("username", username)

	s.Save()

	ctx.Status(http.StatusOK)
}

func GetUser(ctx *gin.Context) *models.Client {
	s := sessions.Default(ctx)

	auth := s.Get("authentificated")

	if auth == nil || !auth.(bool) {
		return nil
	}

	return models.GetClientByUsername(s.Get("username").(string))
}
