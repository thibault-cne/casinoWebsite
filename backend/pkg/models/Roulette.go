package models

type RouletteBets struct {
	ClientId    string `json:"id"`
	BetPosition string `json:"position"`
	BetAmount   int    `json:"amount"`
}

type Broadcast struct {
	DataType string                 `json:"dataType"`
	Data     map[string]interface{} `json:"data"`
}

type FrontBet struct {
	Color  string `json:"color"`
	Amount int    `json:"amount"`
}
