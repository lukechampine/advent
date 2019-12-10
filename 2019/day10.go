package main

import (
	"math"
	"sort"

	"github.com/lukechampine/advent/utils"
)

var input = utils.ReadInput("day10_input.txt")
var ils = utils.Lines(input)

func main() {
	// part 1
	isAsteroid := func(p utils.Pos) bool {
		return ils[p.Y][p.X] == '#'
	}
	var asteroids []utils.Pos
	for y := range ils {
		for x := range ils[y] {
			if p := (utils.Pos{X: x, Y: y}); isAsteroid(p) {
				asteroids = append(asteroids, p)
			}
		}
	}

	numDetectableFrom := func(a utils.Pos) int {
		// for each other asteroid, calculate whether we can "see" it
		return utils.Count(len(asteroids), func(j int) bool {
			b := asteroids[j]
			if a == b {
				return false // asteroids don't "see" themselves
			}
			for p := a.StrideTowards(b); p != b; p = p.StrideTowards(b) {
				if isAsteroid(p) {
					return false
				}
			}
			return true
		})
	}

	stationIndex := utils.MaximumIndex(len(asteroids), func(i int) int {
		return numDetectableFrom(asteroids[i])
	})
	station := asteroids[stationIndex]
	utils.Println(numDetectableFrom(station))

	// part 2

	// remove station from set of asteroids
	asteroids = append(asteroids[:stationIndex], asteroids[stationIndex+1:]...)

	// map asteroid -> number of asteroids between it and station
	precedence := make(map[utils.Pos]int)
	for _, a := range asteroids {
		n := 0
		for p := a.StrideTowards(station); p != station; p = p.StrideTowards(station) {
			if isAsteroid(p) {
				n++
			}
		}
		precedence[a] = n
	}

	radiansUntil := func(a utils.Pos) (phi float64) {
		_, phi = a.Rel(station).Polar()
		// rotate polar axis to point straight "up"
		if phi < -math.Pi/2 {
			phi += 2 * math.Pi
		}
		// add another full rotation for each asteroid "in the way"
		phi += 2 * math.Pi * float64(precedence[a])
		return
	}

	sort.Slice(asteroids, func(i, j int) bool {
		return radiansUntil(asteroids[i]) < radiansUntil(asteroids[j])
	})
	a := asteroids[199]
	utils.Println(a.X*100 + a.Y)
}
