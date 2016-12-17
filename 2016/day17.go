package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const input = "veumntbg"

type pos struct {
	x, y int
	path string
}

func shortest(paths []string) string {
	if len(paths) == 0 {
		return ""
	}
	best := paths[0]
	for _, p := range paths[1:] {
		if len(p) < len(best) {
			best = p
		}
	}
	return best
}
func longest(paths []string) string {
	if len(paths) == 0 {
		return ""
	}
	best := paths[0]
	for _, p := range paths[1:] {
		if len(p) > len(best) {
			best = p
		}
	}
	return best
}

func choices(p pos) []pos {
	h := fmt.Sprintf("%x", md5.Sum([]byte(input+p.path)))
	possible := []pos{
		{p.x, p.y - 1, p.path + "U"},
		{p.x, p.y + 1, p.path + "D"},
		{p.x - 1, p.y, p.path + "L"},
		{p.x + 1, p.y, p.path + "R"},
	}
	filtered := possible[:0]
	for i, pp := range possible {
		if 0 <= pp.x && pp.x <= 3 && 0 <= pp.y && pp.y <= 3 && strings.ContainsAny(h[i:i+1], "bcdef") {
			filtered = append(filtered, pp)
		}
	}
	return filtered
}

func path(p pos, best func([]string) string) string {
	if p.x == 3 && p.y == 3 {
		return p.path
	}
	cs := choices(p)
	var ps []string
	for _, c := range cs {
		if y := path(c, best); y != "" {
			ps = append(ps, y)
		}
	}
	return best(ps)
}

func main() {
	// part 1
	println(path(pos{0, 0, ""}, shortest))

	// part 2
	println(len(path(pos{0, 0, ""}, longest)))
}
