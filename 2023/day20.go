package main

import (
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 20)

type pulse struct {
	src, dst string
	high     bool
}

type module struct {
	name string
	typ  byte
	dsts []string
	on   bool            // for flip-flops
	mem  map[string]bool // for conjuctions
}

func (m *module) send(high bool) []pulse {
	ps := make([]pulse, len(m.dsts))
	for i, dst := range m.dsts {
		ps[i] = pulse{m.name, dst, high}
	}
	return ps
}

func (m *module) recv(p pulse) []pulse {
	switch m.typ {
	case 'b': // broadcast
		return m.send(p.high)
	case '%': // flip-flop
		if p.high {
			return nil
		} else {
			m.on = !m.on
			return m.send(m.on)
		}
	case '&': // conjunction
		m.mem[p.src] = p.high
		for _, v := range m.mem {
			if !v {
				return m.send(true)
			}
		}
		return m.send(false)
	}
	panic("unreachable")
}

func main() {
	modules := make(map[string]*module)
	for _, line := range utils.Lines(input) {
		name, dsts, _ := strings.Cut(line, " -> ")
		modules[strings.Trim(name, "%&")] = &module{
			name: strings.Trim(name, "%&"),
			typ:  name[0],
			dsts: strings.Split(dsts, ", "),
			on:   false,
			mem:  make(map[string]bool),
		}
	}
	// wire up conjunctions
	for _, src := range modules {
		for _, name := range src.dsts {
			if dst, ok := modules[name]; ok && dst.typ == '&' {
				dst.mem[src.name] = false
			}
		}
	}

	var high, low int
	cycles := make(map[string]int)
	for i := 1; ; i++ {
		queue := []pulse{{"button", "broadcaster", false}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.high {
				high++
			} else {
				low++
			}
			if p.dst == "lv" && p.high {
				cycles[p.src] = i - cycles[p.src]
				if len(cycles) == 4 {
					i := 1
					for _, n := range cycles {
						i *= n
					}
					utils.Println(i)
					return
				}
			}

			if dst, ok := modules[p.dst]; ok {
				queue = append(queue, dst.recv(p)...)
			}
		}
		if i == 1000 {
			utils.Println(high * low)
		}
	}
}
