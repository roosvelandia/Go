package main

import "fmt"

// especify input chan<- or output <-chan
func say(text string, c chan<- string) {
	c <- text
	fmt.Println("in function", text)
}
func main() {
	c := make(chan string, 1)
	fmt.Println("Hi")
	go say("Bye", c)

	fmt.Println(<-c)
}
