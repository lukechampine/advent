package main

import (
	"github.com/lukechampine/advent/utils"
)

const input = `1113222113`

func split(str string) []string {
	var groups []string
	for len(str) != 0 {
		var i int
		for i = 0; i < len(str) && str[i] == str[0]; i++ {
		}
		groups = append(groups, str[:i])
		str = str[i:]
	}
	return groups
}

func lookandsay(str string) string {
	var end string
	for _, nums := range split(str) {
		end += utils.Itoa(len(nums)) + nums[:1]
	}
	return end
}

func main() {
	// part 1
	var final string = input
	for i := 0; i < 40; i++ {
		final = lookandsay(final)
	}
	println(len(final))

	// part 2
	final = input
	for i := 0; i < 50; i++ {
		final = lookandsay(final)
	}
	println(len(final))
}
