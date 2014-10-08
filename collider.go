package main

type Collider interface {
	GetCurrentPosition() Position
	GetPreviousPosition() Position
	W() int
	H() int
}
