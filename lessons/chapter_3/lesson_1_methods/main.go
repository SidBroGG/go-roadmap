package main

import (
	"fmt"
	"math"
)

type Point struct {
	x	float64
	y	float64
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow(p.x - q.x, 2) + math.Pow(p.y - q.y, 2))
}

func (p *Point) ChangeX(new_x int) {
	p.x = float64(new_x)
}

func (p *Point) IncrementY() {
	p.y++
}

func main() {
	p1 := Point {
		x: 52,
		y: 10,
	}

	p2 := Point {}
	
	fmt.Println("Result", p1.Distance(p2))

	fmt.Println("Old x", p1.x)
	p1.ChangeX(11)
	fmt.Println("Nex x", p1.x)

	fmt.Println("Old y", p1.y)
	p1.IncrementY()
	fmt.Println("New y", p1.y)
}
