package basic

import "testing"

func TestQuickSort(t *testing.T) {
	input := []int{5, 2, 1, 3}
	quickSort(input)
	expect := []int{1, 2, 3, 5}
	if len(expect) != len(input) {
		t.Fatalf("length is not same")
	}
	for i, _ := range expect {
		if expect[i] != input[i] {
			t.Fatalf("index:%v, expect:%v, but input:%v", i, expect[i], input[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	input := []int{5, 2, 1, 3}
	mergeSort(input)
	expect := []int{1, 2, 3, 5}
	if len(expect) != len(input) {
		t.Fatalf("length is not same")
	}
	for i, _ := range expect {
		if expect[i] != input[i] {
			t.Fatalf("index:%v, expect:%v, but input:%v", i, expect[i], input[i])
		}
	}
}
