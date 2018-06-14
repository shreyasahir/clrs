package main

import (
	"fmt"
)

// a := make([][]uint8, dy)
// for i := range a {
//     a[i] = make([]uint8, dx)
// }
func findSeq(s1, s2 string) int {
	m := len(s2)
	n := len(s1)
	if m > n {
		return 0
	}
	table := make([][]int, m+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}
	fmt.Println(table)
	for i := 0; i <= m; i++ {
		table[i][0] = 0
	}
	for i := 0; i <= n; i++ {
		table[0][i] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Println("s2[i-1]", string(s2[i-1]))
			fmt.Println("s1[j-1]", string(s1[j-1]))
			fmt.Println("table", table)

			if s2[i-1] == s1[j-1] {
				table[i][j] = table[i][j-1] + table[i-1][j-1]
			} else {
				table[i][j] = table[i][j-1]
			}
		}
	}
	fmt.Println("end of func", table)
	return table[m][n]
}

func main() {
	s := "geeksforgeeks"
	t := "ge"
	fmt.Println(findSeq(s, t))
}
