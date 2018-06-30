package main

import (
	"fmt"
)

func findMissingNumber(arr []int) {
	n := len(arr) + 1
	expectedSum := n * (n + 1) / 2
	sum := 0
	for _, v := range arr {
		sum += v
	}
	missingNumber := expectedSum - sum
	fmt.Println("missing", missingNumber)
}
func main() {
	arr := []int{1, 2, 3, 5}
	findMissingNumber(arr)
}
