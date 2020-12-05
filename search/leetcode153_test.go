package search

import "testing"

func TestFindMin(t *testing.T) {
	input := []int{3, 4, 5, 1, 2}
	result := findMin(input)
	if result != 1 {
		t.Fail()
	}
}
