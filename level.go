package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Level struct {
	tiles    [][]Tile
	camera   *Camera
	Renderer *sdl.Renderer
	dimx     int
	dimy     int
}

const Tile_size = 16 * SCALE

type Tile struct {
	sprite Sprite
	health int
	solid  bool
}

func (t Tile) Sprite() Sprite {
	return t.sprite
}

func (t Tile) Solid() bool {
	return t.solid
}

func DummyLevel(spr SpriteManager, renderer *sdl.Renderer, cam *Camera) Level {
	lsize := 20
	tiles := make([][]Tile, lsize, lsize)
	for i := 0; i < lsize; i++ {
		r := make([]Tile, lsize, lsize)
		for j := 0; j < lsize; j++ {
			var tl Tile
			var solid bool
			if i < 9 {
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

	tiles[4][5] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[4][6] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[4][7] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[5][5] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[5][6] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[5][7] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[6][5] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[6][6] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[6][7] = Tile{spr.GetSprite("tile_stone"), 100, true}

	tiles[6][10] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[6][11] = Tile{spr.GetSprite("tile_stone"), 100, true}
	tiles[6][12] = Tile{spr.GetSprite("tile_stone"), 100, true}

	for i := 0; i < 10; i++ {
		tiles[i][19] = Tile{spr.GetSprite("tile_stone"), 100, true}
	}
	return Level{tiles[:][:], cam, renderer,
		lsize, lsize}
}

func (lvl Level) Draw() {
	for y, arr := range lvl.tiles {
		for x, tl := range arr {
			if tl.Solid() {
				xpos := int32(x)*Tile_size - lvl.camera.X()
				ypos := int32(y)*Tile_size - lvl.camera.Y()
				dstRec := sdl.Rect{xpos, ypos, Tile_size, Tile_size}
				lvl.Renderer.Copy(tl.Sprite().Texture, tl.Sprite().Rect,
					&dstRec)
			}
		}
	}
}

func (lvl Level) IsSolid(i, j int) bool {
	_ = fmt.Println
	if i >= lvl.dimy || j >= lvl.dimy {
		return false
	}
	return lvl.tiles[i][j].Solid()
}

func (lvl Level) DimX() int {
	return lvl.dimx * Tile_size
}

func (lvl Level) DimY() int {
	return lvl.dimy * Tile_size
}
