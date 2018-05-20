package main

import (
	"fmt"
)

func multi(a, b int64) int64 {
	var ans int64
	var count uint
	for a > 0 {
		fmt.Println("a", a)
		if a%2 == 1 {
			ans += b << count

		}
		count++
		fmt.Println("ans", ans)
		fmt.Println("count", count)
		fmt.Println("b", b)
		a = a / 2
	}
	return ans
}
func main() {
	var a int64 = 5
	var b int64 = 4
	fmt.Println("result", multi(a, b))
}
