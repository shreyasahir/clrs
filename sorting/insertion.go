package main

import (
	"fmt"
)

func insert(x []int) []int {
	i := 1
	for i < len(x) {
		j := i
		for j > 0 && x[j-1] > x[j] {
			x[j], x[j-1] = x[j-1], x[j]
			j--
			fmt.Println(x)
		}
		i++
	}
	return x
}

func main() {
	var x = []int{5, 2, 4, 6, 1, 3}
	fmt.Println(x)
	insert(x)
	fmt.Println(x)
}
