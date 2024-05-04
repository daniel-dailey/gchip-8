package chip8

import (
	"fmt"
	"math/rand"
	"time"
)

func (c *Chip8) rand() uint8 {
	return uint8(rand.New(rand.NewSource(time.Now().UnixMilli())).Intn(255))
}

func (c *Chip8) fetchOpcode() Opcode {
	addrVal := c.memory.read(c.stack.getProgramCounter())
	addrValInc := c.memory.read(c.stack.getProgramCounter() + 1)
	c.stack.incrementProgramCounter()
	return Opcode(uint16(addrVal)<<8 | uint16(addrValInc))
}

func (c *Chip8) cycle() {
	c.opcode = c.fetchOpcode()
	c.executeCurrentInstruction()
	c.registers.decrementDelay()
	c.registers.decrementSound()
	c.ticks++
	c.logger.Info().Msg(fmt.Sprintf("I: %d, sp: %X pc: %d, op: %X, shift: %X, vx: %X, nn: %X, tick: %d", c.registers.getIRegister(), c.stack.getStackPointer(), c.stack.getProgramCounter(), c.opcode, c.opcode.opDecode(), c.opcode.vx(), c.opcode.nn(), c.ticks))
}

func (c *Chip8) Cycle() {
	c.cycle()
}

func (c *Chip8) GetDisplayBuffer() [2048]uint32 {
	return c.frameBuf.getFrameBuffer()
}

func (c *Chip8) Load(rom []byte) {
	c.memory.loadROM(rom)
}

func (c *Chip8) GetKeys() *[16]uint8 {
	return &[16]uint8{}
}
