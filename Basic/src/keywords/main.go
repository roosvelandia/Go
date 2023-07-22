package main

import "fmt"

func main() {
	// Defer sends at the end
	defer fmt.Println("Hello")
	fmt.Println("World")
	// Continue & break
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		// Continue
		if i == 2 {
			fmt.Println("is 2")
			continue
		}
		// Break
		if i == 8 {
			fmt.Println("is 8")
			break
		}

	}
}
