package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2025, 9)

func main() {
	var tiles []utils.Pos
	for _, line := range utils.Lines(input) {
		var t utils.Pos
		utils.Sscanf(line, "%d,%d", &t.X, &t.Y)
		tiles = append(tiles, t)
	}
	var rects []utils.Rect
	for i, t1 := range tiles {
		for _, t2 := range tiles[i+1:] {
			rects = append(rects, utils.Rect{Min: t1, Max: t2})
		}
	}

	utils.Println(utils.Maximum(len(rects), func(i int) int {
		return rects[i].Area()
	}))

	mids := make([]utils.Pos, len(tiles))
	for i := range tiles {
		t1, t2 := tiles[i], tiles[(i+1)%len(tiles)]
		mids[i] = utils.Pos{X: (t1.X + t2.X) / 2, Y: (t1.Y + t2.Y) / 2}
	}
	// hack, but good enough
	valid := func(r utils.Rect) bool {
		normed := utils.Rect{
			Min: utils.Pos{
				X: utils.Min(r.Min.X, r.Max.X) + 1,
				Y: utils.Min(r.Min.Y, r.Max.Y) + 1,
			},
			Max: utils.Pos{
				X: utils.Max(r.Min.X, r.Max.X) - 1,
				Y: utils.Max(r.Min.Y, r.Max.Y) - 1,
			},
		}
		return utils.All(len(tiles), func(i int) bool {
			t1, t2 := tiles[i], tiles[(i+1)%len(tiles)]
			mid := utils.Pos{X: (t1.X + t2.X) / 2, Y: (t1.Y + t2.Y) / 2}
			return !normed.Contains(t1) &&
				!normed.Contains(mid) &&
				!normed.Contains(t2)
		})
	}

	utils.Println(utils.Maximum(len(rects), func(i int) int {
		return rects[i].Area() * utils.BoolToInt(valid(rects[i]))
	}))
}
