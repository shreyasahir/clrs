// Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

// Examples:
// Given [1, 2, 3, 4, 5],
// return true.
package main

import (
	"fmt"
)

func triplet(n []int) bool {
	x := 1000000
	y := 1000000
	for i := 0; i < len(n); i++ {
		z := n[i]
		if x >= z {
			x = z
		} else if y >= z {
			y = z
		} else {
			return true
		}
		fmt.Println(x, y, z)
	}

	return false
}
func main() {
	n := []int{5, 1, 5, 5, 2, 5, 4}
	fmt.Println("n", n)
	fmt.Println(triplet(n))
}
