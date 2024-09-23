package main

func MinRefils(n, l int, x []int) int {
	numRefils := 0
	currentRefil := 0

	for currentRefil <= n {
		lastRefil := currentRefil
		for currentRefil <= n && x[currentRefil+1]-x[lastRefil] <= l {
			currentRefil = currentRefil + 1
			if currentRefil == lastRefil {
				return 0
			}
			if currentRefil <= n {
				numRefils = numRefils + 1
			}
		}

	}
	return numRefils
}
