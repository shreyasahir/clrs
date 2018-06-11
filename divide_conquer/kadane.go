package main

import (
	"fmt"
)

func kadaneArray(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	var maxSum, maxStartIndex, maxEndIndex int
	currentSum := arr[0]
	currentStartIndex := 0

	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		if currentSum > maxSum {
			maxSum, maxStartIndex, maxEndIndex = currentSum, currentStartIndex, i

		}

		if currentSum < 0 {
			currentSum = 0
			currentStartIndex = i + 1
		}
	}
	fmt.Println("Max sum", maxSum, "MaxIndex", maxStartIndex, "maxEndIndex", maxEndIndex)
	return maxSum
}

func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(kadaneArray(arr))

}
