package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}

func main() {
	x := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(x)
	fmt.Println("value after sorting", x)
}
