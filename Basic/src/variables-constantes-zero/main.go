package main

import "fmt"

func main() {
	// declare constants
	const pi float64 = 3.14
	const pi2 float64 = 3.141516
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)
	// declare integer variable
	base := 12
	var height int = 14
	var area int
	fmt.Println(base, height, area)
	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	// Area cuadrado
	const baseSquare = 10
	areaSquare := baseSquare * baseSquare
	fmt.Println(areaSquare)
	x := 10
	y := 50
	// adition
	result := x + y
	fmt.Println("adition", result)
	// substraction
	result = x - y
	fmt.Println("substraction", result)
	// times
	result = x * y
	fmt.Println("times", result)
	// div
	result = x / y
	fmt.Println("div", result)
	// module
	result = y % x
	fmt.Println("module", result)
	// incremental
	x++
	fmt.Println("incremental", x)
	// decremental
	x--
	fmt.Println("decremental", x)
	// rectangle area
	areaRectangle := baseSquare * height
	fmt.Println("rectangleArea", areaRectangle)
	// area circle
	radius := 5
	areaCirle := pi * float64(radius) * float64(radius)
	fmt.Println("areaCircle", areaCirle)
}
