package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	i := 1

	for i < n {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
		i++
	}
}
func main() {
	x := []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(x)
	fmt.Println("value after sorting", x)
}
