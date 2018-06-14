package main

import (
	"fmt"
)

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

func editDistance(s, t string, m, n int) int {

	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if s[m-1] == t[n-1] {
		return editDistance(s, t, m-1, n-1)
	}
	return 1 + min(editDistance(s, t, m, n-1), editDistance(s, t, m-1, n), editDistance(s, t, m-1, n-1))

}

func dpeditDistance(s, t string, m, n int) int {

	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func main() {
	s := "sunday"
	t := "saturday"
	fmt.Println(dpeditDistance(s, t, len(s), len(t)))
}
