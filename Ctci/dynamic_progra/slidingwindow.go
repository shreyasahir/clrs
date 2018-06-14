package main

import (
	"fmt"
)

func slidingWindow(arr []int, k int) int {
	n := len(arr)
	var maxSum = 0
	if k > n {
		return -1
	}
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	var windowSum = maxSum

	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		maxSum = max(windowSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//−2 ,1,−3,4,−1,2,1,−5,4
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	k := 5
	fmt.Println(slidingWindow(arr, k))
}
