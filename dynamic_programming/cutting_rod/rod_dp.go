package main

import (
	"fmt"
)

func cutRod(arr []int, size int) int {
	result := make([]int, size+1)
	result[0] = 0
	if size <= 0 {
		return 0
	}
	//maxVal := -1

	for i := 1; i <= size; i++ {
		j := 0
		maxVal := -1
		for j < i {
			maxVal = max(maxVal, arr[j]+result[i-j-1])
			j++
		}
		result[i] = maxVal

	}
	return result[size]
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
