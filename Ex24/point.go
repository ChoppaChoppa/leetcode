package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(X, Y float64) *Point {
	return &Point{
		x: X,
		y: Y,
	}
}

func (p *Point) findDistance() float64 {
	return math.Abs(math.Abs(p.x) - math.Abs(p.y))
}

func main() {
	point := NewPoint(10, 20)

	fmt.Println(point.findDistance())
}
