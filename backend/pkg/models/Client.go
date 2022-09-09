package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	AccessType int    `json:"accessType"`
	Wallet     int    `json:"wallet"`
}
