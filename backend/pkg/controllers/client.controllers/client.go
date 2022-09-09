package clientcontrollers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	clientservices "casino.website/pkg/services/client.services"
	oauthservices "casino.website/pkg/services/oauth.services"
	"github.com/gin-gonic/gin"
)

func createNewClient(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	claims, err := oauthservices.DecodeToken(reqToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if !clientservices.CheckClientModerationType(claims.User_id) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		return
	}

	username := ctx.PostForm("username")
	accessType := ctx.PostForm("accessType")

	intAccessType, err := strconv.Atoi(accessType)

	if err != nil {
		fmt.Printf("An error occured while transtyping string to int : %s", err.Error())
		return
	}

	password := clientservices.CreateNewClient(username, intAccessType)

	ctx.JSON(http.StatusOK, gin.H{"password": password})
}

func ClientRouteHandler(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/client")

	routerGroup.POST("/create", createNewClient)
}
