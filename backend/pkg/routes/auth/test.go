package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(ctx *gin.Context) {
	c := GetUser(ctx)

	if c == nil {
		fmt.Printf("nil\n")
	} else {
		fmt.Printf("done\n")
	}

	ctx.Status(http.StatusOK)
}
