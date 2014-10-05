package main

import (
	//	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	previousPos Position
	currentPos  Position
	drawPos     Position
	sprites     []Sprite
	renderer    *sdl.Renderer
	camera      *Camera
	lvl         Level
	running     bool
	jumping     bool
	w           int32
	h           int32
	direction   int
}

func Init_player(spr SpriteManager, renderer *sdl.Renderer,
	level Level, camera *Camera) Player {
	sprites := []Sprite{spr.GetSprite("player_anim_0"),
		spr.GetSprite("player_anim_1"),
		spr.GetSprite("player_anim_2")}
	pos := Init_pos(50, 50,
		int32(level.DimX())-SCALE*sprites[0].W,
		int32(level.DimY())-SCALE*sprites[0].H)
	return Player{pos, pos, pos, sprites, renderer, camera, level, false, true,
		sprites[0].W * SCALE, sprites[0].H * SCALE, DIRECTION_RIGHT}
}

func (p *Player) Draw() {
	var i int
	if p.running {
		i = int((sdl.GetTicks() / 150) % 4)
		if i == 3 {
			i = 1
		}
	} else {
		i = 0
	}
	if p.jumping {
		i = 2
	}
	sprite := p.sprites[i]
	X := p.drawPos.X()
	Y := p.drawPos.Y()
	dstRec := sdl.Rect{X, Y, SCALE * sprite.Rect.W, SCALE * sprite.Rect.H}
	if p.direction == DIRECTION_RIGHT {
		p.renderer.Copy(sprite.Texture, sprite.Rect, &dstRec)
	} else {
		p.renderer.CopyEx(sprite.Texture, sprite.Rect, &dstRec,
			0, nil, sdl.FLIP_HORIZONTAL)
	}
}

func (p *Player) SetDirection(direction int) {
	switch direction {
	case DIRECTION_RIGHT:
		p.running = true
		p.currentPos.SetVelX(RUNSPEED)
		p.direction = DIRECTION_RIGHT
		break
	case DIRECTION_LEFT:
		p.running = true
		p.currentPos.SetVelX(-RUNSPEED)
		p.direction = DIRECTION_LEFT
		break
	case STOP:
		p.currentPos.SetVelX(0)
		p.running = false
		break
	}
}

func (p *Player) Jump() {
	if !p.jumping {
		p.currentPos.SetVelY(JUMPSPEED)
		p.jumping = true
	}
}

func (p *Player) Update() {

	p.previousPos = p.currentPos
	p.currentPos = p.previousPos.Update()

	//collision handling
	left_x_index := p.currentPos.X() / Tile_size
	right_x_index := (p.currentPos.X() + p.w) / Tile_size
	top_y_index := p.currentPos.Y() / Tile_size
	bottom_y_index := (p.currentPos.Y() + p.h) / Tile_size

	for i := top_y_index; i <= bottom_y_index; i++ {
		for j := left_x_index; j <= right_x_index; j++ {
			p.Collide(i, j)
		}
	}

	//fmt.Printf("current velocity: x %.2f y %.2f\n", p.currentPos.GetVelX(), p.currentPos.GetVelY())
}

func (p *Player) Collide(i, j int32) {
	if p.lvl.IsSolid(int(i), int(j)) {
		x, y := getMTV(
			sdl.Rect{p.currentPos.X(), p.currentPos.Y(), p.w, p.h},
			sdl.Rect{j * Tile_size, i * Tile_size, Tile_size, Tile_size})
		if x != 0 {
			p.currentPos.SetVelX(0)
		}
		if y != 0 {
			p.currentPos.SetVelY(0)
		}
		if y < 0 {
			p.jumping = false
		}
		p.currentPos.SetX(p.currentPos.X() + x)
		p.currentPos.SetY(p.currentPos.Y() + y)
	}
}

func getMTV(p, t sdl.Rect) (x, y int32) {
	//player completely out of the box
	if (p.X+p.W < t.X) || (t.X+t.W < p.X) {
		return 0, 0
	}
	if (p.Y+p.H < t.Y) || (t.Y+t.H < p.Y) {
		return 0, 0
	}
	//player is inside the box
	left_x := t.X - (p.X + p.W)
	right_x := (t.X + t.W) - p.X
	if left_x*left_x < right_x*right_x {
		x = left_x
	} else {
		x = right_x
	}
	top_y := t.Y - (p.Y + p.H)
	bottom_y := (t.Y + t.H) - p.Y
	if top_y*top_y < bottom_y*bottom_y {
		y = top_y
	} else {
		y = bottom_y
	}
	if x*x < y*y {
		return x, 0
	} else {
		return 0, y
	}
}

func (p *Player) Interpolate(alpha float64) {
	p.drawPos = InterpolatePos(p.currentPos, p.previousPos, alpha)
}

func (p *Player) setCamera() {
}
