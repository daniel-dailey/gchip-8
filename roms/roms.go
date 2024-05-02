package roms

import (
	"encoding/binary"
	"fmt"
)

func DumpRomInfo(buf []byte) {
	fmt.Println("Dumping ROM info...")
	for i := 0; i < len(buf); i += 2 {
		sl := buf[i : i+2]
		fmt.Printf("0x%X, shift:  %x", sl, binary.BigEndian.Uint16(sl)>>12)
		fmt.Println()
	}
}
