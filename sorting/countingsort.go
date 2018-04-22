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
	//var aux []int
	aux := make([]int, highest+1)
	// for i := 0; i <= highest; i++ {
	// 	aux[i] := 0
	// }

	sortedArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}

	//fmt.Println("Aux array is", aux)

	j := 0
	for i := 0; i < len(aux); i++ {
		tmp := aux[i]
		//fmt.Println(tmp)
		for ; tmp > 0; tmp-- {
			//fmt.Println("sort", sortedArr)
			sortedArr[j] = i
			//fmt.Println("sortedArry", sortedArr)
			j++
		}
	}
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
