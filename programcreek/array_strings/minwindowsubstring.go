package main

import (
	"fmt"
)

func inArray(t string, key rune) bool {
	for _, v := range t {
		if v == key {
			return true
		}
	}
	return false
}
func minWindow(s, t string) {
	var hashDict = make(map[rune]int, len(s))
	var start, end int
	var result string
	result = s
	for k, v := range s {
		//fmt.Println("hashdict", hashDict)
		end++
		if inArray(t, v) {
			if _, ok := hashDict[v]; ok {
				start = k
				end = start

				hashDict = make(map[rune]int, len(s))
			} else {
				hashDict[v] = 1
			}
			fmt.Println("string", string(v))

		} else {
			if len(hashDict) > 0 {
				//end++
			} else {
				start++
			}
		}
		fmt.Println("hashdict", hashDict)
		if len(result) > end-start && len(hashDict) == 3 {
			result = s[start:end]
			fmt.Println("results", result)
			fmt.Println("resultssss", s[start:end])

		}

	}

	fmt.Println("result", result, "start", start, "end", end)
}
func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println("original string", s)
	minWindow(s, t)
}
