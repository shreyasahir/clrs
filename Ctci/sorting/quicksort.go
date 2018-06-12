package main

import (
	"fmt"
)

func quickSort(arr []int, start, end int) {
	if start < end {
		split := quickPartition(arr, start, end)
		quickSort(arr, 0, split-1)
		quickSort(arr, split+1, end)

	}
}

func quickPartition(arr []int, start, end int) int {
	i := start - 1
	pivot := arr[end]

	for j := start; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
func main() {
	x := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(x, 0, len(x)-1)
	fmt.Println("value after sorting", x)
}
