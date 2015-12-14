package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.`

type reindeer struct {
	name             string
	speed, dur, rest int
}

func (r reindeer) dist(secs int) int {
	cycles := secs / (r.dur + r.rest)
	km := cycles * (r.speed * r.dur)
	// remainder
	for i := 0; i < secs%(r.dur+r.rest); i++ {
		if i < r.dur {
			km += r.speed
		}
	}
	return km
}

func parse(str string) reindeer {
	var name string
	var speed, dur, rest int
	utils.Sscanf(str, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &dur, &rest)
	return reindeer{name, speed, dur, rest}
}

func main() {
	// part 1
	var rs []reindeer
	for _, str := range utils.Lines(input) {
		rs = append(rs, parse(str))
	}
	var farthest int
	for _, r := range rs {
		farthest = utils.Max(farthest, r.dist(2503))
	}
	println(farthest)

	// part 2
	points := make(map[string]int)
	for i := 0; i < 2503; i++ {
		var farthest int
		for _, r := range rs {
			farthest = utils.Max(farthest, r.dist(i+1))
		}
		for _, r := range rs {
			if r.dist(i+1) == farthest {
				points[r.name]++
			}
		}
	}
	var mostPoints int
	for _, p := range points {
		mostPoints = utils.Max(mostPoints, p)
	}
	println(mostPoints)
}
