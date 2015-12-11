package main

import (
	"strings"
)

const input = `hxbxwxba`

func increment(str string) string {
	init, last := str[:len(str)-1], str[len(str)-1]
	if last == 'z' {
		return increment(init) + "a"
	} else {
		return init + string(last+1)
	}
}

func rule1(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i+1] == str[i]+1 && str[i+2] == str[i]+2 {
			return true
		}
	}
	return false
}

func rule2(str string) bool {
	return !strings.ContainsAny(str, "iol")
}

func rule3(str string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	n := 0
	for _, l := range alphabet {
		if strings.Contains(str, string(l)+string(l)) {
			n++
		}
	}
	return n >= 2
}

func valid(pw string) bool {
	return rule1(pw) && rule2(pw) && rule3(pw)
}

func nextPassword(pw string) string {
	for !valid(pw) {
		pw = increment(pw)
	}
	return pw
}

func main() {
	// part 1
	println(nextPassword(input))

	// part 2
	println(nextPassword(increment(nextPassword(input))))
}
