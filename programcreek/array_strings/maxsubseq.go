package main

import (
	"fmt"
)

func longSubSeq(s string, n int) int {
	//var charMap map[rune]int
	var start, maxChar, startIndex, endIndex, maxResult int
	var result string
	charMap := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		endIndex++
		if _, ok := charMap[str]; ok {
			charMap[str]++
		} else {
			charMap[str] = 1

		}

		//startIndex = start

		if len(charMap) > n {
			//result = ""
			maxChar = max(maxChar, i-start)
			for len(charMap) > n {
				t := string(s[start])
				count := charMap[t]
				if count > 1 {
					charMap[t]--
				} else {
					delete(charMap, t)
				}
				startIndex++
				start++
			}
			//endIndex = start

		}
		if len(result) < endIndex-start+1 {
			result = s[start:endIndex]

		}

		// fmt.Println("value of startindex", startIndex)
		// fmt.Println("value of endIndex", endIndex)
		// fmt.Println("value of result", result)

		// fmt.Println("value of i", i)

	}
	//fmt.Println("start", charMap)
	maxChar = max(maxChar, len(s)-start)
	//fmt.Println("value of i", i)
	// fmt.Println("start", start)
	fmt.Println("value of result", result)

	fmt.Println("maxresult", maxResult)
	return maxChar
}
func main() {
	fmt.Println("original string is ", "abcbbbbcccbdddadacb")
	fmt.Println("long", longSubSeq("abcbbbbcccbdddadacb", 1))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
