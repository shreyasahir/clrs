package main

import (
	"fmt"
)

func add(a, b int) int {

	for b > 0 {
		carry := a & b
		fmt.Println("carry", carry)
		a ^= b
		fmt.Println("a", a)
		b = carry << 1
		fmt.Println("b", b)
	}
	return a
}
func main() {
	a := 0
	b := 6
	fmt.Println("value of a: ", a, "value of b: ", b)
	fmt.Println(add(a, b))
}
