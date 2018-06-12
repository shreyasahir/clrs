package main

import (
	"fmt"
	"runtime"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return 1
	}
	PrintMemUsage()
	return fib(n-1) + fib(n-2)
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

func main() {
	start := time.Now()
	result := fib(30)
	fmt.Println("Fib is ", result)
	fmt.Println("Time since", time.Since(start))
	fmt.Println("Forcing garbage collection")
	runtime.GC()
	PrintMemUsage()

}
