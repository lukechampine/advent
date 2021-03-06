package main

import "github.com/lukechampine/advent/utils"

const input = `R4, R4, L1, R3, L5, R2, R5, R1, L4, R3, L5, R2, L3, L4, L3, R1, R5, R1, L3, L1, R3, L1, R2, R2, L2, R5, L3, L4, R4, R4, R2, L4, L1, R5, L1, L4, R4, L1, R1, L2, R5, L2, L3, R2, R1, L194, R2, L4, R49, R1, R3, L5, L4, L1, R4, R2, R1, L5, R3, L5, L4, R4, R4, L2, L3, R78, L5, R4, R191, R4, R3, R1, L2, R1, R3, L1, R3, R4, R2, L2, R1, R4, L5, R2, L2, L4, L2, R1, R2, L3, R5, R2, L3, L3, R3, L1, L1, R5, L4, L4, L2, R5, R1, R4, L3, L5, L4, R5, L4, R5, R4, L3, L2, L5, R4, R3, L3, R1, L5, R5, R1, L3, R2, L5, R5, L3, R1, R4, L5, R4, R2, R3, L4, L5, R3, R4, L5, L5, R4, L4, L4, R1, R5, R3, L1, L4, L3, L4, R1, L5, L1, R2, R2, R4, R4, L5, R4, R1, L1, L1, L3, L5, L2, R4, L3, L5, L4, L1, R3`

type dir struct {
	left   bool
	length int
}

func parse(str string) []dir {
	var dirs []dir
	for _, d := range utils.Split(str, ", ") {
		dirs = append(dirs, dir{
			left:   d[0] == 'L',
			length: utils.Atoi(d[1:]),
		})
	}
	return dirs
}

const (
	// turning left
	north = iota
	west
	south
	east
)

func walk(p *utils.Pos, orientation int, length int) {
	switch orientation {
	case north:
		p.Y += length
	case south:
		p.Y -= length
	case east:
		p.X += length
	case west:
		p.X -= length
	}
}

func follow(dirs []dir) utils.Pos {
	var p utils.Pos // starting position
	o := north      // starting orientation
	for _, d := range dirs {
		if d.left {
			o = (o + 1) % 4
		} else {
			o = (o + 3) % 4
		}
		walk(&p, o, d.length)
	}
	return p
}

func visitTwice(dirs []dir) utils.Pos {
	visited := make(map[utils.Pos]bool)
	var p utils.Pos // starting position
	o := north      // starting orientation
	visited[p] = true
	for _, d := range dirs {
		if d.left {
			o = (o + 1) % 4
		} else {
			o = (o + 3) % 4
		}
		// need to record every step
		for i := 0; i < d.length; i++ {
			walk(&p, o, 1)
			if visited[p] {
				return p
			}
			visited[p] = true
		}
	}
	return p
}

func main() {
	// part 1
	dirs := parse(input)
	end := follow(dirs)
	println(end.Dist(utils.Pos{0, 0}))

	// part 2
	end = visitTwice(dirs)
	println(end.Dist(utils.Pos{0, 0}))
}
