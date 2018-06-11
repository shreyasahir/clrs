package main

import (
	"fmt"
)

func gcdextended(a, b, x, y int) int {
	if a == 0 {
		x = 0
		y = 0
		return b
	}
	x1 := 1
	y1 := 1
	gcd := gcdextended(b%a, a, x1, y1)
	x = y1 - (b/a)*x1
	y = x1
	return gcd
}

func main() {
	x := 1
	y := 1
	fmt.Println(gcdextended(60, 36, x, y))
}
