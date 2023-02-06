package utils

import "github.com/bwmarrin/snowflake"

var (
	userGenerator     *snowflake.Node
	rouletteGenerator *snowflake.Node
	betGenerator      *snowflake.Node
	claimsGenerator   *snowflake.Node
)

func init() {
	// Epoch is at 8/12/2022 18:50:00
	snowflake.Epoch = 1670521800000

	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	userGenerator = node

	node, err = snowflake.NewNode(2)

	if err != nil {
		panic(err)
	}

	rouletteGenerator = node

	node, err = snowflake.NewNode(3)

	if err != nil {
		panic(err)
	}

	betGenerator = node

	node, err = snowflake.NewNode(4)

	if err != nil {
		panic(err)
	}

	claimsGenerator = node
}

func GenerateUserId() string {
	return userGenerator.Generate().String()
}

func GenerateRouletteId() string {
	return rouletteGenerator.Generate().String()
}

func GenerateBetId() string {
	return betGenerator.Generate().String()
}

func GenerateClaimsCode() string {
	return claimsGenerator.Generate().String()
}
