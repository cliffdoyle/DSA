package main

import (
	"fmt"
	"os"
	"strconv"
)

func NaiveGcd(a, b int) int {
	best := 0 // remembers the biggest thing we've seen so far

	for d := 1; d <= a+b; d++ {
		if a%d == 0 && b%d == 0 {
			best = d
		}
	}

	return best
}

func EfficientGcd(a, b int) int {
	if b == 0 {
		return a
	}

	d := a % b
	return EfficientGcd(b, d)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	m, _ := strconv.Atoi(os.Args[2])

	fmt.Println(EfficientGcd(n, m))
}
