package clientcontrollers

import (
	"fmt"
	"net/http"

	clientservices "casino.website/pkg/services/client.services"
	"github.com/gin-gonic/gin"
)

func getUserWallet(ctx *gin.Context) {
	userIdInterface, _ := ctx.Get("user_id")
	userId := userIdInterface.(int)

	c := clientservices.GetClientById(userId)

	ctx.JSON(http.StatusOK, gin.H{"Wallet": c.Wallet})
}

func getAllUsers(ctx *gin.Context) {
	clients := clientservices.RetrieveAllClients()

	fmt.Printf("%+v\n", clients)

	ctx.JSON(http.StatusOK, gin.H{"clients": clients})
}

func ClientDataHandler(rg *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	routerGroup := rg.Group("/client/data", handlers...)

	routerGroup.GET("/wallet", getUserWallet)
}

func ClientAdminDataHandler(rg *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	routerGroup := rg.Group("/client/data/admin", handlers...)

	routerGroup.GET("/all", getAllUsers)
}
