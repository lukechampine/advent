package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 10)

func minWiringPresses(goal string, buttons [][]int) int {
	applyButton := func(lights string, button []int) string {
		runes := []rune(lights)
		for _, i := range button {
			i++
			if runes[i] == '#' {
				runes[i] = '.'
			} else {
				runes[i] = '#'
			}
		}
		return string(runes)
	}

	type stackEntry struct {
		lights  string
		presses int
	}
	stack := []stackEntry{{strings.ReplaceAll(goal, "#", "."), 0}}
	dist := make(map[string]int)

	for len(stack) > 0 {
		entry := stack[0]
		stack = stack[1:]
		if entry.lights == goal {
			return entry.presses
		} else if d, ok := dist[entry.lights]; ok && d <= entry.presses {
			continue
		}
		dist[entry.lights] = entry.presses

		for _, button := range buttons {
			stack = append(stack, stackEntry{
				lights:  applyButton(entry.lights, button),
				presses: entry.presses + 1,
			})
		}
	}
	panic("unreachable")
}

func minJoltagePresses(goal []int, buttons [][]int) int {
	// Build integer matrix
	matrix := make([][]int, len(goal))
	for i := range matrix {
		matrix[i] = make([]int, len(buttons)+1)
	}
	for j, button := range buttons {
		for _, i := range button {
			matrix[i][j] = 1
		}
	}
	for i := range goal {
		matrix[i][len(buttons)] = goal[i]
	}

	// Gaussian elimination
	numRows := len(matrix)
	numCols := len(matrix[0])
	row := 0
	for col := 0; col < numCols-1 && row < numRows; col++ {
		pivotRow := -1
		for r := row; r < numRows; r++ {
			if matrix[r][col] != 0 {
				pivotRow = r
				break
			}
		}
		if pivotRow == -1 {
			continue
		}
		matrix[row], matrix[pivotRow] = matrix[pivotRow], matrix[row]

		for r := 0; r < numRows; r++ {
			if r != row && matrix[r][col] != 0 {
				a := matrix[r][col]
				b := matrix[row][col]
				for c := 0; c < numCols; c++ {
					matrix[r][c] = matrix[r][c]*b - matrix[row][c]*a
				}
				// simplify
				g := 0
				for c := 0; c < len(matrix[r]); c++ {
					g = utils.GCD(g, matrix[r][c])
				}
				if g > 1 {
					for c := 0; c < len(matrix[r]); c++ {
						matrix[r][c] /= g
					}
				}
			}
		}
		row++
	}

	var freeVars []int
	for c := 0; c < numCols-1; c++ {
		nonZero := 0
		for r := 0; r < numRows; r++ {
			if matrix[r][c] != 0 {
				nonZero++
			}
		}
		if nonZero > 1 {
			freeVars = append(freeVars, c)
		}
	}

	totalPressed := func(freePresses []int) (int, bool) {
		presses := make([]int, numCols-1)
		for i, v := range freeVars {
			presses[v] = freePresses[i]
		}
		for r := numRows - 1; r >= 0; r-- {
			sum := matrix[r][numCols-1]
			pivotCol := -1
			for c := 0; c < numCols-1; c++ {
				if matrix[r][c] != 0 {
					if pivotCol == -1 {
						pivotCol = c
					} else {
						sum -= matrix[r][c] * presses[c]
					}
				}
			}
			if pivotCol != -1 {
				if sum%matrix[r][pivotCol] != 0 {
					return 0, false
				}
				presses[pivotCol] = sum / matrix[r][pivotCol]
				if presses[pivotCol] < 0 {
					return 0, false
				}
			}
		}
		total := 0
		for _, p := range presses {
			total += p
		}
		return total, true
	}

	// Brute force over free variables
	perm := make([]int, len(freeVars))
	max := utils.Maximum(len(goal), func(i int) int { return goal[i] })
	min := int(1e9)
	for {
		total, ok := totalPressed(perm)
		if ok && total < min {
			min = total
		}
		i := 0
		for i < len(perm) && perm[i] == max {
			perm[i] = 0
			i++
		}
		if i == len(perm) {
			break
		}
		perm[i]++
	}
	if min == int(1e9) {
		panic("no solution")
	}
	return min
}

func main() {
	var part1, part2 int
	for _, line := range utils.Lines(input) {
		lights, rest, _ := strings.Cut(line, " ")
		parts := strings.Split(rest, " ")
		buttons := make([][]int, len(parts)-1)
		for i, part := range parts[:len(parts)-1] {
			buttons[i] = utils.ExtractInts(part)
		}
		joltage := utils.ExtractInts(parts[len(parts)-1])
		part1 += minWiringPresses(lights, buttons)
		part2 += minJoltagePresses(joltage, buttons)
	}
	utils.Println(part1)
	utils.Println(part2)
}
