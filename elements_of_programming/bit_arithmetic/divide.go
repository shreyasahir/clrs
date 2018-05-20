package main

import (
	"fmt"
)

//Abs returns abs
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func divide(a, b int64) int64 {
	var sign int64
	var quot int64
	if a < 0 && b < 0 {
		sign = 1
	} else if a < 0 && b > 0 {
		sign = -1
	} else if a > 0 && b < 0 {
		sign = -1
	} else {
		sign = 1
	}
	if b != 0 {
		for a >= b {
			a = a - b
			quot++
		}
	}

	return sign * quot
}
func main() {
	fmt.Println("divide", divide(8, 4))
}
