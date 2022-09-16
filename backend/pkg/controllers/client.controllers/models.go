package clientcontrollers

type NewClient struct {
	Username   string `json:"username" form:"username"`
	AccessType int    `json:"accessType" form:"accessType"`
}
