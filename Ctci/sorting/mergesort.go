package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	fmt.Println("left", left)
	fmt.Println("right", right)
	return merge(left, right)

}

func merge(a, b []int) []int {
	var result []int
	//i := 0
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			result = append(result, b[0])
			b = b[1:]
		} else {
			result = append(result, a[0])
			a = a[1:]
		}
	}

	if len(a) > 0 {
		result = append(result, a...)
	}

	if len(b) > 0 {
		result = append(result, b...)

	}
	return result
}
func main() {
	x := []int{64, 34, 25, 12, 22, 11, 90}
	y := mergeSort(x)
	fmt.Println("value after sorting", y)
}
