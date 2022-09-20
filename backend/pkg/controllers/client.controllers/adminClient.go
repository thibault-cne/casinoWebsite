package clientcontrollers

import (
	"fmt"
	"net/http"

	clientservices "casino.website/pkg/services/client.services"
	"github.com/gin-gonic/gin"
)

func getAllUsers(ctx *gin.Context) {
	clients := clientservices.RetrieveAllClients()

	fmt.Printf("%+v\n", clients)

	ctx.JSON(http.StatusOK, gin.H{"clients": clients})
}

func updateUser(ctx *gin.Context) {
	var client clientservices.ShortClient

	if err := ctx.Bind(&client); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client.UpdateClient()
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func updateWallet(ctx *gin.Context) {
	var client clientservices.ShortClient

	if err := ctx.Bind(&client); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client.UpdateWallet()
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func createNewClient(ctx *gin.Context) {
	var newClient NewClient

	if err := ctx.Bind(&newClient); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pass, check := clientservices.CreateNewClient(newClient.Username, newClient.AccessType)

	if !check {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "This username already exists in the database. Please change the username."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "pass": pass})
}

func ClientAdminDataHandler(rg *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	routerGroup := rg.Group("/client/data/admin", handlers...)

	routerGroup.GET("/all", getAllUsers)
	routerGroup.POST("/update/wallet", updateWallet)
	routerGroup.POST("/create", createNewClient)
}

func ClientSuperAdminDataHandler(rg *gin.RouterGroup, handlers ...gin.HandlerFunc) {
	routerGroup := rg.Group("/client/data/admin", handlers...)

	routerGroup.POST("/update/user", updateUser)
}
