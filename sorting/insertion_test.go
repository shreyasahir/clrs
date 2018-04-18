package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	x := []int{5, 2, 4, 6, 1, 3}
	y := []int{1, 2, 3, 4, 5, 6}
	sorted := insert(x)

	if !reflect.DeepEqual(sorted, y) {
		t.Errorf("Sorting was incorrect got %d", sorted)
	}
}
