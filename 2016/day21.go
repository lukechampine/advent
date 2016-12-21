package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `swap position 7 with position 1
swap letter e with letter d
swap position 7 with position 6
move position 4 to position 0
move position 1 to position 4
move position 6 to position 5
rotate right 1 step
swap letter e with letter b
reverse positions 3 through 7
swap position 2 with position 6
reverse positions 2 through 4
reverse positions 1 through 4
reverse positions 5 through 7
rotate left 2 steps
swap letter g with letter f
rotate based on position of letter a
swap letter b with letter h
swap position 0 with position 3
move position 4 to position 7
rotate based on position of letter g
swap letter f with letter e
move position 1 to position 5
swap letter d with letter e
move position 5 to position 2
move position 6 to position 5
rotate right 6 steps
rotate left 4 steps
reverse positions 0 through 3
swap letter g with letter c
swap letter f with letter e
reverse positions 6 through 7
move position 6 to position 1
rotate left 2 steps
rotate left 5 steps
swap position 3 with position 6
reverse positions 1 through 5
rotate right 6 steps
swap letter a with letter b
reverse positions 3 through 4
rotate based on position of letter f
swap position 2 with position 6
reverse positions 5 through 6
swap letter h with letter e
reverse positions 0 through 4
rotate based on position of letter g
rotate based on position of letter d
rotate based on position of letter b
swap position 5 with position 1
rotate based on position of letter f
move position 1 to position 5
rotate right 0 steps
rotate based on position of letter e
move position 0 to position 1
swap position 7 with position 2
rotate left 3 steps
reverse positions 0 through 1
rotate right 7 steps
rotate right 5 steps
swap position 2 with position 0
swap letter g with letter a
rotate left 0 steps
rotate based on position of letter f
swap position 5 with position 1
rotate right 0 steps
rotate left 5 steps
swap letter e with letter a
swap position 5 with position 4
reverse positions 2 through 5
swap letter e with letter a
swap position 3 with position 7
reverse positions 0 through 2
swap letter a with letter b
swap position 7 with position 1
move position 1 to position 6
rotate right 1 step
reverse positions 2 through 6
rotate based on position of letter b
move position 1 to position 0
swap position 7 with position 3
move position 6 to position 5
rotate right 4 steps
reverse positions 2 through 7
reverse positions 3 through 4
reverse positions 4 through 5
rotate based on position of letter f
reverse positions 0 through 5
reverse positions 3 through 4
move position 1 to position 2
rotate left 4 steps
swap position 7 with position 6
rotate right 1 step
move position 5 to position 2
rotate right 5 steps
swap position 7 with position 4
swap letter a with letter e
rotate based on position of letter e
swap position 7 with position 1
swap position 7 with position 3
move position 7 to position 1
swap position 7 with position 4`

type str []byte

func (s str) swapPos(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func (s str) swapLetter(x, y byte) {
	var xpos, ypos int
	for i, b := range s {
		if b == x {
			xpos = i
		}
		if b == y {
			ypos = i
		}
	}
	s.swapPos(xpos, ypos)
}

func (s str) rotateLeft(x int) {
	c := make(str, len(s))
	for i := range s {
		j := (i + x) % len(s)
		c[i] = s[j]
	}
	copy(s, c)
}

func (s str) rotateRight(x int) {
	c := make(str, len(s))
	for i := range s {
		j := (i + x) % len(s)
		c[j] = s[i]
	}
	copy(s, c)
}

func (s str) rotatePos(x byte) {
	var xpos int
	for xpos = range s {
		if s[xpos] == x {
			break
		}
	}
	if xpos >= 4 {
		xpos++
	}
	s.rotateRight(xpos + 1)
}

func (s str) undoRotatePos(x byte) {
	c := make(str, len(s))
	rot := 0
	for rot = range s {
		copy(c, s)
		c.rotateLeft(rot)
		c.rotatePos(x)
		if string(c) == string(s) {
			break
		}
	}
	s.rotateLeft(rot)
}

func (s str) reverse(x, y int) {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func (s str) move(x, y int) {
	c := s[x]
	if x > y {
		copy(s[y+1:], s[y:x])
	} else {
		copy(s[x:], s[x+1:y+1])
	}
	s[y] = c
}

func scramble(s str, input string) {
	for _, line := range utils.Lines(input) {
		words := strings.Fields(line)
		switch {
		case strings.HasPrefix(line, "swap position"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[5])
			s.swapPos(x, y)
		case strings.HasPrefix(line, "swap letter"):
			x := words[2][0]
			y := words[5][0]
			s.swapLetter(x, y)
		case strings.HasPrefix(line, "rotate left"):
			x := utils.Atoi(words[2])
			s.rotateLeft(x)
		case strings.HasPrefix(line, "rotate right"):
			x := utils.Atoi(words[2])
			s.rotateRight(x)
		case strings.HasPrefix(line, "rotate based"):
			x := words[len(words)-1][0]
			s.rotatePos(x)
		case strings.HasPrefix(line, "reverse"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[4])
			s.reverse(x, y)
		case strings.HasPrefix(line, "move"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[5])
			s.move(x, y)
		}
	}
}

func unscramble(s str, input string) {
	lines := utils.Lines(input)
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		words := strings.Fields(line)
		switch {
		case strings.HasPrefix(line, "swap position"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[5])
			s.swapPos(x, y)
		case strings.HasPrefix(line, "swap letter"):
			x := words[2][0]
			y := words[5][0]
			s.swapLetter(x, y)
		case strings.HasPrefix(line, "rotate left"):
			x := utils.Atoi(words[2])
			s.rotateRight(x)
		case strings.HasPrefix(line, "rotate right"):
			x := utils.Atoi(words[2])
			s.rotateLeft(x)
		case strings.HasPrefix(line, "rotate based"):
			x := words[len(words)-1][0]
			s.undoRotatePos(x)
		case strings.HasPrefix(line, "reverse"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[4])
			s.reverse(x, y)
		case strings.HasPrefix(line, "move"):
			x := utils.Atoi(words[2])
			y := utils.Atoi(words[5])
			s.move(y, x)
		}
	}
}

func main() {
	// part 1
	s := str("abcdefgh")
	scramble(s, input)
	println(string(s))

	s = str("fbgdceah")
	unscramble(s, input)
	println(string(s))
}
