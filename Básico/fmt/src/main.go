package main

import "fmt"

func main() {
	helloMessage := "hello, "
	worldMessage := "world"
	// PrintLn
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	// Printf
	name := "Roosvell"
	lastName := "Velandia"
	age := 33
	// v if unknown data type
	fmt.Printf("name: %s lastname: %v Is %d years old\n", name, lastName, age)
	// Sprintf save in a String variable
	message := fmt.Sprintf("name: %s lastname: %v Is %d years old", name, lastName, age)
	fmt.Println(message)
	// dataType of variable
	fmt.Printf("name Type: %T\n", name)
	fmt.Printf("age Type: %T", age)
}
