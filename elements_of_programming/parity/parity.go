package main

import (
	"fmt"
	"runtime"
)

func findParity(n int64) int64 {
	count := 0

	for n > 0 {
		n = n & (n - 1)
		count++
		PrintMemUsage()
	}
	if count%2 == 0 {
		return 0
	}
	return 1
}

func main() {
	//fmt.Println(findParity(13))

	fmt.Println(findParity(1345678964))
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
