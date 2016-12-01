package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `jio a, +19
inc a
tpl a
inc a
tpl a
inc a
tpl a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
jmp +23
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`

type computer struct {
	a, b  uint
	insts []string
}

func (c *computer) run() {
	for i := 0; i < len(c.insts); i++ {
		switch inst, arg := c.insts[i][:3], c.insts[i][4:]; inst {
		case "hlf":
			if arg == "a" {
				c.a /= 2
			} else {
				c.b /= 2
			}
		case "tpl":
			if arg == "a" {
				c.a *= 3
			} else {
				c.b *= 3
			}
		case "inc":
			if arg == "a" {
				c.a++
			} else {
				c.b++
			}
		case "jmp":
			var off int
			utils.Sscanf(arg, "%d", &off)
			i += off
			i-- // to offset i++ on loop
		case "jie":
			if reg := arg[:1]; reg == "a" && c.a%2 != 0 {
				continue
			} else if reg == "b" && c.b%2 != 0 {
				continue
			}
			var off int
			utils.Sscanf(arg[3:], "%d", &off)
			i += off
			i--
		case "jio":
			if reg := arg[:1]; reg == "a" && c.a != 1 {
				continue
			} else if reg == "b" && c.b != 1 {
				continue
			}
			var off int
			utils.Sscanf(arg[3:], "%d", &off)
			i += off
			i--
		}
	}
}

func main() {
	// part 1
	c := computer{insts: utils.Lines(input)}
	c.run()
	println(c.b)

	// part 2
	c.a, c.b = 1, 0
	c.run()
	println(c.b)
}
