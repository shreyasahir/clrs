package main

import (
	"fmt"
)

func longSubSeq(s string) int {
	var ans, n int
	n = len(s)
	arr := make(map[string]int)
	s1 := []rune(s)
	for i, j := 0, 0; j < n; j++ {
		s = string(s1[j])
		if arr[s] != 0 {
			i = max(arr[s], i)
		}
		fmt.Println("i", i, " j", j)
		ans = max(ans, j-i+1)
		fmt.Println("ans", ans)
		arr[s] = j + 1
	}
	fmt.Println("arr", arr)
	return ans
}
func main() {
	fmt.Println("long", longSubSeq("abcabcbb"))
}

func max(a, b int) int {
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
