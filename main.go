package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"os"
)

const SCALE = 3
const DIRECTION_RIGHT = 0
const DIRECTION_LEFT = 1
const STOP = -1
const RUNSPEED = SCALE * 5
const JUMPSPEED = -(SCALE * 8)
const GRAVITY = SCALE * 0.6
const TOPSPEED = SCALE * 8
const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600
const HALF_SCREEN_WIDTH = SCREEN_WIDTH / 2
const HALF_SCREEN_HEIGHT = SCREEN_HEIGHT / 2

type GameData struct {
	Spr      SpriteManager
	Lvl      Level
	Ply      Player
	renderer *sdl.Renderer
	gameOver bool
}

type GameObject interface {
	Update()
	Interp()
	Draw()
}

type Drawer interface {
	Draw()
}

type Updater interface {
	Update()
}

func main() {

	if 0 != sdl.Init(sdl.INIT_EVERYTHING) {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", sdl.GetError())
		os.Exit(2)
	}
	ttf.Init()
	/*
		        window := sdl.CreateWindow("goplot", sdl.WINDOWPOS_UNDEFINED,
		            sdl.WINDOWPOS_UNDEFINED,
					1600, 900, sdl.WINDOW_SHOWN|sdl.WINDOW_FULLSCREEN)
	*/
	window := sdl.CreateWindow("jumper", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(2)
	}
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n",
			sdl.GetError())
		os.Exit(2)
	}

	gd := Game_Init(renderer)

	currentTime := sdl.GetTicks()

	var frameTime uint32
	var accumulator uint32
	var dt uint32
	var alpha float64

	dt = 33 //time for a single logic frame.

	var rendertime uint32

	for {
		/*
			begin mainloop
			implemented as described in
			http://gafferongames.com/game-physics/fix-your-timestep/
		*/
		newTime := sdl.GetTicks()
		frameTime = newTime - currentTime
		if frameTime > 250 {
			frameTime = 250
		}
		currentTime = newTime
		accumulator += frameTime
		for {
			if accumulator < dt {
				break
			}
			gd.Update()
			accumulator -= dt
		}

		alpha = float64(accumulator) / float64(dt)
		gd.Interpolate(alpha)

		rendertime = sdl.GetTicks()
		gd.Draw()
		rendertime = sdl.GetTicks() - rendertime
		fmt.Printf("Rendertime: %d \n", rendertime)

		/*
			end mainloop
		*/

		if gd.gameOver {
			break
		}
	}
	gd.Spr.TearDown()
	ttf.Quit()
	sdl.Quit()
}

func Game_Init(renderer *sdl.Renderer) GameData {
	spr := Init_from_json(GetDataPath()+"sprites.json", renderer)
	cam := Camera{0, 0}
	lvl := DummyLevel(spr, renderer, &cam)
	ply := Init_player(spr, renderer, &lvl, &cam)
	return GameData{spr, lvl, ply, renderer, false}
}

func (gd *GameData) Draw() {
	gd.renderer.Clear()
	//draw the sky
	sky := gd.Spr.GetSprite("sky")
	gd.renderer.Copy(sky.Texture, sky.Rect, nil)

	gd.Lvl.Draw()
	gd.Ply.Draw()
	gd.renderer.Present()
}

func (gd *GameData) Update() {
	gd.handleEvents()
	gd.Ply.Update()
}

func (gd *GameData) Interpolate(alpha float64) {
	gd.Ply.Interpolate(alpha)
	gd.Ply.SetCamera()
}

func (gd *GameData) handleKeyDown(sym sdl.Keysym) {
	switch sym.Scancode {
	case sdl.GetScancodeFromName("SPACE"):
		gd.Ply.Jump()
		break
	case sdl.GetScancodeFromName("RIGHT"):
		gd.Ply.SetDirection(DIRECTION_RIGHT)
		break
	case sdl.GetScancodeFromName("LEFT"):
		gd.Ply.SetDirection(DIRECTION_LEFT)
		break
	}
}

func (gd *GameData) handleKeyUp(sym sdl.Keysym) {
	switch sym.Scancode {
	case sdl.GetScancodeFromName("SPACE"):
		break
	case sdl.GetScancodeFromName("LEFT"):
		gd.Ply.SetDirection(STOP)
		break
	case sdl.GetScancodeFromName("RIGHT"):
		gd.Ply.SetDirection(STOP)
		break
	}
}

func (gd *GameData) handleEvents() {
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			_ = t
			gd.gameOver = true
			break
		case *sdl.KeyDownEvent:
			gd.handleKeyDown(t.Keysym)
			break
		case *sdl.KeyUpEvent:
			gd.handleKeyUp(t.Keysym)
			break
		}
	}
	kbstate := sdl.GetKeyboardState()
	if kbstate[sdl.GetScancodeFromName("LEFT")] == 1 {
		gd.Ply.SetDirection(DIRECTION_LEFT)
	}
	if kbstate[sdl.GetScancodeFromName("RIGHT")] == 1 {
		gd.Ply.SetDirection(DIRECTION_RIGHT)
	}
}

func GetDataPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/rtmb/jumper/data/"
}
