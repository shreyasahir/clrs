package main

import (
	"fmt"
)

func maxSubArray(arr []int) {
	var maxsum, indexi, indexj int

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxsum {
				maxsum = sum
				indexi = i
				indexj = j
			}
		}
	}
	fmt.Println("maxsum", maxsum, "indexi", indexi, "indexj", indexj)
}

func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	maxSubArray(arr)
}
