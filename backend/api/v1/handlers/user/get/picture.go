package get

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getPicture(ctx *gin.Context) {
	userId := ctx.Param("userId")
	path, err := os.Getwd()

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	filePath := fmt.Sprintf("%s/images/webp/%s.webp", path, userId)

	if _, err := os.Stat(filePath); err != nil {
		filePath = fmt.Sprintf("%s/images/webp/default.webp", path)
	}

	fmt.Printf("%s", filePath)

	ctx.File(filePath)
}
