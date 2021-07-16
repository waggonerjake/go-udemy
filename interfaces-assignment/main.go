package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

func main() {
	tri := triangle{2, 3}
	sq := square{4}
	printArea(sq)
	printArea(tri)
}

func printArea(s shape) {
	fmt.Printf("Area of shape is: %f\n", s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * .5 * t.height
}
