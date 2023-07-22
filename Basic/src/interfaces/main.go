package main

import "fmt"

type figures2D interface {
	area() float64
}
type square struct {
	base float64
}
type rectangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}
func (r rectangle) area() float64 {
	return r.base * r.height
}

func calculate(f figures2D) {
	fmt.Println("the area is: ", f.area())
}
func main() {
	mySquare := square{base: 4}
	myRectangle := rectangle{base: 4, height: 8}
	calculate(mySquare)
	calculate(myRectangle)
	// list of interfaces
	myInterface := []interface{}{"hello", 35, 4.903}
	fmt.Println(myInterface...)
}
