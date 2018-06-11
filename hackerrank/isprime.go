package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math"
)

func isPrime(n int32) bool{
	var i int32
	for i =2;i <= int32(math.Sqrt(float64(n)));i++ {
		if n % i == 0 {
			//fmt.Printf("%s","Prime")
			return true
		}
	}
	return false
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    p := int32(pTemp)

    for pItr := 0; pItr < int(p); pItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)
		if n == 1 {
			fmt.Println("Not prime")
		} else if n < 4 {
			fmt.Println("Prime")
		} else {
			ans := isPrime(n)
			if !ans {
				fmt.Println("Prime")
			} else {
				fmt.Println("Not prime")
			}
		}
    }
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
