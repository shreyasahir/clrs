package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	text       string
	inputSlice []string
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text = scanner.Text()
		//fmt.Println("text", text)
		break
	}
	inputSlice = strings.Fields(text)

}

func main() {
	var noOfCases int32
	var k int64
	var inputArr []int64
	fmt.Scanf("%d", &noOfCases)
	scanner()
	fmt.Scanf("%d", &k)
	for _, v := range inputSlice {
		v1, _ := strconv.ParseInt(v, 10, 32)
		inputArr = append(inputArr, v1)

	}
	sor := countingSort(inputArr)

	var count, prev int64
	prev = -1
	//fmt.Println("sor", sor)
	for _, v := range sor {
		if v != prev {
			count = 1
		}
		if v == prev {
			count++
		}
		if count == k {
			break
		}
		prev = v

	}
	fmt.Println(prev)
}

func countingSort(arr []int64) []int64 {

	highest := findHighest(arr)

	//fmt.Println("Highest", highest)
	aux := make([]int, highest+1)

	sortedArr := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}

	var j int64
	j = 0
	for i := 0; i < len(aux); i++ {
		tmp := aux[i]
		for ; tmp > 0; tmp-- {
			sortedArr[j] = int64(i)
			j++
		}
	}
	return sortedArr
	//fmt.Println("Sorted Array", sortedArr)
}

func findHighest(arr []int64) int64 {
	highest := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > highest {
			highest = arr[i]
		}
	}
	return highest
}
