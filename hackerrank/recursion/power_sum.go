package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findPower(i, N int32) int32 {
	if i == 1 {
		return 1
	}
	if i%2 == 0 {
		return findPower(i, N/2) * findPower(i, N/2)
	}
	return N * findPower(i, N/2) * findPower(i, N/2)

}

func recursiveCheck(X int32, N int32, i int32, currSum int32) int32 {
	var power, result int32
	i = 1
	power = findPower(i, N)
	for {
		if power+currSum < X {
			currSum += power
			i++
			result += recursiveCheck(X, N, i, (power + currSum))
			power = findPower(i, N)
		} else {
			break
		}
	}
	if power+currSum == X {
		result++
	}
	return result
}

// Complete the powerSum function below.
func powerSum(X int32, N int32) int32 {
	return recursiveCheck(X, N, 1, 0)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	XTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	X := int32(XTemp)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	result := powerSum(X, N)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
