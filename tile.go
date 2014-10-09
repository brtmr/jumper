package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

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

/*
 * The following types & functions are for automatically parsing the
 * data/tiles/tiles.json files and creating the prototypes
 */

type TileCreator struct {
	tlmap map[string]*Tile
	idmap map[int]string
}

func InitTileCreator(spr *SpriteManager) TileCreator {
	json_bytes, err := ioutil.ReadFile(GetDataPath() + "tiles/tiles.json")
	if err != nil {
		fmt.Println("Error while loading JSON Tile Description")
		panic(err)
	}
	var json_data_arr []JSON_prototype
	err = json.Unmarshal(json_bytes, &json_data_arr)
	if err != nil {
		fmt.Println("Error while unmarshalling Tile data")
		panic(err)
	}

	tlmap := make(map[string]*Tile)
	idmap := make(map[int]string)

	for _, prototype := range json_data_arr {
		sprite := spr.GetSprite(prototype.Name)
		tl := Tile{sprite, prototype.Health, prototype.Solid}
		if _, contains := tlmap[prototype.Name]; contains {
			errstr := fmt.Sprintf("%s is not a unique tile name, please check tile.json",
				prototype.Name)
			err := errors.New(errstr)
			panic(err)
		}
		tlmap[prototype.Name] = &tl
		if _, contains := idmap[prototype.Id]; contains {
			errstr := fmt.Sprintf("%d is not a unique tile id, please check tile.json",
				prototype.Id)
			err := errors.New(errstr)
			panic(err)
		}
		idmap[prototype.Id] = prototype.Name
	}

	return TileCreator{tlmap, idmap}
}

/* for programmatically creating levels */
func TileFromPrototypeByName(protoname string) *Tile {
	return nil //TODO
}

/* for loading tiles from a file */
func TileFromPrototypeById(id int) *Tile {
	return nil //TODO
}

type JSON_prototype struct {
	Health int
	Id     int
	Name   string
	Sprite string
	Solid  bool
}
