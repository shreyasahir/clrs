package main

import (
	"fmt"
)

func coinChange(s []int, m int, n int) int {

	if m <= 0 {
		return 0
	}
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return coinChange(s, m, n-s[m-1]) + coinChange(s, m-1, n)
}

func dpcoinChange(s []int, m int, n int) int {
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		dp[0][i] = 1
	}
	var x, y int

	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			if i-s[j] >= 0 {
				x = dp[i-s[j]][j]
			} else {
				x = 0
			}
			if j >= 1 {
				y = dp[i][j-1]
			} else {
				y = 0
			}
			dp[i][j] = x + y
		}
	}

	return dp[n][m-1]

}
func main() {
	n := 4
	s := []int{1, 2, 3}
	fmt.Println("ways", dpcoinChange(s, len(s), n))
}
