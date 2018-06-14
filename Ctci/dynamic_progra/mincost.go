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

func minCost(cost [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 100000000
	} else if i == 0 && j == 0 {
		return cost[i][j]
	} else {
		return cost[i][j] + min(minCost(cost, i-1, j), minCost(cost, i-1, j-1), minCost(cost, i, j-1))
	}
}

func dpminCost(cost [][]int, m, n int) {
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = cost[0][0]
	for i := 1; i < m+1; i++ {
		dp[i][0] = cost[i][0] + dp[i-1][0]
	}

	for i := 1; i < m+1; i++ {
		dp[0][i] = cost[0][i] + dp[0][i-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			dp[i][j] = cost[i][j] + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])

		}
	}
	fmt.Println("dp", dp)
	fmt.Println("dp value", dp[m][n])
}

func main() {
	cost := [][]int{{1, 2, 3},
		{4, 8, 2},
		{1, 5, 3}}

	//fmt.Println("mincost", minCost(cost, 2, 2))
	dpminCost(cost, 2, 2)
}
