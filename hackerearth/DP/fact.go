package main

import (
	"fmt"
)

var (
	result []int64
	i      int64
)

func fact(n int64) int64 {
	if n < 2 {
		result[n] = 1
		return result[n]
	}
	if result[n] != 0 {
		return result[n] % 1000000007
	} else {
		result[n] = n * fact(n-1)
		return result[n] % 1000000007
	}

}
func main() {
	var num, numOfCases, y int64
	result = make([]int64, 100000)
	fmt.Scanf("%d", &numOfCases)

	for y < numOfCases {
		fmt.Scanf("%d", &num)
		res := fact(num)
		fmt.Println(res)
		y++
	}
}
