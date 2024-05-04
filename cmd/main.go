package main

import (
	"gochip8/internal/chip8"
	"gochip8/internal/clog"
	"gochip8/internal/ui"
	"gochip8/roms"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func getRomBytes() []byte {
	f, err := os.ReadFile("../roms/pong1p.ch8")
	if err != nil {
		return nil
	}
	return f
}

func main() {
	debug := false
	block := make(chan bool)
	c8 := chip8.Init()
	c8.Load(getRomBytes())
	roms.DumpRomInfo(getRomBytes())
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

	if debug {
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
		if debug {
			continue
		}
		ui.Update(c8.GetDisplayBuffer())
		c8.Cycle()
		sdl.Delay(10)
	}
	logger.Info().Msg("Exiting...")
}
