package utils

import (
	"fmt"
	"math/rand"
)

func GenerateRandomId() (randId string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	randId = fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

func GenerateRouletteColor() (int, string) {
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
