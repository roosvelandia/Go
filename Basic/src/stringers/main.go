package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("i have %d GB RAM, %d GB Dik and is a %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 100, brand: "mac"}
	fmt.Println(myPc)
}
