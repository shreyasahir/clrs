package main

import (
	"fmt"
)

func main() {
	var x = []int{5, 2, 4, 6, 1, 3}
	fmt.Println(x)
	fmt.Println(insertSort(x))
}

func insertSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		key := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		i++
	}
	return arr
}
