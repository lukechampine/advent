package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const input = "bgvyzdsv"

func main() {
	// part 1
	for i := 0; ; i++ {
		data := input + strconv.Itoa(i)
		sum := fmt.Sprintf("%x", md5.Sum([]byte(data)))
		if strings.HasPrefix(sum, "00000") {
			println(i)
			break
		}
	}

	// part 2
	for i := 0; ; i++ {
		data := input + strconv.Itoa(i)
		sum := fmt.Sprintf("%x", md5.Sum([]byte(data)))
		if strings.HasPrefix(sum, "000000") {
			println(i)
			break
		}
	}
}
