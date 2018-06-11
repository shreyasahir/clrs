package main

import (
	"fmt"
	"sort"
)

func twoSum(arr []int, n int, key int) int {
	fmt.Println("sorted array", arr)
	sumDiff := make(map[int]int)
	for k, v := range arr {

		if val, ok := sumDiff[n-v]; ok {
			fmt.Println("pairs are", val, k, key)
			return -1
		} else {
			if n > v {
				sumDiff[v] = k
			}
		}
		fmt.Println("sumdiff", sumDiff)
	}
	return -1
}

func threeSum(arr []int) {
	sort.Ints(arr)
	for k, v := range arr {
		twoSum(arr, -1*v, k)
	}
}

func main() {
	arr := []int{4, 3, -1, 2, -2, 10}
	threeSum(arr)
}
