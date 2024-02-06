package main

import "fmt"

func BinarySearch() int {
	list := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	val := 23
	lo := 0
	hi := len(list) - 1
	for lo < hi {

		m := lo + (hi-lo)/2
		if val == (list)[m] {
			return m
		} else if val < (list)[m] {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}
	return -1
}

func main(){
    fmt.Println(BinarySearch())
}
