package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}
type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.side * s.side
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	s := square{5}
	t := triangle{height: 10, base: 50}
	printArea(s)
	printArea(t)
}
