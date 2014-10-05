package main

import (
	"encoding/json"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"io/ioutil"
	"os"
)

//all the necessary data to create a Sprite
//will be imported using the sprites.json config file

type JSON_SpriteData struct {
	Id_string string
	File      string
	Width     int32
	Height    int32
	Offset_x  int32
	Offset_y  int32
}

type Sprite struct {
	Texture *sdl.Texture
	Rect    *sdl.Rect
	Name    string
	W       int32
	H       int32
}

type SpriteManager struct {
	sprmap  map[string]Sprite
	fontMap map[string]*ttf.Font
}

//methods
func (spr SpriteManager) GetSprite(id string) Sprite {
	return spr.sprmap[id]
}

func (spr SpriteManager) GetFont(id string) *ttf.Font {
	return spr.fontMap[id]
}

//functions

/*
 * inits the tilemanager given a json config file.
 * TODO: when reading from the same file, do not create the same texture
 * again.
 */
func Init_from_json(json_path string, renderer *sdl.Renderer) SpriteManager {
	json_bytes, err := ioutil.ReadFile(json_path)
	if err != nil {
		fmt.Println("Error while loading JSON Sprite Description")
		panic(err)
	}
	var json_data_arr []JSON_SpriteData
	err = json.Unmarshal(json_bytes, &json_data_arr)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON data")
		panic(err)
	}
	//load all unique files
	unique_textures := make(map[string]*sdl.Texture)
	for _, tile_object := range json_data_arr {
		if unique_textures[tile_object.File] == nil {
			surface := img.Load(GetDataPath() + tile_object.File)
			//create a surface
			if surface == nil {
				fmt.Println(
					"Unable to load sprite " + tile_object.Id_string + "from file " + tile_object.File)
				os.Exit(1)
			}
			//create a texture
			texture := renderer.CreateTextureFromSurface(surface)
			if texture == nil {
				fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", sdl.GetError())
				os.Exit(1)
			}
			//put it in the map, using the filename as key
			unique_textures[tile_object.File] = texture
			surface.Free()
		}
	}
	//the map that will represent the tilemanager
	m := make(map[string]Sprite)
	//iterate over the json objects, and create sprites
	for _, tile_object := range json_data_arr {
		//create the clipping rectangle for the texture
		rect := sdl.Rect{tile_object.Offset_x, tile_object.Offset_y,
			tile_object.Width, tile_object.Height}
		spr := Sprite{unique_textures[tile_object.File], &rect,
			tile_object.Id_string, rect.W, rect.H}
		m[tile_object.Id_string] = spr
	}

	fontMap := Init_Fonts()

	return SpriteManager{m, fontMap}
}

func (spr *SpriteManager) TearDown() {
	for _, v := range spr.sprmap {
		v.Texture.Destroy()
	}

}

func Init_Fonts() map[string]*ttf.Font {
	font, err := ttf.OpenFont(FPATH, SMALLSIZE)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Font couldn't be loaded")
		os.Exit(-1)
	}
	m := make(map[string]*ttf.Font)
	m["LiberationMono5"] = font
	return m
}
