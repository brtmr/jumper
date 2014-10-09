package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Level struct {
	tiles    [][]Tile
	camera   *Camera
	Renderer *sdl.Renderer
	spr      *SpriteManager
	//dimension in tiles
	dimx int
	dimy int
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

func DummyLevel(spr *SpriteManager, renderer *sdl.Renderer, cam *Camera) Level {
	lsize := 40
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
			tl = RandomTile(spr, solid)
			r[j] = tl
		}
		tiles[i] = r
	}

	tiles[4][5] = RandomTile(spr, true)
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
		tiles[i][lsize-1] = Tile{spr.GetSprite("tile_stone"), 100, true}
	}
	return Level{tiles[:][:], cam, renderer, spr, lsize, lsize}
}

func RandomTile(spr *SpriteManager, solid bool) Tile {
	var tl Tile
	ra := rand.Int() % 5
	switch ra {
	case 0:
		tl = Tile{spr.GetSprite("tile_stone"), 100, solid}
		break
	case 1:
		tl = Tile{spr.GetSprite("tile_stone"), 100, solid}
		break
	case 2:
		tl = Tile{spr.GetSprite("tile_stone_skull"), 100, solid}
		break
	case 3:
		tl = Tile{spr.GetSprite("tile_diamond"), 100, solid}
		break
	case 4:
		tl = Tile{spr.GetSprite("tile_gold"), 100, solid}
		break
	}
	return tl
}

func (lvl Level) Draw() {
	minX := (lvl.camera.X() / Tile_size)
	maxX := (lvl.camera.X() + SCREEN_WIDTH/Tile_size) + 1
	minY := (lvl.camera.Y() / Tile_size)
	maxY := (lvl.camera.Y() + SCREEN_HEIGHT/Tile_size) + 1

	BoundsInt32(0, int32(lvl.dimx), &minX)
	BoundsInt32(0, int32(lvl.dimx), &maxX)
	BoundsInt32(0, int32(lvl.dimy), &minY)
	BoundsInt32(0, int32(lvl.dimy), &maxY)

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			tl := lvl.tiles[y][x]
			xpos := int32(x)*Tile_size - lvl.camera.X()
			ypos := int32(y)*Tile_size - lvl.camera.Y()
			dstRec := sdl.Rect{xpos, ypos, Tile_size, Tile_size}
			if tl.Solid() {
				lvl.Renderer.Copy(tl.Sprite().Texture, tl.Sprite().Rect,
					&dstRec)
			}
			if DRAW_DEBUG {
				lvl.Renderer.SetDrawColor(254, 0, 0, 255)
				lvl.Renderer.DrawRect(&dstRec)
				//font := lvl.spr.GetFont("LiberationMono5")
				text := fmt.Sprintf("%02d,%02d", y, x)
				//	DrawTextAt(font, text, xpos+2, ypos+2, lvl.Renderer)
				DrawBitmapTextAtUnscaled(lvl.Renderer, lvl.spr, text, xpos+2, ypos+2)
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
