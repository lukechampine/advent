package main

const input = 1362

type pos struct {
	x, y int
}

func isWall(p pos) bool {
	n := uint64(p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + input)
	n -= (n >> 1) & 0x5555555555555555
	n = (n>>2)&0x3333333333333333 + n&0x3333333333333333
	n += n >> 4
	n &= 0x0f0f0f0f0f0f0f0f
	n *= 0x0101010101010101
	return byte(n>>56)%2 != 0
}

func distances(start pos, max pos) map[pos]int {
	dist := make(map[pos]int)
	recdistances(dist, 0, start, max)
	return dist
}

func recdistances(distances map[pos]int, dist int, p pos, max pos) {
	distances[p] = dist

	possible := []pos{
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x - 1, p.y},
		{p.x, p.y - 1},
	}
	for _, next := range possible {
		if isWall(next) || next.x > max.x || next.y > max.y {
			continue
		}
		if d, ok := distances[next]; !ok || d > dist {
			recdistances(distances, dist+1, next, max)
		}
	}
}

func main() {
	// part 1
	dist := distances(pos{1, 1}, pos{50, 50})
	println(dist[pos{31, 39}])

	// part 2
	var total int
	for y := 0; y < 30; y++ {
		for x := 0; x < 30; x++ {
			if d, ok := dist[pos{x, y}]; ok && d <= 50 {
				total++
			}
		}
	}
	println(total)
}
