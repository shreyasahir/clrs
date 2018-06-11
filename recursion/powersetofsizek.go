package main

import (
	"fmt"
	"math"
)

func powerSet(s []string) [][]string {
	powerSetSize := math.Pow(2, float64(len(s)))
	var result = make([][]string, 0, int(powerSetSize))
	var index int
	for index < int(powerSetSize) {
		var subSet []string

		for j, v := range s {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, string(v))
			}
		}
		if len(subSet) == 4 {
			result = append(result, subSet)
		}
		index++
	}
	return result
}
func main() {
	s := []string{"{", "}"}
	//s := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(powerSet(s))
}
