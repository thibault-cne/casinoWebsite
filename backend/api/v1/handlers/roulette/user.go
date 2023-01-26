package roulette

import (
	"errors"

	"casino.website/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/gosf"
)

func getUser(c *gosf.Client) (*models.User, error) {
	w := NewCustomResponseWriter()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = c.GetChannel().Request()
	Engine.HandleContext(ctx)

	session := sessions.Default(ctx)
	userID := session.Get("user")

	switch userID.(type) {
	case string:
	default:
		return nil, errors.New("not a string id")
	}

	return models.GetUserByID(userID.(string))
}
