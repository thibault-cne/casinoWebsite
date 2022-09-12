package clientservices

func NewClient(username string, accessType int) *Client {
	return &Client{Username: username, Password: generatePassword(), AccessType: accessType, Wallet: 0}
}
