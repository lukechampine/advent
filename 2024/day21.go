package main

import (
	"fmt"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2024, 21)

var numpad = utils.ToByteGrid(utils.Lines(`
789
456
123
x0A
`))

var dirpad = utils.ToByteGrid(utils.Lines(`
x^A
<v>
`))

var moveProgs = map[string][]string{
	"A^": {"<A"}, "^A": {">A"},
	"A>": {"vA"}, ">A": {"^A"},
	"<v": {">A"}, "v<": {"<A"},
	">v": {"<A"}, "v>": {">A"},
	"^<": {"v<A"}, "<^": {">^A"},
	"A<": {"v<<A"}, "<A": {">>^A"},
	"Av": {"<vA", "v<A"}, "vA": {">^A", "^>A"},
	"^>": {">vA", "v>A"}, ">^": {"<^A", "^<A"},
}

func dirArrow(d utils.Dir) string {
	return []string{"^", ">", "v", "<"}[d]
}

var memo = make(map[string]int)

func moveLen(a, b byte, dirpads int) int {
	key := fmt.Sprintf("%c-%c-%v", a, b, dirpads)
	if v, ok := memo[key]; ok {
		return v
	} else if a == b {
		return 1
	}
	progs := moveProgs[string(a)+string(b)]
	if dirpads == 0 {
		return len(progs[0])
	}
	memo[key] = utils.Minimum(len(progs), func(i int) int {
		return progLen(progs[i], dirpads-1)
	})
	return memo[key]
}

func progLen(code string, dirpads int) int {
	sum := moveLen('A', code[0], dirpads)
	for i := 1; i < len(code); i++ {
		x := moveLen(code[i-1], code[i], dirpads)
		sum += x
	}
	return sum
}

func shortestProg(code string, dirpads int) int {
	var rec func(int, utils.Pos, string) int
	rec = func(pressed int, pos utils.Pos, prog string) int {
		if pressed == len(code) {
			return progLen(prog, dirpads)
		}
		next := utils.Locate(numpad, code[pressed])
		dx, dy := utils.Signum(next.X-pos.X), utils.Signum(next.Y-pos.Y)
		if dx == 0 && dy == 0 {
			return rec(pressed+1, pos, prog+"A")
		} else if dx == 0 || dy == 0 {
			d := pos.DirTo(next)
			n := pos.Dist(next)
			return rec(pressed+1, pos.MoveArray(d, n), prog+strings.Repeat(dirArrow(d), n)+"A")
		} else {
			// multiple paths
			c1 := utils.Pos{X: next.X, Y: pos.Y}
			c2 := utils.Pos{X: pos.X, Y: next.Y}
			d1, n1 := pos.DirTo(c1), c1.Dist(pos)
			d2, n2 := pos.DirTo(c2), c2.Dist(pos)
			best := int(100e9)
			if numpad[c1.Y][c1.X] != 'x' {
				best = utils.Min(best, rec(pressed, pos.MoveArray(d1, n1), prog+strings.Repeat(dirArrow(d1), n1)))
			}
			if numpad[c2.Y][c2.X] != 'x' {
				best = utils.Min(best, rec(pressed, pos.MoveArray(d2, n2), prog+strings.Repeat(dirArrow(d2), n2)))
			}
			return best
		}
	}
	return rec(0, utils.Locate(numpad, 'A'), "")
}

func main() {
	codes := utils.Lines(input)
	utils.Println(utils.Sum(len(codes), func(i int) int {
		return utils.ExtractInts(codes[i])[0] * shortestProg(codes[i], 1)
	}))
	utils.Println(utils.Sum(len(codes), func(i int) int {
		return utils.ExtractInts(codes[i])[0] * shortestProg(codes[i], 24)
	}))
}
