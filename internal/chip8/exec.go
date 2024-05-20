package chip8

import "log"

const (
	T0             = 0x00
	JUMP           = 0x01
	SUBROUTINE     = 0x02
	SKIP_EQ        = 0x03
	SKIP_NEQ       = 0x04
	SKIP_VX_EQ_VY  = 0x05
	SET_VX_NN      = 0x06
	VX_INC_NN      = 0x07
	T8             = 0x08
	SKIP_VX_NEQ_VY = 0x09
	SET_I_NNN      = 0x0A
	JMP_NNN_V0     = 0x0B
	RAND_NN_MASK   = 0x0C
	DRAW           = 0x0D
	TE             = 0x0E
	TF             = 0x0F

	CLEAR  = 0xE0
	RETURN = 0xEE

	COPY_V_REGISTER = 0x0
	OR_V_REGISTER   = 0x1
	AND_V_REGISTER  = 0x2
	XOR_V_REGISTER  = 0x3
	SUM_V_REGISTER  = 0x4
	DECREMENT_VX    = 0x5
	SHIFT_RIGHT     = 0x6
	DIFF_V_REGISTER = 0x7
	SHIFT_LEFT      = 0xE

	SKIP_ON_KEY_PRESSED  = 0x9E
	SKIP_ON_KEY_RELEASED = 0xA1

	SET_VX_DELAY_TIMER = 0x07
	WAIT_FOR_KEY       = 0x0A
	SET_DELAY_TIMER_VX = 0x15
	SET_SOUND_TIMER    = 0x18
	ADD_VX_TO_I        = 0x1E
	SET_I_TO_SPRITE    = 0x29
	SET_BCD            = 0x33
	REG_DUMP           = 0x55
	READ_REGISTERS     = 0x65
)

// executeCurrentInstruction decodes & executes the current opcode
func (c *Chip8) executeCurrentInstruction() {
	op := c.opcode.opDecode()
	vx := c.opcode.vx()
	vy := c.opcode.vy()
	n := c.opcode.n()
	nn := c.opcode.nn()
	nnn := c.opcode.nnn()
	switch op {
	case T0:
		c.executeInstructionType0(nn)
	case JUMP:
		//Jump to location nnn
		c.stack.setProgramCounter(nnn)
	case SUBROUTINE:
		//Call subroutine at nnn
		c.stack.incrementStackPointer()
		c.stack.setCurStackVal(c.stack.getProgramCounter())
		c.stack.setProgramCounter(nnn)
	case SKIP_EQ:
		//Skip next instruction if Vx = nn
		if c.registers.getVRegisterVal(vx) == nn {
			c.stack.incrementProgramCounter()
		}
	case SKIP_NEQ:
		//Skip next instruction if Vx != nn
		if c.registers.getVRegisterVal(vx) != nn {
			c.stack.incrementProgramCounter()
		}
	case SKIP_VX_EQ_VY:
		//Skip next instruction if Vx = Vy
		if c.registers.areVRegistersEqual(vx, vy) {
			c.stack.incrementProgramCounter()
		}
	case SET_VX_NN:
		//Set Vx = nn
		c.registers.setVRegister(vx, nn)
	case VX_INC_NN:
		//Set Vx = Vx + nn
		c.registers.incrementVRegister(vx, nn)
	case T8:
		// fmt.Println(vx, vy, c.opcode&0xFF)
		c.executeInstructionType8(vx, vy, n)
	case SKIP_VX_NEQ_VY:
		//Skip next instruction if Vx != Vy
		if !c.registers.areVRegistersEqual(vx, vy) {
			c.stack.incrementProgramCounter()
		}
	case SET_I_NNN:
		//Set I = nnn
		c.registers.setIRegister(nnn)
	case JMP_NNN_V0:
		//Jump to location nnn + V0
		c.stack.setProgramCounter(nnn + uint16(c.registers.getVRegisterVal(0)))
	case RAND_NN_MASK:
		//Set Vx = random byte AND nn
		c.registers.setVRegister(vx, c.rand()&nn)
	case DRAW:
		//Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
		height := n
		xPos := c.registers.getVRegisterVal(vx) % VideoBufferWidth
		yPos := c.registers.getVRegisterVal(vy) % VideoBufferHeight
		c.registers.clearVRegister(VF)
		fb := c.frameBuf.getFrameBuffer()
		for row := uint16(0); row < uint16(height); row++ {
			p := c.memory.read(c.registers.getIRegister() + row)
			for col := uint16(0); col < 8; col++ {
				if (p & (0x80 >> col)) == 0 {
					continue
				}
				displayIndex := (uint16(yPos)+row)*VideoBufferWidth + (uint16(xPos) + col)
				screenPixel := fb[displayIndex]
				if screenPixel == 0xFFFFFFFF {
					c.registers.setVRegister(VF, 1)
				}
				c.frameBuf.setPixel(displayIndex)
			}
		}
	case TE:
		c.executeInstructionTypeE(vx, nn)
	case TF:
		c.executeInstructionTypeF(vx, nn)
	default:
		c.no_op()
	}

}

