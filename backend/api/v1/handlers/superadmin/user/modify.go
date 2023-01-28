package user

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

	u, err := models.GetUserByID(dataUser.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "wrong user id"})
		return
	}

	// Get authorized fields
	u.Username = dataUser.Username
	u.Wallet = dataUser.Wallet
	u.Status = dataUser.Status

	err = u.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
