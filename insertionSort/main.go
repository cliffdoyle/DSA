package main

import (
	"fmt"
	"sort"
	// "sort"
)

func InsertionSort(a []int, pos, num int) {
	size := len(a)

	a[size-1] = a[pos-1]
	a[pos-1] = num

	fmt.Println(a)

	// a = append(a, 0)
	// for i := size-1; i >= 0; i-- {
	// 	a[i+1] = a[i]

	// }
	// fmt.Println(a)

	// a[0] = num
	// // size++

	// fmt.Println(a)
}

func DeletionArray(a []int, pos, num int) {
	size := len(a)
	sort.Ints(a)

	for i := 0; i < size-1; i++ {
		a[i] = a[i+1]
	}
	// a[0] = a[size-1]

	a = a[:size-1]
	fmt.Println(a)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	pos := 3
	num := 245
	DeletionArray(a, pos, num)
	// 	a[7]=2367

	// a=append(a, 234)
	// fmt.Println(a)
}
