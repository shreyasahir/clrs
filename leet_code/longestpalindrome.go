package main


func longestPalindrome(s string) string {
    start := 0
	end := len(s)-1
	count := 0
	for i:=0;i<end/2;i++ {
		if s[i] == s[end-i] {
			count +=2
		} else {
			count = 0
		}
	}
	print(count)
	return count
}

func main() {
longestPalindrome("babad")
}