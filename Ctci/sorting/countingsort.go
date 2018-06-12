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
	fmt.Println("highest", highest)
	aux := make([]int, highest+1)

	sortedArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}
	fmt.Println("aux", aux)

	j := 0

	for i := 0; i <= highest; i++ {

		for tmp := aux[i]; tmp > 0; tmp-- {
			sortedArr[j] = i
			j++
		}
	}
	fmt.Println("sortedarr", sortedArr)
}

func findHighest(arr []int) int {
	highest := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > highest {
			highest = arr[i]
		}
	}
	return highest
}
