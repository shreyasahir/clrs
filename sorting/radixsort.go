package main

import (
	"fmt"
)

func main() {

	x := []int{64, 34, 25, 12, 22, 11, 90}
	countingSort(x)
	fmt.Println("value after sorting", x)
}

func countingSort(arr []int, exp int) {

	outputArr := make([]int, len(arr))

	count := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		index := arr[i] / exp
		count[index%10]++
	}
}

func findHighest(arr []int) int {
	highest := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > highest {
			highest = arr[i]
		}
	}
	return highest
}
