package roulette

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"time": Roulette.CreatedAt.Add(time.Second * 30).UnixMilli()})
}
