package main

import (
	"lukechampine.com/advent/utils"
)

var input = utils.Input(2020, 25)

func bezout(a, b int) int {
	s0, s1 := 1, 0
	for b != 0 {
		s0, s1 = s1, s0-(a/b)*s1
		a, b = b, a%b
	}
	return s0
}

func main() {
	pubs := utils.ExtractInts(input)
	cardPub, doorPub := pubs[0], pubs[1]

	inv7 := bezout(7, 20201227) + 20201227 // modular inverse
	secret := 1
	for cardPub != 1 {
		cardPub = (cardPub * inv7) % 20201227
		secret = (secret * doorPub) % 20201227
	}
	utils.Println(secret)
}
