package main

import (
	"fmt"
)

func cutRod(arr []int, size int) int {
	if size <= 0 {
		return 0
	}
	maxVal := -1
	for i := 0; i < size; i++ {
		maxVal = max(maxVal, arr[i]+cutRod(arr, size-i-1))
	}
	return maxVal
}
func max(a, b int) int {
	//fmt.Println("a", a, "b", b)
	if a > b {
		return a
	}
	return b

}
func main() {
	arr := []int{1, 5, 8, 9, 10, 17, 17, 20}

	size := len(arr)
	fmt.Println("Maximum Obtainable Value is", cutRod(arr, size))
}
