package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	myPc.ram *= 2
	fmt.Println(myPc.ram)
}

func main() {
	a := 50
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 10
	fmt.Println(a)
	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)
	myPc.ping()
	myPc.duplicateRam()
	myPc.duplicateRam()
	fmt.Println(myPc)
}
