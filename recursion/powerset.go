//power set of s =[1,2,3] is [1],[2],3
package main

import (
	"fmt"
	"math"
)

func powerSet(original []string) [][]string {
	powerSetSize := int(math.Pow(2, float64(len(original))))
	result := make([][]string, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		var subSet []string

		for j, elem := range original {
			// fmt.Println("index", index)
			// fmt.Println("index&j", index&(1<<uint(j)))
			// fmt.Println("elem", elem)
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(powerSet(s))
}
