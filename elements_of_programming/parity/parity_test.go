package main

import (
	"testing"
)

type testData struct {
	input int64
	out   int64
}

func testTable() []testData {
	return []testData{
		{8, 1},
		{0, 0},
		{5, 0},
		{13, 1},
	}
}
func TestFindParity(t *testing.T) {
	data := testTable()

	for _, v := range data {
		actual := findParity(v.input)
		if actual != v.out {
			t.Errorf("findParity(%d): expected %d, actual %d", v.input, v.out, actual)
		}
	}
}
