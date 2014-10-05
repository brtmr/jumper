package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
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
	Spr           SpriteManager
	Lvl           Level
	Ply           Player
	renderer      *sdl.Renderer
	spaceReleased bool
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

	gameOver := false

	dt := 0.03
	accumulator := 0.0

	currentTime := float64(sdl.GetTicks()) / 1000.0

	frameTime := 0.0

	for {
		/*
			begin mainloop
			implemented as described in
			http://gafferongames.com/game-physics/fix-your-timestep/
		*/
		newTime := float64(sdl.GetTicks()) / 1000.0
		frameTime = newTime - currentTime
		if frameTime > 0.25 {
			frameTime = 0.25
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

		alpha := accumulator / dt
		gd.Interpolate(alpha)

		gd.Draw()
		/*
			end mainloop
		*/

		if gameOver {
			break
		}

		var event sdl.Event
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				_ = t
				gameOver = true
				break
			}
		}
	}
	gd.Spr.TearDown()
	sdl.Quit()
}

func Game_Init(renderer *sdl.Renderer) GameData {
	spr := Init_from_json(GetDataPath()+"sprites.json", renderer)
	cam := Camera{0, 0}
	lvl := DummyLevel(spr, renderer, &cam)
	ply := Init_player(spr, renderer, &lvl, &cam)
	return GameData{spr, lvl, ply, renderer, true}
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
	gd.handleKeys()
	gd.Ply.Update()
}

func (gd *GameData) Interpolate(alpha float64) {
	gd.Ply.Interpolate(alpha)
	gd.Ply.SetCamera()
}

func (gd *GameData) handleKeys() {
	keystate := sdl.GetKeyboardState()
	if keystate[sdl.GetScancodeFromName("LEFT")] == 1 {
		gd.Ply.SetDirection(DIRECTION_LEFT)
	}
	if keystate[sdl.GetScancodeFromName("RIGHT")] == 1 {
		gd.Ply.SetDirection(DIRECTION_RIGHT)
	}
	if keystate[sdl.GetScancodeFromName("RIGHT")]+
		keystate[sdl.GetScancodeFromName("LEFT")] == 0 {
		gd.Ply.SetDirection(STOP)
	}
	if keystate[sdl.GetScancodeFromName("SPACE")] == 0 {
		gd.spaceReleased = true
	}
	if keystate[sdl.GetScancodeFromName("SPACE")] == 1 {
		if gd.spaceReleased {
			gd.Ply.Jump()
		}
		gd.spaceReleased = false
	}

}

func GetDataPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/rtmb/jumper/data/"
}
