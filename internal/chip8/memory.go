package chip8

// loadROM loads the ROM into memory starting at 0x200,
// the start address all chip8 programs are loaded into
func (m *Memory) loadROM(raw []byte) {
	for i, b := range raw {
		m.buf[StartAddr+i] = b
	}
}

// loadFontset loads the fontset into memory starting at 0x50
func (m *Memory) loadFontset() {
	for i, b := range fontset {
		m.buf[FontsetStartAddr+i] = b
	}
}

// write writes a byte to the memory buffer at the given address
func (m *Memory) write(address uint16, value uint8) {
	m.buf[address] = value
}

// read reads a byte from the memory buffer at the given address
func (m *Memory) read(address uint16) uint8 {
	return m.buf[address]
}

// InitMemory initializes the memory buffer and loads the fontset
func InitMemory() *Memory {
	mem := &Memory{
		buf: [MemoryBufferSize]uint8{},
	}
	mem.loadFontset()
	return mem
}
