package main

import "math"

const input = 3004953

func closestpow(base float64, n int) int {
	return int(math.Pow(base, math.Floor(math.Log(float64(n))/math.Log(base))))
}

func next(n int) int {
	i := closestpow(2, n)
	return 2*n - 2*i + 1
}

func across(n int) int {
	if n == 1 {
		return 1
	}
	i := closestpow(3, n-1)
	dist := n - i
	if dist < i {
		return dist
	} else {
		return 2*n - 3*i
	}
}

func main() {
	// part 1
	println(next(input))

	// part 2
	println(across(input))
}
