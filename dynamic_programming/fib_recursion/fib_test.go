package main

import (
	"testing"
)

type fibo struct {
	key   int
	value int
}

func testTables() []fibo {
	var fibotests = []fibo{
		{5, 8},
		{6, 13},
		{1, 1},
		{0, 1},
	}
	return fibotests
}
func TestFib(t *testing.T) {
	testData := testTables()
	for _, tt := range testData {
		actual := fib(tt.key)
		if actual != tt.value {
			t.Errorf("fib(%d): expected %d, actual %d", tt.key, tt.value, actual)
		}
	}
}
