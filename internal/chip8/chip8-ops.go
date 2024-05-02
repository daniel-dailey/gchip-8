package chip8

func (c *Chip8) op00E0() {
	c.display = [2048]uint32{}
}

func (c *Chip8) op00EE() {
	c.pc = c.stack[c.sp-1]
}

func (c *Chip8) op1nnn() {
	c.pc = c.opcode & 0x0FFF
}

func (c *Chip8) op2nnn() {
	addr := c.opcode & 0x0FFF
	c.stack[c.sp] = c.pc
	c.sp++
	c.pc = addr
}

func (c *Chip8) op3xkk() {
	if c.registers[c.vx()] == c.kk() {
		c.pc += 2
	}
}

func (c *Chip8) op4xkk() {
	if c.registers[c.vx()] != c.kk() {
		c.pc += 2
	}
}

func (c *Chip8) op5xy0() {
	if c.registers[c.vx()] == c.registers[c.vy()] {
		c.pc += 2
	}
}

func (c *Chip8) op6xkk() {
	c.registers[c.vx()] = c.kk()
}

func (c *Chip8) op7xkk() {
	c.registers[c.vx()] += c.kk()
}

func (c *Chip8) op8xy0() {
	c.registers[c.vx()] = c.registers[c.vy()]
}

func (c *Chip8) op8xy1() {
	c.registers[c.vx()] |= c.registers[c.vy()]
}

func (c *Chip8) op8xy2() {
	c.registers[c.vx()] &= c.registers[c.vy()]
}

func (c *Chip8) op8xy3() {
	c.registers[c.vx()] ^= c.registers[c.vy()]
}

func (c *Chip8) op8xy4() {
	sum := c.registers[c.vx()] + c.registers[c.vx()]
	if sum > 255 {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}
	c.registers[c.vx()] = byte(sum) & 0xFF
}

func (c *Chip8) op8xy5() {
	if c.registers[c.vx()] > c.registers[c.vy()] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}
	c.registers[c.vx()] -= c.registers[c.vy()]
}

func (c *Chip8) op8xy6() {
	c.registers[0xF] = c.registers[c.vx()] & 0x1
	c.registers[c.vx()] >>= 1
}

func (c *Chip8) op8xy7() {
	if c.registers[c.vy()] > c.registers[c.vx()] {
		c.registers[0xF] = 1
	} else {
		c.registers[0xF] = 0
	}
	c.registers[c.vx()] = c.registers[c.vy()] - c.registers[c.vx()]
}

func (c *Chip8) op8xyE() {
	c.registers[0xF] = c.registers[c.vx()] >> 7
	c.registers[c.vx()] <<= 1
}

func (c *Chip8) op9xy0() {
	if c.registers[c.vx()] != c.registers[c.vy()] {
		c.pc += 2
	}
}

func (c *Chip8) opAnnn() {
	c.i = c.opcode & 0x0FFF
}

func (c *Chip8) opBnnn() {
	addr := c.opcode & 0x0FFF
	c.pc = addr + uint16(c.registers[0])
}

func (c *Chip8) opCxkk() {
	c.registers[c.vx()] = c.rand() & c.kk()
}

func (c *Chip8) opDxyn() {
	height := c.opcode & 0x000F
	xPos := c.registers[c.vx()] % VideoBufferWidth
	yPos := c.registers[c.vy()] % VideoBufferHeight
	c.registers[0xF] = 0
	for row := uint16(0); row < height; row++ {
		sb := c.memory[c.i+row]
		for col := uint16(0); col < 8; col++ {
			sp := sb & (0x80 >> col)
			displayIndex := (uint16(yPos)+row)*VideoBufferWidth + (uint16(xPos) + col)
			if displayIndex >= 2048 {
				continue
			}
			screenPixel := &c.display[displayIndex]
			if sp != 0 {
				if *screenPixel == 0xFFFFFFFF {
					c.registers[0xF] = 1
				}
				*screenPixel ^= 0xFFFFFFFF
			}
		}
	}
}

func (c *Chip8) opEx9E() {
	if c.keys[c.registers[c.vx()]] != 0 {
		c.pc += 2
	}
}

func (c *Chip8) opExA1() {
	if c.keys[c.registers[c.vx()]] == 0 {
		c.pc += 2
	}
}

func (c *Chip8) opFx07() {
	c.registers[c.vx()] = c.delay
}

func (c *Chip8) opFx0A() {
	Vx := c.vx()
	if c.keys[0] != 0 {
		c.registers[Vx] = 0
	} else if c.keys[1] != 0 {
		c.registers[Vx] = 1
	} else if c.keys[2] != 0 {
		c.registers[Vx] = 2
	} else if c.keys[3] != 0 {
		c.registers[Vx] = 3
	} else if c.keys[4] != 0 {
		c.registers[Vx] = 4
	} else if c.keys[5] != 0 {
		c.registers[Vx] = 5
	} else if c.keys[6] != 0 {
		c.registers[Vx] = 6
	} else if c.keys[7] != 0 {
		c.registers[Vx] = 7
	} else if c.keys[8] != 0 {
		c.registers[Vx] = 8
	} else if c.keys[9] != 0 {
		c.registers[Vx] = 9
	} else if c.keys[10] != 0 {
		c.registers[Vx] = 10
	} else if c.keys[11] != 0 {
		c.registers[Vx] = 11
	} else if c.keys[12] != 0 {
		c.registers[Vx] = 12
	} else if c.keys[13] != 0 {
		c.registers[Vx] = 13
	} else if c.keys[14] != 0 {
		c.registers[Vx] = 14
	} else if c.keys[15] != 0 {
		c.registers[Vx] = 15
	} else {
		c.pc -= 2
	}
}

func (c *Chip8) opFx15() {
	c.delay = c.registers[c.vx()]
}

func (c *Chip8) opFx18() {
	c.sound = c.registers[c.vx()]
}

func (c *Chip8) opFx1E() {
	c.i += uint16(c.registers[c.vx()])
}

func (c *Chip8) opFx29() {
	c.i = FontsetStartAddr + uint16(c.registers[c.vx()])*5
}

func (c *Chip8) opFx33() {
	val := c.registers[c.vx()]
	c.memory[c.i+2] = val % 10
	val /= 10
	c.memory[c.i+1] = val % 10
	val /= 10
	c.memory[c.i] = val % 10
}

func (c *Chip8) opFx55() {
	for i := uint16(0); i <= c.vx(); i++ {
		c.memory[c.i+i] = c.registers[i]
	}
}

func (c *Chip8) opFx65() {
	for i := uint16(0); i <= c.vx(); i++ {
		c.registers[i] = c.memory[c.i+i]
	}
}

func (c *Chip8) null_op() {
	// Do nothing
}

func (c *Chip8) Table0() {
	c.funcTable0[c.opcode&0x000F]()
}

func (c *Chip8) Table8() {
	c.funcTable8[c.opcode&0x000F]()
}

func (c *Chip8) TableE() {
	c.funcTableE[c.opcode&0x000F]()
}

func (c *Chip8) TableF() {
	c.funcTableF[c.opcode&0x00FF]()
}
