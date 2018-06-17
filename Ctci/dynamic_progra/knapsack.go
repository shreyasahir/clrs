package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func knapsack(value, weight []int, n int, w int) int {
	if n == 0 || w == 0 {
		return 0
	}
	if weight[n-1] > w {
		return knapsack(value, weight, n-1, w)
	}
	return max(value[n-1]+knapsack(value, weight, n-1, w-weight[n-1]), knapsack(value, weight, n-1, w))
}
func main() {
	value := []int{60, 100, 120}
	weight := []int{10, 20, 30}
	w := 50
	n := len(value)
	fmt.Println("knapsack", knapsack(value, weight, n, w))
}
