package main

import (
	"fmt"
)

func paren(left, right int, s string) {
	if left == 0 && right == 0 {
		fmt.Println(s)
	}
	if left > right {
		return
	}
	if left > 0 {
		paren(left-1, right, s+"(")
	}
	if right > 0 {
		paren(left, right-1, s+")")
	}
}
func validParanthesis(n int) {
	paren(n/2, n/2, "")
}
func main() {
	validParanthesis(8)
}
