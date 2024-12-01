package main

import (
	"sort"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 22)

type Pos3 struct {
	X, Y, Z int
}

type Brick struct {
	bot, top Pos3
}

func (b Brick) shadow() utils.Rect {
	return utils.Rect{
		Min: utils.Pos{X: b.bot.X, Y: b.bot.Y},
		Max: utils.Pos{X: b.top.X, Y: b.top.Y},
	}
}

func drop(tops map[int][]utils.Rect, b Brick) Brick {
	for b.bot.Z > 0 {
		if utils.Any(len(tops[b.bot.Z-1]), func(i int) bool {
			return b.shadow().Overlaps(tops[b.bot.Z-1][i])
		}) {
			break
		}
		b.top.Z--
		b.bot.Z--
	}
	tops[b.top.Z] = append(tops[b.top.Z], b.shadow())
	return b
}

func dropAll(bricks []Brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].top.Z < bricks[j].top.Z
	})
	tops := make(map[int][]utils.Rect)
	for i := range bricks {
		bricks[i] = drop(tops, bricks[i])
	}
}

func remove(bricks []Brick, i int) []Brick {
	bricks = append([]Brick(nil), bricks...)
	bricks[0], bricks[i] = bricks[i], bricks[0]
	bricks = bricks[1:]
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].top.Z < bricks[j].top.Z
	})
	return bricks
}

func numFalls(bricks []Brick) int {
	tops := make(map[int][]utils.Rect)
	return utils.Count(len(bricks), func(i int) bool {
		return drop(tops, bricks[i]) != bricks[i]
	})
}

func main() {
	var bricks []Brick
	for _, line := range utils.Lines(input) {
		ints := utils.ExtractInts(line)
		bricks = append(bricks, Brick{Pos3{ints[0], ints[1], ints[2]}, Pos3{ints[3], ints[4], ints[5]}})
	}

	dropAll(bricks)
	var canDisintegrate int
	var totalFalls int
	for i := range bricks {
		falls := numFalls(remove(bricks, i))
		canDisintegrate += utils.BoolToInt(falls == 0)
		totalFalls += falls
	}
	utils.Println(canDisintegrate)
	utils.Println(totalFalls)
}
