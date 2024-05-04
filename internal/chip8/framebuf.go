package chip8

func (fb *FrameBuf) getFrameBuffer() [2048]uint32 {
	return fb.buf
}

func (fb *FrameBuf) clear() {
	fb.buf = [2048]uint32{}
}

func (fb *FrameBuf) setPixel(displayIndex uint16) {
	fb.buf[displayIndex] ^= 0xFFFFFFFF
}

func InitFrameBuf() *FrameBuf {
	return &FrameBuf{
		buf: [2048]uint32{},
	}
}
