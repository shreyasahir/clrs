package main

import (
	"fmt"
)

func minCandies(ratings []int) int {
	if len(ratings) == 0 || ratings == nil {
		return 0
	}
	var curr = make([]int, len(ratings))
	//var result int
	curr[0] = 1
	fmt.Println("curr", curr)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			curr[i] = curr[i-1] + 1
		} else {
			curr[i] = 1
		}
	}
	totalCandies := len(curr)
	fmt.Println("totalCandies", totalCandies)
	for i := totalCandies - 2; i >= 0; i-- {
		fmt.Println("i", i)
		if ratings[i] > ratings[i+1] {
			curr[i] = curr[i+1] + 1
		}
	}

	fmt.Println("curr", curr)
	return 1
}
func main() {
	ratings := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	minCandies(ratings)
}
