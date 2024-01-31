package main

import "fmt"

func main() {
	list := []int{11, 16, 2, 8, 1, 9, 4, 7}
	bubbleSort(&list)
	fmt.Println(list)
}

func bubbleSort(list *[]int) {
	for true {
		c := 0
		for i := 0; i < len(*list)-1; i++ {
			if (*list)[i] > (*list)[i+1] {
				temp := (*list)[i]
				(*list)[i] = (*list)[i+1]
				(*list)[i+1] = temp
				c++
			}
		}
		if c == 0 {
			break
		}
	}
}
