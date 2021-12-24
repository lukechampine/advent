package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2021, 16)

func parseInt(bits string) int {
	v, err := strconv.ParseUint(bits, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

type bitReader struct {
	bits string
}

func (br *bitReader) read(n int) string {
	s := br.bits[:n]
	br.bits = br.bits[n:]
	return s
}

func (br *bitReader) readInt(n int) int {
	return parseInt(br.read(n))
}

func (br *bitReader) readLiteral() int {
	var lit string
	for last := false; !last; {
		last = br.readInt(1) == 0
		lit += br.read(4)
	}
	return parseInt(lit)
}

func versionSum(br *bitReader) (v int) {
	v += br.readInt(3)
	id := br.readInt(3)
	if id == 4 {
		_ = br.readLiteral()
	} else if br.readInt(1) == 0 {
		for sub := (&bitReader{bits: br.read(br.readInt(15))}); len(sub.bits) > 0; {
			v += versionSum(sub)
		}
	} else {
		for sub := br.readInt(11); sub > 0; sub-- {
			v += versionSum(br)
		}
	}
	return
}

func eval(br *bitReader) (result int) {
	br.readInt(3)
	id := br.readInt(3)
	if id == 4 {
		return br.readLiteral()
	}

	var subExprs []int
	if br.readInt(1) == 0 {
		for sub := (&bitReader{bits: br.read(br.readInt(15))}); len(sub.bits) > 0; {
			subExprs = append(subExprs, eval(sub))
		}
	} else {
		for sub := br.readInt(11); sub > 0; sub-- {
			subExprs = append(subExprs, eval(br))
		}
	}

	switch id {
	case 0:
		return utils.Sum(len(subExprs), func(i int) int { return subExprs[i] })
	case 1:
		return utils.Product(len(subExprs), func(i int) int { return subExprs[i] })
	case 2:
		return utils.Minimum(len(subExprs), func(i int) int { return subExprs[i] })
	case 3:
		return utils.Maximum(len(subExprs), func(i int) int { return subExprs[i] })
	case 5:
		return utils.BoolToInt(subExprs[0] > subExprs[1])
	case 6:
		return utils.BoolToInt(subExprs[0] < subExprs[1])
	case 7:
		return utils.BoolToInt(subExprs[0] == subExprs[1])
	default:
		panic(id)
	}
}

func main() {
	b, _ := hex.DecodeString(input)
	var bits string
	for i := range b {
		bits += fmt.Sprintf("%08b", b[i])
	}
	utils.Println(versionSum(&bitReader{bits}))
	utils.Println(eval(&bitReader{bits}))
}
