package clientservices

import "casino.website/pkg/models"

func NewClient(username string, accessType int) *models.Client {
	return &models.Client{Username: username, Password: generatePassword(), AccessType: accessType, Wallet: 0}
}
