package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `position=<-31065, -31102> velocity=< 3,  3>
position=<-31036, -31096> velocity=< 3,  3>
position=<-31057,  41827> velocity=< 3, -4>
position=<-20631,  41822> velocity=< 2, -4>
position=<-41478, -31096> velocity=< 4,  3>
position=< 21061,  52246> velocity=<-2, -5>
position=<-31060, -20681> velocity=< 3,  2>
position=<-41455, -31096> velocity=< 4,  3>
position=<-10229, -31100> velocity=< 1,  3>
position=<-20594, -51939> velocity=< 2,  5>
position=< 10655,  20991> velocity=<-1, -2>
position=<-51856, -51941> velocity=< 5,  5>
position=<-20631, -51936> velocity=< 2,  5>
position=<-10181, -31096> velocity=< 1,  3>
position=< 21018,  52239> velocity=<-2, -5>
position=< 10623,  20991> velocity=<-1, -2>
position=< 21017,  31410> velocity=<-2, -3>
position=<-20594,  10568> velocity=< 2, -1>
position=< 10656, -51937> velocity=<-1,  5>
position=< 21052,  10576> velocity=<-2, -1>
position=< 31484,  10570> velocity=<-3, -1>
position=<-20627,  41825> velocity=< 2, -4>
position=< 21029,  52247> velocity=<-2, -5>
position=<-41459,  31411> velocity=< 4, -3>
position=< 41866, -20685> velocity=<-4,  2>
position=< 41869, -10265> velocity=<-4,  1>
position=< 52320,  20991> velocity=<-5, -2>
position=<-31061, -41516> velocity=< 3,  4>
position=< 31495, -10269> velocity=<-3,  1>
position=< 41869, -10267> velocity=<-4,  1>
position=< 52305, -41518> velocity=<-5,  4>
position=<-20622,  20994> velocity=< 2, -2>
position=< 10623, -41519> velocity=<-1,  4>
position=< 21073, -41515> velocity=<-2,  4>
position=<-51872, -51941> velocity=< 5,  5>
position=< 52284, -20679> velocity=<-5,  2>
position=<-51906,  20985> velocity=< 5, -2>
position=< 21038, -31098> velocity=<-2,  3>
position=< 52321, -10264> velocity=<-5,  1>
position=< 10623,  10576> velocity=<-1, -1>
position=<-20611,  10569> velocity=< 2, -1>
position=< 41912, -10265> velocity=<-4,  1>
position=< 41866, -41517> velocity=<-4,  4>
position=<-10221, -51933> velocity=< 1,  5>
position=<-41472, -41517> velocity=< 4,  4>
position=< 21078, -41520> velocity=<-2,  4>
position=< 41863, -20684> velocity=<-4,  2>
position=<-20639, -51937> velocity=< 2,  5>
position=< 21037, -51933> velocity=<-2,  5>
position=< 31464, -10260> velocity=<-3,  1>
position=< 41874,  41830> velocity=<-4, -4>
position=<-41486, -20678> velocity=< 4,  2>
position=< 31479, -31098> velocity=<-3,  3>
position=<-10236,  52245> velocity=< 1, -5>
position=< 31456,  52246> velocity=<-3, -5>
position=<-51873,  20985> velocity=< 5, -2>
position=<-51896, -10268> velocity=< 5,  1>
position=< 41890,  41821> velocity=<-4, -4>
position=<-10201,  41824> velocity=< 1, -4>
position=< 52327,  10576> velocity=<-5, -1>
position=< 52279, -10262> velocity=<-5,  1>
position=<-10193,  20992> velocity=< 1, -2>
position=< 31495,  41825> velocity=<-3, -4>
position=< 10651, -10267> velocity=<-1,  1>
position=<-20639,  10574> velocity=< 2, -1>
position=<-41472, -31100> velocity=< 4,  3>
position=< 31443,  31409> velocity=<-3, -3>
position=<-20618, -31104> velocity=< 2,  3>
position=< 41874, -31104> velocity=<-4,  3>
position=< 31472,  31405> velocity=<-3, -3>
position=<-20644, -10263> velocity=< 2,  1>
position=< 41862, -31103> velocity=<-4,  3>
position=<-41463, -20680> velocity=< 4,  2>
position=<-51856,  52247> velocity=< 5, -5>
position=<-10217,  41829> velocity=< 1, -4>
position=< 52308, -10260> velocity=<-5,  1>
position=<-51885, -41514> velocity=< 5,  4>
position=< 41874, -31104> velocity=<-4,  3>
position=<-41435,  10574> velocity=< 4, -1>
position=< 52271, -10261> velocity=<-5,  1>
position=< 10623, -41515> velocity=<-1,  4>
position=< 31488,  31412> velocity=<-3, -3>
position=<-10235, -20682> velocity=< 1,  2>
position=< 52279,  10574> velocity=<-5, -1>
position=< 52298, -20683> velocity=<-5,  2>
position=<-20599, -31102> velocity=< 2,  3>
position=< 10615, -31096> velocity=<-1,  3>
position=< 52330,  31403> velocity=<-5, -3>
position=<-51899, -31102> velocity=< 5,  3>
position=< 31476, -31096> velocity=<-3,  3>
position=< 10612,  31406> velocity=<-1, -3>
position=< 52295, -10269> velocity=<-5,  1>
position=< 31443, -31101> velocity=<-3,  3>
position=<-41490,  41830> velocity=< 4, -4>
position=< 21041,  52245> velocity=<-2, -5>
position=< 21033,  31403> velocity=<-2, -3>
position=<-51891, -41520> velocity=< 5,  4>
position=< 41906, -41522> velocity=<-4,  4>
position=< 21061,  31405> velocity=<-2, -3>
position=< 52315,  52243> velocity=<-5, -5>
position=<-41447,  31406> velocity=< 4, -3>
position=< 31487,  41828> velocity=<-3, -4>
position=< 52327, -31105> velocity=<-5,  3>
position=<-31041, -31098> velocity=< 3,  3>
position=<-10210, -41518> velocity=< 1,  4>
position=<-51861,  20986> velocity=< 5, -2>
position=< 31478, -51932> velocity=<-3,  5>
position=<-41448,  31412> velocity=< 4, -3>
position=<-41490,  52245> velocity=< 4, -5>
position=< 41897,  10572> velocity=<-4, -1>
position=< 31454,  41826> velocity=<-3, -4>
position=< 10659,  52239> velocity=<-1, -5>
position=< 21061, -20686> velocity=<-2,  2>
position=< 52282,  20991> velocity=<-5, -2>
position=<-20599, -41522> velocity=< 2,  4>
position=< 21050, -51941> velocity=<-2,  5>
position=< 21045,  31409> velocity=<-2, -3>
position=< 41857, -20684> velocity=<-4,  2>
position=<-41475, -51936> velocity=< 4,  5>
position=< 10620, -20679> velocity=<-1,  2>
position=< 52281, -51937> velocity=<-5,  5>
position=<-31024, -51938> velocity=< 3,  5>
position=<-20599,  52247> velocity=< 2, -5>
position=< 31445, -31101> velocity=<-3,  3>
position=<-41467, -31103> velocity=< 4,  3>
position=< 31446, -41517> velocity=<-3,  4>
position=< 52296, -10265> velocity=<-5,  1>
position=< 10620,  20992> velocity=<-1, -2>
position=<-20639,  41829> velocity=< 2, -4>
position=< 21037, -31098> velocity=<-2,  3>
position=<-51898,  31408> velocity=< 5, -3>
position=<-20611, -51941> velocity=< 2,  5>
position=< 41858,  10567> velocity=<-4, -1>
position=< 41888,  31407> velocity=<-4, -3>
position=< 10624,  20985> velocity=<-1, -2>
position=<-51909,  20992> velocity=< 5, -2>
position=< 52304, -31105> velocity=<-5,  3>
position=<-41459, -41516> velocity=< 4,  4>
position=<-31036, -41522> velocity=< 3,  4>
position=< 41901,  52247> velocity=<-4, -5>
position=<-10181, -41521> velocity=< 1,  4>
position=<-31072, -10260> velocity=< 3,  1>
position=< 52331,  52243> velocity=<-5, -5>
position=< 31491, -20686> velocity=<-3,  2>
position=<-31057,  52247> velocity=< 3, -5>
position=<-51881, -10263> velocity=< 5,  1>
position=< 41877,  20994> velocity=<-4, -2>
position=<-20621, -10264> velocity=< 2,  1>
position=<-41458,  31403> velocity=< 4, -3>
position=< 10631, -51933> velocity=<-1,  5>
position=< 10607, -20683> velocity=<-1,  2>
position=<-51856, -10269> velocity=< 5,  1>
position=< 41909,  20993> velocity=<-4, -2>
position=< 31478, -41514> velocity=<-3,  4>
position=< 41877, -41515> velocity=<-4,  4>
position=<-41433, -10265> velocity=< 4,  1>
position=< 10623,  52239> velocity=<-1, -5>
position=<-20611,  20993> velocity=< 2, -2>
position=< 41897, -41523> velocity=<-4,  4>
position=< 10634,  31407> velocity=<-1, -3>
position=< 10650, -20682> velocity=<-1,  2>
position=<-41475,  20985> velocity=< 4, -2>
position=<-31036, -51940> velocity=< 3,  5>
position=<-41451, -41516> velocity=< 4,  4>
position=< 52295,  10574> velocity=<-5, -1>
position=<-20606,  41827> velocity=< 2, -4>
position=< 10633, -10269> velocity=<-1,  1>
position=< 10644, -41523> velocity=<-1,  4>
position=< 52331,  31407> velocity=<-5, -3>
position=< 10615,  10568> velocity=<-1, -1>
position=< 31460,  41821> velocity=<-3, -4>
position=< 52315,  52245> velocity=<-5, -5>
position=< 10599, -20678> velocity=<-1,  2>
position=<-41475,  31408> velocity=< 4, -3>
position=< 10599, -41516> velocity=<-1,  4>
position=< 21021, -31102> velocity=<-2,  3>
position=<-41455,  41821> velocity=< 4, -4>
position=< 21066,  41823> velocity=<-2, -4>
position=< 31485,  41825> velocity=<-3, -4>
position=<-31052,  20993> velocity=< 3, -2>
position=<-10188,  52246> velocity=< 1, -5>
position=<-20639, -10262> velocity=< 2,  1>
position=<-10180, -10265> velocity=< 1,  1>
position=<-10202,  20985> velocity=< 1, -2>
position=< 21066,  41823> velocity=<-2, -4>
position=< 31463, -31099> velocity=<-3,  3>
position=<-51896,  20986> velocity=< 5, -2>
position=<-20598, -31105> velocity=< 2,  3>
position=< 31491, -20685> velocity=<-3,  2>
position=< 41877, -51938> velocity=<-4,  5>
position=< 21018,  52239> velocity=<-2, -5>
position=<-10216, -31099> velocity=< 1,  3>
position=<-10193, -10264> velocity=< 1,  1>
position=<-31060, -31098> velocity=< 3,  3>
position=< 52292,  10567> velocity=<-5, -1>
position=< 10647, -20687> velocity=<-1,  2>
position=< 52295,  41822> velocity=<-5, -4>
position=<-31068, -10268> velocity=< 3,  1>
position=< 52304,  52248> velocity=<-5, -5>
position=<-10220,  10568> velocity=< 1, -1>
position=< 41897,  41824> velocity=<-4, -4>
position=<-51885, -51938> velocity=< 5,  5>
position=< 10651, -20681> velocity=<-1,  2>
position=< 31436, -31105> velocity=<-3,  3>
position=< 31487,  31409> velocity=<-3, -3>
position=<-51888,  41830> velocity=< 5, -4>
position=<-20647,  31412> velocity=< 2, -3>
position=<-31072,  20991> velocity=< 3, -2>
position=< 52319, -51941> velocity=<-5,  5>
position=<-20607,  41821> velocity=< 2, -4>
position=< 21020,  52248> velocity=<-2, -5>
position=<-20599,  41825> velocity=< 2, -4>
position=<-51853,  31408> velocity=< 5, -3>
position=<-20646,  52240> velocity=< 2, -5>
position=<-10189,  41829> velocity=< 1, -4>
position=< 10604,  31412> velocity=<-1, -3>
position=<-41488,  10571> velocity=< 4, -1>
position=< 41914, -31103> velocity=<-4,  3>
position=< 52291, -41516> velocity=<-5,  4>
position=< 31462,  31407> velocity=<-3, -3>
position=<-41470, -31103> velocity=< 4,  3>
position=< 21059, -10260> velocity=<-2,  1>
position=<-31060,  10575> velocity=< 3, -1>
position=<-41472,  20990> velocity=< 4, -2>
position=< 41903, -41518> velocity=<-4,  4>
position=<-51880, -10268> velocity=< 5,  1>
position=< 10623,  41822> velocity=<-1, -4>
position=<-41435,  31408> velocity=< 4, -3>
position=< 21070, -31104> velocity=<-2,  3>
position=<-31069, -20678> velocity=< 3,  2>
position=< 21043,  10567> velocity=<-2, -1>
position=< 21035, -20683> velocity=<-2,  2>
position=<-51877,  41829> velocity=< 5, -4>
position=< 10648, -10266> velocity=<-1,  1>
position=< 31479,  52240> velocity=<-3, -5>
position=< 21049,  52246> velocity=<-2, -5>
position=<-31041,  52248> velocity=< 3, -5>
position=< 52306,  20994> velocity=<-5, -2>
position=< 52307,  10570> velocity=<-5, -1>
position=< 31443, -51940> velocity=<-3,  5>
position=< 41874,  10572> velocity=<-4, -1>
position=< 41902, -10262> velocity=<-4,  1>
position=< 41858,  31403> velocity=<-4, -3>
position=< 31456, -31104> velocity=<-3,  3>
position=< 41893,  20993> velocity=<-4, -2>
position=< 52274, -20678> velocity=<-5,  2>
position=<-41435, -51939> velocity=< 4,  5>
position=<-51909,  41830> velocity=< 5, -4>
position=<-10227, -20683> velocity=< 1,  2>
position=<-31049,  20990> velocity=< 3, -2>
position=< 21034, -10267> velocity=<-2,  1>
position=< 31437,  31403> velocity=<-3, -3>
position=<-31065, -10262> velocity=< 3,  1>
position=<-10217, -20679> velocity=< 1,  2>
position=< 21075, -41523> velocity=<-2,  4>
position=< 52279, -10266> velocity=<-5,  1>
position=<-51861,  20994> velocity=< 5, -2>
position=< 31479,  52239> velocity=<-3, -5>
position=< 10633, -41514> velocity=<-1,  4>
position=<-51876,  52245> velocity=< 5, -5>
position=<-20647,  31403> velocity=< 2, -3>
position=<-51896, -51935> velocity=< 5,  5>
position=<-31060,  20989> velocity=< 3, -2>
position=<-41483,  52239> velocity=< 4, -5>
position=<-31029,  41826> velocity=< 3, -4>
position=<-41462,  10569> velocity=< 4, -1>
position=<-20607, -10268> velocity=< 2,  1>
position=< 21027,  20988> velocity=<-2, -2>
position=<-41438,  10567> velocity=< 4, -1>
position=< 21025,  31409> velocity=<-2, -3>
position=<-41483, -20682> velocity=< 4,  2>
position=< 21057, -41515> velocity=<-2,  4>
position=< 52292,  20985> velocity=<-5, -2>
position=< 31460,  10571> velocity=<-3, -1>
position=<-10200, -51941> velocity=< 1,  5>
position=<-41486,  10568> velocity=< 4, -1>
position=<-51851,  10567> velocity=< 5, -1>
position=< 21026,  31404> velocity=<-2, -3>
position=<-20634,  20988> velocity=< 2, -2>
position=< 10628,  10575> velocity=<-1, -1>
position=<-20612, -31105> velocity=< 2,  3>
position=<-31070,  20989> velocity=< 3, -2>
position=< 52324,  10568> velocity=<-5, -1>
position=< 31437, -20687> velocity=<-3,  2>
position=< 52313,  52248> velocity=<-5, -5>
position=<-20629,  20985> velocity=< 2, -2>
position=<-20604, -20683> velocity=< 2,  2>
position=< 21022, -51939> velocity=<-2,  5>
position=< 10655, -31096> velocity=<-1,  3>
position=<-20655, -20687> velocity=< 2,  2>
position=< 52327, -41518> velocity=<-5,  4>
position=< 21018,  41830> velocity=<-2, -4>
position=< 10641,  41830> velocity=<-1, -4>
position=<-10224,  52241> velocity=< 1, -5>
position=<-31029,  10569> velocity=< 3, -1>
position=< 21028,  52244> velocity=<-2, -5>
position=< 10602,  52243> velocity=<-1, -5>
position=<-20642,  10572> velocity=< 2, -1>
position=< 52295, -20683> velocity=<-5,  2>
position=< 31437, -41514> velocity=<-3,  4>
position=< 21044, -51941> velocity=<-2,  5>
position=<-31041, -31105> velocity=< 3,  3>
position=< 31435,  41829> velocity=<-3, -4>
position=<-51880,  31406> velocity=< 5, -3>
position=<-41439, -31103> velocity=< 4,  3>
position=<-51888, -41521> velocity=< 5,  4>
position=< 52306, -20678> velocity=<-5,  2>
position=<-41479,  20993> velocity=< 4, -2>
position=< 10612,  20993> velocity=<-1, -2>
position=< 31448, -10269> velocity=<-3,  1>
position=<-31064,  10569> velocity=< 3, -1>
position=<-10221, -31102> velocity=< 1,  3>
position=<-41439, -20684> velocity=< 4,  2>
position=< 52287,  52245> velocity=<-5, -5>
position=< 41877,  10569> velocity=<-4, -1>
position=<-10193, -10265> velocity=< 1,  1>
position=<-20621, -51932> velocity=< 2,  5>
position=<-10216, -20683> velocity=< 1,  2>
position=< 41878, -20687> velocity=<-4,  2>
position=< 52299,  52239> velocity=<-5, -5>
position=<-41475, -10268> velocity=< 4,  1>
position=<-41483, -10261> velocity=< 4,  1>
position=< 41897,  10570> velocity=<-4, -1>
position=<-10201, -31096> velocity=< 1,  3>
position=<-31052, -10265> velocity=< 3,  1>
position=<-10232,  52248> velocity=< 1, -5>
position=< 21065,  41829> velocity=<-2, -4>
position=< 31471,  31412> velocity=<-3, -3>
position=<-51889,  31410> velocity=< 5, -3>
position=<-31021, -31103> velocity=< 3,  3>
position=<-41483,  20993> velocity=< 4, -2>
position=<-20651,  10576> velocity=< 2, -1>
position=< 10643, -10265> velocity=<-1,  1>
position=<-41486,  20987> velocity=< 4, -2>
position=< 31435,  20985> velocity=<-3, -2>
position=<-31054, -41517> velocity=< 3,  4>
position=<-51907,  20994> velocity=< 5, -2>
position=<-51883,  52243> velocity=< 5, -5>
position=< 31478, -20687> velocity=<-3,  2>
position=< 10616, -10267> velocity=<-1,  1>
position=< 31467, -51941> velocity=<-3,  5>
position=< 31475,  52247> velocity=<-3, -5>
position=<-20642, -51936> velocity=< 2,  5>
position=<-41462,  31404> velocity=< 4, -3>
position=< 41902, -41516> velocity=<-4,  4>
position=< 31435,  41821> velocity=<-3, -4>
position=<-10229,  31405> velocity=< 1, -3>
position=<-41475, -10263> velocity=< 4,  1>
position=<-41475,  41826> velocity=< 4, -4>
position=<-41487,  10567> velocity=< 4, -1>
position=< 31448, -51937> velocity=<-3,  5>
position=<-20622,  20991> velocity=< 2, -2>
position=< 10632, -51932> velocity=<-1,  5>
position=<-51885,  20990> velocity=< 5, -2>
position=<-31057, -20684> velocity=< 3,  2>
position=<-20655, -20679> velocity=< 2,  2>`

