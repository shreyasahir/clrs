package main

import (
	"fmt"
)

var (
	result []int
	sum    int
)

func evenNo(n string) {
	result = make([]int, len(n))
	sum := 0
	for k, v := range n {
		y := v - 48
		if y%2 == 0 {
			result[k] = 1
			sum++
		}
	}
	//fmt.Println("result", result)

	for i := 0; i < len(result); i++ {
		if i <= 0 {
			fmt.Printf("%d", sum)
			fmt.Printf(" ")
		} else {
			fmt.Printf("%d", sum-result[i-1])
			fmt.Printf(" ")
			sum = sum - result[i-1]
		}
	}
}
func main() {
	var num string
	fmt.Scanf("%s", &num)
	evenNo(num)
}
