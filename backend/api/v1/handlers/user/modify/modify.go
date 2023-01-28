package modify

import (
	"net/http"

	"casino.website/internal/models"
	"github.com/gin-gonic/gin"
)

func modify(ctx *gin.Context) {
	var dataUser models.User

	err := ctx.ShouldBindJSON(&dataUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	user := ctx.MustGet("user").(*models.User)

	// Retrieve authorized fields
	user.Username = dataUser.Username

	err = user.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
