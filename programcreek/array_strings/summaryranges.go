// Given a sorted integer array without duplicates, return the summary of its ranges for consecutive numbers.
// For example, given [0,1,2,4,5,7], return ["0->2","4->5","7"].

package main

import (
	"fmt"
)

func summaryRanges(s []int) []int {
	if len(s) == 0 || s == nil {
		return nil
	}
	result := []int{}
	var prev = s[0]
	result = append(result, prev)
	for i := 1; i < len(s); i++ {
		if prev+1 != s[i] {
			//end = prev
			result = append(result, prev)
			result = append(result, s[i])

			prev = s[i]
		}
		prev = s[i]

	}
	fmt.Println("result", result)
	return result
}
func main() {
	s := []int{0, 1, 2, 4, 5, 7}
	summaryRanges(s)
}
