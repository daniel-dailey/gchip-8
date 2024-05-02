package chip8

import (
	"fmt"
	"math/rand"
	"time"
)

func (c *Chip8) rand() byte {
	return byte(rand.New(rand.NewSource(time.Now().UnixMilli())).Intn(255))
}

func (c *Chip8) vx() uint16 {
	return (c.opcode & 0x0F00) >> 8
}

func (c *Chip8) vy() uint16 {
	return (c.opcode & 0x00F0) >> 4
}

func (c *Chip8) kk() byte {
	return byte(c.opcode & 0x00FF)
}

func (c *Chip8) cycle() {
	c.opcode = uint16(c.memory[c.pc])<<8 | uint16(c.memory[c.pc+1])
	c.pc += 2
	c.funcTable[c.opcode>>12]()
	if c.delay > 0 {
		c.delay--
	}
	if c.sound > 0 {
		c.sound--
	}
	c.logger.Info().Msg(fmt.Sprintf("Index: %d, PC: %d, Opcode: %X, Opcode (shift): %X, vx: %x, kk: %x", c.i, c.pc, c.opcode, c.opcode>>12, c.vx(), c.kk()))
}

func (c *Chip8) Cycle() {
	t1 := time.Now().UnixMilli()
	c.cycle()
	t2 := time.Now().UnixMilli()
	c.cycleTimes = append(c.cycleTimes, t2-t1)
}

func (c *Chip8) GetDisplayBuffer() [2048]uint32 {
	return c.display
}

func (c *Chip8) GetCycleTimes() []int64 {
	return c.cycleTimes
}

func (c *Chip8) GetKeys() *[16]byte {
	return &c.registers
}
