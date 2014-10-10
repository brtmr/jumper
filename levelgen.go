package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
)

const N = 10
const AMPLITUDE = 10
const MIN_DILATION = 1
const MAX_DILATION = 10
const MAX_PHASE = math.Pi
const MAPWIDTH = 400
const MAPHEIGHT = 100

func genLevel(tc *TileCreator, cam *Camera, renderer *sdl.Renderer,
	spr *SpriteManager) Level {
	_ = fmt.Printf
	//create the surface boundaries:
	var boundaries [MAPWIDTH]int
	allparams := genParams()
	for i := 0; i < MAPWIDTH; i++ {
		boundaries[i] = 20
		for _, params := range allparams {
			boundaries[i] += int(float64(params.amplitude) *
				math.Sin(0.01*(math.Log(float64(params.stretch)))*(float64(i)+params.phase)))
		}
	}
	var tiles [MAPHEIGHT][MAPWIDTH]Tile
	for i := 0; i < MAPHEIGHT; i++ {
		for j := 0; j < MAPWIDTH; j++ {
			if boundaries[j] > i {
				tiles[i][j] = tc.TileByName("empty")
			} else {
				tiles[i][j] = tc.TileById(1 + rand.Int()%4)
			}
		}
	}
	return Level{tiles, cam, renderer, spr, MAPWIDTH, MAPHEIGHT}
}

func genParams() [N]Params {
	var result [N]Params
	for i := 0; i < N; i++ {
		amplitude := (rand.Int() % (2 * AMPLITUDE)) - AMPLITUDE
		phase := ((2 * rand.Float64()) - 1.0) * MAX_PHASE
		stretch := MIN_DILATION + (rand.Int() % (MAX_DILATION - MIN_DILATION))
		result[i] = Params{amplitude, stretch, phase}
	}
	return result
}

type Params struct {
	amplitude int
	stretch   int
	phase     float64
}
