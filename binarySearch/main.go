package main

import (
	"fmt"
	"sort"
)

func BinarySearch(a []int, target int) int {
	sort.Ints(a)
	start := 0
	end := len(a) - 1

	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			start = mid + 1
		} else if a[mid] > target {
			end = mid - 1
		}
	}
	return -1
}

func LinearSearch(a []int, target int) int {
	sort.Ints(a)

	for i := 0; i < len(a); i++ {
		if a[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	a := []int{5, 6, 8, 9, 11, 13, 17}
	sort.Ints(a)
	targ := 9
	fmt.Println(LinearSearch(a, targ))
}
