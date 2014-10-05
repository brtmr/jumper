package main

type Position struct {
	x     int32
	y     int32
	rem_x float64
	rem_y float64
	vel_x float64
	vel_y float64
	max_x int32
	max_y int32
}

func capVelocity(x float64) float64 {
	if x > TOPSPEED {
		x = TOPSPEED
	}
	if x < (-TOPSPEED) {
		x = -TOPSPEED
	}
	return x
}

func (p *Position) SetX(x int32) {
	p.x = x
}

func (p *Position) SetY(y int32) {
	p.y = y
}

func (p *Position) SetVelX(vel_x float64) {
	p.vel_x = vel_x
}

func (p *Position) SetVelY(vel_y float64) {
	p.vel_y = vel_y
}

func (p Position) X() int32 {
	return p.x
}

func (p Position) Y() int32 {
	return p.y
}

func (p Position) RemX() float64 {
	return p.rem_x
}

func (p Position) RemY() float64 {
	return p.rem_y
}

func (p Position) VelX() float64 {
	return p.vel_x
}

func (p Position) VelY() float64 {
	return p.vel_y
}

func (p Position) Update() Position {
	vel_y := p.vel_y + GRAVITY
	vel_y = capVelocity(vel_y)
	vel_x := capVelocity(p.vel_x)

	var diffx int32
	diffx, rem_x := Round_diff(vel_x + p.rem_x)
	x := p.x + diffx

	var diffy int32
	diffy, rem_y := Round_diff(vel_y + p.rem_y)
	y := p.y + diffy

	if x < 0 {
		x = 0
	}
	if x > p.max_x {
		x = p.max_x
	}
	if y < 0 {
		y = 0
	}
	if y > p.max_y {
		y = p.max_y
	}

	return Position{x, y, rem_x, rem_y, vel_x, vel_y, p.max_x, p.max_y}
}

func InterpolatePos(current, previous Position, alpha float64) Position {
	old_x := previous.X()
	new_x := current.X()
	old_rem_x := previous.RemX()
	new_rem_x := current.RemX()
	diff_x := (float64(new_x-old_x) + (new_rem_x - old_rem_x)) * (alpha)
	int_x, _ := Round_diff(old_rem_x + diff_x)

	old_y := previous.Y()
	new_y := current.Y()
	old_rem_y := previous.RemY()
	new_rem_y := current.RemY()
	diff_y := (float64(new_y-old_y) + (new_rem_y - old_rem_y)) * (alpha)
	int_y, _ := Round_diff(old_rem_y + diff_y)

	return Init_pos(old_x+int_x, old_y+int_y, previous.max_x, previous.max_y)
}

func Init_pos(x, y, max_x, max_y int32) Position {
	return Position{x, y, 0, 0, 0, 0, max_x, max_y}
}
