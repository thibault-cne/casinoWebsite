package get

import (
	"fmt"
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func getPicture(ctx *gin.Context) {
	userId := ctx.Query("userId")

	fmt.Printf("%s\n", userId)

	user, err := models.GetUserByID(userId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	filePath, err := user.GetPicture()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cannot retrive picture"})
		return
	}

	ctx.File(filePath)
}
