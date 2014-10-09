package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

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
	prototypemap map[string]JSON_prototype
	idmap        map[int]string
	spr          *SpriteManager
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

	prototypemap := make(map[string]JSON_prototype)
	idmap := make(map[int]string)

	for _, prototype := range json_data_arr {
		if _, contains := prototypemap[prototype.Name]; contains {
			errstr := fmt.Sprintf("%s is not a unique tile name, please check tile.json",
				prototype.Name)
			err := errors.New(errstr)
			panic(err)
		}
		prototypemap[prototype.Name] = prototype
		if _, contains := idmap[prototype.Id]; contains {
			errstr := fmt.Sprintf("%d is not a unique tile id, please check tile.json",
				prototype.Id)
			err := errors.New(errstr)
			panic(err)
		}
		idmap[prototype.Id] = prototype.Name
	}

	return TileCreator{prototypemap, idmap, spr}
}

func (tc TileCreator) tileFromPrototype(prototype JSON_prototype) Tile {
	sprite := tc.spr.GetSprite(prototype.Sprite)
	return Tile{sprite, prototype.Health, prototype.Solid}
}

/* for programmatically creating levels */
func (tc TileCreator) TileByName(protoname string) Tile {
	if prototype, contains := tc.prototypemap[protoname]; contains {
		return tc.tileFromPrototype(prototype)
	} else {
		errstr := fmt.Sprintf("TileCreator: requested prototype %s does not exist\n",
			protoname)
		err := errors.New(errstr)
		panic(err)
	}
}

/* for loading tiles from a file */
func (tc TileCreator) TileById(id int) Tile {
	if protoname, contains := tc.idmap[id]; contains {
		return tc.TileByName(protoname)
	} else {
		errstr := fmt.Sprintf("TileCreator: requested id %d does not exist\n",
			id)
		err := errors.New(errstr)
		panic(err)
	}

}

type JSON_prototype struct {
	Health int
	Id     int
	Name   string
	Sprite string
	Solid  bool
}
