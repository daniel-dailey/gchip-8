package chip8

func (r *Registers) getVRegisterVal(register uint16) uint8 {
	return r.vRegister[register]
}

func (r *Registers) setVRegister(register uint16, value uint8) {
	r.vRegister[register] = value
}

func (r *Registers) clearVRegister(register uint16) {
	r.vRegister[register] = 0
}

func (r *Registers) incrementVRegister(register uint16, val uint8) {
	r.vRegister[register] += val
}

func (r *Registers) decrementVRegister(register uint16, val uint8) {
	r.vRegister[register] -= val
}

func (r *Registers) copyVRegister(register1, register2 uint16) {
	r.vRegister[register1] = r.vRegister[register2]
}

func (r *Registers) orVRegister(register1, register2 uint16) {
	r.vRegister[register1] |= r.vRegister[register2]
}

func (r *Registers) andVRegister(register1, register2 uint16) {
	r.vRegister[register1] &= r.vRegister[register2]
}

func (r *Registers) xorVRegister(register1, register2 uint16) {
	r.vRegister[register1] ^= r.vRegister[register2]
}

func (r *Registers) sumVRegister(register1 uint16, register2 uint16) uint8 {
	return r.vRegister[register1] + r.vRegister[register2]
}

func (r *Registers) diffVRegister(register1 uint16, register2 uint16) uint8 {
	return r.vRegister[register1] - r.vRegister[register2]
}

func (r *Registers) shiftRightVRegister(register uint16) {
	r.vRegister[register] >>= 1
}

func (r *Registers) shiftLeftVRegister(register uint16) {
	r.vRegister[register] <<= 1
}

func (r *Registers) isGreaterThan(register1 uint16, register2 uint16) bool {
	return r.vRegister[register1] > r.vRegister[register2]
}

func (r *Registers) areVRegistersEqual(register1 uint16, register2 uint16) bool {
	return r.vRegister[register1] == r.vRegister[register2]
}

func (r *Registers) getIRegister() uint16 {
	return r.iRegister
}

func (r *Registers) setIRegister(value uint16) {
	r.iRegister = value
}

func (r *Registers) setDelay(value uint8) {
	r.delay = value
}

func (r *Registers) getDelay() uint8 {
	return r.delay
}

func (r *Registers) decrementDelay() {
	if r.delay > 0 {
		r.delay--
	}
}

func (r *Registers) setSound(value uint8) {
	r.sound = value
}

func (r *Registers) getSound() uint8 {
	return r.sound
}

func (r *Registers) decrementSound() {
	if r.sound > 0 {
		r.sound--
	}
}

func InitRegisters() *Registers {
	return &Registers{
		vRegister: [16]uint8{},
	}
}
