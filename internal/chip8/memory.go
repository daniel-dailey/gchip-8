package chip8

func (m *Memory) loadROM(raw []byte) {
	// Load the ROM into memory starting at 0x200
	for i, b := range raw {
		m.buf[StartAddr+i] = b
	}
}

func (m *Memory) getROM() []uint8 {
	return m.buf[StartAddr:]
}

func (m *Memory) loadFontset() {
	for i, b := range fontset {
		m.buf[FontsetStartAddr+i] = b
	}
}

func (m *Memory) write(address uint16, value uint8) {
	m.buf[address] = value
}

func (m *Memory) read(address uint16) uint8 {
	return m.buf[address]
}

func InitMemory() *Memory {
	mem := &Memory{
		buf: [MemoryBufferSize]uint8{},
	}
	mem.loadFontset()
	return mem
}
