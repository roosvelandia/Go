package main

import (
	"fmt"
	"sync"
	"time"
)

func saySimple(text string) {
	fmt.Println(text)
}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
func main() {
	// not recommended
	saySimple("hi")
	go saySimple("you")

	time.Sleep(time.Second * 1)

	// use wiat group
	var wg sync.WaitGroup

	fmt.Println("Hi better")
	wg.Add(1)
	go say("you", &wg)
	wg.Wait()

	// Anonymous function
	go func(text string) {
		fmt.Println(text, " anonymous function")
	}("bye")
	time.Sleep(time.Second * 1)
}
