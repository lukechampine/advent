package main

import (
	"fmt"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2022, 25)

func SNAFUdigit(c byte) int {
	switch c {
	case '=':
		return -2
	case '-':
		return -1
	default:
		return int(c - '0')
	}
}

func toDigit(i int) byte {
	switch i {
	case -2:
		return '='
	case -1:
		return '-'
	default:
		return byte(i + '0')
	}
}

func toSNAFU(x int) string {
	var s []byte
	for x > 0 {
		s = append(s, byte((x%5)+'0'))
		x /= 5
	}
	s = append(s, '0')
	for i := range s {
		if s[i] >= '3' {
			s[i+1]++
			s[i] = toDigit(int(s[i]-'0') - 5)
		}
	}

	return strings.TrimLeft(utils.ReverseString(string(s)), "0")
}

func toSNAFU2(x int) string {
	// convert to base 5
	var b5 []int
	for x > 0 {
		b5 = append(b5, (x%5)+2)
		x /= 5
	}
	// add 2222...
	var carry int
	for i := range b5 {
		d := b5[i] + carry
		b5[i] = d % 5
		carry = d / 5
	}
	if carry > 0 {
		b5 = append(b5, carry)
	}

	// render
	var s string
	for i := range b5 {
		s += string(toDigit(b5[i] - 2))
	}

	return utils.ReverseString(s)
}

func fromSNAFU(s string) (x int) {
	s = utils.ReverseString(s)
	pow := 1
	for i := range s {
		x += pow * SNAFUdigit(s[i])
		pow *= 5
	}
	return
}

func main() {
	var sum int
	for _, n := range utils.Lines(input) {
		sum += fromSNAFU(n)
	}
	fmt.Println(toSNAFU(sum))
	fmt.Println(toSNAFU2(sum))
}
