package clientservices

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Username   string `json:"username"`
	Password   string `json:"password"`
	AccessType int    `json:"accessType"`
	Wallet     int    `json:"wallet"`
}

type ShortClient struct {
	UserId     uint   `json:"userId"`
	Username   string `json:"username"`
	AccessType int    `json:"accessType"`
	Wallet     int    `json:"wallet"`
}
