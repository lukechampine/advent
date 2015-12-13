package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Min returns the greater of x and y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min returns the lesser of x and y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// And returns true if all of its arguments are true.
func And(preds ...bool) bool {
	for _, pred := range preds {
		if !pred {
			return false
		}
	}
	return true
}

// And returns true if any of its arguments are true.
func Or(preds ...bool) bool {
	for _, pred := range preds {
		if pred {
			return true
		}
	}
	return false
}

// Lines splits a string by newlines.
func Lines(input string) []string {
	return strings.Split(input, "\n")
}

// Itoa is a passthrough for strconv.Itoa
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// Itoa is a passthrough for strconv.Atoi that panics upon failure.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Sscanf is a passthrough for fmt.Sscanf that panics upon failure.
func Sscanf(str, format string, args ...interface{}) {
	_, err := fmt.Sscanf(str, format, args...)
	if err != nil {
		panic(err)
	}
}

func interleave(n int, perms [][]int) [][]int {
	leaved := make([][]int, 0, len(perms)*(n+1))
	for _, perm := range perms {
		for i := 0; i <= len(perm); i++ {
			withN := make([]int, len(perm)+1)
			copy(withN[:i], perm[:i])
			withN[i] = n
			copy(withN[i+1:], perm[i:])
			leaved = append(leaved, withN)
		}
	}
	return leaved
}

// Perms returns all possible permutations of the numbers [0,n).
func Perms(n int) [][]int {
	if n == 1 {
		return [][]int{{0}}
	}
	perms := Perms(n - 1)
	return interleave(n-1, perms)
}
