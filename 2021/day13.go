package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 13)

func fold(paper *utils.LightBoard, isY bool, crease int) {
	paper.ForEach(func(p utils.Pos, b bool) {
		paper.Set(p.X, p.Y, false ||
			(isY && p.Y < crease && (b || paper.Get(p.X, crease+(crease-p.Y)))) ||
			(!isY && p.X < crease && (b || paper.Get(crease+(crease-p.X), p.Y))))
	})
}

func main() {
	parts := utils.Split(input, "\n\n")
	var coords []utils.Pos
	for _, line := range utils.Lines(parts[0]) {
		is := utils.ExtractInts(line)
		coords = append(coords, utils.Pos{is[0], is[1]})
	}
	_, max := utils.BoundingBox(coords)
	paper := utils.NewLightBoard(max.X+10, max.Y+10)
	for _, p := range coords {
		paper.Set(p.X, p.Y, true)
	}

	for i, line := range utils.Lines(parts[1]) {
		fold(paper, strings.Contains(line, "y"), utils.ExtractInts(line)[0])
		if i == 0 {
			utils.Println(paper.Count(true))
		}
	}
	var rem []utils.Pos
	paper.ForEach(func(p utils.Pos, b bool) {
		if b {
			rem = append(rem, p)
		}
	})
	_, max = utils.BoundingBox(rem)
	paper.Board = paper.Board[:max.Y+1]
	for y := range paper.Board {
		paper.Board[y] = paper.Board[y][:max.X+1]
	}
	paper.Print()
}
