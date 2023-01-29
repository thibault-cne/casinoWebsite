package get

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getPicture(ctx *gin.Context) {
	userId := ctx.Param("userId")
	filePath := fmt.Sprintf("./images/webp/%s.webp", userId)

	if _, err := os.Stat(filePath); err != nil {
		filePath = "./images/webp/default.webp"
	}

	ctx.Header("Content-Type", "image/webp")
	ctx.File(filePath)
	ctx.Status(http.StatusOK)
}
