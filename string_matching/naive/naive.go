package main

import (
//"fmt"
)

func findPattern(txt, pat string) []int {
	t := len(txt)
	p := len(pat)
	result := []int{}
	for i := 0; i < (t-p)+1; i++ {
		for j := 0; j < p; j++ {
			if txt[i+j] != pat[j] {
				break
			}

			if j == p-1 {
				//fmt.Println("Pattern found at i", i)
				result = append(result, i)
			}
		}
	}
	return result
}
func main() {
	var txt = "AABAACAADAABAAABAA"
	var pat = "AABA"
	findPattern(txt, pat)
}
