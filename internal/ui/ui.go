package ui

import (
	"image/color"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type UI struct {
	window              *sdl.Window
	renderer            *sdl.Renderer
	surface             *sdl.Surface
	lastUpdateCycleTime []int64
}

func (ui *UI) Clear() {
	ui.renderer.SetDrawColor(0, 0, 0, 255)
	ui.renderer.Clear()
}

func (ui *UI) GetWindow() *sdl.Window {
	return ui.window
}

func (ui *UI) GetRenderer() *sdl.Renderer {
	return ui.renderer
}

func (ui *UI) GetLastUpdateCycleTime() []int64 {
	return ui.lastUpdateCycleTime
}

func Init() (*UI, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow("Chip8", 0, 0, 640, 320, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		return nil, err
	}
	surface, err := sdl.CreateRGBSurface(0, 64, 32, 32, 0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
	if err != nil {
		return nil, err
	}
	return &UI{window, renderer, surface, []int64{}}, nil
}

func (ui *UI) makeColor(incomingColor uint32) color.Color {
	r := incomingColor >> 24 & 0xFF
	g := incomingColor >> 16 & 0xFF
	b := incomingColor >> 8 & 0xFF
	a := incomingColor & 0xFF
	return color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func (ui *UI) putPixel(x, y int, color uint32) {
	ui.surface.Lock()
	ui.surface.Set(x, y, ui.makeColor(color))
	ui.surface.Unlock()
}

func (ui *UI) Update(buf [2048]uint32) {
	t1 := time.Now().UnixMilli()
	for i := 0; i < 2048; i++ {
		x := i % 64
		y := i / 64
		ui.putPixel(x, y, buf[i])
	}
	tex, err := ui.renderer.CreateTextureFromSurface(ui.surface)
	if err != nil {
		log.Println("Error creating texture from surface: ", err)
		return
	}
	ui.renderer.Clear()
	ui.renderer.Copy(tex, nil, nil)
	ui.renderer.Present()
	t2 := time.Now().UnixMilli()
	ui.lastUpdateCycleTime = append(ui.lastUpdateCycleTime, t2-t1)
}

func (ui *UI) ProcessInput(keys *[16]byte) bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			typ := event.Type
			switch typ {
			case sdl.KEYDOWN:
				key := event.Keysym.Sym
				switch key {
				case sdl.K_ESCAPE:
					return false
				case sdl.K_x:
					keys[0] = 1
				case sdl.K_1:
					keys[1] = 1
				case sdl.K_2:
					keys[2] = 1
				case sdl.K_3:
					keys[3] = 1
				case sdl.K_q:
					keys[4] = 1
				case sdl.K_w:
					keys[5] = 1
				case sdl.K_e:
					keys[6] = 1
				case sdl.K_a:
					keys[7] = 1
				case sdl.K_s:
					keys[8] = 1
				case sdl.K_d:
					keys[9] = 1
				case sdl.K_z:
					keys[0xA] = 1
				case sdl.K_c:
					keys[0xB] = 1
				case sdl.K_4:
					keys[0xC] = 1
				case sdl.K_r:
					keys[0xD] = 1
				case sdl.K_f:
					keys[0xE] = 1
				case sdl.K_v:
					keys[0xF] = 1
				}
			case sdl.KEYUP:
				key := event.Keysym.Sym
				switch key {
				case sdl.K_x:
					keys[0] = 0
				case sdl.K_1:
					keys[1] = 0
				case sdl.K_2:
					keys[2] = 0
				case sdl.K_3:
					keys[3] = 0
				case sdl.K_q:
					keys[4] = 0
				case sdl.K_w:
					keys[5] = 0
				case sdl.K_e:
					keys[6] = 0
				case sdl.K_a:
					keys[7] = 0
				case sdl.K_s:
					keys[8] = 0
				case sdl.K_d:
					keys[9] = 0
				case sdl.K_z:
					keys[0xA] = 0
				case sdl.K_c:
					keys[0xB] = 0
				case sdl.K_4:
					keys[0xC] = 0
				case sdl.K_r:
					keys[0xD] = 0
				case sdl.K_f:
					keys[0xE] = 0
				case sdl.K_v:
					keys[0xF] = 0
				}
			}
		}
	}
	return true
}
