package main

import (
	"fmt"
)

func main() {
	x := []int{64, 34, 25, 12, 22, 11, 90}
	countingSort(x)
	fmt.Println("value after sorting", x)
}

func countingSort(arr []int) {

	k := findHighest(arr)
	aux := make([]int, k+1)
	sortedArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}

	j := 0
	for i := 0; i <= k; i++ {

		for tmp := aux[i]; tmp > 0; tmp-- {
			sortedArr[j] = i
			j++
		}
	}

	fmt.Println("sortedArr", sortedArr)
}

func findHighest(arr []int) int {
	n := len(arr)
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
