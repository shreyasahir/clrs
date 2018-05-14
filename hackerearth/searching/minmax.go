package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	inputSlice []string
	text       string
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()
		break
	}
	inputSlice = strings.Fields(text)

}

func main() {
	var num, i int32
	var y, sum, min, max int64
	min = -1
	i = 1
	fmt.Scanf("%d", &num)
	scanner()
	for _, v := range inputSlice {
		//fmt.Println(v)
		y, _ = strconv.ParseInt(v, 10, 64)
		if i == 1 {
			min = y
		}
		if min > y {
			min = y
		}
		if max < y {
			max = y
		}
		sum += y
		i++
		if i > num {
			break
		}
	}
	fmt.Printf("%d %d", sum-max, sum-min)

}
