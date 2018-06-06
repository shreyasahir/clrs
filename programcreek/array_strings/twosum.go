package main

import (
	"fmt"
)

func twoSum(arr []int, n int) {

	sumDiff := make(map[int]int)
	for k, v := range arr {

		if val, ok := sumDiff[n-v]; ok {
			fmt.Println("pairs are", val, k)
		} else {
			if n > v {
				sumDiff[v] = k
			}
		}
		fmt.Println("sumdiff", sumDiff)
	}
}
func main() {
	arr := []int{2, 7, 11, 15}
	n := 6
	twoSum(arr, n)
}
