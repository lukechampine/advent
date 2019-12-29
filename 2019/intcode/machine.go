package intcode

import (
	"strings"

	"github.com/lukechampine/advent/utils"
)

func EncodeASCII(s string) []int {
	ints := make([]int, len(s))
	for i := range s {
		ints[i] = int(s[i])
	}
	return ints
}

func DecodeASCII(ints []int) string {
	s := make([]byte, len(ints))
	for i := range s {
		s[i] = byte(ints[i])
	}
	return string(s)
}

type Machine struct {
	prog          []int
	i             int
	rel           int
	Halted        bool
	AwaitingInput bool
}

func (m *Machine) Run(inputs ...int) (output []int) {
	m.AwaitingInput = false
	p := m.prog
	for m.i < len(p) {
		op, flags := p[m.i]%100, utils.Digits(utils.Itoa(p[m.i]/100))
		getArg := func(n int) *int {
			imm := p[m.i+n]
			if len(flags) < n || flags[len(flags)-n] == 0 {
				return &p[imm]
			} else if flags[len(flags)-n] == 1 {
				return &imm
			} else if flags[len(flags)-n] == 2 {
				return &p[m.rel+imm]
			}
			panic("unreachable")
		}
		switch op {
		case 1:
			*getArg(3) = *getArg(1) + *getArg(2)
			m.i += 4
		case 2:
			*getArg(3) = *getArg(1) * *getArg(2)
			m.i += 4
		case 3:
			if len(inputs) == 0 {
				m.AwaitingInput = true
				return
			}
			*getArg(1) = inputs[0]
			inputs = inputs[1:]
			m.i += 2
		case 4:
			x := *getArg(1)
			m.i += 2
			output = append(output, x)
		case 5:
			if *getArg(1) != 0 {
				m.i = *getArg(2)
			} else {
				m.i += 3
			}
		case 6:
			if *getArg(1) == 0 {
				m.i = *getArg(2)
			} else {
				m.i += 3
			}
		case 7:
			*getArg(3) = utils.BoolToInt(*getArg(1) < *getArg(2))
			m.i += 4
		case 8:
			*getArg(3) = utils.BoolToInt(*getArg(1) == *getArg(2))
			m.i += 4
		case 9:
			m.rel += *getArg(1)
			m.i += 2
		case 99:
			m.Halted = true
			return
		default:
			panic("illegal opcode")
		}
	}
	panic("unreachable")
}

func (m *Machine) RunASCII(input string) (output []int) {
	input = strings.TrimPrefix(input, "\n") // for more readable `` literals
	return m.Run(EncodeASCII(input)...)
}

func New(prog []int) *Machine {
	return &Machine{
		prog: append(prog[:len(prog):len(prog)], make([]int, 30000)...),
	}
}
