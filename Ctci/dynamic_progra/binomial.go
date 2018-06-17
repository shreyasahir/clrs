package main

import (
	"fmt"
)

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
func binomial(n, k int) int {
	if k == n {
		return 1
	}
	if k == 0 {
		return 1
	}
	return binomial(n-1, k-1) + binomial(n-1, k)
}
func main() {
	n := 5
	k := 2
	fmt.Println("binomial", dpbinomial(n, k))
}

func dpbinomial(n, k int) int {
	var table = make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, k+1)
	}
	for i := 0; i <= n; i++ {
		table[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= min(i, k); j++ {
			if i == j || j == 0 {
				table[i][j] = 1
			} else {
				table[i][j] = table[i-1][j-1] + table[i-1][j]
			}
		}
	}

	return table[n][k]
}
