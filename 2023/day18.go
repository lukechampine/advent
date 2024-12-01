package main

import (
	"strconv"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 18)

func bsp(border []utils.Pos) []utils.Rect {
	var borderRects []utils.Rect
	for i := range border {
		b := utils.Rect{Min: border[i], Max: border[(i+1)%len(border)]}
		if b.Min.X > b.Max.X || b.Min.Y > b.Max.Y {
			b.Min, b.Max = b.Max, b.Min
		}
		borderRects = append(borderRects, b)
	}

	var rects []utils.Rect
	var visit func(r utils.Rect)
	visit = func(r utils.Rect) {
		if r.Width() < 0 || r.Height() < 0 {
			return
		}
		isBorder := utils.Any(len(borderRects), func(i int) bool { return borderRects[i].ContainsRect(r) })
		if isBorder {
			return // skip
		}
		var overlaps []utils.Rect
		for _, b := range borderRects {
			if b.Overlaps(r) {
				overlaps = append(overlaps, b)
			}
		}
		if len(overlaps) == 0 {
			rects = append(rects, r)
			return
		}
		for _, b := range overlaps {
			if b.Width() == 0 && r.Width() != 0 {
				// split vertically
				left, middle, right := r, r, r
				left.Max.X = b.Min.X - 1
				middle.Min.X = b.Min.X
				middle.Max.X = b.Max.X
				right.Min.X = b.Min.X + 1
				visit(left)
				visit(middle)
				visit(right)
				return
			} else if b.Height() == 0 && r.Height() != 0 {
				// split horizontally
				left, middle, right := r, r, r
				left.Max.Y = b.Min.Y - 1
				middle.Min.Y = b.Min.Y
				middle.Max.Y = b.Max.Y
				right.Min.Y = b.Min.Y + 1
				visit(left)
				visit(middle)
				visit(right)
				return
			}
		}
		panic("unreachable")
	}

	visit(utils.BoundingRect(border))
	return rects
}

func flood(rects []utils.Rect) int {
	var points []utils.Pos
	for _, r := range rects {
		points = append(points, r.Corners()...)
	}
	r := utils.BoundingRect(points)

	// add outer border
	grow := func(r utils.Rect) utils.Rect {
		return utils.Rect{Min: r.Min.Add(-1, -1), Max: r.Max.Add(1, 1)}
	}
	r = grow(r)
	c := r.Corners()
	rects = append(rects,
		utils.Rect{Min: c[0], Max: c[1].Add(-1, 0)},
		utils.Rect{Min: c[1], Max: c[2].Add(0, -1)},
		utils.Rect{Min: c[3], Max: c[2].Add(-1, 0)},
		utils.Rect{Min: c[0], Max: c[3].Add(0, -1)})

	var outside int
	seen := make(map[utils.Rect]bool)
	queue := []utils.Rect{rects[len(rects)-1]}
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if seen[r] {
			continue
		}

		seen[r] = true
		outside += r.Area()
		for _, n := range rects {
			if grow(r).Overlaps(n) {
				queue = append(queue, n)
			}
		}
	}
	return r.Area() - outside
}

func main() {
	p := utils.Origin
	var border []utils.Pos
	for _, line := range utils.Lines(input) {
		dir := utils.DirFromUDLR(line[0])
		steps := utils.Atoi(strings.Fields(line)[1])
		p = p.MoveArray(dir, steps)
		border = append(border, p)
	}
	utils.Println(flood(bsp(border)))

	p = utils.Origin
	border = nil
	for _, line := range utils.Lines(input) {
		code := strings.Trim(strings.Fields(line)[2], "(#)")
		steps, _ := strconv.ParseInt(code[:len(code)-1], 16, 32)
		dir := utils.Dir(code[len(code)-1] - '0').TurnRight()
		p = p.MoveArray(dir, int(steps))
		border = append(border, p)
	}
	utils.Println(flood(bsp(border)))
}
