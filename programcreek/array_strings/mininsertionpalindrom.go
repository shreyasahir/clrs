package main

import (
	"fmt"
)

func isPalin(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func shortestPalindrom(s string, n int) int {
	for i := n - 1; i >= 0; i-- {
		if isPalin(s, 0, i) {
			return n - 1 - i
		}
	}
	return 0
}
func main() {
	s := "aacecaaa"
	count := shortestPalindrom(s, len(s))
	fmt.Println("count is", count)
	fmt.Println("string to be added", s[len(s)-count:])
}
