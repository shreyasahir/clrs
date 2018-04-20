package main

import (
	"fmt"
)

func main() {

	x := []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(x)
	fmt.Println("value after sorting", x)
}

func selectionSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
}
