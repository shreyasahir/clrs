package main

import (
	"fmt"
)

func longcomm(x, y string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if x[m-1] == y[n-1] {
		return 1 + longcomm(x, y, m-1, n-1)
	}
	return max(longcomm(x, y, m, n-1), longcomm(x, y, m-1, n))

}

func dplongcomm(x, y string, m, n int) {
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	fmt.Println("dp", dp)
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if x[i-1] == y[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				fmt.Println("i&j", i, j)
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println("dp", dp)
	fmt.Println("dp", dp[m][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := "AGGTAB"
	y := "GXTXAYB"
	m := len(x)
	n := len(y)
	//	fmt.Println("longcomm", longcomm(x, y, m, n))
	dplongcomm(x, y, m, n)
}
