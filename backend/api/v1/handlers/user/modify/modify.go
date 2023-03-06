package modify

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"casino.website/internal/models"
	"casino.website/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ModifyForm struct {
	Picture *multipart.FileHeader `form:"picture"`
	User    *models.User          `form:"user" binding:"required"`
}

func modify(ctx *gin.Context) {
	var data ModifyForm

	err := ctx.ShouldBindWith(&data, binding.FormMultipart)

	if err != nil {
		fmt.Printf("Bind error : %+v\n", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	user := ctx.MustGet("user").(*models.User)

	// Retrieve authorized fields
	user.Username = data.User.Username

	err = user.Save()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = utils.ImageToWebp(data.Picture, data.User.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
