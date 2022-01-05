package main

import "fmt"

type shapes interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tA := triangle{height: 10, base: 10}
	sA := square{sideLength: 10}
	printArea(tA)
	printArea(sA)

}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(s shapes) {
	fmt.Println(s.getArea())
}
