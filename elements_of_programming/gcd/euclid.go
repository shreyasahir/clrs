package main

import (
	"fmt"
)

func gcd(a, b int) int {
	fmt.Println("a", a, "b", b)
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	fmt.Println(gcd(60, 36))
}
