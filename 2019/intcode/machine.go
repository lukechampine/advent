package intcode

import "github.com/lukechampine/advent/utils"

type Machine struct {
	prog   []int
	i      int
	rel    int
	Halted bool
}

func (m *Machine) Run(inputs ...int) (output int) {
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
			*getArg(1) = inputs[0]
			inputs = inputs[1:]
			m.i += 2
		case 4:
			x := *getArg(1)
			m.i += 2
			return x
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
			return -1
		default:
			m.i++
		}
	}
	panic("unreachable")
}

func New(prog []int) *Machine {
	return &Machine{
		prog: append(prog[:len(prog):len(prog)], make([]int, 30000)...),
	}
}
