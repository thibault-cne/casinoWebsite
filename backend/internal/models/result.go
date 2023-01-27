package models

import (
	"github.com/yyewolf/gosf"
)

type Result struct {
	Color  string `json:"color"`
	Number int    `json:"number"`
}

func (r *Result) Message() *gosf.Message {
	msg := gosf.NewSuccessMessage()
	msg.Body = gosf.StructToMap(r)

	return msg
}
