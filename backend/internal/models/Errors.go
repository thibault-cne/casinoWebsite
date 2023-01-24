package models

type Error struct {
	ErrorType    int    `json:"errorType"`
	ErrorMessage string `json:"errorMessage"`
}
