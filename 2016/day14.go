package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/lukechampine/advent/utils"
)

const input = "yjdafjpo"

type hasher interface {
	hash(int) string
}

type memoHash map[int]string

func (memo memoHash) hash(i int) string {
	if h, ok := memo[i]; ok {
		return h
	}
	h := md5.Sum([]byte(input + utils.Itoa(i)))
	s := hex.EncodeToString(h[:])
	memo[i] = s
	return s
}

type memoStretchHash map[int]string

func (memo memoStretchHash) hash(i int) string {
	if h, ok := memo[i]; ok {
		return h
	}

	var buf [16]byte
	m := md5.New()
	m.Write([]byte(input + utils.Itoa(i)))
	s := []byte(hex.EncodeToString(m.Sum(buf[:0])))
	for i := 0; i < 2016; i++ {
		m.Reset()
		m.Write(s)
		hex.Encode(s, m.Sum(buf[:0]))
	}
	memo[i] = string(s)
	return string(s)
}

func triple(h string) string {
	for i := 2; i < len(h); i++ {
		if h[i] == h[i-1] && h[i] == h[i-2] {
			return string(h[i])
		}
	}
	return ""
}

func hasQuint(h, t string) bool {
	return strings.Contains(h, t+t+t+t+t)
}

func isKey(i int, h hasher) bool {
	hash := h.hash(i)
	t := triple(hash)
	if t == "" {
		return false
	}
	for j := 1; j <= 1000; j++ {
		if hasQuint(h.hash(i+j), t) {
			return true
		}
	}
	return false
}

func findIndex(h hasher) int {
	keys := 0
	for i := 0; ; i++ {
		if isKey(i, h) {
			keys++
			if keys == 64 {
				return i
			}
		}
	}
}

func main() {
	// part 1
	println(findIndex(make(memoHash)))

	// part 2
	println(findIndex(make(memoStretchHash)))
}
