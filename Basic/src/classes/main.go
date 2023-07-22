package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
}

func main() {
	// instantiate
	myCar := car{brand: "ford", year: 2023}
	fmt.Println(myCar)
	// other way
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
