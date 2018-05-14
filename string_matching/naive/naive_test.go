package main

import (
	"reflect"
	"testing"
)

type stringMatching struct {
	text    string
	pattern string
	matches []int
}

func testTables() []stringMatching {
	var stringmatch = []stringMatching{
		{"AABAACAADAABAAABAA", "AABA", []int{0, 9, 13}},
		{"AABAACAADAABAAABAA", "AAWA", []int{}},
	}
	return stringmatch
}

func TestfindPattern(t *testing.T) {
	testData := testTables()
	for _, tt := range testData {
		actual := findPattern(tt.text, tt.pattern)
		if !reflect.DeepEqual(actual, tt.matches) {
			t.Errorf("findPattern(%s,%s): expected %v, actual %v", tt.text, tt.pattern, tt.matches, actual)
		}
	}
}
