package main

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/lukechampine/advent/utils"
)

const input = `uqwqemis`

func fivezeros(h string) bool {
	for _, b := range h[:5] {
		if b != '0' {
			return false
		}
	}
	return true
}

func nextHash(str string, count int) string {
	hash := md5.Sum([]byte(str + utils.Itoa(count)))
	return hex.EncodeToString(hash[:])
}

func password(str string) string {
	pass := make([]byte, 8)
	var count int
	for i := 0; i < 8; i++ {
		h := "xxxxxx"
		for ; !fivezeros(h); count++ {
			h = nextHash(str, count)
		}
		pass[i] = h[5]
	}
	return string(pass)
}

func password2(str string) string {
	pass := make([]byte, 8)
	var count int
	for i := 0; i < 8; {
		h := "xxxxxx"
		for ; !(fivezeros(h) && h[5] < '8'); count++ {
			h = nextHash(str, count)
		}
		pos := utils.Atoi(string(h[5]))
		if pass[pos] == 0 {
			pass[pos] = h[6]
			i++
		}
	}
	return string(pass)
}

func main() {
	// part 1
	println(password(input))

	println(password2(input))
}
