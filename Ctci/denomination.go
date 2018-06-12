package main

import (
	"fmt"
)

func makeChange(n int, deno int) int {
	var nextDeno, ways int

	switch deno {
	case 25:
		nextDeno = 10
	case 10:
		nextDeno = 5
	case 5:
		nextDeno = 1
	case 1:
		return 1
	}
	for i := 0; i*deno <= n; i++ {
		ways += makeChange(n-i*deno, nextDeno)
	}
	return ways
}
func main() {
	fmt.Println("MakeChange", makeChange(20, 1))
}
