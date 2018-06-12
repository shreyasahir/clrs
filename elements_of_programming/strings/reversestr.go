package main

import (
	"fmt"
)

func reverse(s string) []string{
	var result []string
	for i:=len(s)-1;i>=0;i-- {
		result = append(result,string(s[i]))
	}
	return result
}
func reverseStr(s1 string) {
s := reverse(s1)
//fmt.Println(s)
result := ""
subresult := ""
for _,v := range s {
//fmt.Println("value",v)
	if v == " " {
		result += subresult
		result += " "
		subresult = ""
		//fmt.Println("result in if",result)
	} else {
		subresult = v + subresult
	}
}
result += subresult
fmt.Println(result)
}

func main() {
	s:= "ram is costly"
	reverseStr(s)
}