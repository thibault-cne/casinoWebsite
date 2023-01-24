package utils

import "github.com/bwmarrin/snowflake"

var (
	generator *snowflake.Node
)

func init() {
	// Epoch is at 8/12/2022 18:50:00
	snowflake.Epoch = 1670521800000

	// Node is Yewolf
	node, err := snowflake.NewNode(89101119111108102 % 1023)
	if err != nil {
		panic(err)
	}

	generator = node
}

func Generate() string {
	return generator.Generate().String()
}
