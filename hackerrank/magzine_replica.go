package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


func matchMagzine(a, b []string) {
	arr := make(map[string]int,len(a)+len(b))
	for _, v := range b{
		arr[v]++
	}
	//fmt.Println("arr",arr)
		for _, v := range a{
		arr[v]--
	}
		//fmt.Println("arr",arr)
		count := 0
		for _,v := range arr {
			if v > 0 {
				count++
			}
		}
		if count > 0 {
			fmt.Printf("%s","No")
		} else {
			fmt.Printf("%s","Yes")
		}

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    mn := strings.Split(readLine(reader), " ")

    mTemp, err := strconv.ParseInt(mn[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(mn[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    magazineTemp := strings.Split(readLine(reader), " ")

    var magazine []string

    for i := 0; i < int(m); i++ {
        magazineItem := magazineTemp[i]
        magazine = append(magazine, magazineItem)
    }

    ransomTemp := strings.Split(readLine(reader), " ")

    var ransom []string

    for i := 0; i < int(n); i++ {
        ransomItem := ransomTemp[i]
        ransom = append(ransom, ransomItem)
    }
	//fmt.Println("magzine",magazine)
	//fmt.Println("ransom",ransom)
	matchMagzine(magazine,ransom)
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
