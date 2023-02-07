package main

import (
	"fmt"

	"casino.website/internal/models"
	"casino.website/internal/utils"
)

func main() {
	models.Migrate()

	c := &models.Claims{
		Code: utils.GenerateClaimsCode(),
		Use:  5,
	}
	c.Save()

	fmt.Printf("%+v", c.Code)
}
