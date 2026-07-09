package main

import (
	"slices"
	"testing"
)

func TestSortDesc(t *testing.T) {
	input := []int{2, 5, 6, 3, 1, 4}
	want := []int{6, 5, 4, 3, 2, 1}

	SortDesc(input)

	if !slices.Equal(input, want) {
		t.Errorf("SortDesc() = %v; expect %v", input, want)
	}
}
