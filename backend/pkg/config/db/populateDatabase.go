package db

import (
	"fmt"

	clientservices "casino.website/pkg/services/client.services"
)

func populateDatabase() {
	password := clientservices.CreateNewClient("test", 3)
	clientservices.CreateNewClient("John Doe", 1)
	clientservices.CreateNewClient("Gerald Doe", 1)
	clientservices.CreateNewClient("Christophe Doe", 1)

	c := clientservices.GetClientByUsername("test")

	c.AddMoneyToClient(500)

	fmt.Printf("Admin account : \n username : `test` password : `%s`.\n", password)
}
