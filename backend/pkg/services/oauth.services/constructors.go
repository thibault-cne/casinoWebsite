package oauthservices

import (
	"fmt"
	"time"

	"casino.website/pkg/models"
	"github.com/golang-jwt/jwt"
)

func NewAccessClaims(user_id int) *models.Claims {
	standard_claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	}

	return &models.Claims{User_id: user_id, StandardClaims: standard_claims}
}

func NewRefreshClaims(user_id int) *models.Claims {
	standard_claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	return &models.Claims{User_id: user_id, StandardClaims: standard_claims}
}

func NewUserClaims(userId int) *models.UserClaims {
	accessToken, err := createToken(NewAccessClaims(userId))

	if err != nil {
		fmt.Printf("An error occured while creating the accessToken : %s", err.Error())
		return nil
	}

	refreshToken, err := createToken(NewRefreshClaims(userId))

	if err != nil {
		fmt.Printf("An error occured while creating the refreshToken : %s", err.Error())
		return nil
	}

	return &models.UserClaims{Access_token: accessToken, Refresh_token: refreshToken}
}
