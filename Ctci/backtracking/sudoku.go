package main

import (
	"fmt"
)

var (
	grid = [][]int{}
)

func printGrid() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Println()
	}
}

func usedInRow(row, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func usedInColumn(col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(row, col, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+row][j+col] == num {
				return true
			}
		}
	}
	return false
}
func findEmptyLocation(l []int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				l[0] = i
				l[1] = j
				return true
			}
		}

	}
	return false
}
func checkLocationSafe(row, col, num int) bool {
	return !usedInRow(row, num) && !usedInColumn(col, num) && !usedInBox(row-row%3, col-col%3, num)
}
func solveSudoku() bool {
	l := []int{0, 0}
	if !findEmptyLocation(l) {
		return true
	}
	row := l[0]
	col := l[1]

	for k := 0; k < 10; k++ {
		if checkLocationSafe(row, col, k) {
			grid[row][col] = k
			if solveSudoku() {
				return true
			}
		}
		grid[row][col] = 0
	}
	return false
}
func main() {
	// grid = make([][]int, 9)
	// for i := range grid {
	// 	grid[i] = make([]int, len(grid))
	// }
	grid = [][]int{{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0}}

	if solveSudoku() {
		printGrid()
	} else {
		fmt.Println("No solution exists")
	}
}
