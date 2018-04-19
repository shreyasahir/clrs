package main

import (
	"fmt"
)

func mergeSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	var left, right []int
	middle := len(x) / 2
	left = mergeSort(x[:middle])
	right = mergeSort(x[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
func main() {
	var x = []int{5, 2, 4, 6, 1, 3}
	result := mergeSort(x)
	fmt.Println(result)
}