type vector struct {
	x, y int
}

type point struct {
	pos, vel vector
}

func (p *point) update() {
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
}

func (p *point) rewind() {
	p.pos.x -= p.vel.x
	p.pos.y -= p.vel.y
}

func boundingBox(points []point) (x1, y1, x2, y2 int) {
	minX := points[0].pos.x
	minY := points[0].pos.y
	maxX, maxY := minX, minY
	for _, p := range points[1:] {
		if p.pos.x < minX {
			minX = p.pos.x
		}
		if p.pos.x > maxX {
			maxX = p.pos.x
		}
		if p.pos.y < minY {
			minY = p.pos.y
		}
		if p.pos.y > maxY {
			maxY = p.pos.y
		}
	}
	return minX, minY, maxX, maxY
}

func boxSize(points []point) int {
	x1, y1, x2, y2 := boundingBox(points)
	return (x2 - x1) * (y2 - y1)
}

func printPoints(points []point) {
	x1, y1, x2, y2 := boundingBox(points)
	grid := utils.NewLightBoard((x2-x1)+1, (y2-y1)+1)
	for _, p := range points {
		grid.Set(p.pos.x-x1, p.pos.y-y1, true)
	}
	grid.Print()
}

func main() {
	// part 1
	var points []point
	for _, line := range utils.Lines(input) {
		var p point
		utils.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &p.pos.x, &p.pos.y, &p.vel.x, &p.vel.y)
		points = append(points, p)
	}

	var secs int
	size := boxSize(points)
	for ; ; secs++ {
		for i := range points {
			points[i].update()
		}
		newSize := boxSize(points)
		if newSize > size {
			for i := range points {
				points[i].rewind()
			}
			break
		}
		size = newSize
	}
	printPoints(points)

	// part 2
	utils.Println(secs)
}
