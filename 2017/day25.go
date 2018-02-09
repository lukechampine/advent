package main

type machine struct {
	tapeLeft  []int
	curVal    int
	tapeRight []int
}

func moveLeft(m *machine) {
	m.tapeRight = append(m.tapeRight, m.curVal)
	if len(m.tapeLeft) == 0 {
		m.curVal = 0
	} else {
		m.curVal = m.tapeLeft[len(m.tapeLeft)-1]
		m.tapeLeft = m.tapeLeft[:len(m.tapeLeft)-1]
	}
}

func moveRight(m *machine) {
	m.tapeLeft = append(m.tapeLeft, m.curVal)
	if len(m.tapeRight) == 0 {
		m.curVal = 0
	} else {
		m.curVal = m.tapeRight[len(m.tapeRight)-1]
		m.tapeRight = m.tapeRight[:len(m.tapeRight)-1]
	}
}

type state func(*machine) state

func metaState(w0 int, m0 func(*machine), s0 state, w1 int, m1 func(*machine), s1 state) state {
	return func(m *machine) state {
		if m.curVal == 0 {
			m.curVal = w0
			m0(m)
			return s0
		} else {
			m.curVal = w1
			m1(m)
			return s1
		}
	}
}

func stateA(m *machine) state { return metaState(1, moveRight, stateB, 0, moveRight, stateC)(m) }
func stateB(m *machine) state { return metaState(0, moveLeft, stateA, 0, moveRight, stateD)(m) }
func stateC(m *machine) state { return metaState(1, moveRight, stateD, 1, moveRight, stateA)(m) }
func stateD(m *machine) state { return metaState(1, moveLeft, stateE, 0, moveLeft, stateD)(m) }
func stateE(m *machine) state { return metaState(1, moveRight, stateF, 1, moveLeft, stateB)(m) }
func stateF(m *machine) state { return metaState(1, moveRight, stateA, 1, moveRight, stateE)(m) }

func main() {
	// part 1
	var m machine
	s := stateA
	for i := 0; i < 12399302; i++ {
		s = s(&m)
	}
	n := m.curVal
	for _, v := range m.tapeLeft {
		n += v
	}
	for _, v := range m.tapeRight {
		n += v
	}
	println(n)
}
