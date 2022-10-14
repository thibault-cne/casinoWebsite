package auth

import (
	"fmt"
	"net/http"

	"casino.website/pkg/models"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	fmt.Printf("Username %s", username)

	if !models.CheckUserPassword(username, password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		return
	}

	session, _ := store.Get(ctx.Request, "cookie-name")

	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(ctx.Request, ctx.Writer)

	ctx.Status(http.StatusOK)
}

func GetUser(ctx *gin.Context) *models.Client {
	session, _ := store.Get(ctx.Request, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return nil
	}

	return models.GetClientByUsername(session.Values["username"].(string))
}