// executeInstructionType0 executes sub instruction for type 0 instructions
func (c *Chip8) executeInstructionType0(instruction uint8) {
	switch instruction {
	case CLEAR:
		//Clear the display
		c.frameBuf.clear()
	case RETURN:
		//Return from a subroutine
		c.stack.setProgramCounter(c.stack.getCurStackVal())
		c.stack.decrementStackPointer()
		log.Println("Returning to location: ", c.stack.getCurStackVal())
	default:
		c.no_op()
	}
}

// executeInstructionType8 executes sub instruction for type 8 instructions
func (c *Chip8) executeInstructionType8(vx, vy uint16, instruction uint8) {
	switch instruction {
	case COPY_V_REGISTER:
		//Set Vx = Vy
		c.registers.copyVRegister(vx, vy)
	case OR_V_REGISTER:
		//Set Vx = Vx OR Vy
		c.registers.orVRegister(vx, vy)
	case AND_V_REGISTER:
		//Set Vx = Vx AND Vy
		c.registers.andVRegister(vx, vy)
	case XOR_V_REGISTER:
		//Set Vx = Vx XOR Vy
		c.registers.xorVRegister(vx, vy)
	case SUM_V_REGISTER:
		//Set Vx = Vx + Vy, set VF = carry
		sum := c.registers.sumVRegister(vx, vy)
		c.registers.setVRegister(vx, sum&0xFF)
		if sum > 255 {
			c.registers.setVRegister(VF, 1)
			return
		}
		c.registers.clearVRegister(VF)
	case DECREMENT_VX:
		//Set Vx = Vx - Vy, set VF = NOT borrow
		c.registers.decrementVRegister(vx, c.registers.getVRegisterVal(vy))
		if c.registers.isGreaterThan(vx, vy) {
			c.registers.setVRegister(VF, 1)
			return
		}
		c.registers.clearVRegister(VF)
	case SHIFT_RIGHT:
		//Set Vx = Vx SHR 1
		maskedVxRegisterValue := c.registers.getVRegisterVal(vx) & 0x1
		c.registers.setVRegister(VF, maskedVxRegisterValue)
		c.registers.shiftRightVRegister(vx)
	case DIFF_V_REGISTER:
		//Set Vx = Vy - Vx, set VF = NOT borrow
		diff := c.registers.diffVRegister(vy, vx)
		c.registers.setVRegister(vx, diff)
		if c.registers.isGreaterThan(vy, vx) {
			c.registers.setVRegister(VF, 1)
			return
		}
		c.registers.clearVRegister(VF)
	case SHIFT_LEFT:
		//Set Vx = Vx SHL 1
		c.registers.setVRegister(VF, c.registers.getVRegisterVal(vx)>>7)
		c.registers.shiftLeftVRegister(vx)
	default:
		c.no_op()
	}
}

// executeInstructionTypeE executes sub instruction for type E instructions
func (c *Chip8) executeInstructionTypeE(vx uint16, instruction uint8) {
	switch instruction {
	case SKIP_ON_KEY_PRESSED:
		//Skip next instruction if key with the value of Vx is pressed
		key := c.registers.getVRegisterVal(vx)
		if c.keys[key] != 0 {
			c.stack.incrementProgramCounter()
		}
	case SKIP_ON_KEY_RELEASED:
		//Skip next instruction if key with the value of Vx is not pressed
		key := c.registers.getVRegisterVal(vx)
		if c.keys[key] == 0 {
			c.stack.incrementProgramCounter()
		}
	default:
		c.no_op()
	}
}

// executeInstructionTypeF executes sub instruction for type F instructions
func (c *Chip8) executeInstructionTypeF(vx uint16, instruction uint8) {
	switch instruction {
	case SET_VX_DELAY_TIMER:
		//Set Vx = delay timer value
		c.registers.setVRegister(vx, c.registers.getDelay())
	case WAIT_FOR_KEY:
		c.stack.decrementProgramCounter()
	case SET_DELAY_TIMER_VX:
		//Set delay timer = Vx
		c.registers.setDelay(c.registers.getVRegisterVal(vx))
	case SET_SOUND_TIMER:
		//Set sound timer = Vx
		c.registers.setSound(c.registers.getVRegisterVal(vx))
	case ADD_VX_TO_I:
		//Set I = I + Vx
		i := c.registers.getIRegister()
		c.registers.setIRegister(i + uint16(c.registers.getVRegisterVal(vx)))
	case SET_I_TO_SPRITE:
		//Set I = location of sprite for digit Vx
		addr := FontsetStartAddr + c.registers.getVRegisterVal(vx)
		c.registers.setIRegister(uint16(addr) * 5)
	case SET_BCD:
		//Store BCD representation of Vx in memory locations I, I+1, and I+2
		val := c.registers.getVRegisterVal(vx)
		iReg := c.registers.getIRegister()
		c.memory.write(iReg+2, val%10)
		val /= 10
		c.memory.write(iReg+1, val%10)
		val /= 10
		c.memory.write(iReg, val%10)
	case REG_DUMP:
		for i := uint16(0); i <= vx; i++ {
			iReg := c.registers.getIRegister()
			c.memory.write(iReg+i, c.registers.getVRegisterVal(i))
		}
	case READ_REGISTERS:
		for i := uint16(0); i <= vx; i++ {
			mem := c.memory.read(c.registers.getIRegister() + i)
			c.registers.setVRegister(i, mem)
		}
	default:
		c.no_op()
	}
}

func (c *Chip8) no_op() {
	//Do nothing
}
