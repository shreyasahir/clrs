package main

import (
	"fmt"
)

func findDistance(s1, s2 string, m, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if s1[m-1] == s2[n-1] {
		return findDistance(s1, s2, m-1, n-1)
	}
	return 1 + min(findDistance(s1, s2, m-1, n), findDistance(s1, s2, m, n-1), findDistance(s1, s2, m-1, n-1))

}

func min(x, y, z int) int {
	var greater int
	if x >= y {
		greater = y
	} else {
		greater = x
	}
	if greater > z {
		return z
	}
	return greater
}

func main() {
	s1 := "sunday"
	s2 := "saturday"
	fmt.Println(findDistance(s1, s2, len(s1), len(s2)))
}
