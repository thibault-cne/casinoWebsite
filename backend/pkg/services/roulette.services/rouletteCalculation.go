package rouletteservices

import (
	"fmt"
	"math/rand"
)

func calculateRouletteColor() (int, string) {
	result := randInt(0, 14)
	color := ""

	if result == 0 {
		color = "green"
	} else if result < 8 {
		color = "red"
	} else {
		color = "black"
	}

	fmt.Printf("Rand number : %d\n", result)

	return result, color
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
