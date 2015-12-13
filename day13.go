package main

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 81 happiness units by sitting next to Carol.
Alice would lose 42 happiness units by sitting next to David.
Alice would gain 89 happiness units by sitting next to Eric.
Alice would lose 89 happiness units by sitting next to Frank.
Alice would gain 97 happiness units by sitting next to George.
Alice would lose 94 happiness units by sitting next to Mallory.
Bob would gain 3 happiness units by sitting next to Alice.
Bob would lose 70 happiness units by sitting next to Carol.
Bob would lose 31 happiness units by sitting next to David.
Bob would gain 72 happiness units by sitting next to Eric.
Bob would lose 25 happiness units by sitting next to Frank.
Bob would lose 95 happiness units by sitting next to George.
Bob would gain 11 happiness units by sitting next to Mallory.
Carol would lose 83 happiness units by sitting next to Alice.
Carol would gain 8 happiness units by sitting next to Bob.
Carol would gain 35 happiness units by sitting next to David.
Carol would gain 10 happiness units by sitting next to Eric.
Carol would gain 61 happiness units by sitting next to Frank.
Carol would gain 10 happiness units by sitting next to George.
Carol would gain 29 happiness units by sitting next to Mallory.
David would gain 67 happiness units by sitting next to Alice.
David would gain 25 happiness units by sitting next to Bob.
David would gain 48 happiness units by sitting next to Carol.
David would lose 65 happiness units by sitting next to Eric.
David would gain 8 happiness units by sitting next to Frank.
David would gain 84 happiness units by sitting next to George.
David would gain 9 happiness units by sitting next to Mallory.
Eric would lose 51 happiness units by sitting next to Alice.
Eric would lose 39 happiness units by sitting next to Bob.
Eric would gain 84 happiness units by sitting next to Carol.
Eric would lose 98 happiness units by sitting next to David.
Eric would lose 20 happiness units by sitting next to Frank.
Eric would lose 6 happiness units by sitting next to George.
Eric would gain 60 happiness units by sitting next to Mallory.
Frank would gain 51 happiness units by sitting next to Alice.
Frank would gain 79 happiness units by sitting next to Bob.
Frank would gain 88 happiness units by sitting next to Carol.
Frank would gain 33 happiness units by sitting next to David.
Frank would gain 43 happiness units by sitting next to Eric.
Frank would gain 77 happiness units by sitting next to George.
Frank would lose 3 happiness units by sitting next to Mallory.
George would lose 14 happiness units by sitting next to Alice.
George would lose 12 happiness units by sitting next to Bob.
George would lose 52 happiness units by sitting next to Carol.
George would gain 14 happiness units by sitting next to David.
George would lose 62 happiness units by sitting next to Eric.
George would lose 18 happiness units by sitting next to Frank.
George would lose 17 happiness units by sitting next to Mallory.
Mallory would lose 36 happiness units by sitting next to Alice.
Mallory would gain 76 happiness units by sitting next to Bob.
Mallory would lose 34 happiness units by sitting next to Carol.
Mallory would gain 37 happiness units by sitting next to David.
Mallory would gain 40 happiness units by sitting next to Eric.
Mallory would gain 18 happiness units by sitting next to Frank.
Mallory would gain 7 happiness units by sitting next to George.`

type person struct {
	name string
	happ map[string]int
}

func parse(str string, m map[string]person) {
	var name1, sign, name2 string
	var mag int
	str = strings.TrimSuffix(str, ".") // otherwise . is added to name2
	utils.Sscanf(str, "%s would %s %d happiness units by sitting next to %s", &name1, &sign, &mag, &name2)
	if _, ok := m[name1]; !ok {
		m[name1] = person{
			name: name1,
			happ: make(map[string]int),
		}
	}
	if sign == "lose" {
		mag = -mag
	}
	m[name1].happ[name2] = mag
}

func calcHappiness(seating []person) int {
	// add wraparound for easier loop logic
	circle := make([]person, len(seating)+2)
	circle[0] = seating[len(seating)-1]
	copy(circle[1:], seating)
	circle[len(circle)-1] = seating[0]

	var total int
	for i := 1; i < len(circle)-2; i++ {
		total += circle[i].happ[circle[i-1].name]
		total += circle[i].happ[circle[i+1].name]
	}
	return total
}

func main() {
	// part 1
	m := make(map[string]person)
	for _, str := range utils.Lines(input) {
		parse(str, m)
	}
	var guests []person
	for _, guest := range m {
		guests = append(guests, guest)
	}

	var optimal int
	for _, perm := range utils.Perms(len(guests)) {
		seating := make([]person, len(guests))
		for si, gi := range perm {
			seating[si] = guests[gi]
		}
		happiness := calcHappiness(seating)
		if happiness > optimal {
			optimal = happiness
		}
	}
	println(optimal)

	// part 2
	myHapp := make(map[string]int)
	for name := range m {
		m[name].happ["Me"] = 0
		myHapp[name] = 0
	}
	m["Me"] = person{
		name: "Me",
		happ: myHapp,
	}
	guests = append(guests, m["Me"])

	optimal = 0
	for _, perm := range utils.Perms(len(guests)) {
		seating := make([]person, len(guests))
		for si, gi := range perm {
			seating[si] = guests[gi]
		}
		happiness := calcHappiness(seating)
		if happiness > optimal {
			optimal = happiness
		}
	}
	println(optimal)
}
