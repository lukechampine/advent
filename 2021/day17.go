package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 17)

func fire(xv, yv, x1, x2, y1, y2 int) (maxY int, within bool) {
	xp, yp := 0, 0
	for (xv >= 0 || xp >= x1) && (yv >= 0 || yp >= y1) {
		if x1 <= xp && xp <= x2 && y1 <= yp && yp <= y2 {
			return maxY, true
		}
		xp += xv
		yp += yv
		xv -= utils.Signum(xv) // drag
		yv--                   // gravity
		maxY = utils.Max(maxY, yp)
	}
	return 0, false
}

func main() {
	parts := utils.ExtractInts(input)
	x1, x2, y1, y2 := parts[0], parts[1], parts[2], parts[3]

	var max int
	var n int
	for xv := 0; xv <= x2; xv++ {
		for yv := -200; yv < 200; yv++ {
			maxY, ok := fire(xv, yv, x1, x2, y1, y2)
			max = utils.Max(max, maxY)
			n += utils.BoolToInt(ok)
		}
	}
	utils.Println(max)
	utils.Println(n)
}
