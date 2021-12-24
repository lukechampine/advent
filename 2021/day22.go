package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 22)

type Range struct {
	Start, End int
}

func (r Range) overlaps(o Range) bool {
	return (r.Start <= o.Start && o.Start <= r.End) ||
		(r.Start <= o.End && o.End <= r.End) ||
		(o.Start <= r.Start && r.Start <= o.End) ||
		(o.Start <= r.End && r.End <= o.End)
}

type Cmd struct {
	on         bool
	rX, rY, rZ Range
}

func (c Cmd) size() int {
	mag := (c.rX.End - c.rX.Start + 1) *
		(c.rY.End - c.rY.Start + 1) *
		(c.rZ.End - c.rZ.Start + 1)
	if !c.on {
		mag *= -1
	}
	return mag
}

func overlap(c, d Cmd) (Cmd, bool) {
	if c.rX.overlaps(d.rX) && c.rY.overlaps(d.rY) && c.rZ.overlaps(d.rZ) {
		return Cmd{
			on: !d.on,
			rX: Range{utils.Max(c.rX.Start, d.rX.Start), utils.Min(c.rX.End, d.rX.End)},
			rY: Range{utils.Max(c.rY.Start, d.rY.Start), utils.Min(c.rY.End, d.rY.End)},
			rZ: Range{utils.Max(c.rZ.Start, d.rZ.Start), utils.Min(c.rZ.End, d.rZ.End)},
		}, true
	}
	return Cmd{}, false
}

func main() {
	var cmds []Cmd
	for _, line := range utils.Lines(input) {
		cur := Cmd{on: strings.Contains(line, "on")}
		utils.Sscanf(line[4-utils.BoolToInt(cur.on):], "x=%d..%d,y=%d..%d,z=%d..%d",
			&cur.rX.Start, &cur.rX.End, &cur.rY.Start, &cur.rY.End, &cur.rZ.Start, &cur.rZ.End)
		for _, prev := range cmds {
			if o, ok := overlap(cur, prev); ok {
				cmds = append(cmds, o)
			}
		}
		if cur.on {
			cmds = append(cmds, cur)
		}
	}

	utils.Println(utils.Sum(len(cmds), func(i int) int {
		c := cmds[i]
		inPart1 := c.rX.Start <= 50 && c.rX.Start >= -50 &&
			c.rY.Start <= 50 && c.rY.Start >= -50 &&
			c.rZ.Start <= 50 && c.rZ.Start >= -50
		return utils.BoolToInt(inPart1) * c.size()
	}))

	utils.Println(utils.Sum(len(cmds), func(i int) int { return cmds[i].size() }))
}
