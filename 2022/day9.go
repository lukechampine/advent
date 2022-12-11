package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 9)

func moveToward(src, dst utils.Pos) utils.Pos {
	if utils.Abs(dst.X-src.X) < 2 && utils.Abs(dst.Y-src.Y) < 2 {
		return src
	}
	return utils.Pos{
		X: src.X + utils.Signum(dst.X-src.X),
		Y: src.Y + utils.Signum(dst.Y-src.Y),
	}
}

func main() {
	h := utils.Pos{0, 0}
	rope := make([]utils.Pos, 9)
	v1 := make(map[utils.Pos]bool)
	v1[rope[0]] = true
	v2 := make(map[utils.Pos]bool)
	v2[rope[8]] = true
	for _, line := range utils.Lines(input) {
		d := utils.DirFromUDLR(line[0])
		n := utils.Atoi(line[2:])
		h = h.Tread(d, n, func(p utils.Pos) {
			for i := range rope {
				rope[i] = moveToward(rope[i], p)
				p = rope[i]
			}
			v1[rope[0]] = true
			v2[rope[8]] = true
		})
	}
	utils.Println(len(v1))
	utils.Println(len(v2))
}
