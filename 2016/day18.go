package main

import "strings"

const input = `^.^^^.^..^....^^....^^^^.^^.^...^^.^.^^.^^.^^..^.^...^.^..^.^^.^..^.....^^^.^.^^^..^^...^^^...^...^.`

func tile(prev []byte, t int) byte {
	left := t-1 >= 0 && prev[t-1] == '^'
	center := t > 0 && prev[t] == '^'
	right := t+1 < len(prev) && prev[t+1] == '^'
	isTrap := (left && center && !right) ||
		(!left && center && right) ||
		(left && !center && !right) ||
		(!left && !center && right)
	if isTrap {
		return '^'
	} else {
		return '.'
	}
}

func count(first string, n int) int {
	total := strings.Count(first, ".")
	prev := []byte(first)
	next := []byte(first)
	for i := 1; i < n; i++ {
		for j := range next {
			next[j] = tile(prev, j)
			if next[j] == '.' {
				total++
			}
		}
		copy(prev, next)
	}
	return total
}

func main() {
	// part 1
	println(count(input, 40))

	// part 2
	println(count(input, 400000))
}
