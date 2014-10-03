package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Level struct {
	Tiles    [][]Tile
	Cam      Camera
	Renderer *sdl.Renderer
}

const Tile_size = 32

type Tile struct {
	Sprite Sprite
	Health int
	Solid  bool
}

func DummyLevel(spr SpriteManager, renderer *sdl.Renderer) Level {
	tiles := make([][]Tile, 200, 200)
	for i := 0; i < 200; i++ {
		r := make([]Tile, 200, 200)
		for j := 0; j < 200; j++ {
			var tl Tile
			var solid bool
			if i < 15 {
				solid = false
			} else {
				solid = true
			}
			if i%2 == 0 {
				tl = Tile{spr.GetSprite("tile_stone"), 100, solid}
			} else {
				tl = Tile{spr.GetSprite("tile_stone_skull"), 100, solid}
			}
			r[j] = tl
		}
		tiles[i] = r
	}
	return Level{tiles[:][:], Camera{0, 0}, renderer}
}

type Camera struct {
	X int32
	Y int32
}

func (lvl Level) Draw() {
	for y, arr := range lvl.Tiles {
		for x, tl := range arr {
			if tl.Solid {
				xpos := int32(x)*Tile_size + lvl.Cam.X
				ypos := int32(y)*Tile_size + lvl.Cam.Y
				dstRec := sdl.Rect{xpos, ypos, Tile_size, Tile_size}
				lvl.Renderer.Copy(tl.Sprite.Texture, tl.Sprite.Rect, &dstRec)
			}
		}
	}
}

func (lvl Level) IsSolid(i, j int32) bool {
	_ = fmt.Println
	return lvl.Tiles[i][j].Solid
}
