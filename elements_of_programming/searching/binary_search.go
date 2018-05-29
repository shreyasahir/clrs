package main

import (
	"fmt"
)

func binarySearch(arr []int, start, end, x int) int {
	var middle int
	for start <= end {
		middle = start + (end-start)/2
		if arr[middle] == x {
			return middle
		} else if arr[middle] < x {
			start = middle + 1
		} else {
			end = middle - 1
		}

	}

	return -1
}
func main() {
	arr := []int{2, 3, 4, 10, 40}
	fmt.Println(binarySearch(arr, 0, len(arr), 10))
}
