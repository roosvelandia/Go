package main

import "fmt"

func normalFunction(message string) {
	fmt.Printf("%s in function \n", message)
}

func triplArguments(a, b int, c string) {
	fmt.Printf("primero %v segundo %v tercero %s \n", a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func Doublereturn(a int) (c, d int) {
	return a, a * 2
}
func main() {
	// print 3 times
	fmt.Println("Hello World")
	fmt.Println("Hello World 2")
	fmt.Println("Hello World 3")
	//use a function
	normalFunction("hi there")
	// use a function with several parameters
	triplArguments(1, 2, "triple")
	// function that returns something
	fmt.Printf("2 times 2 is %v \n", returnValue(2))
	// Function with several return values
	value1, value2 := Doublereturn(3)
	fmt.Printf("2 times %v is %v \n", value1, value2)
	// Just one value from function
	value3, _ := Doublereturn(3)
	fmt.Printf("check only one value %v ", value3)

}
