package main

import (
	"fmt"
)

func kadane(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var maxSum, maxStartIndex, maxEndIndex int
	currentIndex := 0
	currentSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currentSum += arr[i]
		if currentSum > maxSum {
			maxSum, maxStartIndex, maxEndIndex = currentSum, currentIndex, i
		}

		if currentSum < 0 {
			currentSum = 0
			currentIndex = i + 1
		}
	}
	fmt.Println("Max sum", maxSum, "MaxIndex", maxStartIndex, "maxEndIndex", maxEndIndex)
	return maxSum

}

func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(kadane(arr))

}
