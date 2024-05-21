package chip8

// getFrameBuffer returns the frame buffer
func (fb *FrameBuf) getFrameBuffer() [2048]uint32 {
	return fb.buf
}

// clear clears the frame buffer
func (fb *FrameBuf) clear() {
	fb.buf = [2048]uint32{}
}

// setPixel sets the pixel at the given display index
func (fb *FrameBuf) setPixel(displayIndex uint16) {
	fb.buf[displayIndex] ^= 0xFFFFFFFF
}

// initFrameBuf initializes the frame buffer
func InitFrameBuf() *FrameBuf {
	return &FrameBuf{
		buf: [2048]uint32{},
	}
}
