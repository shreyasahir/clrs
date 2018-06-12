package main

import (
	"fmt"
)

func partition(arr []int, low, high int) int {
	i := low - 1
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		fmt.Println("for loop", arr)
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	fmt.Println("array in partition", arr)
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		split := partition(arr, low, high)
		fmt.Println("split", split)
		quickSort(arr, low, split-1)
		fmt.Println("left", arr[low:])
		quickSort(arr, split+1, high)
		fmt.Println("right", arr[split+1:])

	}
}
func main() {
	x := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("value before sorting", x)
	quickSort(x, 0, len(x)-1)
	fmt.Println("value after sorting", x)
}
