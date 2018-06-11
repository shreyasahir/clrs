package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	count := 0
	maxtill := 0
    subhash := make(map[rune]int)
	for _,v := range s {
		if subhash[v] == 0 {
			subhash[v] += 1
			count++
		} else {
			if maxtill < count {
				maxtill = count
				subhash =  make(map[rune]int)
				subhash[v] += 1
				count = 1
			}
		}
	}
	if maxtill < count {
		maxtill = count
	}
	return maxtill
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}