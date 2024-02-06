package sorting_test

import (
	. "algo/internals/sorting"
	"testing"
)

func TestSorting(t *testing.T) {

	arr := []int{1, 23, 5, 5, 67, 34, 34, 69, 0}
	testList(BubbleSort(arr), []int{0, 1, 5, 5, 23, 34, 34, 67, 69}, t)
}

func testList(ls []int, arr []int, t *testing.T) {
	if len(ls) != len(arr) {
		t.Fatalf("list %v got %v", arr, ls)
	}

	for index, i := range arr {
		if ls[index] != i {
			t.Fatalf("%v got %v", arr, ls)
		}
	}
}
