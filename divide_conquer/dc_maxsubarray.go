package main

import (
	"fmt"
	"math"
)

func maxSum(a, b float64) float64 {
	fmt.Println("a", a, "b", b)
	if a > b {
		return a
	}
	return b

}

func findCrossSubArray(nums []int, left int, right int) int {

	mid := left + (right-left)/2
	leftSum := math.MinInt32
	rightSum := math.MinInt32
	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

func findSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	mid := left + (right-left)/2

	a := findSubArray(nums, left, mid)
	b := findSubArray(nums, mid+1, right)

	c := findCrossSubArray(nums, left, right)

	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	arr := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	//arr := []int{2, 3, 4, 5, 7}
	fmt.Println(findSubArray(arr, 0, len(arr)-1))

}
