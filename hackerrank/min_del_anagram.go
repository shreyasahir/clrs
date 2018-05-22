package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)


func mindelAnagram(a, b string){
	//var arr map[rune]int
	arr := make(map[int]int,26)
	for _,v:= range a {
		arr[int(v- 'a')]++
	}
	for _,v:= range b{
		arr[int(v- 'a')]--
	}
	//fmt.Println("arr",arr)
	var ans int
	for i:=0;i<26;i++ {
		if arr[i] < 0 {
			arr[i] *= -1
		}
		ans += arr[i]
	}
	fmt.Println(ans)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    a := readLine(reader)

    b := readLine(reader)
	//fmt.Println("a",a)
	//fmt.Println("b",b)
	mindelAnagram(a,b)

}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
