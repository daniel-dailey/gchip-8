package main

import (
	"flag"
	"gochip8/internal/chip8"
	"gochip8/internal/clog"
	"gochip8/internal/ui"
	"gochip8/roms"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func getRomBytes(romLocation string) []byte {
	f, err := os.ReadFile(romLocation)
	if err != nil {
		return nil
	}
	return f
}

func main() {

	romLocation := flag.String("rom", "", "Location of the ROM file")
	testRom := flag.Bool("test", false, "Use the test ROM")
	debug := flag.Bool("debug", false, "Debug mode")

	flag.Parse()
	var rom []byte

	if *romLocation == "" {
		if !*testRom {
			panic("No ROM file specified, use -rom or -test")
		}
		rom = roms.TestRomRaw
	} else {
		rom = getRomBytes(*romLocation)
	}
	if len(rom) == 0 {
		panic("Failed to load ROM")
	}
	block := make(chan bool)
	c8 := chip8.Init()
	c8.Load(rom)
	ui, err := ui.Init()
	if err != nil {
		panic(err)
	}
	if ui == nil {
		panic("ui is nil")
	}
	logger := clog.NewLog(0, "MAIN", "c8-emulator")
	logger.Info().Msg("Starting...")
	defer sdl.Quit()
	defer ui.GetRenderer().Destroy()
	defer ui.GetWindow().Destroy()

	if *debug {
		roms.DumpRomInfo(rom)
		go func() {
			for range block {
				c8.Cycle()
				ui.Update(c8.GetDisplayBuffer())
			}
		}()
	}

	for {
		if running := ui.ProcessInput(c8.GetKeys(), block); !running {
			break
		}
		if *debug {
			continue
		}
		ui.Update(c8.GetDisplayBuffer())
		c8.Cycle()
		sdl.Delay(10)
	}
	logger.Info().Msg("Exiting...")
}
