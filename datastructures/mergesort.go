package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	//var left, right []int
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
		//left = left[1:]
	}
	if len(right) > 0 {
		result = append(result, right...)
		//right = right[1:]
	}

	return result
}

func main() {
	var x = []int{5, 2, 4, 6, 1, 3, 56}
	result := mergeSort(x)
	fmt.Println(result)
}
