package main

import (
	"fmt"
	"math"
)

func intVal(v rune) int {
	if int(v-'0') > 0 && int(v-'0') < 9 {
		return int(v - '0')
	} else {
		return int(v-'A') + 10
	}
}

func convertToDec(input string, b1 int) {
	power := len(input) - 1
	result := 0
	for _, v := range input {
		fmt.Println(intVal(v))
		result += intVal(v) * int(math.Pow(float64(b1), float64(power)))
		power--
	}
	fmt.Println("result", result)
}

func main() {

	input := "11A"
	convertToDec(input, 16)
}
