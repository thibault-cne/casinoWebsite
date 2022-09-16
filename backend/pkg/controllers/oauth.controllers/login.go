package oauthcontrollers

import (
	"fmt"
	"net/http"
	"strings"

	clientservices "casino.website/pkg/services/client.services"
	oauthservices "casino.website/pkg/services/oauth.services"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	fmt.Printf("Username %s", username)

	if !oauthservices.CheckUserPassword(username, password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bad credentials"})
		return
	}

	user := clientservices.GetClientByUsername(username)

	userClaims := oauthservices.NewUserClaims(int(user.ID))

	if userClaims == nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error occured while creating the jwt tokens."})
		return
	}

	ctx.JSON(http.StatusOK, userClaims)
}

func refreshToken(ctx *gin.Context) {
	reqToken := ctx.Request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	new_access_token, err := oauthservices.RefreshToken(reqToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accessToken": new_access_token})
}

func LoginRoutesHandler(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/oauth")

	routerGroup.POST("/login", login)
	routerGroup.GET("/refresh", refreshToken)
}
