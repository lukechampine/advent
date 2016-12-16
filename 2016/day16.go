package main

const input = "01111001100111011"

func dragon(data string, n int) []byte {
	buf := make([]byte, 0, n)
	buf = append(buf, data...)
	for len(buf) < n {
		buf = append(buf, '0')
		for i := len(buf) - 2; i >= 0; i-- {
			if buf[i] == '0' {
				buf = append(buf, '1')
			} else {
				buf = append(buf, '0')
			}
		}
	}
	return buf[:n]
}

func checksum(data []byte) string {
	chk := data
	for len(chk)%2 == 0 {
		for i := 0; i < len(chk); i += 2 {
			if chk[i] == chk[i+1] {
				chk[i/2] = '1'
			} else {
				chk[i/2] = '0'
			}
		}
		chk = chk[:len(chk)/2]
	}
	return string(chk)
}

func main() {
	// part 1
	println(checksum(dragon(input, 272)))

	// part 2
	println(checksum(dragon(input, 35651584)))
}
