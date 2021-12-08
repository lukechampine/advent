package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 8)
var inputLines = utils.Lines(input)

var segments = map[string]byte{
	"abcefg":  '0',
	"cf":      '1',
	"acdeg":   '2',
	"acdfg":   '3',
	"bcdf":    '4',
	"abdfg":   '5',
	"abdefg":  '6',
	"acf":     '7',
	"abcdefg": '8',
	"abcdfg":  '9',
}

var perms = func() []string {
	all := "abcdefg"
	var perms []string
	for _, p := range utils.Perms(7) {
		b := make([]byte, 7)
		for i, j := range p {
			b[i] = all[j]
		}
		perms = append(perms, string(b))
	}
	return perms
}()

func sub(seq string, mapping string) string {
	b := []byte(seq)
	for i, c := range b {
		if c != ' ' {
			b[i] = mapping[c-'a']
		}
	}
	return string(b)
}

func isValid(seq string) bool {
	digits := utils.Split(seq, " ")
	return utils.All(len(digits), func(i int) bool {
		return segments[utils.SortString(digits[i])] != 0
	})
}

func solve(seq string) string {
	for _, p := range perms {
		if isValid(sub(seq, p)) {
			return p
		}
	}
	panic("unsolvable")
}

func atoi(output string) int {
	var digits []byte
	for _, o := range utils.Split(output, " ") {
		digits = append(digits, segments[utils.SortString(o)])
	}
	return utils.Atoi(string(digits))
}

func main() {
	var sum int
	for _, line := range inputLines {
		seq := utils.Split(line, " | ")[1]
		for _, w := range utils.Split(seq, " ") {
			switch len(w) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	utils.Println(sum)

	sum = 0
	for _, line := range inputLines {
		parts := utils.Split(line, " | ")
		sum += atoi(sub(parts[1], solve(parts[0])))
	}
	utils.Println(sum)
}
