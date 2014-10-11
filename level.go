package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Level struct {
	tiles    [MAPHEIGHT][MAPWIDTH]Tile
	camera   *Camera
	Renderer *sdl.Renderer
	spr      *SpriteManager
	//dimension in tiles
	dimx int
	dimy int
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
				ret := lvl.Renderer.Copy(tl.Sprite().Texture, tl.Sprite().Rect, &dstRec)
				if ret != 0 {
					SdlPanic()
				}
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
	if i >= lvl.dimy || j >= lvl.dimx {
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
