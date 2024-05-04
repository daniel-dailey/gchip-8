package chip8

import "gochip8/internal/clog"

const (
	StartAddr        = 0x200
	FontsetStartAddr = 0x50
	VF               = 0xF

	VideoBufferWidth  = 64
	VideoBufferHeight = 32
	MemoryBufferSize  = 4096
)

var fontset = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type Registers struct {
	iRegister uint16
	vRegister [16]uint8
	delay     uint8
	sound     uint8
}

type Memory struct {
	buf [MemoryBufferSize]uint8
}

type Stack struct {
	stack          [16]uint16
	programCounter uint16
	stackPointer   uint16
}

type FrameBuf struct {
	buf [2048]uint32
}

type Opcode uint16

func (o Opcode) opDecode() uint16 {
	return uint16((o & 0xF000) >> 12)
}

func (o Opcode) vx() uint16 {
	return uint16(o&0x0F00) >> 8
}

func (o Opcode) vy() uint16 {
	return uint16((o & 0x00F0) >> 4)
}

func (o Opcode) n() uint8 {
	return uint8(o & 0x000F)
}

func (o Opcode) nn() uint8 {
	return uint8(o & 0x00FF)
}

func (o Opcode) nnn() uint16 {
	return uint16(o & 0x0FFF)
}

type Chip8 struct {
	registers *Registers
	memory    *Memory
	stack     *Stack
	frameBuf  *FrameBuf
	keys      [16]uint8
	opcode    Opcode

	//For testing
	logger *clog.Log
	ticks  int64
}

func Init() *Chip8 {
	c := &Chip8{
		registers: InitRegisters(),
		stack:     InitStack(),
		memory:    InitMemory(),
		frameBuf:  InitFrameBuf(),
		keys:      [16]uint8{},
		logger:    clog.NewLog(int(clog.LogLevelInfo), "Chip8", "c8-cpu"),
	}
	return c
}
