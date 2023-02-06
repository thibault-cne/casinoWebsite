package models

import "casino.website/internal/db"

type Claims struct {
	Code string `json:"code" gorm:"primaryKey"`
	Use  int    `json:"use"`
}

func (c *Claims) Create() error {
	return db.DB.Create(c).Error
}

func (c *Claims) Save() error {
	return db.DB.Save(c).Error
}

func GetAllClaims() ([]*Claims, error) {
	var claims []*Claims

	if err := db.DB.Find(&claims).Error; err != nil {
		return nil, err
	}

	return claims, nil
}

func GetClaim(code string) (*Claims, error) {
	claim := &Claims{
		Code: code,
	}

	err := db.DB.Find(claim).Error

	return claim, err
}
