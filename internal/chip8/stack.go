package chip8

// decrements the stack pointer by 1
func (s *Stack) decrementStackPointer() {
	s.stackPointer -= 1
}

// increments the stack pointer by 1
func (s *Stack) incrementStackPointer() {
	s.stackPointer += 1
}

// sets program counter to the given value
func (s *Stack) setProgramCounter(pc uint16) {
	s.programCounter = pc
}

// returns the program counter
func (s *Stack) getProgramCounter() uint16 {
	return s.programCounter
}

// increments the program counter by 2, as we consume
// 2 bytes for each opcode
func (s *Stack) incrementProgramCounter() {
	s.programCounter += 2
}

// decrements the program counter by 2 for the same reason
func (s *Stack) decrementProgramCounter() {
	s.programCounter -= 2
}

// returns the value at the current stack pointer
func (s *Stack) getCurStackVal() uint16 {
	return s.stack[s.stackPointer]
}

// sets the value at the current stack pointer
func (s *Stack) setCurStackVal(val uint16) {
	s.stack[s.stackPointer] = val
}

// returns the stack pointer
func (s *Stack) getStackPointer() uint16 {
	return s.stackPointer
}

// initializes the stack
func InitStack() *Stack {
	return &Stack{
		stack:          [16]uint16{},
		programCounter: StartAddr,
	}
}
