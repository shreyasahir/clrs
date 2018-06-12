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

	highest := findHighest(arr)

	fmt.Println("Highest", highest)
	aux := make([]int, highest+1)

	sortedArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}

	j := 0
	for i := 0; i < len(arr); i++ {
		tmp := aux[i]
		for ; tmp > 0; tmp-- {
			sortedArr[j] = i
			j++
		}
	}

	fmt.Println("sortedarr", sortedArr)
	fmt.Println("Sorted Array", sortedArr)
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
