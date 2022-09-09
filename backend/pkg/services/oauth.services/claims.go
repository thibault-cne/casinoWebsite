package oauthservices

import (
	"os"
	"strings"

	"casino.website/pkg/models"
	"github.com/golang-jwt/jwt"
)

func createToken(claims *models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	loadEnv()
	JWT_SECRET := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(JWT_SECRET))
}

func RetrieveUserClaims(reqToken string) (*models.Claims, error) {
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := DecodeToken(reqToken)

	return claims, err
}

func DecodeToken(token string) (*models.Claims, error) {
	loadEnv()
	JWT_SECRET := os.Getenv("JWT_SECRET")

	tkn, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	claims := tkn.Claims.(*models.Claims)
	return claims, nil
}

func RefreshToken(refresh_token string) (string, error) {
	claims, err := DecodeToken(refresh_token)

	if err != nil {
		return "", err
	}

	return createToken(claims)
}
