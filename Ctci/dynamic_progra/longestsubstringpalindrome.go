package main

import (
	"fmt"
)

func longPalin(s string) string {
	n := len(s)
	var table = make([][]int, n)
	for i := range table {
		table[i] = make([]int, len(table))
	}
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	//longest := s[:1]
	for i := 0; i < n; i++ {
		table[i][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 0; j < n-i+1; j++ {
			k := j + i - 1

			if s[j] == s[k] && i == 2 {
				table[j][k] = 2
			} else if s[j] == s[k] {
				table[j][k] = table[j+1][k-1] + 2
			} else {
				table[j][k] = max(table[j][k-1], table[j+1][k])
			}
		}
	}
	fmt.Println(table)
	fmt.Println(table[0][n-1])

	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "BBABCBCAB"
	longPalin(s)
}
