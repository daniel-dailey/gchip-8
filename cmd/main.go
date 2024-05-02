package main

import (
	"gochip8/internal/chip8"
	"gochip8/internal/clog"
	"gochip8/internal/ui"
	"gochip8/roms"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	c8 := chip8.Init()
	c8.LoadROM("../roms/maze.ch8")
	ui, err := ui.Init()
	if err != nil {
		panic(err)
	}
	if ui == nil {
		panic("ui is nil")
	}
	logger := clog.NewLog(0, "MAIN", "c8-emulator")
	logger.Info().Msg("Starting...")
	roms.DumpRomInfo(c8.GetROM())
	defer sdl.Quit()
	defer ui.GetRenderer().Destroy()
	defer ui.GetWindow().Destroy()
	for {
		if running := ui.ProcessInput(c8.GetKeys()); !running {
			break
		}
		c8.Cycle()
		ui.Update(c8.GetDisplayBuffer())
		sdl.Delay(1000)
	}
	logger.Info().Any(c8.GetCycleTimes()).Msg("Exiting...")
}
