package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var FPATH = GetDataPath() + "ttf/LiberationMono-Regular.ttf"

const SMALLSIZE = 10

func DrawTextAt(font *ttf.Font, text string, x, y int32,
	renderer *sdl.Renderer) {
	surface := CreateSurfaceFromString(font, text, renderer)
	texture := renderer.CreateTextureFromSurface(surface)
	if texture == nil {
		fmt.Println("Could not create Texture of String " + text)
	}
	dst := sdl.Rect{x, y, surface.W, surface.H}
	renderer.Copy(texture, nil, &dst)
	surface.Free()
	texture.Destroy()
}

func CreateSurfaceFromString(font *ttf.Font, text string,
	renderer *sdl.Renderer) *sdl.Surface {
	surface := font.RenderText_Solid(text, sdl.Color{255, 0, 0, 255})
	return surface
}

/* bitmap font. */
func DrawBitmapTextAt(renderer *sdl.Renderer, spr *SpriteManager, text string,
	x, y int32) {
	for i, char := range text {
		str := string(char)
		sprite := spr.GetSprite(str)
		dst := sdl.Rect{x + int32(i*SCALE*6), y, SCALE * 6, SCALE * 7}
		renderer.Copy(sprite.Texture, sprite.Rect, &dst)
	}
}

func DrawBitmapTextAtUnscaled(renderer *sdl.Renderer, spr *SpriteManager, text string,
	x, y int32) {
	for i, char := range text {
		str := string(char)
		sprite := spr.GetSprite(str)
		dst := sdl.Rect{x + int32(i*6), y, 6, 7}
		renderer.Copy(sprite.Texture, sprite.Rect, &dst)
	}
}
