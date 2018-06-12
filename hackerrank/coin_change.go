package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)


func coinChange(coins []int32,n,m int32) int32{

	var i,j,x,y int32
	result := make([][]int32,n+1)
	for k,_ := range result {
		result[k] = make([]int32,m)
	}
	for i=0;i<m;i++ {
		result[0][i] = 1
	}
	fmt.Println("result",result)

//var x,y int

	for i=0;i<n;i++ {
		for j=0;j<m;j++{
			if i -coins[j] > 0 {
				x = result[i-coins[j]][j]
			} else {
				x = 0
			}
			if j >= 1 {
				y = result[i][j-1]
			} else {
				y = 0
			}
			result[i][j] = x + y
		}
	}

	fmt.Println("result",result)
	return result[n][m-1]
}


func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    coinsTemp := strings.Split(readLine(reader), " ")

    var coins []int32

    for i := 0; i < int(m); i++ {
        coinsItemTemp, err := strconv.ParseInt(coinsTemp[i], 10, 64)
        checkError(err)
        coinsItem := int32(coinsItemTemp)
        coins = append(coins, coinsItem)
    }
	
	fmt.Println("coinchange",coinChange(coins,n,m))
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
