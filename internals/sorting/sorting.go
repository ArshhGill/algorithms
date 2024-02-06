package sorting

func BubbleSort(list []int) []int {
	for {
		c := 0
		for i := 0; i < len(list)-1; i++ {
			if (list)[i] > (list)[i+1] {
				temp := (list)[i]
				(list)[i] = (list)[i+1]
				(list)[i+1] = temp
				c++
			}
		}
		if c == 0 {
			break
		}
	}
	return list
}

func QuickSort(lst []int) {
	list := lst
	quickSortRunner(list, 0, len(list)-1)
}

func quickSortRunner(list []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotidx := pivot(&list, lo, hi)
	quickSortRunner(list, lo, pivotidx-1)
	quickSortRunner(list, pivotidx+1, hi)
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
