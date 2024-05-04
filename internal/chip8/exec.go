package chip8

import "fmt"

func (c *Chip8) executeCurrentInstruction() {
	op := (c.opcode & 0xF000) >> 12
	vx := c.opcode.vx()
	vy := c.opcode.vy()
	nn := c.opcode.nn()
	nnn := c.opcode.nnn()
	switch op {
	case 0x0:
		c.executeInstructionType0()
	case 0x1:
		//Jump to location nnn
		c.stack.setProgramCounter(nnn)
	case 0x2:
		//Call subroutine at nnn
		c.stack.incrementStackPointer()
		c.stack.setCurStackVal(c.stack.getProgramCounter())
		c.stack.setProgramCounter(nnn)
	case 0x3:
		//Skip next instruction if Vx = nn
		if c.registers.getVRegisterVal(vx) == nn {
			c.stack.incrementProgramCounter()
		}
	case 0x4:
		//Skip next instruction if Vx != nn
		if c.registers.getVRegisterVal(vx) != nn {
			c.stack.incrementProgramCounter()
		}
	case 0x5:
		//Skip next instruction if Vx = Vy
		if c.registers.areVRegistersEqual(vx, vy) {
			c.stack.incrementProgramCounter()
		}
	case 0x6:
		//Set Vx = nn
		c.registers.setVRegister(vx, nn)
	case 0x7:
		//Set Vx = Vx + nn
		c.registers.incrementVRegister(vx, nn)
	case 0x8:
		// fmt.Println(vx, vy, c.opcode&0xFF)
		c.executeInstructionType8(vx, vy)
	case 0x9:
		//Skip next instruction if Vx != Vy
		if !c.registers.areVRegistersEqual(vx, vy) {
			c.stack.incrementProgramCounter()
		}
	case 0xA:
		//Set I = nnn
		c.registers.setIRegister(nnn)
	case 0xB:
		//Jump to location nnn + V0
		c.stack.setProgramCounter(nnn + uint16(c.registers.getVRegisterVal(0)))
	case 0xC:
		//Set Vx = random byte AND nn
		c.registers.setVRegister(vx, c.rand()&nn)
	case 0xD:
		//Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision
		height := c.opcode.n()
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
	case 0xE:
		c.executeInstructionTypeE(vx)
	case 0xF:
		c.executeInstructionTypeF(vx)
	}

}

func (c *Chip8) executeInstructionType0() {
	switch c.opcode & 0xFF {
	case 0xE0:
		//Clear the display
		c.frameBuf.clear()
	case 0xEE:
		//Return from a subroutine
		c.stack.setProgramCounter(c.stack.getCurStackVal())
		c.stack.decrementStackPointer()
	default:
		c.no_op()
	}
}

func (c *Chip8) executeInstructionType8(vx, vy uint16) {
	fmt.Printf("executeInstructionType8: %X\n ", c.opcode&0xFF)
	switch c.opcode & 0xFF {
	case 0x00:
		//Set Vx = Vy
		c.registers.copyVRegister(vx, vy)
	case 0x01:
		//Set Vx = Vx OR Vy
		c.registers.orVRegister(vx, vy)
	case 0x02:
		//Set Vx = Vx AND Vy
		c.registers.andVRegister(vx, vy)
	case 0x03:
		//Set Vx = Vx XOR Vy
		c.registers.xorVRegister(vx, vy)
	case 0x04:
		//Set Vx = Vx + Vy, set VF = carry
		sum := c.registers.sumVRegister(vx, vy)
		if sum > 255 {
			c.registers.setVRegister(VF, 1)
		} else {
			c.registers.clearVRegister(VF)
		}
		c.registers.setVRegister(vx, uint8(sum&0xFF))
	case 0x05:
		//Set Vx = Vx - Vy, set VF = NOT borrow
		if c.registers.isGreaterThan(vx, vy) {
			c.registers.setVRegister(VF, 1)
		} else {
			c.registers.clearVRegister(VF)
		}
		c.registers.decrementVRegister(vx, c.registers.getVRegisterVal(vy))
	case 0x06:
		//Set Vx = Vx SHR 1
		maskedVxRegisterValue := c.registers.getVRegisterVal(vx) & 0x1
		c.registers.setVRegister(VF, maskedVxRegisterValue)
		c.registers.shiftRightVRegister(vx)
	case 0x07:
		//Set Vx = Vy - Vx, set VF = NOT borrow
		if c.registers.isGreaterThan(vy, vx) {
			c.registers.setVRegister(VF, 1)
		} else {
			c.registers.clearVRegister(VF)
		}
		diff := c.registers.diffVRegister(vy, vx)
		c.registers.setVRegister(vx, diff)
	case 0x0E:
		//Set Vx = Vx SHL 1
		c.registers.setVRegister(VF, c.registers.getVRegisterVal(vx)>>7)
		c.registers.shiftLeftVRegister(vx)
	default:
		fmt.Printf("executeInstructionType8: %X\n ", c.opcode&0xFF)
		c.no_op()
	}
}

func (c *Chip8) executeInstructionTypeE(vx uint16) {
	switch c.opcode & 0xFF {
	case 0x9E:
		//Skip next instruction if key with the value of Vx is pressed
		key := c.registers.getVRegisterVal(vx)
		if c.keys[key] != 0 {
			c.stack.incrementProgramCounter()
		}
	case 0xA1:
		//Skip next instruction if key with the value of Vx is not pressed
		key := c.registers.getVRegisterVal(vx)
		if c.keys[key] == 0 {
			c.stack.incrementProgramCounter()
		}
	default:
		c.no_op()
	}
}

func (c *Chip8) executeInstructionTypeF(vx uint16) {
	switch c.opcode & 0xFF {
	case 0x07:
		//Set Vx = delay timer value
		c.registers.setVRegister(vx, c.registers.getDelay())
	case 0x0A:
		c.stack.decrementProgramCounter()
	case 0x15:
		//Set delay timer = Vx
		c.registers.setDelay(c.registers.getVRegisterVal(vx))
	case 0x18:
		//Set sound timer = Vx
		c.registers.setSound(c.registers.getVRegisterVal(vx))
	case 0x1E:
		//Set I = I + Vx
		i := c.registers.getIRegister()
		c.registers.setIRegister(i + uint16(c.registers.getVRegisterVal(vx)))
	case 0x29:
		//Set I = location of sprite for digit Vx
		addr := FontsetStartAddr + c.registers.getVRegisterVal(vx)
		c.registers.setIRegister(uint16(addr) * 5)
	case 0x33:
		//Store BCD representation of Vx in memory locations I, I+1, and I+2
		val := c.registers.getVRegisterVal(vx)
		iReg := c.registers.getIRegister()
		c.memory.write(iReg+2, val%10)
		val /= 10
		c.memory.write(iReg+1, val%10)
		val /= 10
		c.memory.write(iReg, val%10)
	case 0x55:
		for i := uint16(0); i <= vx; i++ {
			iReg := c.registers.getIRegister()
			c.memory.write(iReg+i, c.registers.getVRegisterVal(i))
		}
	case 0x65:
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
