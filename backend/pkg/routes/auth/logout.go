package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func logout(ctx *gin.Context) {
	c, err := ctx.Cookie("session")

	if err != nil {
		return
	}

	fmt.Printf("Cookie : %s\n", c)
}
