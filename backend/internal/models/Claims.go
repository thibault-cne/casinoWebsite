package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}

type UserClaims struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}
