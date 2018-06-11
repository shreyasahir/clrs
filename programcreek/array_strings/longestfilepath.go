package main

import (
	"fmt"
	"strings"
)

func findLength(result map[int]int, key int) int {
	var length int
	for k, v := range result {
		if k < key {
			length += v
		}
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLongestLength(s string) {
	var result map[int]int
	var maxLength int
	paths := strings.Split(s, "\n")
	result = make(map[int]int, len(paths))
	for _, v := range paths {
		if strings.Count(v, ".") == 0 {
			key := strings.Count(v, "\t")
			value := len(strings.Replace(v, "\t", "", -1))
			result[key] = value
		} else {
			key := strings.Count(v, "\t")
			length := findLength(result, key) + len(strings.Replace(v, "\t", "", -1)) + key
			maxLength = max(maxLength, length)
		}
	}
	fmt.Println("result", result)
	fmt.Println("length", maxLength)
}
func main() {
	path := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Printf("%s", path)
	fmt.Println()
	findLongestLength(path)
}
