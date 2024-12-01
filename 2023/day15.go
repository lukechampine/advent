package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 15)

func hash(s string) (h uint8) {
	for i := range s {
		h += s[i]
		h *= 17
	}
	return h
}

func main() {
	var sum int
	for _, label := range utils.Split(strings.TrimSpace(input), ",") {
		sum += int(hash(label))
	}
	utils.Println(sum)

	type entry struct {
		label string
		focal int
	}
	m := make(map[uint8][]entry)
	for _, label := range utils.Split(strings.TrimSpace(input), ",") {
		if l, f, ok := strings.Cut(label, "="); ok {
			e := entry{l, utils.Atoi(f)}
			es := m[hash(l)]
			i := 0
			for i < len(es) && es[i].label != l {
				i++
			}
			if i < len(es) {
				es[i] = e
			} else {
				m[hash(l)] = append(es, e)
			}
		} else {
			l := strings.TrimSuffix(label, "-")
			es := m[hash(l)]
			for i, e := range es {
				if e.label == l {
					m[hash(l)] = append(es[:i], es[i+1:]...)
					break
				}
			}
		}
	}
	sum = 0
	for i := 0; i < 256; i++ {
		for j, es := range m[uint8(i)] {
			sum += (i + 1) * (j + 1) * es.focal
		}
	}
	utils.Println(sum)
}
