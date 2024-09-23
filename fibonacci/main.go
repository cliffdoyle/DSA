package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibRecurs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibRecurs(n-1) + fibRecurs(n-2)
}

func fibList(n int) int {
	// if n == 0 {
	// 	return []int{0}
	// }

	// if n == 1 {
	// 	return []int{1}
	// }
	arr := make([]int, n+1)

	arr[0] = 0

	arr[1] = 1
	// if n > 0 {
	// 	arr[1] = 1
	// }

	for i := 2; i <=n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("The %d fibonacci number is %d\n", n, fibList(n))
	// fmt.Println(fibRecurs(8))
}
