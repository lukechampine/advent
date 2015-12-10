package main

const input = `1113222113`

func run(b []byte, c byte) int // day10_amd64.s

func lookandsay(buf []byte) []byte {
	var i int
	end := make([]byte, 0, len(buf)*2)
	for len(buf) != 0 {
		i = run(buf, buf[0])
		end = append(end, byte(i), buf[0])
		buf = buf[i:]
	}
	return end
}

func main() {
	// part 1
	final := make([]byte, len(input))
	for i := range final {
		final[i] = input[i] - '0'
	}
	for i := 0; i < 40; i++ {
		final = lookandsay(final)
	}
	println(len(final))

	// part 2
	for i := 0; i < 10; i++ {
		final = lookandsay(final)
	}
	println(len(final))
}
