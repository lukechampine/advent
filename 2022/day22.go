package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 22)

func main() {
	parts := utils.Split(input, "\n\n")
	g := utils.ToByteGrid(strings.Split(parts[0], "\n"))

	var cmds []string
	start := 0
	for i, c := range strings.TrimSpace(parts[1]) {
		if c == 'L' || c == 'R' {
			cmds = append(cmds, parts[1][start:i])
			start = i
			cmds = append(cmds, parts[1][start:i+1])
			start++
		}
	}
	cmds = append(cmds, strings.TrimSpace(parts[1][start:]))

	type face struct {
		pos utils.Pos
		adj [4]*face
	}
	onFace := func(p utils.Pos, f *face) bool {
		return (0 <= p.X && p.X < 50) && (0 <= p.Y && p.Y < 50)
	}
	tile := func(p utils.Pos, f *face) byte {
		return g[f.pos.Y+p.Y][f.pos.X+p.X]
	}
	crossEdge := func(p utils.Pos, d utils.Dir, f *face) (utils.Pos, utils.Dir, *face) {
		nf := f.adj[d]
		nd := d
		np := map[utils.Dir]utils.Pos{
			utils.Up:    {p.X, 49},
			utils.Right: {0, p.Y},
			utils.Down:  {p.X, 0},
			utils.Left:  {49, p.Y},
		}[d]
		return np, nd, nf
	}

	var faces [7]face
	faces[1].pos = utils.Pos{50, 0}
	faces[2].pos = utils.Pos{100, 0}
	faces[3].pos = utils.Pos{50, 50}
	faces[4].pos = utils.Pos{0, 100}
	faces[5].pos = utils.Pos{50, 100}
	faces[6].pos = utils.Pos{0, 150}
	for i, adj := range [][]int{
		1: {utils.Up: 5, utils.Right: 2, utils.Down: 3, utils.Left: 2},
		2: {utils.Up: 2, utils.Right: 1, utils.Down: 2, utils.Left: 1},
		3: {utils.Up: 1, utils.Right: 3, utils.Down: 5, utils.Left: 3},
		4: {utils.Up: 6, utils.Right: 5, utils.Down: 6, utils.Left: 5},
		5: {utils.Up: 3, utils.Right: 4, utils.Down: 1, utils.Left: 4},
		6: {utils.Up: 4, utils.Right: 6, utils.Down: 4, utils.Left: 6},
	} {
		for d, j := range adj {
			faces[i].adj[d] = &faces[j]
		}
	}

	score := func() int {
		curFace := &faces[1]
		a := &utils.Agent{
			Pos: utils.Pos{0, 0},
			Dir: utils.Right,
		}
		for tile(a.Pos, curFace) != '.' {
			a.Pos.X++
		}
		for _, c := range cmds {
			switch c {
			case "L":
				a.TurnLeft()
			case "R":
				a.TurnRight()
			default:
				for i := 0; i < utils.Atoi(c); i++ {
					p := a.Pos.MoveArray(a.Dir, 1)
					d := a.Dir
					cf := curFace
					if !onFace(p, cf) {
						p, d, cf = crossEdge(a.Pos, a.Dir, curFace)
					}
					if tile(p, cf) != '#' {
						a.Pos = p
						a.Dir = d
						curFace = cf
					}
				}
			}
		}
		a.Pos.X += curFace.pos.X
		a.Pos.Y += curFace.pos.Y

		return 1000*(a.Y+1) + 4*(a.X+1) + map[utils.Dir]int{
			utils.Right: 0,
			utils.Down:  1,
			utils.Left:  2,
			utils.Up:    3,
		}[a.Dir]
	}

	utils.Println(score())

	// part 2
	for i, adj := range [][]int{
		1: {utils.Up: 6, utils.Right: 2, utils.Down: 3, utils.Left: 4},
		2: {utils.Up: 6, utils.Right: 5, utils.Down: 3, utils.Left: 1},
		3: {utils.Up: 1, utils.Right: 2, utils.Down: 5, utils.Left: 4},
		4: {utils.Up: 3, utils.Right: 5, utils.Down: 6, utils.Left: 1},
		5: {utils.Up: 3, utils.Right: 2, utils.Down: 6, utils.Left: 4},
		6: {utils.Up: 4, utils.Right: 5, utils.Down: 2, utils.Left: 1},
	} {
		for d, j := range adj {
			faces[i].adj[d] = &faces[j]
		}
	}

	crossEdge = func(p utils.Pos, d utils.Dir, f *face) (utils.Pos, utils.Dir, *face) {
		nf := f.adj[d]
		var nd utils.Dir
		for i := range nf.adj {
			if nf.adj[i] == f {
				nd = utils.Dir(i).TurnAround()
				break
			}
		}
		np := map[[2]utils.Dir]utils.Pos{
			{utils.Up, utils.Up}:       {p.X, 49},
			{utils.Up, utils.Right}:    {0, p.X},
			{utils.Right, utils.Up}:    {p.Y, 49},
			{utils.Right, utils.Right}: {0, p.Y},
			{utils.Right, utils.Left}:  {49, 49 - p.Y},
			{utils.Down, utils.Down}:   {p.X, 0},
			{utils.Down, utils.Left}:   {49, p.X},
			{utils.Left, utils.Right}:  {0, 49 - p.Y},
			{utils.Left, utils.Down}:   {p.Y, 0},
			{utils.Left, utils.Left}:   {49, p.Y},
		}[[2]utils.Dir{d, nd}]
		return np, nd, nf
	}
	utils.Println(score())
}
