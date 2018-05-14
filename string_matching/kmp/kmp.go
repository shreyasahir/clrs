package main

import (
	"fmt"
)

func computeLPS(pat string, m int, lps []int) {
	length := 0
	i := 1

	for i < m {
		if pat[i] == pat[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
}

func kmp(txt, pat string) []int {
	m := len(pat)
	n := len(txt)
	j := 0
	lps := make([]int, m)
	result := []int{}
	computeLPS(pat, m, lps)
	for i := 0; i < n; {
		if pat[j] == txt[i] {
			i++
			j++
		}

		if j == m {
			fmt.Println("Pattern matched", i-j)
			result = append(result, i-j)
			j = lps[j-1]
		} else if i < n && pat[j] != txt[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return result
}
func main() {
	var txt = "AABAACAADAABAAABAA"
	var pat = "AABA"
	kmp(txt, pat)
}
