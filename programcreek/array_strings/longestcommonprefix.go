package main

import (
	"fmt"
)

func findMin(s []string) int {
	var length = len(s[0])
	for i := 1; i < len(s); i++ {
		if length > len(s[i]) {
			length = len(s[i])
		}
	}
	return length
}

func containsPrefix(s string, arr []string, low, mid int) bool {
	fmt.Println(s[low:mid])
	for _, v := range arr {
		if v[low:mid] != s[low:mid] {
			return false
		}
	}
	return true
}
func findLongComm(s []string) string {
	minLen := findMin(s)
	var result string
	for i := 0; i < minLen; i++ {
		curr := s[0][i]
		for j := 1; j < len(s); j++ {
			if s[j][i] != curr {
				return result
			}
		}
		result += string(curr)
	}
	return result
}

func main() {
	s := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	fmt.Println(findLongComm(s))
}
