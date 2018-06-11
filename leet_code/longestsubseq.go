package main

import (
	"fmt"
)
func longSubSeq(s string) int {
	var ans,n int
	n = len(s)
	arr := make(map[rune]int)
	s1 := []rune(s)
	for i,j:=0,0;j<n;j++ {
		if arr[s1[j]] != 0 {
			i = max(arr[s1[j]],i)
		}
		ans = max(ans,j-i+1)
		arr[s1[j]] = j+1
	}
	fmt.Println("arr",arr)
	return ans
}
func main () {
	fmt.Println("long",longSubSeq("abcde"))
}

func max(a,b int) int{
	if a >= b {
		return a
	}
	return b
}

    // public int lengthOfLongestSubstring(String s) {
    //     int n = s.length(), ans = 0;
    //     Map<Character, Integer> map = new HashMap<>(); // current index of character
    //     // try to extend the range [i, j]
    //     for (int j = 0, i = 0; j < n; j++) {
    //         if (map.containsKey(s.charAt(j))) {
    //             i = Math.max(map.get(s.charAt(j)), i);
    //         }
    //         ans = Math.max(ans, j - i + 1);
    //         map.put(s.charAt(j), j + 1);
    //     }
    //     return ans;
    // }