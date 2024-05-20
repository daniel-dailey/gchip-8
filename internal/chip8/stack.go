package chip8

func (s *Stack) decrementStackPointer() {
	s.stackPointer -= 1
}

func (s *Stack) incrementStackPointer() {
	s.stackPointer += 1
}

func (s *Stack) setProgramCounter(pc uint16) {
	s.programCounter = pc
}

func (s *Stack) getProgramCounter() uint16 {
	return s.programCounter
}

func (s *Stack) incrementProgramCounter() {
	s.programCounter += 2
}

func (s *Stack) decrementProgramCounter() {
	s.programCounter -= 2
}

func (s *Stack) getCurStackVal() uint16 {
	return s.stack[s.stackPointer]
}

func (s *Stack) setCurStackVal(val uint16) {
	s.stack[s.stackPointer] = val
}

func (s *Stack) getStackPointer() uint16 {
	return s.stackPointer
}

func InitStack() *Stack {
	return &Stack{
		stack:          [16]uint16{},
		programCounter: StartAddr,
	}
}
