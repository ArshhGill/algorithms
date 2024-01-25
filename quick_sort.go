package main

import "fmt"

func main() {
	list := []int{11, 16, 2, 8, 1, 9, 4, 7}
	quickSort(&list, 0, len(list)-1)
	fmt.Println(list)
}

func quickSort(list *[]int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotidx := pivot(list, lo, hi)
	quickSort(list, lo, pivotidx-1)
	quickSort(list, pivotidx+1, hi)
}

func pivot(list *[]int, lo, hi int) int {
	idx := lo
	pivotVal := (*list)[hi]

	for i := lo; i < hi; i++ {
		if (*list)[i] < pivotVal {
			temp := (*list)[idx]
			(*list)[idx] = (*list)[i]
			(*list)[i] = temp
			idx++
		}
	}
	(*list)[hi] = (*list)[idx]
	(*list)[idx] = pivotVal
	return idx
}
