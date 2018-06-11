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
		if !isPalin(s, 0, i) {
			fmt.Println("longest non palin", i)
			return i
		}
	}
	return 0
}
func main() {
	s := "aacecaaa"
	fmt.Println("s", s)
	count := shortestPalindrom(s, len(s))
	fmt.Println("count is", count)
	fmt.Println("missing string", s[len(s)-count:])
	str := ""
	for i := count; i >= 0; i-- {
		str += string(s[i])
	}
	fmt.Println("str", str)
	//fmt.Println("string to be added", s[len(s)-count:])
}
