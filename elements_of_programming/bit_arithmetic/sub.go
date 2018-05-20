package main

import (
	"fmt"
)

func sub(a, b int) int {

	for b != 0 {
		borrow := (^a) & b
		fmt.Println("carry", borrow)
		a ^= b
		fmt.Println("a", a)
		b = borrow << 1
		fmt.Println("b", b)
	}
	return a
}
func main() {
	a := 8
	b := 6
	fmt.Println("value of a: ", a, "value of b: ", b)
	fmt.Println(sub(a, b))
}
