package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 5)

func move(stack *[9][]byte, n, from, to int) {
	f, t := &stack[from-1], &stack[to-1]
	*t = append(*t, (*f)[len(*f)-n:]...)
	*f = (*f)[:len(*f)-n]
}

func main() {
	groups := utils.Split(input, "\n\n")
	lines := utils.Lines(groups[0])
	// first line may be misaligned due to trimming spaces; realign it
	for len(lines[0]) < len(lines[1]) {
		lines[0] = " " + lines[0]
	}

	var stacks1, stacks2 [9][]byte
	for i := range lines[1:] {
		for j := range stacks1 {
			c := lines[i][j*4+1]
			if c != ' ' {
				stacks1[j] = append([]byte{c}, stacks1[j]...)
				stacks2[j] = append([]byte{c}, stacks2[j]...)
			}
		}
	}

	for _, line := range utils.Lines(groups[1]) {
		var n, from, to int
		utils.Sscanf(line, "move %d from %d to %d", &n, &from, &to)
		for i := 0; i < n; i++ {
			move(&stacks1, 1, from, to)
		}
		move(&stacks2, n, from, to)
	}
	var r1, r2 string
	for i := range stacks1 {
		r1 += string(stacks1[i][len(stacks1[i])-1])
		r2 += string(stacks2[i][len(stacks2[i])-1])
	}
	utils.Println(r1)
	utils.Println(r2)
}
