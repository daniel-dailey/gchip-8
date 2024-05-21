package chip8

import (
	"fmt"
	"math/rand"
	"time"
)

// random number generator for the chip8
func (c *Chip8) rand() uint8 {
	return uint8(rand.New(rand.NewSource(time.Now().UnixMilli())).Intn(255))
}

// grabs opcode from combining the current and next memory addresses
func (c *Chip8) fetchOpcode() {
	addrVal := c.memory.read(c.stack.getProgramCounter())
	addrValInc := c.memory.read(c.stack.getProgramCounter() + 1)
	c.opcode = Opcode((uint16(addrVal)<<8 | uint16(addrValInc)))
}

// cycles through the chip8
func (c *Chip8) cycle() {
	c.fetchOpcode()
	c.stack.incrementProgramCounter()
	c.executeCurrentInstruction()
	c.registers.decrementDelay()
	c.registers.decrementSound()
	c.ticks++
	c.logger.Info().Msg(fmt.Sprintf("Frame end: I: %d, sp: %X pc: %d, op: %X, shift: %X, vx: %X, vy: %d, nn: %X, tick: %d", c.registers.getIRegister(), c.stack.getStackPointer(), c.stack.getProgramCounter(), c.opcode, c.opcode.opDecode(), c.opcode.vx(), c.opcode.vy(), c.opcode.nn(), c.ticks))
}

// public method for external pkg to call chip8 cycle
func (c *Chip8) Cycle() {
	c.cycle()
}

// public method for external pkg to get display buffer
func (c *Chip8) GetDisplayBuffer() [2048]uint32 {
	return c.frameBuf.getFrameBuffer()
}

// public method for external pkg to load ROM into memory
func (c *Chip8) Load(rom []byte) {
	c.memory.loadROM(rom)
}

// public method for external pkg to get keys
func (c *Chip8) GetKeys() *[16]uint8 {
	return &[16]uint8{}
}
