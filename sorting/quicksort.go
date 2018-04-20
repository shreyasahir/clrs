package main

import (
	"fmt"
)

func main() {

	x := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	quickSort(x, 0, len(x)-1)
	fmt.Println("value after sorting", x)
}

func quickSort(arr []int, first, last int) {
	if first < last {
		split := quickPartition(arr, first, last)
		quickSort(arr, first, split-1)
		quickSort(arr, split+1, len(arr)-1)
	}
}

func quickPartition(arr []int, low, high int) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
