package main

import (
	"fmt"
)

func recursive(arr []int) int {
	n := len(arr)
	var sum = make([]int, n+1)

	for i := range sum {
		sum[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && sum[i] < sum[j]+1 {
				sum[i] = sum[j] + 1
			}
		}
	}
	fmt.Println("sum", sum)
	return 1
}
func main() {
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60}
	fmt.Println(recursive(arr))
}
