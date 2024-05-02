package chip8

import (
	"gochip8/internal/clog"
	"os"
)

func (c *Chip8) loadFontset() {
	for i, b := range fontset {
		c.memory[FontsetStartAddr+i] = b
	}
}

func (c *Chip8) LoadROM(fn string) {
	// Load the ROM into memory starting at 0x200
	f, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	for i, b := range f {
		c.romSize++
		c.memory[StartAddr+i] = b
	}
}

func (c *Chip8) GetROM() []byte {
	return c.memory[StartAddr : StartAddr+c.romSize]
}

func (c *Chip8) initOpTables() {
	c.funcTable = map[uint16]func(){
		0x0000: c.Table0,
		0x0001: c.op1nnn,
		0x0002: c.op2nnn,
		0x0003: c.op3xnn,
		0x0004: c.op4xnn,
		0x0005: c.op5xy0,
		0x0006: c.op6xnn,
		0x0007: c.op7xnn,
		0x0008: c.Table8,
		0x0009: c.op9xy0,
		0x000A: c.opAnnn,
		0x000B: c.opBnnn,
		0x000C: c.opCxnn,
		0x000D: c.opDxyn,
		0x000E: c.TableE,
		0x000F: c.TableF,
	}

	c.funcTable0 = make(map[uint16]func())
	c.funcTable8 = make(map[uint16]func())
	c.funcTableE = make(map[uint16]func())
	c.funcTableF = make(map[uint16]func())

	for i := 0; i <= 0x0E; i++ {
		c.funcTable0[uint16(i)] = c.null_op
		c.funcTable8[uint16(i)] = c.null_op
		c.funcTableE[uint16(i)] = c.null_op
	}

	c.funcTable0[0x0000] = c.op00E0
	c.funcTable0[0x000E] = c.op00EE

	c.funcTable8[0x0000] = c.op8xy0
	c.funcTable8[0x0001] = c.op8xy1
	c.funcTable8[0x0002] = c.op8xy2
	c.funcTable8[0x0003] = c.op8xy3
	c.funcTable8[0x0004] = c.op8xy4
	c.funcTable8[0x0005] = c.op8xy5
	c.funcTable8[0x0006] = c.op8xy6
	c.funcTable8[0x0007] = c.op8xy7
	c.funcTable8[0x000E] = c.op8xyE

	c.funcTableE[0x0001] = c.opEx9E
	c.funcTableE[0x000E] = c.opExA1

	for i := 0; i <= 0x65; i++ {
		c.funcTableF[uint16(i)] = c.null_op
	}

	c.funcTableF[0x0007] = c.opFx07
	c.funcTableF[0x000A] = c.opFx0A
	c.funcTableF[0x0015] = c.opFx15
	c.funcTableF[0x0018] = c.opFx18
	c.funcTableF[0x001E] = c.opFx1E
	c.funcTableF[0x0029] = c.opFx29
	c.funcTableF[0x0033] = c.opFx33
	c.funcTableF[0x0055] = c.opFx55
	c.funcTableF[0x0065] = c.opFx65
}

func Init() *Chip8 {
	c := &Chip8{
		pc:        StartAddr,
		memory:    [MemoryBufferSize]byte{},
		display:   [VideoBufferWidth * VideoBufferHeight]uint32{},
		stack:     [16]uint16{},
		keys:      [16]byte{},
		registers: [16]byte{},
		logger:    clog.NewLog(int(clog.LogLevelInfo), "Chip8", "c8-cpu"),
	}
	c.loadFontset()
	c.initOpTables()
	return c
}
