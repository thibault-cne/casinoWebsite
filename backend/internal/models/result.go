package models

import (
	"encoding/json"

	"github.com/yyewolf/gosf"
)

type Result struct {
	Color  string `json:"color"`
	Number int    `json:"number"`
}

func (r *Result) Unwrap() string {
	b, err := json.Marshal(r)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func (r *Result) Message() *gosf.Message {
	msg := gosf.NewSuccessMessage()
	msg.Text = r.Unwrap()

	return msg
}
