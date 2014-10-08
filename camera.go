package main

type Camera struct {
	//Points to the upper right corner of the screen
	x int32
	y int32
}

func (c Camera) X() int32 {
	return c.x
}

func (c Camera) Y() int32 {
	return c.y
}

func (c *Camera) SetX(x int32) {
	c.x = x
}

func (c *Camera) SetY(y int32) {
	c.y = y
}
